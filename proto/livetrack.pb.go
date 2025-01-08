// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v5.29.2
// source: livetrack.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type GetLivetrackRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetLivetrackRequest) Reset() {
	*x = GetLivetrackRequest{}
	mi := &file_livetrack_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetLivetrackRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetLivetrackRequest) ProtoMessage() {}

func (x *GetLivetrackRequest) ProtoReflect() protoreflect.Message {
	mi := &file_livetrack_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetLivetrackRequest.ProtoReflect.Descriptor instead.
func (*GetLivetrackRequest) Descriptor() ([]byte, []int) {
	return file_livetrack_proto_rawDescGZIP(), []int{0}
}

type GetLivetrackResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Polyline *string `protobuf:"bytes,1,opt,name=polyline" json:"polyline,omitempty"`
}

func (x *GetLivetrackResponse) Reset() {
	*x = GetLivetrackResponse{}
	mi := &file_livetrack_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetLivetrackResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetLivetrackResponse) ProtoMessage() {}

func (x *GetLivetrackResponse) ProtoReflect() protoreflect.Message {
	mi := &file_livetrack_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetLivetrackResponse.ProtoReflect.Descriptor instead.
func (*GetLivetrackResponse) Descriptor() ([]byte, []int) {
	return file_livetrack_proto_rawDescGZIP(), []int{1}
}

func (x *GetLivetrackResponse) GetPolyline() string {
	if x != nil && x.Polyline != nil {
		return *x.Polyline
	}
	return ""
}

var File_livetrack_proto protoreflect.FileDescriptor

var file_livetrack_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x6c, 0x69, 0x76, 0x65, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x0c, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x15, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x76, 0x65, 0x74, 0x72, 0x61, 0x63, 0x6b,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x32, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x4c, 0x69,
	0x76, 0x65, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x1a, 0x0a, 0x08, 0x70, 0x6f, 0x6c, 0x79, 0x6c, 0x69, 0x6e, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x70, 0x6f, 0x6c, 0x79, 0x6c, 0x69, 0x6e, 0x65, 0x42, 0x15, 0x5a, 0x13, 0x65,
	0x72, 0x6e, 0x69, 0x65, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x6f, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x08, 0x65, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x70, 0xe8, 0x07,
}

var (
	file_livetrack_proto_rawDescOnce sync.Once
	file_livetrack_proto_rawDescData = file_livetrack_proto_rawDesc
)

func file_livetrack_proto_rawDescGZIP() []byte {
	file_livetrack_proto_rawDescOnce.Do(func() {
		file_livetrack_proto_rawDescData = protoimpl.X.CompressGZIP(file_livetrack_proto_rawDescData)
	})
	return file_livetrack_proto_rawDescData
}

var file_livetrack_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_livetrack_proto_goTypes = []any{
	(*GetLivetrackRequest)(nil),  // 0: GetLivetrackRequest
	(*GetLivetrackResponse)(nil), // 1: GetLivetrackResponse
}
var file_livetrack_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_livetrack_proto_init() }
func file_livetrack_proto_init() {
	if File_livetrack_proto != nil {
		return
	}
	file_points_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_livetrack_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_livetrack_proto_goTypes,
		DependencyIndexes: file_livetrack_proto_depIdxs,
		MessageInfos:      file_livetrack_proto_msgTypes,
	}.Build()
	File_livetrack_proto = out.File
	file_livetrack_proto_rawDesc = nil
	file_livetrack_proto_goTypes = nil
	file_livetrack_proto_depIdxs = nil
}
