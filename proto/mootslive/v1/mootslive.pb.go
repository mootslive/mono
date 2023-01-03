// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: mootslive/v1/mootslive.proto

package mootslivepbv1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
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

var File_mootslive_v1_mootslive_proto protoreflect.FileDescriptor

var file_mootslive_v1_mootslive_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x6d, 0x6f, 0x6f, 0x74, 0x73, 0x6c, 0x69, 0x76, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x6d,
	0x6f, 0x6f, 0x74, 0x73, 0x6c, 0x69, 0x76, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c,
	0x6d, 0x6f, 0x6f, 0x74, 0x73, 0x6c, 0x69, 0x76, 0x65, 0x2e, 0x76, 0x31, 0x22, 0x12, 0x0a, 0x10,
	0x47, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x22, 0x3f, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2a, 0x0a, 0x11, 0x78, 0x5f, 0x63, 0x6c, 0x61, 0x63, 0x6b,
	0x73, 0x5f, 0x6f, 0x76, 0x65, 0x72, 0x68, 0x65, 0x61, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0f, 0x78, 0x43, 0x6c, 0x61, 0x63, 0x6b, 0x73, 0x4f, 0x76, 0x65, 0x72, 0x68, 0x65, 0x61,
	0x64, 0x32, 0x5e, 0x0a, 0x0c, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x4e, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1e,
	0x2e, 0x6d, 0x6f, 0x6f, 0x74, 0x73, 0x6c, 0x69, 0x76, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65,
	0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f,
	0x2e, 0x6d, 0x6f, 0x6f, 0x74, 0x73, 0x6c, 0x69, 0x76, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65,
	0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x42, 0x22, 0x5a, 0x20, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6d, 0x6f, 0x6f, 0x74, 0x73,
	0x6c, 0x69, 0x76, 0x65, 0x2f, 0x76, 0x31, 0x3b, 0x6d, 0x6f, 0x6f, 0x74, 0x73, 0x6c, 0x69, 0x76,
	0x65, 0x70, 0x62, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
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

var file_mootslive_v1_mootslive_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_mootslive_v1_mootslive_proto_goTypes = []interface{}{
	(*GetStatusRequest)(nil),  // 0: mootslive.v1.GetStatusRequest
	(*GetStatusResponse)(nil), // 1: mootslive.v1.GetStatusResponse
}
var file_mootslive_v1_mootslive_proto_depIdxs = []int32{
	0, // 0: mootslive.v1.AdminService.GetStatus:input_type -> mootslive.v1.GetStatusRequest
	1, // 1: mootslive.v1.AdminService.GetStatus:output_type -> mootslive.v1.GetStatusResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
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
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_mootslive_v1_mootslive_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
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
