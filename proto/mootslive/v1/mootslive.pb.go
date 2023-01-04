// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: mootslive/v1/mootslive.proto

package mootslivepbv1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type GetStatusRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetStatusRequest) Reset() {
	*x = GetStatusRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mootslive_v1_mootslive_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetStatusRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetStatusRequest) ProtoMessage() {}

func (x *GetStatusRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mootslive_v1_mootslive_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetStatusRequest.ProtoReflect.Descriptor instead.
func (*GetStatusRequest) Descriptor() ([]byte, []int) {
	return file_mootslive_v1_mootslive_proto_rawDescGZIP(), []int{0}
}

type GetStatusResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	XClacksOverhead string `protobuf:"bytes,1,opt,name=x_clacks_overhead,json=xClacksOverhead,proto3" json:"x_clacks_overhead,omitempty"`
}

func (x *GetStatusResponse) Reset() {
	*x = GetStatusResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mootslive_v1_mootslive_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetStatusResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetStatusResponse) ProtoMessage() {}

func (x *GetStatusResponse) ProtoReflect() protoreflect.Message {
	mi := &file_mootslive_v1_mootslive_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetStatusResponse.ProtoReflect.Descriptor instead.
func (*GetStatusResponse) Descriptor() ([]byte, []int) {
	return file_mootslive_v1_mootslive_proto_rawDescGZIP(), []int{1}
}

func (x *GetStatusResponse) GetXClacksOverhead() string {
	if x != nil {
		return x.XClacksOverhead
	}
	return ""
}

type GetMeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetMeRequest) Reset() {
	*x = GetMeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mootslive_v1_mootslive_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMeRequest) ProtoMessage() {}

func (x *GetMeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mootslive_v1_mootslive_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMeRequest.ProtoReflect.Descriptor instead.
func (*GetMeRequest) Descriptor() ([]byte, []int) {
	return file_mootslive_v1_mootslive_proto_rawDescGZIP(), []int{2}
}

type GetMeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
}

func (x *GetMeResponse) Reset() {
	*x = GetMeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mootslive_v1_mootslive_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMeResponse) ProtoMessage() {}

func (x *GetMeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_mootslive_v1_mootslive_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMeResponse.ProtoReflect.Descriptor instead.
func (*GetMeResponse) Descriptor() ([]byte, []int) {
	return file_mootslive_v1_mootslive_proto_rawDescGZIP(), []int{3}
}

func (x *GetMeResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *GetMeResponse) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

// OAuth2State contains bits of state we need the client to hold during the
// OAuth2 3-legged flow. These values aren't something that need to be kept
// secret from the client.
type OAuth2State struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	State            string `protobuf:"bytes,1,opt,name=state,proto3" json:"state,omitempty"`
	PkceCodeVerifier string `protobuf:"bytes,2,opt,name=pkce_code_verifier,json=pkceCodeVerifier,proto3" json:"pkce_code_verifier,omitempty"`
}

func (x *OAuth2State) Reset() {
	*x = OAuth2State{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mootslive_v1_mootslive_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OAuth2State) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OAuth2State) ProtoMessage() {}

func (x *OAuth2State) ProtoReflect() protoreflect.Message {
	mi := &file_mootslive_v1_mootslive_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OAuth2State.ProtoReflect.Descriptor instead.
func (*OAuth2State) Descriptor() ([]byte, []int) {
	return file_mootslive_v1_mootslive_proto_rawDescGZIP(), []int{4}
}

func (x *OAuth2State) GetState() string {
	if x != nil {
		return x.State
	}
	return ""
}

func (x *OAuth2State) GetPkceCodeVerifier() string {
	if x != nil {
		return x.PkceCodeVerifier
	}
	return ""
}

type BeginTwitterAuthRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *BeginTwitterAuthRequest) Reset() {
	*x = BeginTwitterAuthRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mootslive_v1_mootslive_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BeginTwitterAuthRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BeginTwitterAuthRequest) ProtoMessage() {}

func (x *BeginTwitterAuthRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mootslive_v1_mootslive_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BeginTwitterAuthRequest.ProtoReflect.Descriptor instead.
func (*BeginTwitterAuthRequest) Descriptor() ([]byte, []int) {
	return file_mootslive_v1_mootslive_proto_rawDescGZIP(), []int{5}
}

type BeginTwitterAuthResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RedirectUrl string       `protobuf:"bytes,1,opt,name=redirect_url,json=redirectUrl,proto3" json:"redirect_url,omitempty"`
	State       *OAuth2State `protobuf:"bytes,2,opt,name=state,proto3" json:"state,omitempty"`
}

func (x *BeginTwitterAuthResponse) Reset() {
	*x = BeginTwitterAuthResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mootslive_v1_mootslive_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BeginTwitterAuthResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BeginTwitterAuthResponse) ProtoMessage() {}

