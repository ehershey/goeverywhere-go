// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.28.3
// source: polylines.proto

package proto

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

type GetPolylinesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// google.protobuf.Timestamp oldest_polyline_timestamp = 1;
	// google.protobuf.Timestamp newest_polyline_timestamp = 2;
	// uint32 polyline_count = 3;
	Polylines []*ActivityPolyline `protobuf:"bytes,1,rep,name=polylines,proto3" json:"polylines,omitempty"`
}

func (x *GetPolylinesResponse) Reset() {
	*x = GetPolylinesResponse{}
	mi := &file_polylines_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetPolylinesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPolylinesResponse) ProtoMessage() {}

func (x *GetPolylinesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_polylines_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPolylinesResponse.ProtoReflect.Descriptor instead.
func (*GetPolylinesResponse) Descriptor() ([]byte, []int) {
	return file_polylines_proto_rawDescGZIP(), []int{0}
}

func (x *GetPolylinesResponse) GetPolylines() []*ActivityPolyline {
	if x != nil {
		return x.Polylines
	}
	return nil
}

type ActivityPolyline struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Polyline string `protobuf:"bytes,1,opt,name=polyline,proto3" json:"polyline,omitempty"`
}

func (x *ActivityPolyline) Reset() {
	*x = ActivityPolyline{}
	mi := &file_polylines_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ActivityPolyline) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ActivityPolyline) ProtoMessage() {}

func (x *ActivityPolyline) ProtoReflect() protoreflect.Message {
	mi := &file_polylines_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ActivityPolyline.ProtoReflect.Descriptor instead.
func (*ActivityPolyline) Descriptor() ([]byte, []int) {
	return file_polylines_proto_rawDescGZIP(), []int{1}
}

func (x *ActivityPolyline) GetPolyline() string {
	if x != nil {
		return x.Polyline
	}
	return ""
}

type GetPolylinesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetPolylinesRequest) Reset() {
	*x = GetPolylinesRequest{}
	mi := &file_polylines_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetPolylinesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPolylinesRequest) ProtoMessage() {}

func (x *GetPolylinesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_polylines_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPolylinesRequest.ProtoReflect.Descriptor instead.
func (*GetPolylinesRequest) Descriptor() ([]byte, []int) {
	return file_polylines_proto_rawDescGZIP(), []int{2}
}

var File_polylines_proto protoreflect.FileDescriptor

var file_polylines_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x70, 0x6f, 0x6c, 0x79, 0x6c, 0x69, 0x6e, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x47, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x6c, 0x79, 0x6c, 0x69, 0x6e, 0x65,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2f, 0x0a, 0x09, 0x70, 0x6f, 0x6c,
	0x79, 0x6c, 0x69, 0x6e, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x41,
	0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x50, 0x6f, 0x6c, 0x79, 0x6c, 0x69, 0x6e, 0x65, 0x52,
	0x09, 0x70, 0x6f, 0x6c, 0x79, 0x6c, 0x69, 0x6e, 0x65, 0x73, 0x22, 0x2e, 0x0a, 0x10, 0x41, 0x63,
	0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x50, 0x6f, 0x6c, 0x79, 0x6c, 0x69, 0x6e, 0x65, 0x12, 0x1a,
	0x0a, 0x08, 0x70, 0x6f, 0x6c, 0x79, 0x6c, 0x69, 0x6e, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x70, 0x6f, 0x6c, 0x79, 0x6c, 0x69, 0x6e, 0x65, 0x22, 0x15, 0x0a, 0x13, 0x47, 0x65,
	0x74, 0x50, 0x6f, 0x6c, 0x79, 0x6c, 0x69, 0x6e, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x42, 0x15, 0x5a, 0x13, 0x65, 0x72, 0x6e, 0x69, 0x65, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67,
	0x6f, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_polylines_proto_rawDescOnce sync.Once
	file_polylines_proto_rawDescData = file_polylines_proto_rawDesc
)

func file_polylines_proto_rawDescGZIP() []byte {
	file_polylines_proto_rawDescOnce.Do(func() {
		file_polylines_proto_rawDescData = protoimpl.X.CompressGZIP(file_polylines_proto_rawDescData)
	})
	return file_polylines_proto_rawDescData
}

var file_polylines_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_polylines_proto_goTypes = []any{
	(*GetPolylinesResponse)(nil), // 0: GetPolylinesResponse
	(*ActivityPolyline)(nil),     // 1: ActivityPolyline
	(*GetPolylinesRequest)(nil),  // 2: GetPolylinesRequest
}
var file_polylines_proto_depIdxs = []int32{
	1, // 0: GetPolylinesResponse.polylines:type_name -> ActivityPolyline
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_polylines_proto_init() }
func file_polylines_proto_init() {
	if File_polylines_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_polylines_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_polylines_proto_goTypes,
		DependencyIndexes: file_polylines_proto_depIdxs,
		MessageInfos:      file_polylines_proto_msgTypes,
	}.Build()
	File_polylines_proto = out.File
	file_polylines_proto_rawDesc = nil
	file_polylines_proto_goTypes = nil
	file_polylines_proto_depIdxs = nil
}
