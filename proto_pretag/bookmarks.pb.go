// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v5.29.3
// source: bookmarks.proto

package proto

import (
	_ "github.com/srikrsna/protoc-gen-gotag/tagger"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type GetBookmarksResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Bookmark      *Bookmark              `protobuf:"bytes,1,opt,name=bookmark" json:"bookmark,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetBookmarksResponse) Reset() {
	*x = GetBookmarksResponse{}
	mi := &file_bookmarks_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetBookmarksResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBookmarksResponse) ProtoMessage() {}

func (x *GetBookmarksResponse) ProtoReflect() protoreflect.Message {
	mi := &file_bookmarks_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBookmarksResponse.ProtoReflect.Descriptor instead.
func (*GetBookmarksResponse) Descriptor() ([]byte, []int) {
	return file_bookmarks_proto_rawDescGZIP(), []int{0}
}

func (x *GetBookmarksResponse) GetBookmark() *Bookmark {
	if x != nil {
		return x.Bookmark
	}
	return nil
}

type GetBookmarksRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	MinLon        *float32               `protobuf:"fixed32,1,opt,name=min_lon,json=minLon" json:"min_lon,omitempty"`
	MaxLon        *float32               `protobuf:"fixed32,2,opt,name=max_lon,json=maxLon" json:"max_lon,omitempty"`
	MinLat        *float32               `protobuf:"fixed32,3,opt,name=min_lat,json=minLat" json:"min_lat,omitempty"`
	MaxLat        *float32               `protobuf:"fixed32,4,opt,name=max_lat,json=maxLat" json:"max_lat,omitempty"`
	BoundString   *string                `protobuf:"bytes,5,opt,name=bound_string,json=boundString" json:"bound_string,omitempty"`
	Rind          *string                `protobuf:"bytes,6,opt,name=rind" json:"rind,omitempty"`
	Ts            *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=ts" json:"ts,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetBookmarksRequest) Reset() {
	*x = GetBookmarksRequest{}
	mi := &file_bookmarks_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetBookmarksRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBookmarksRequest) ProtoMessage() {}

func (x *GetBookmarksRequest) ProtoReflect() protoreflect.Message {
	mi := &file_bookmarks_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBookmarksRequest.ProtoReflect.Descriptor instead.
func (*GetBookmarksRequest) Descriptor() ([]byte, []int) {
	return file_bookmarks_proto_rawDescGZIP(), []int{1}
}

func (x *GetBookmarksRequest) GetMinLon() float32 {
	if x != nil && x.MinLon != nil {
		return *x.MinLon
	}
	return 0
}

func (x *GetBookmarksRequest) GetMaxLon() float32 {
	if x != nil && x.MaxLon != nil {
		return *x.MaxLon
	}
	return 0
}

func (x *GetBookmarksRequest) GetMinLat() float32 {
	if x != nil && x.MinLat != nil {
		return *x.MinLat
	}
	return 0
}

func (x *GetBookmarksRequest) GetMaxLat() float32 {
	if x != nil && x.MaxLat != nil {
		return *x.MaxLat
	}
	return 0
}

func (x *GetBookmarksRequest) GetBoundString() string {
	if x != nil && x.BoundString != nil {
		return *x.BoundString
	}
	return ""
}

func (x *GetBookmarksRequest) GetRind() string {
	if x != nil && x.Rind != nil {
		return *x.Rind
	}
	return ""
}

func (x *GetBookmarksRequest) GetTs() *timestamppb.Timestamp {
	if x != nil {
		return x.Ts
	}
	return nil
}

type Bookmark struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            *string                `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Loc           *Geometry              `protobuf:"bytes,2,opt,name=loc" json:"loc,omitempty"`
	Label         *string                `protobuf:"bytes,3,opt,name=label" json:"label,omitempty"`
	CreationDate  *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=creation_date,json=creationDate" json:"creation_date,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Bookmark) Reset() {
	*x = Bookmark{}
	mi := &file_bookmarks_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Bookmark) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Bookmark) ProtoMessage() {}