func (x *BeginTwitterAuthResponse) ProtoReflect() protoreflect.Message {
	mi := &file_mootslive_v1_mootslive_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BeginTwitterAuthResponse.ProtoReflect.Descriptor instead.
func (*BeginTwitterAuthResponse) Descriptor() ([]byte, []int) {
	return file_mootslive_v1_mootslive_proto_rawDescGZIP(), []int{6}
}

func (x *BeginTwitterAuthResponse) GetRedirectUrl() string {
	if x != nil {
		return x.RedirectUrl
	}
	return ""
}

func (x *BeginTwitterAuthResponse) GetState() *OAuth2State {
	if x != nil {
		return x.State
	}
	return nil
}

type FinishTwitterAuthRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	State *OAuth2State `protobuf:"bytes,1,opt,name=state,proto3" json:"state,omitempty"`
	// received_state is the state url parameter received during the callback to
	// the webapp from the relying party.
	ReceivedState string `protobuf:"bytes,2,opt,name=received_state,json=receivedState,proto3" json:"received_state,omitempty"`
	ReceivedCode  string `protobuf:"bytes,3,opt,name=received_code,json=receivedCode,proto3" json:"received_code,omitempty"`
}

func (x *FinishTwitterAuthRequest) Reset() {
	*x = FinishTwitterAuthRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mootslive_v1_mootslive_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FinishTwitterAuthRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FinishTwitterAuthRequest) ProtoMessage() {}

func (x *FinishTwitterAuthRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mootslive_v1_mootslive_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FinishTwitterAuthRequest.ProtoReflect.Descriptor instead.
func (*FinishTwitterAuthRequest) Descriptor() ([]byte, []int) {
	return file_mootslive_v1_mootslive_proto_rawDescGZIP(), []int{7}
}

func (x *FinishTwitterAuthRequest) GetState() *OAuth2State {
	if x != nil {
		return x.State
	}
	return nil
}

func (x *FinishTwitterAuthRequest) GetReceivedState() string {
	if x != nil {
		return x.ReceivedState
	}
	return ""
}

func (x *FinishTwitterAuthRequest) GetReceivedCode() string {
	if x != nil {
		return x.ReceivedCode
	}
	return ""
}

type FinishTwitterAuthResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Me string `protobuf:"bytes,1,opt,name=me,proto3" json:"me,omitempty"`
}

func (x *FinishTwitterAuthResponse) Reset() {
	*x = FinishTwitterAuthResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mootslive_v1_mootslive_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FinishTwitterAuthResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FinishTwitterAuthResponse) ProtoMessage() {}

