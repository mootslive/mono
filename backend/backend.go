package backend

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/mootslive/mono/backend/trace"
	"time"

	"github.com/jackc/pgx/v4"
	"github.com/mootslive/mono/backend/db"
	"github.com/segmentio/ksuid"
	"github.com/zmb3/spotify/v2"
	spotifyauth "github.com/zmb3/spotify/v2/auth"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"golang.org/x/exp/slog"
	"golang.org/x/oauth2"
)

type SpotifyPoller struct {
	queries db.TXQuerier
	log     *slog.Logger
}

func NewSpotifyPoller(log *slog.Logger, queries db.TXQuerier) *SpotifyPoller {
	return &SpotifyPoller{
		log:     log,
		queries: queries,
	}
}

func (sp *SpotifyPoller) Run(ctx context.Context) error {
	sp.log.Info("starting poller")

	for {
		sp.log.Info("running account scan")
		accounts, err := sp.queries.GetSpotifyAccountsForScanning(ctx)
		if err != nil {
			return fmt.Errorf("fetching accounts: %w", err)
		}

		for _, account := range accounts {
			if err := sp.ScanAccount(ctx, account.SpotifyUserID); err != nil {
				return fmt.Errorf("scanning account: %w", err)
			}
		}

		select {
		// Simple ten second backoff
		case <-time.After(time.Second * 10):
			continue
		case <-ctx.Done():
			sp.log.Info("context cancelled, stopping poller")
			return ctx.Err()
		}
	}
}

const (
	sourceSpotify = "spotify"
)

func (sp *SpotifyPoller) ScanAccount(
	ctx context.Context, spotifyUserID string,
) error {
	ctx, span := trace.Start(ctx, "backend/SpotifyPoller.ScanAccount")
	defer span.End()

	commit, rollback, tx, err := sp.queries.BeginTx(ctx)
	if err != nil {
		return fmt.Errorf("opening tx: %w", err)
	}
	defer func() {
		if err := rollback(context.Background()); err != nil {
			if !errors.Is(err, pgx.ErrTxClosed) {
				sp.log.Error("failed to rollback", err)
			}
		}
	}()

	account, err := tx.SelectSpotifyAccountForUpdate(ctx, spotifyUserID)
	if err != nil {
		return fmt.Errorf("locking account: %w", err)
	}

	var afterEpochMs int64 = 0
	if account.LastListenedAt.Valid {
		afterEpochMs = (account.LastListenedAt.Time.Add(time.Second).Unix()) * 1000
	}

	client := clientForSpotifyAccount(ctx, account)
	played, err := client.PlayerRecentlyPlayedOpt(ctx, &spotify.RecentlyPlayedOptions{
		Limit:        50,
		AfterEpochMs: afterEpochMs,
	})
	if err != nil {
		return fmt.Errorf("fetching recently played: %w", err)
	}

	var listenedAt *time.Time
	for _, track := range played {
		// TODO: Batch insert these :)
		track := track
		sp.log.Debug("recording listen", "user_id", account.UserID, "track_name", track.Track.Name, "listened_at", track.PlayedAt)
		err := tx.CreateListen(ctx, db.CreateListenParams{
			ID:         ksuid.New().String(),
			UserID:     account.UserID,
			CreatedAt:  time.Now(),
			Source:     sourceSpotify,
			Isrc:       track.Track.ExternalIDs.ISRC,
			ListenedAt: track.PlayedAt,
		})
		if err != nil {
			return fmt.Errorf("recording listen: %w", err)
		}
		if listenedAt == nil {
			listenedAt = &track.PlayedAt
		}
	}

	if listenedAt != nil {
		err := tx.UpdateSpotifyAccountListenedAt(ctx, db.UpdateSpotifyAccountListenedAtParams{
			SpotifyUserID: account.SpotifyUserID,
			LastListenedAt: sql.NullTime{
				Valid: true,
				Time:  *listenedAt,
			},
		})
		if err != nil {
			return fmt.Errorf("updating listened at: %w", err)
		}
	}

	if err := commit(ctx); err != nil {
		return err
	}

	sp.log.Info("recorded listens for user",
		slog.String("user_id", account.UserID),
		slog.Int("count", len(played)),
	)

	return nil
}

func clientForSpotifyAccount(
	ctx context.Context, account db.SpotifyAccount,
) *spotify.Client {
	token := oauth2.Token(account.OauthToken)
	httpClient := spotifyauth.New().Client(ctx, &token)
	httpClient.Transport = otelhttp.NewTransport(httpClient.Transport)
	client := spotify.New(httpClient)
	return client
}