func (x *Bookmark) ProtoReflect() protoreflect.Message {
	mi := &file_bookmarks_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Bookmark.ProtoReflect.Descriptor instead.
func (*Bookmark) Descriptor() ([]byte, []int) {
	return file_bookmarks_proto_rawDescGZIP(), []int{2}
}

func (x *Bookmark) GetId() string {
	if x != nil && x.Id != nil {
		return *x.Id
	}
	return ""
}

func (x *Bookmark) GetLoc() *Geometry {
	if x != nil {
		return x.Loc
	}
	return nil
}

func (x *Bookmark) GetLabel() string {
	if x != nil && x.Label != nil {
		return *x.Label
	}
	return ""
}

func (x *Bookmark) GetCreationDate() *timestamppb.Timestamp {
	if x != nil {
		return x.CreationDate
	}
	return nil
}

type OldBookmark struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            *string                `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Loc           *OldGeometry           `protobuf:"bytes,2,opt,name=loc" json:"loc,omitempty"`
	Label         *string                `protobuf:"bytes,3,opt,name=label" json:"label,omitempty"`
	CreationDate  *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=creation_date,json=creationDate" json:"creation_date,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *OldBookmark) Reset() {
	*x = OldBookmark{}
	mi := &file_bookmarks_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *OldBookmark) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OldBookmark) ProtoMessage() {}

func (x *OldBookmark) ProtoReflect() protoreflect.Message {
	mi := &file_bookmarks_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OldBookmark.ProtoReflect.Descriptor instead.
func (*OldBookmark) Descriptor() ([]byte, []int) {
	return file_bookmarks_proto_rawDescGZIP(), []int{3}
}

func (x *OldBookmark) GetId() string {
	if x != nil && x.Id != nil {
		return *x.Id
	}
	return ""
}

func (x *OldBookmark) GetLoc() *OldGeometry {
	if x != nil {
		return x.Loc
	}
	return nil
}

func (x *OldBookmark) GetLabel() string {
	if x != nil && x.Label != nil {
		return *x.Label
	}
	return ""
}

func (x *OldBookmark) GetCreationDate() *timestamppb.Timestamp {
	if x != nil {
		return x.CreationDate
	}
	return nil
}

type SaveBookmarkRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Loc           *Geometry              `protobuf:"bytes,1,opt,name=loc" json:"loc,omitempty"`
	Label         *string                `protobuf:"bytes,2,opt,name=label" json:"label,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SaveBookmarkRequest) Reset() {
	*x = SaveBookmarkRequest{}
	mi := &file_bookmarks_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SaveBookmarkRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SaveBookmarkRequest) ProtoMessage() {}