func (x *FinishTwitterAuthResponse) ProtoReflect() protoreflect.Message {
	mi := &file_mootslive_v1_mootslive_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FinishTwitterAuthResponse.ProtoReflect.Descriptor instead.
func (*FinishTwitterAuthResponse) Descriptor() ([]byte, []int) {
	return file_mootslive_v1_mootslive_proto_rawDescGZIP(), []int{8}
}

func (x *FinishTwitterAuthResponse) GetMe() string {
	if x != nil {
		return x.Me
	}
	return ""
}

var File_mootslive_v1_mootslive_proto protoreflect.FileDescriptor

var file_mootslive_v1_mootslive_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x6d, 0x6f, 0x6f, 0x74, 0x73, 0x6c, 0x69, 0x76, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x6d,
	0x6f, 0x6f, 0x74, 0x73, 0x6c, 0x69, 0x76, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c,
	0x6d, 0x6f, 0x6f, 0x74, 0x73, 0x6c, 0x69, 0x76, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x12, 0x0a,
	0x10, 0x47, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x22, 0x3f, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2a, 0x0a, 0x11, 0x78, 0x5f, 0x63, 0x6c, 0x61, 0x63,
	0x6b, 0x73, 0x5f, 0x6f, 0x76, 0x65, 0x72, 0x68, 0x65, 0x61, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0f, 0x78, 0x43, 0x6c, 0x61, 0x63, 0x6b, 0x73, 0x4f, 0x76, 0x65, 0x72, 0x68, 0x65,
	0x61, 0x64, 0x22, 0x0e, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x22, 0x5a, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x51,
	0x0a, 0x0b, 0x4f, 0x41, 0x75, 0x74, 0x68, 0x32, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x74,
	0x61, 0x74, 0x65, 0x12, 0x2c, 0x0a, 0x12, 0x70, 0x6b, 0x63, 0x65, 0x5f, 0x63, 0x6f, 0x64, 0x65,
	0x5f, 0x76, 0x65, 0x72, 0x69, 0x66, 0x69, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x10, 0x70, 0x6b, 0x63, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x56, 0x65, 0x72, 0x69, 0x66, 0x69, 0x65,
	0x72, 0x22, 0x19, 0x0a, 0x17, 0x42, 0x65, 0x67, 0x69, 0x6e, 0x54, 0x77, 0x69, 0x74, 0x74, 0x65,
	0x72, 0x41, 0x75, 0x74, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x6e, 0x0a, 0x18,
	0x42, 0x65, 0x67, 0x69, 0x6e, 0x54, 0x77, 0x69, 0x74, 0x74, 0x65, 0x72, 0x41, 0x75, 0x74, 0x68,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x72, 0x65, 0x64, 0x69,
	0x72, 0x65, 0x63, 0x74, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x72, 0x65, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x55, 0x72, 0x6c, 0x12, 0x2f, 0x0a, 0x05, 0x73,
	0x74, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x6d, 0x6f, 0x6f,
	0x74, 0x73, 0x6c, 0x69, 0x76, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x41, 0x75, 0x74, 0x68, 0x32,
	0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x22, 0x97, 0x01, 0x0a,
	0x18, 0x46, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x54, 0x77, 0x69, 0x74, 0x74, 0x65, 0x72, 0x41, 0x75,
	0x74, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2f, 0x0a, 0x05, 0x73, 0x74, 0x61,
	0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x6d, 0x6f, 0x6f, 0x74, 0x73,
	0x6c, 0x69, 0x76, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x41, 0x75, 0x74, 0x68, 0x32, 0x53, 0x74,
	0x61, 0x74, 0x65, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x72, 0x65,
	0x63, 0x65, 0x69, 0x76, 0x65, 0x64, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0d, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x64, 0x53, 0x74, 0x61, 0x74,
	0x65, 0x12, 0x23, 0x0a, 0x0d, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x64, 0x5f, 0x63, 0x6f,
	0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76,
	0x65, 0x64, 0x43, 0x6f, 0x64, 0x65, 0x22, 0x2b, 0x0a, 0x19, 0x46, 0x69, 0x6e, 0x69, 0x73, 0x68,
	0x54, 0x77, 0x69, 0x74, 0x74, 0x65, 0x72, 0x41, 0x75, 0x74, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x6d, 0x65, 0x32, 0x5e, 0x0a, 0x0c, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x4e, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x1e, 0x2e, 0x6d, 0x6f, 0x6f, 0x74, 0x73, 0x6c, 0x69, 0x76, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x47, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1f, 0x2e, 0x6d, 0x6f, 0x6f, 0x74, 0x73, 0x6c, 0x69, 0x76, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x47, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x32, 0x9e, 0x02, 0x0a, 0x0b, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x42, 0x0a, 0x05, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x12, 0x1a, 0x2e, 0x6d,
	0x6f, 0x6f, 0x74, 0x73, 0x6c, 0x69, 0x76, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x4d,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x6d, 0x6f, 0x6f, 0x74, 0x73,
	0x6c, 0x69, 0x76, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x63, 0x0a, 0x10, 0x42, 0x65, 0x67, 0x69, 0x6e,
	0x54, 0x77, 0x69, 0x74, 0x74, 0x65, 0x72, 0x41, 0x75, 0x74, 0x68, 0x12, 0x25, 0x2e, 0x6d, 0x6f,
	0x6f, 0x74, 0x73, 0x6c, 0x69, 0x76, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x65, 0x67, 0x69, 0x6e,
	0x54, 0x77, 0x69, 0x74, 0x74, 0x65, 0x72, 0x41, 0x75, 0x74, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x26, 0x2e, 0x6d, 0x6f, 0x6f, 0x74, 0x73, 0x6c, 0x69, 0x76, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x42, 0x65, 0x67, 0x69, 0x6e, 0x54, 0x77, 0x69, 0x74, 0x74, 0x65, 0x72, 0x41, 0x75,
	0x74, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x66, 0x0a, 0x11,
	0x46, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x54, 0x77, 0x69, 0x74, 0x74, 0x65, 0x72, 0x41, 0x75, 0x74,
	0x68, 0x12, 0x26, 0x2e, 0x6d, 0x6f, 0x6f, 0x74, 0x73, 0x6c, 0x69, 0x76, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x46, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x54, 0x77, 0x69, 0x74, 0x74, 0x65, 0x72, 0x41, 0x75,
	0x74, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x6d, 0x6f, 0x6f, 0x74,
	0x73, 0x6c, 0x69, 0x76, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x54,
	0x77, 0x69, 0x74, 0x74, 0x65, 0x72, 0x41, 0x75, 0x74, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x42, 0x3c, 0x5a, 0x3a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x6d, 0x6f, 0x6f, 0x74, 0x73, 0x6c, 0x69, 0x76, 0x65, 0x2f, 0x6d, 0x6f, 0x6e,
	0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6d, 0x6f, 0x6f, 0x74, 0x73, 0x6c, 0x69, 0x76,
	0x65, 0x2f, 0x76, 0x31, 0x3b, 0x6d, 0x6f, 0x6f, 0x74, 0x73, 0x6c, 0x69, 0x76, 0x65, 0x70, 0x62,
	0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_mootslive_v1_mootslive_proto_rawDescOnce sync.Once
	file_mootslive_v1_mootslive_proto_rawDescData = file_mootslive_v1_mootslive_proto_rawDesc
)

