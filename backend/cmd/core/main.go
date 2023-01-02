package main

import (
	"context"
	"fmt"
	"io"
	"os"

	"github.com/jackc/pgx/v4"
	"github.com/mootslive/mono/backend"
	"github.com/mootslive/mono/backend/db"
	"golang.org/x/exp/slog"
)

func run(out io.Writer) error {
	ctx := context.Background()
	log := slog.New(slog.NewJSONHandler(out))
	log.Info("starting mootslive backend")

	pgxCfg, err := pgx.ParseConfig("postgres://mootslive:mootslive@localhost:5432/mootslive")
	if err != nil {
		return fmt.Errorf("parsing cfg: %w", err)
	}

	conn, err := pgx.ConnectConfig(ctx, pgxCfg)
	if err != nil {
		return fmt.Errorf("connecting to pgx: %w", err)
	}

	poller := &backend.SpotifyPoller{
		DB:  db.New(conn),
		Log: log.WithGroup("poller"),
	}

	if err := poller.Run(ctx); err != nil {
		return fmt.Errorf("polling: %w", err)
	}

	return nil
}

func main() {
	if err := run(os.Stdout); err != nil {
		fmt.Println("Exiting with fatal err: ", err)
	}
}