func (x *SaveBookmarkRequest) ProtoReflect() protoreflect.Message {
	mi := &file_bookmarks_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SaveBookmarkRequest.ProtoReflect.Descriptor instead.
func (*SaveBookmarkRequest) Descriptor() ([]byte, []int) {
	return file_bookmarks_proto_rawDescGZIP(), []int{4}
}

func (x *SaveBookmarkRequest) GetLoc() *Geometry {
	if x != nil {
		return x.Loc
	}
	return nil
}

func (x *SaveBookmarkRequest) GetLabel() string {
	if x != nil && x.Label != nil {
		return *x.Label
	}
	return ""
}

type SaveBookmarkResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Bookmark      *Bookmark              `protobuf:"bytes,1,opt,name=bookmark" json:"bookmark,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SaveBookmarkResponse) Reset() {
	*x = SaveBookmarkResponse{}
	mi := &file_bookmarks_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SaveBookmarkResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SaveBookmarkResponse) ProtoMessage() {}

func (x *SaveBookmarkResponse) ProtoReflect() protoreflect.Message {
	mi := &file_bookmarks_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SaveBookmarkResponse.ProtoReflect.Descriptor instead.
func (*SaveBookmarkResponse) Descriptor() ([]byte, []int) {
	return file_bookmarks_proto_rawDescGZIP(), []int{5}
}

func (x *SaveBookmarkResponse) GetBookmark() *Bookmark {
	if x != nil {
		return x.Bookmark
	}
	return nil
}

var File_bookmarks_proto protoreflect.FileDescriptor

var file_bookmarks_proto_rawDesc = string([]byte{
	0x0a, 0x0f, 0x62, 0x6f, 0x6f, 0x6b, 0x6d, 0x61, 0x72, 0x6b, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x0c, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x13, 0x74, 0x61, 0x67, 0x67, 0x65, 0x72, 0x2f, 0x74, 0x61, 0x67, 0x67, 0x65, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x3d, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x42, 0x6f, 0x6f, 0x6b,
	0x6d, 0x61, 0x72, 0x6b, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x25, 0x0a,
	0x08, 0x62, 0x6f, 0x6f, 0x6b, 0x6d, 0x61, 0x72, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x09, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x6d, 0x61, 0x72, 0x6b, 0x52, 0x08, 0x62, 0x6f, 0x6f, 0x6b,
	0x6d, 0x61, 0x72, 0x6b, 0x22, 0xdc, 0x01, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x42, 0x6f, 0x6f, 0x6b,
	0x6d, 0x61, 0x72, 0x6b, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07,
	0x6d, 0x69, 0x6e, 0x5f, 0x6c, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x02, 0x52, 0x06, 0x6d,
	0x69, 0x6e, 0x4c, 0x6f, 0x6e, 0x12, 0x17, 0x0a, 0x07, 0x6d, 0x61, 0x78, 0x5f, 0x6c, 0x6f, 0x6e,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x06, 0x6d, 0x61, 0x78, 0x4c, 0x6f, 0x6e, 0x12, 0x17,
	0x0a, 0x07, 0x6d, 0x69, 0x6e, 0x5f, 0x6c, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52,
	0x06, 0x6d, 0x69, 0x6e, 0x4c, 0x61, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x6d, 0x61, 0x78, 0x5f, 0x6c,
	0x61, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x02, 0x52, 0x06, 0x6d, 0x61, 0x78, 0x4c, 0x61, 0x74,
	0x12, 0x21, 0x0a, 0x0c, 0x62, 0x6f, 0x75, 0x6e, 0x64, 0x5f, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x62, 0x6f, 0x75, 0x6e, 0x64, 0x53, 0x74, 0x72,
	0x69, 0x6e, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x69, 0x6e, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x72, 0x69, 0x6e, 0x64, 0x12, 0x2a, 0x0a, 0x02, 0x74, 0x73, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x02, 0x74, 0x73, 0x22, 0xd9, 0x01, 0x0a, 0x08, 0x42, 0x6f, 0x6f, 0x6b, 0x6d, 0x61, 0x72, 0x6b,
	0x12, 0x3e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x2e, 0x9a, 0x84,
	0x9e, 0x03, 0x29, 0x62, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x5f, 0x69, 0x64, 0x2c, 0x6f, 0x6d, 0x69,
	0x74, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x20, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x5f, 0x69,
	0x64, 0x2c, 0x6f, 0x6d, 0x69, 0x74, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x1b, 0x0a, 0x03, 0x6c, 0x6f, 0x63, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e,
	0x47, 0x65, 0x6f, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x52, 0x03, 0x6c, 0x6f, 0x63, 0x12, 0x14, 0x0a,
	0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x61,
	0x62, 0x65, 0x6c, 0x12, 0x5a, 0x0a, 0x0d, 0x63, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x64, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x19, 0x9a, 0x84, 0x9e, 0x03, 0x14, 0x62, 0x73, 0x6f,
	0x6e, 0x3a, 0x22, 0x63, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x64, 0x61, 0x74, 0x65,
	0x22, 0x52, 0x0c, 0x63, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x65, 0x22,
	0xdf, 0x01, 0x0a, 0x0b, 0x4f, 0x6c, 0x64, 0x42, 0x6f, 0x6f, 0x6b, 0x6d, 0x61, 0x72, 0x6b, 0x12,
	0x3e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x2e, 0x9a, 0x84, 0x9e,
	0x03, 0x29, 0x62, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x5f, 0x69, 0x64, 0x2c, 0x6f, 0x6d, 0x69, 0x74,
	0x65, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x20, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x5f, 0x69, 0x64,
	0x2c, 0x6f, 0x6d, 0x69, 0x74, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x1e, 0x0a, 0x03, 0x6c, 0x6f, 0x63, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x4f,
	0x6c, 0x64, 0x47, 0x65, 0x6f, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x52, 0x03, 0x6c, 0x6f, 0x63, 0x12,
	0x14, 0x0a, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x6c, 0x61, 0x62, 0x65, 0x6c, 0x12, 0x5a, 0x0a, 0x0d, 0x63, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x19, 0x9a, 0x84, 0x9e, 0x03, 0x14, 0x62,
	0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x63, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x64, 0x61,
	0x74, 0x65, 0x22, 0x52, 0x0c, 0x63, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74,
	0x65, 0x22, 0x48, 0x0a, 0x13, 0x53, 0x61, 0x76, 0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x6d, 0x61, 0x72,
	0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x03, 0x6c, 0x6f, 0x63, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x47, 0x65, 0x6f, 0x6d, 0x65, 0x74, 0x72, 0x79,
	0x52, 0x03, 0x6c, 0x6f, 0x63, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x22, 0x3d, 0x0a, 0x14, 0x53,
	0x61, 0x76, 0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x6d, 0x61, 0x72, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x25, 0x0a, 0x08, 0x62, 0x6f, 0x6f, 0x6b, 0x6d, 0x61, 0x72, 0x6b, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x6d, 0x61, 0x72, 0x6b,
	0x52, 0x08, 0x62, 0x6f, 0x6f, 0x6b, 0x6d, 0x61, 0x72, 0x6b, 0x42, 0x15, 0x5a, 0x13, 0x65, 0x72,
	0x6e, 0x69, 0x65, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x6f, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x08, 0x65, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x70, 0xe8, 0x07,
})

var (
	file_bookmarks_proto_rawDescOnce sync.Once
	file_bookmarks_proto_rawDescData []byte
)

func file_bookmarks_proto_rawDescGZIP() []byte {
	file_bookmarks_proto_rawDescOnce.Do(func() {
		file_bookmarks_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_bookmarks_proto_rawDesc), len(file_bookmarks_proto_rawDesc)))
	})
	return file_bookmarks_proto_rawDescData
}

var file_bookmarks_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_bookmarks_proto_goTypes = []any{
	(*GetBookmarksResponse)(nil),  // 0: GetBookmarksResponse
	(*GetBookmarksRequest)(nil),   // 1: GetBookmarksRequest
	(*Bookmark)(nil),              // 2: Bookmark
	(*OldBookmark)(nil),           // 3: OldBookmark
	(*SaveBookmarkRequest)(nil),   // 4: SaveBookmarkRequest
	(*SaveBookmarkResponse)(nil),  // 5: SaveBookmarkResponse
	(*timestamppb.Timestamp)(nil), // 6: google.protobuf.Timestamp
	(*Geometry)(nil),              // 7: Geometry
	(*OldGeometry)(nil),           // 8: OldGeometry
}
var file_bookmarks_proto_depIdxs = []int32{
	2, // 0: GetBookmarksResponse.bookmark:type_name -> Bookmark
	6, // 1: GetBookmarksRequest.ts:type_name -> google.protobuf.Timestamp
	7, // 2: Bookmark.loc:type_name -> Geometry
	6, // 3: Bookmark.creation_date:type_name -> google.protobuf.Timestamp
	8, // 4: OldBookmark.loc:type_name -> OldGeometry
	6, // 5: OldBookmark.creation_date:type_name -> google.protobuf.Timestamp
	7, // 6: SaveBookmarkRequest.loc:type_name -> Geometry
	2, // 7: SaveBookmarkResponse.bookmark:type_name -> Bookmark
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_bookmarks_proto_init() }
func file_bookmarks_proto_init() {
	if File_bookmarks_proto != nil {
		return
	}
	file_points_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_bookmarks_proto_rawDesc), len(file_bookmarks_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_bookmarks_proto_goTypes,
		DependencyIndexes: file_bookmarks_proto_depIdxs,
		MessageInfos:      file_bookmarks_proto_msgTypes,
	}.Build()
	File_bookmarks_proto = out.File
	file_bookmarks_proto_goTypes = nil
	file_bookmarks_proto_depIdxs = nil
}