func file_mootslive_v1_mootslive_proto_rawDescGZIP() []byte {
	file_mootslive_v1_mootslive_proto_rawDescOnce.Do(func() {
		file_mootslive_v1_mootslive_proto_rawDescData = protoimpl.X.CompressGZIP(file_mootslive_v1_mootslive_proto_rawDescData)
	})
	return file_mootslive_v1_mootslive_proto_rawDescData
}

var file_mootslive_v1_mootslive_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_mootslive_v1_mootslive_proto_goTypes = []interface{}{
	(*GetStatusRequest)(nil),          // 0: mootslive.v1.GetStatusRequest
	(*GetStatusResponse)(nil),         // 1: mootslive.v1.GetStatusResponse
	(*GetMeRequest)(nil),              // 2: mootslive.v1.GetMeRequest
	(*GetMeResponse)(nil),             // 3: mootslive.v1.GetMeResponse
	(*OAuth2State)(nil),               // 4: mootslive.v1.OAuth2State
	(*BeginTwitterAuthRequest)(nil),   // 5: mootslive.v1.BeginTwitterAuthRequest
	(*BeginTwitterAuthResponse)(nil),  // 6: mootslive.v1.BeginTwitterAuthResponse
	(*FinishTwitterAuthRequest)(nil),  // 7: mootslive.v1.FinishTwitterAuthRequest
	(*FinishTwitterAuthResponse)(nil), // 8: mootslive.v1.FinishTwitterAuthResponse
	(*timestamppb.Timestamp)(nil),     // 9: google.protobuf.Timestamp
}
var file_mootslive_v1_mootslive_proto_depIdxs = []int32{
	9, // 0: mootslive.v1.GetMeResponse.created_at:type_name -> google.protobuf.Timestamp
	4, // 1: mootslive.v1.BeginTwitterAuthResponse.state:type_name -> mootslive.v1.OAuth2State
	4, // 2: mootslive.v1.FinishTwitterAuthRequest.state:type_name -> mootslive.v1.OAuth2State
	0, // 3: mootslive.v1.AdminService.GetStatus:input_type -> mootslive.v1.GetStatusRequest
	2, // 4: mootslive.v1.UserService.GetMe:input_type -> mootslive.v1.GetMeRequest
	5, // 5: mootslive.v1.UserService.BeginTwitterAuth:input_type -> mootslive.v1.BeginTwitterAuthRequest
	7, // 6: mootslive.v1.UserService.FinishTwitterAuth:input_type -> mootslive.v1.FinishTwitterAuthRequest
	1, // 7: mootslive.v1.AdminService.GetStatus:output_type -> mootslive.v1.GetStatusResponse
	3, // 8: mootslive.v1.UserService.GetMe:output_type -> mootslive.v1.GetMeResponse
	6, // 9: mootslive.v1.UserService.BeginTwitterAuth:output_type -> mootslive.v1.BeginTwitterAuthResponse
	8, // 10: mootslive.v1.UserService.FinishTwitterAuth:output_type -> mootslive.v1.FinishTwitterAuthResponse
	7, // [7:11] is the sub-list for method output_type
	3, // [3:7] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_mootslive_v1_mootslive_proto_init() }
func file_mootslive_v1_mootslive_proto_init() {
	if File_mootslive_v1_mootslive_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_mootslive_v1_mootslive_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetStatusRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_mootslive_v1_mootslive_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetStatusResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_mootslive_v1_mootslive_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMeRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_mootslive_v1_mootslive_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMeResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_mootslive_v1_mootslive_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OAuth2State); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_mootslive_v1_mootslive_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BeginTwitterAuthRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_mootslive_v1_mootslive_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BeginTwitterAuthResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_mootslive_v1_mootslive_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FinishTwitterAuthRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_mootslive_v1_mootslive_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FinishTwitterAuthResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_mootslive_v1_mootslive_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_mootslive_v1_mootslive_proto_goTypes,
		DependencyIndexes: file_mootslive_v1_mootslive_proto_depIdxs,
		MessageInfos:      file_mootslive_v1_mootslive_proto_msgTypes,
	}.Build()
	File_mootslive_v1_mootslive_proto = out.File
	file_mootslive_v1_mootslive_proto_rawDesc = nil
	file_mootslive_v1_mootslive_proto_goTypes = nil
	file_mootslive_v1_mootslive_proto_depIdxs = nil
}
