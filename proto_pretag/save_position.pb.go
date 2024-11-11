// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.28.3
// source: save_position.proto

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

type SavePositionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status     string `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	SavedPoint *Point `protobuf:"bytes,2,opt,name=saved_point,json=savedPoint,proto3" json:"saved_point,omitempty"`
}

func (x *SavePositionResponse) Reset() {
	*x = SavePositionResponse{}
	mi := &file_save_position_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SavePositionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SavePositionResponse) ProtoMessage() {}

func (x *SavePositionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_save_position_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SavePositionResponse.ProtoReflect.Descriptor instead.
func (*SavePositionResponse) Descriptor() ([]byte, []int) {
	return file_save_position_proto_rawDescGZIP(), []int{0}
}

func (x *SavePositionResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *SavePositionResponse) GetSavedPoint() *Point {
	if x != nil {
		return x.SavedPoint
	}
	return nil
}

type SavePositionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Coords *Point `protobuf:"bytes,1,opt,name=coords,proto3" json:"coords,omitempty"`
}

func (x *SavePositionRequest) Reset() {
	*x = SavePositionRequest{}
	mi := &file_save_position_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SavePositionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SavePositionRequest) ProtoMessage() {}

func (x *SavePositionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_save_position_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SavePositionRequest.ProtoReflect.Descriptor instead.
func (*SavePositionRequest) Descriptor() ([]byte, []int) {
	return file_save_position_proto_rawDescGZIP(), []int{1}
}

func (x *SavePositionRequest) GetCoords() *Point {
	if x != nil {
		return x.Coords
	}
	return nil
}

var File_save_position_proto protoreflect.FileDescriptor

var file_save_position_proto_rawDesc = []byte{
	0x0a, 0x13, 0x73, 0x61, 0x76, 0x65, 0x5f, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0c, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x57, 0x0a, 0x14, 0x53, 0x61, 0x76, 0x65, 0x50, 0x6f, 0x73, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x12, 0x27, 0x0a, 0x0b, 0x73, 0x61, 0x76, 0x65, 0x64, 0x5f, 0x70, 0x6f,
	0x69, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x06, 0x2e, 0x50, 0x6f, 0x69, 0x6e,
	0x74, 0x52, 0x0a, 0x73, 0x61, 0x76, 0x65, 0x64, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x22, 0x35, 0x0a,
	0x13, 0x53, 0x61, 0x76, 0x65, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x06, 0x63, 0x6f, 0x6f, 0x72, 0x64, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x06, 0x2e, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x06, 0x63, 0x6f,
	0x6f, 0x72, 0x64, 0x73, 0x42, 0x15, 0x5a, 0x13, 0x65, 0x72, 0x6e, 0x69, 0x65, 0x2e, 0x6f, 0x72,
	0x67, 0x2f, 0x67, 0x6f, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_save_position_proto_rawDescOnce sync.Once
	file_save_position_proto_rawDescData = file_save_position_proto_rawDesc
)

func file_save_position_proto_rawDescGZIP() []byte {
	file_save_position_proto_rawDescOnce.Do(func() {
		file_save_position_proto_rawDescData = protoimpl.X.CompressGZIP(file_save_position_proto_rawDescData)
	})
	return file_save_position_proto_rawDescData
}

var file_save_position_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_save_position_proto_goTypes = []any{
	(*SavePositionResponse)(nil), // 0: SavePositionResponse
	(*SavePositionRequest)(nil),  // 1: SavePositionRequest
	(*Point)(nil),                // 2: Point
}
var file_save_position_proto_depIdxs = []int32{
	2, // 0: SavePositionResponse.saved_point:type_name -> Point
	2, // 1: SavePositionRequest.coords:type_name -> Point
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_save_position_proto_init() }
func file_save_position_proto_init() {
	if File_save_position_proto != nil {
		return
	}
	file_points_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_save_position_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_save_position_proto_goTypes,
		DependencyIndexes: file_save_position_proto_depIdxs,
		MessageInfos:      file_save_position_proto_msgTypes,
	}.Build()
	File_save_position_proto = out.File
	file_save_position_proto_rawDesc = nil
	file_save_position_proto_goTypes = nil
	file_save_position_proto_depIdxs = nil
}
