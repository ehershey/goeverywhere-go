// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.4
// 	protoc        v5.29.3
// source: points.proto

package proto

import (
	_ "github.com/srikrsna/protoc-gen-gotag/tagger"
	latlng "google.golang.org/genproto/googleapis/type/latlng"
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

type GetPointsResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	MinTimestamp  *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=min_timestamp,json=minTimestamp" json:"min_timestamp,omitempty"`
	MaxTimestamp  *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=max_timestamp,json=maxTimestamp" json:"max_timestamp,omitempty"`
	PointCount    *uint32                `protobuf:"varint,3,opt,name=point_count,json=pointCount" json:"point_count,omitempty"`
	EntrySources  []string               `protobuf:"bytes,4,rep,name=entry_sources,json=entrySources" json:"entry_sources,omitempty"`
	MinLon        *float32               `protobuf:"fixed32,5,opt,name=min_lon,json=minLon" json:"min_lon,omitempty"`
	MinLat        *float32               `protobuf:"fixed32,6,opt,name=min_lat,json=minLat" json:"min_lat,omitempty"`
	MaxLon        *float32               `protobuf:"fixed32,7,opt,name=max_lon,json=maxLon" json:"max_lon,omitempty"`
	MaxLat        *float32               `protobuf:"fixed32,8,opt,name=max_lat,json=maxLat" json:"max_lat,omitempty"`
	SkippedCount  *uint32                `protobuf:"varint,9,opt,name=skipped_count,json=skippedCount" json:"skipped_count,omitempty"`
	Rid           *string                `protobuf:"bytes,10,opt,name=rid" json:"rid,omitempty"`
	BoundString   *string                `protobuf:"bytes,11,opt,name=bound_string,json=boundString" json:"bound_string,omitempty"`
	EntrySource   map[string]uint32      `protobuf:"bytes,12,rep,name=entry_source,json=entrySource" json:"entry_source,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"varint,2,opt,name=value"`
	Setsize       *uint32                `protobuf:"varint,13,opt,name=setsize" json:"setsize,omitempty"`
	Limit         *uint32                `protobuf:"varint,14,opt,name=limit" json:"limit,omitempty"`
	Point         *Point                 `protobuf:"bytes,15,opt,name=point" json:"point,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetPointsResponse) Reset() {
	*x = GetPointsResponse{}
	mi := &file_points_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetPointsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPointsResponse) ProtoMessage() {}

func (x *GetPointsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_points_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPointsResponse.ProtoReflect.Descriptor instead.
func (*GetPointsResponse) Descriptor() ([]byte, []int) {
	return file_points_proto_rawDescGZIP(), []int{0}
}

func (x *GetPointsResponse) GetMinTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.MinTimestamp
	}
	return nil
}

func (x *GetPointsResponse) GetMaxTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.MaxTimestamp
	}
	return nil
}

func (x *GetPointsResponse) GetPointCount() uint32 {
	if x != nil && x.PointCount != nil {
		return *x.PointCount
	}
	return 0
}

func (x *GetPointsResponse) GetEntrySources() []string {
	if x != nil {
		return x.EntrySources
	}
	return nil
}

func (x *GetPointsResponse) GetMinLon() float32 {
	if x != nil && x.MinLon != nil {
		return *x.MinLon
	}
	return 0
}

func (x *GetPointsResponse) GetMinLat() float32 {
	if x != nil && x.MinLat != nil {
		return *x.MinLat
	}
	return 0
}

func (x *GetPointsResponse) GetMaxLon() float32 {
	if x != nil && x.MaxLon != nil {
		return *x.MaxLon
	}
	return 0
}

func (x *GetPointsResponse) GetMaxLat() float32 {
	if x != nil && x.MaxLat != nil {
		return *x.MaxLat
	}
	return 0
}

func (x *GetPointsResponse) GetSkippedCount() uint32 {
	if x != nil && x.SkippedCount != nil {
		return *x.SkippedCount
	}
	return 0
}

func (x *GetPointsResponse) GetRid() string {
	if x != nil && x.Rid != nil {
		return *x.Rid
	}
	return ""
}

func (x *GetPointsResponse) GetBoundString() string {
	if x != nil && x.BoundString != nil {
		return *x.BoundString
	}
	return ""
}

func (x *GetPointsResponse) GetEntrySource() map[string]uint32 {
	if x != nil {
		return x.EntrySource
	}
	return nil
}

func (x *GetPointsResponse) GetSetsize() uint32 {
	if x != nil && x.Setsize != nil {
		return *x.Setsize
	}
	return 0
}

func (x *GetPointsResponse) GetLimit() uint32 {
	if x != nil && x.Limit != nil {
		return *x.Limit
	}
	return 0
}

func (x *GetPointsResponse) GetPoint() *Point {
	if x != nil {
		return x.Point
	}
	return nil
}

type GetPointsRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	MinLon        *float32               `protobuf:"fixed32,1,opt,name=min_lon,json=minLon" json:"min_lon,omitempty"`
	MaxLon        *float32               `protobuf:"fixed32,2,opt,name=max_lon,json=maxLon" json:"max_lon,omitempty"`
	MinLat        *float32               `protobuf:"fixed32,3,opt,name=min_lat,json=minLat" json:"min_lat,omitempty"`
	MaxLat        *float32               `protobuf:"fixed32,4,opt,name=max_lat,json=maxLat" json:"max_lat,omitempty"`
	BoundString   *string                `protobuf:"bytes,5,opt,name=bound_string,json=boundString" json:"bound_string,omitempty"`
	Rind          *string                `protobuf:"bytes,6,opt,name=rind" json:"rind,omitempty"`
	From          *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=from" json:"from,omitempty"`
	To            *timestamppb.Timestamp `protobuf:"bytes,8,opt,name=to" json:"to,omitempty"`
	Limit         *uint32                `protobuf:"varint,9,opt,name=limit" json:"limit,omitempty"`
	NoSkip        *bool                  `protobuf:"varint,10,opt,name=no_skip,json=noSkip" json:"no_skip,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetPointsRequest) Reset() {
	*x = GetPointsRequest{}
	mi := &file_points_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetPointsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPointsRequest) ProtoMessage() {}

func (x *GetPointsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_points_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPointsRequest.ProtoReflect.Descriptor instead.
func (*GetPointsRequest) Descriptor() ([]byte, []int) {
	return file_points_proto_rawDescGZIP(), []int{1}
}

func (x *GetPointsRequest) GetMinLon() float32 {
	if x != nil && x.MinLon != nil {
		return *x.MinLon
	}
	return 0
}

func (x *GetPointsRequest) GetMaxLon() float32 {
	if x != nil && x.MaxLon != nil {
		return *x.MaxLon
	}
	return 0
}

func (x *GetPointsRequest) GetMinLat() float32 {
	if x != nil && x.MinLat != nil {
		return *x.MinLat
	}
	return 0
}

func (x *GetPointsRequest) GetMaxLat() float32 {
	if x != nil && x.MaxLat != nil {
		return *x.MaxLat
	}
	return 0
}

func (x *GetPointsRequest) GetBoundString() string {
	if x != nil && x.BoundString != nil {
		return *x.BoundString
	}
	return ""
}

func (x *GetPointsRequest) GetRind() string {
	if x != nil && x.Rind != nil {
		return *x.Rind
	}
	return ""
}

func (x *GetPointsRequest) GetFrom() *timestamppb.Timestamp {
	if x != nil {
		return x.From
	}
	return nil
}

func (x *GetPointsRequest) GetTo() *timestamppb.Timestamp {
	if x != nil {
		return x.To
	}
	return nil
}

func (x *GetPointsRequest) GetLimit() uint32 {
	if x != nil && x.Limit != nil {
		return *x.Limit
	}
	return 0
}

func (x *GetPointsRequest) GetNoSkip() bool {
	if x != nil && x.NoSkip != nil {
		return *x.NoSkip
	}
	return false
}

type Point struct {
	state     protoimpl.MessageState `protogen:"open.v1"`
	Id        *string                `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	EntryDate *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=entry_date,json=entryDate" json:"entry_date,omitempty"`
	Speed     *float32               `protobuf:"fixed32,3,opt,name=speed" json:"speed,omitempty"`
	// float speed = 3;
	EntrySource      *string   `protobuf:"bytes,4,opt,name=entry_source,json=entrySource" json:"entry_source,omitempty"`
	Altitude         *float32  `protobuf:"fixed32,5,opt,name=altitude" json:"altitude,omitempty"`
	Loc              *Geometry `protobuf:"bytes,8,opt,name=loc" json:"loc,omitempty"`
	ActivityType     *string   `protobuf:"bytes,9,opt,name=activity_type,json=activityType" json:"activity_type,omitempty"`
	Heading          *float32  `protobuf:"fixed32,11,opt,name=heading" json:"heading,omitempty"`
	Accuracy         *float32  `protobuf:"fixed32,12,opt,name=accuracy" json:"accuracy,omitempty"`
	AltitudeAccuracy *float32  `protobuf:"fixed32,13,opt,name=altitude_accuracy,json=altitudeAccuracy" json:"altitude_accuracy,omitempty"`
	unknownFields    protoimpl.UnknownFields
	sizeCache        protoimpl.SizeCache
}

func (x *Point) Reset() {
	*x = Point{}
	mi := &file_points_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Point) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Point) ProtoMessage() {}

func (x *Point) ProtoReflect() protoreflect.Message {
	mi := &file_points_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Point.ProtoReflect.Descriptor instead.
func (*Point) Descriptor() ([]byte, []int) {
	return file_points_proto_rawDescGZIP(), []int{2}
}

func (x *Point) GetId() string {
	if x != nil && x.Id != nil {
		return *x.Id
	}
	return ""
}

func (x *Point) GetEntryDate() *timestamppb.Timestamp {
	if x != nil {
		return x.EntryDate
	}
	return nil
}

func (x *Point) GetSpeed() float32 {
	if x != nil && x.Speed != nil {
		return *x.Speed
	}
	return 0
}

func (x *Point) GetEntrySource() string {
	if x != nil && x.EntrySource != nil {
		return *x.EntrySource
	}
	return ""
}

func (x *Point) GetAltitude() float32 {
	if x != nil && x.Altitude != nil {
		return *x.Altitude
	}
	return 0
}

func (x *Point) GetLoc() *Geometry {
	if x != nil {
		return x.Loc
	}
	return nil
}

func (x *Point) GetActivityType() string {
	if x != nil && x.ActivityType != nil {
		return *x.ActivityType
	}
	return ""
}

func (x *Point) GetHeading() float32 {
	if x != nil && x.Heading != nil {
		return *x.Heading
	}
	return 0
}

func (x *Point) GetAccuracy() float32 {
	if x != nil && x.Accuracy != nil {
		return *x.Accuracy
	}
	return 0
}

func (x *Point) GetAltitudeAccuracy() float32 {
	if x != nil && x.AltitudeAccuracy != nil {
		return *x.AltitudeAccuracy
	}
	return 0
}

type Geometry struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Coordinates   *latlng.LatLng         `protobuf:"bytes,1,opt,name=coordinates" json:"coordinates,omitempty"`
	Type          *string                `protobuf:"bytes,2,opt,name=type" json:"type,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Geometry) Reset() {
	*x = Geometry{}
	mi := &file_points_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Geometry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Geometry) ProtoMessage() {}

func (x *Geometry) ProtoReflect() protoreflect.Message {
	mi := &file_points_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Geometry.ProtoReflect.Descriptor instead.
func (*Geometry) Descriptor() ([]byte, []int) {
	return file_points_proto_rawDescGZIP(), []int{3}
}

func (x *Geometry) GetCoordinates() *latlng.LatLng {
	if x != nil {
		return x.Coordinates
	}
	return nil
}

func (x *Geometry) GetType() string {
	if x != nil && x.Type != nil {
		return *x.Type
	}
	return ""
}

type OldGeometry struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Coordinates   []float32              `protobuf:"fixed32,1,rep,packed,name=coordinates" json:"coordinates,omitempty"`
	Type          *string                `protobuf:"bytes,2,opt,name=type" json:"type,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *OldGeometry) Reset() {
	*x = OldGeometry{}
	mi := &file_points_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *OldGeometry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OldGeometry) ProtoMessage() {}

func (x *OldGeometry) ProtoReflect() protoreflect.Message {
	mi := &file_points_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OldGeometry.ProtoReflect.Descriptor instead.
func (*OldGeometry) Descriptor() ([]byte, []int) {
	return file_points_proto_rawDescGZIP(), []int{4}
}

func (x *OldGeometry) GetCoordinates() []float32 {
	if x != nil {
		return x.Coordinates
	}
	return nil
}

func (x *OldGeometry) GetType() string {
	if x != nil && x.Type != nil {
		return *x.Type
	}
	return ""
}

var File_points_proto protoreflect.FileDescriptor

var file_points_proto_rawDesc = string([]byte{
	0x0a, 0x0c, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x18, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x6c, 0x61, 0x74,
	0x6c, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x13, 0x74, 0x61, 0x67, 0x67, 0x65,
	0x72, 0x2f, 0x74, 0x61, 0x67, 0x67, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xef,
	0x04, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3f, 0x0a, 0x0d, 0x6d, 0x69, 0x6e, 0x5f, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0c, 0x6d, 0x69, 0x6e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x3f, 0x0a, 0x0d, 0x6d, 0x61, 0x78, 0x5f, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0c, 0x6d, 0x61, 0x78, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x5f,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x70, 0x6f, 0x69,
	0x6e, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x65, 0x6e, 0x74, 0x72, 0x79,
	0x5f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0c,
	0x65, 0x6e, 0x74, 0x72, 0x79, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x12, 0x17, 0x0a, 0x07,
	0x6d, 0x69, 0x6e, 0x5f, 0x6c, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x02, 0x52, 0x06, 0x6d,
	0x69, 0x6e, 0x4c, 0x6f, 0x6e, 0x12, 0x17, 0x0a, 0x07, 0x6d, 0x69, 0x6e, 0x5f, 0x6c, 0x61, 0x74,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x02, 0x52, 0x06, 0x6d, 0x69, 0x6e, 0x4c, 0x61, 0x74, 0x12, 0x17,
	0x0a, 0x07, 0x6d, 0x61, 0x78, 0x5f, 0x6c, 0x6f, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x02, 0x52,
	0x06, 0x6d, 0x61, 0x78, 0x4c, 0x6f, 0x6e, 0x12, 0x17, 0x0a, 0x07, 0x6d, 0x61, 0x78, 0x5f, 0x6c,
	0x61, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x02, 0x52, 0x06, 0x6d, 0x61, 0x78, 0x4c, 0x61, 0x74,
	0x12, 0x23, 0x0a, 0x0d, 0x73, 0x6b, 0x69, 0x70, 0x70, 0x65, 0x64, 0x5f, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0c, 0x73, 0x6b, 0x69, 0x70, 0x70, 0x65, 0x64,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x72, 0x69, 0x64, 0x18, 0x0a, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x72, 0x69, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x62, 0x6f, 0x75, 0x6e, 0x64,
	0x5f, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x62,
	0x6f, 0x75, 0x6e, 0x64, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x12, 0x46, 0x0a, 0x0c, 0x65, 0x6e,
	0x74, 0x72, 0x79, 0x5f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x0c, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x23, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0b, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x53, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x65, 0x74, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x0d, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x07, 0x73, 0x65, 0x74, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x6c, 0x69, 0x6d,
	0x69, 0x74, 0x12, 0x1c, 0x0a, 0x05, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0x0f, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x06, 0x2e, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x05, 0x70, 0x6f, 0x69, 0x6e, 0x74,
	0x1a, 0x3e, 0x0a, 0x10, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01,
	0x22, 0xb8, 0x02, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x6d, 0x69, 0x6e, 0x5f, 0x6c, 0x6f, 0x6e,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x02, 0x52, 0x06, 0x6d, 0x69, 0x6e, 0x4c, 0x6f, 0x6e, 0x12, 0x17,
	0x0a, 0x07, 0x6d, 0x61, 0x78, 0x5f, 0x6c, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52,
	0x06, 0x6d, 0x61, 0x78, 0x4c, 0x6f, 0x6e, 0x12, 0x17, 0x0a, 0x07, 0x6d, 0x69, 0x6e, 0x5f, 0x6c,
	0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x06, 0x6d, 0x69, 0x6e, 0x4c, 0x61, 0x74,
	0x12, 0x17, 0x0a, 0x07, 0x6d, 0x61, 0x78, 0x5f, 0x6c, 0x61, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x02, 0x52, 0x06, 0x6d, 0x61, 0x78, 0x4c, 0x61, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x62, 0x6f, 0x75,
	0x6e, 0x64, 0x5f, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x62, 0x6f, 0x75, 0x6e, 0x64, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x12, 0x12, 0x0a, 0x04,
	0x72, 0x69, 0x6e, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x72, 0x69, 0x6e, 0x64,
	0x12, 0x2e, 0x0a, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x04, 0x66, 0x72, 0x6f, 0x6d,
	0x12, 0x2a, 0x0a, 0x02, 0x74, 0x6f, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x02, 0x74, 0x6f, 0x12, 0x14, 0x0a, 0x05,
	0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x6c, 0x69, 0x6d,
	0x69, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x6e, 0x6f, 0x5f, 0x73, 0x6b, 0x69, 0x70, 0x18, 0x0a, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x06, 0x6e, 0x6f, 0x53, 0x6b, 0x69, 0x70, 0x22, 0xf8, 0x04, 0x0a, 0x05,
	0x50, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x3e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x2e, 0x9a, 0x84, 0x9e, 0x03, 0x29, 0x62, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x5f, 0x69,
	0x64, 0x2c, 0x6f, 0x6d, 0x69, 0x74, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x20, 0x6a, 0x73, 0x6f,
	0x6e, 0x3a, 0x22, 0x5f, 0x69, 0x64, 0x2c, 0x6f, 0x6d, 0x69, 0x74, 0x65, 0x6d, 0x70, 0x74, 0x79,
	0x22, 0x52, 0x02, 0x69, 0x64, 0x12, 0x39, 0x0a, 0x0a, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x5f, 0x64,
	0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x44, 0x61, 0x74, 0x65,
	0x12, 0x3e, 0x0a, 0x05, 0x73, 0x70, 0x65, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x42,
	0x28, 0x9a, 0x84, 0x9e, 0x03, 0x23, 0x62, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x2c, 0x6f, 0x6d, 0x69,
	0x74, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x20, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x2c, 0x6f,
	0x6d, 0x69, 0x74, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x52, 0x05, 0x73, 0x70, 0x65, 0x65, 0x64,
	0x12, 0x21, 0x0a, 0x0c, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x5f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x53, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x12, 0x44, 0x0a, 0x08, 0x61, 0x6c, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x02, 0x42, 0x28, 0x9a, 0x84, 0x9e, 0x03, 0x23, 0x62, 0x73, 0x6f, 0x6e,
	0x3a, 0x22, 0x2c, 0x6f, 0x6d, 0x69, 0x74, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x20, 0x6a, 0x73,
	0x6f, 0x6e, 0x3a, 0x22, 0x2c, 0x6f, 0x6d, 0x69, 0x74, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x52,
	0x08, 0x61, 0x6c, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x12, 0x1b, 0x0a, 0x03, 0x6c, 0x6f, 0x63,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x47, 0x65, 0x6f, 0x6d, 0x65, 0x74, 0x72,
	0x79, 0x52, 0x03, 0x6c, 0x6f, 0x63, 0x12, 0x4d, 0x0a, 0x0d, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69,
	0x74, 0x79, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x42, 0x28, 0x9a,
	0x84, 0x9e, 0x03, 0x23, 0x62, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x2c, 0x6f, 0x6d, 0x69, 0x74, 0x65,
	0x6d, 0x70, 0x74, 0x79, 0x22, 0x20, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x2c, 0x6f, 0x6d, 0x69,
	0x74, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x52, 0x0c, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74,
	0x79, 0x54, 0x79, 0x70, 0x65, 0x12, 0x42, 0x0a, 0x07, 0x68, 0x65, 0x61, 0x64, 0x69, 0x6e, 0x67,
	0x18, 0x0b, 0x20, 0x01, 0x28, 0x02, 0x42, 0x28, 0x9a, 0x84, 0x9e, 0x03, 0x23, 0x62, 0x73, 0x6f,
	0x6e, 0x3a, 0x22, 0x2c, 0x6f, 0x6d, 0x69, 0x74, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x20, 0x6a,
	0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x2c, 0x6f, 0x6d, 0x69, 0x74, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x22,
	0x52, 0x07, 0x68, 0x65, 0x61, 0x64, 0x69, 0x6e, 0x67, 0x12, 0x44, 0x0a, 0x08, 0x61, 0x63, 0x63,
	0x75, 0x72, 0x61, 0x63, 0x79, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x02, 0x42, 0x28, 0x9a, 0x84, 0x9e,
	0x03, 0x23, 0x62, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x2c, 0x6f, 0x6d, 0x69, 0x74, 0x65, 0x6d, 0x70,
	0x74, 0x79, 0x22, 0x20, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x2c, 0x6f, 0x6d, 0x69, 0x74, 0x65,
	0x6d, 0x70, 0x74, 0x79, 0x22, 0x52, 0x08, 0x61, 0x63, 0x63, 0x75, 0x72, 0x61, 0x63, 0x79, 0x12,
	0x55, 0x0a, 0x11, 0x61, 0x6c, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x5f, 0x61, 0x63, 0x63, 0x75,
	0x72, 0x61, 0x63, 0x79, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x02, 0x42, 0x28, 0x9a, 0x84, 0x9e, 0x03,
	0x23, 0x62, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x2c, 0x6f, 0x6d, 0x69, 0x74, 0x65, 0x6d, 0x70, 0x74,
	0x79, 0x22, 0x20, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x2c, 0x6f, 0x6d, 0x69, 0x74, 0x65, 0x6d,
	0x70, 0x74, 0x79, 0x22, 0x52, 0x10, 0x61, 0x6c, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x41, 0x63,
	0x63, 0x75, 0x72, 0x61, 0x63, 0x79, 0x22, 0x55, 0x0a, 0x08, 0x47, 0x65, 0x6f, 0x6d, 0x65, 0x74,
	0x72, 0x79, 0x12, 0x35, 0x0a, 0x0b, 0x63, 0x6f, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x61, 0x74, 0x65,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x4c, 0x61, 0x74, 0x4c, 0x6e, 0x67, 0x52, 0x0b, 0x63, 0x6f,
	0x6f, 0x72, 0x64, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x22, 0x6b, 0x0a,
	0x0b, 0x4f, 0x6c, 0x64, 0x47, 0x65, 0x6f, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x12, 0x48, 0x0a, 0x0b,
	0x63, 0x6f, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x02, 0x42, 0x26, 0x9a, 0x84, 0x9e, 0x03, 0x21, 0x62, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x2c, 0x74,
	0x72, 0x75, 0x6e, 0x63, 0x61, 0x74, 0x65, 0x22, 0x20, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x2c,
	0x74, 0x72, 0x75, 0x6e, 0x63, 0x61, 0x74, 0x65, 0x22, 0x52, 0x0b, 0x63, 0x6f, 0x6f, 0x72, 0x64,
	0x69, 0x6e, 0x61, 0x74, 0x65, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x42, 0x15, 0x5a, 0x13, 0x65, 0x72,
	0x6e, 0x69, 0x65, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x6f, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x08, 0x65, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x70, 0xe8, 0x07,
})

var (
	file_points_proto_rawDescOnce sync.Once
	file_points_proto_rawDescData []byte
)

func file_points_proto_rawDescGZIP() []byte {
	file_points_proto_rawDescOnce.Do(func() {
		file_points_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_points_proto_rawDesc), len(file_points_proto_rawDesc)))
	})
	return file_points_proto_rawDescData
}

var file_points_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_points_proto_goTypes = []any{
	(*GetPointsResponse)(nil),     // 0: GetPointsResponse
	(*GetPointsRequest)(nil),      // 1: GetPointsRequest
	(*Point)(nil),                 // 2: Point
	(*Geometry)(nil),              // 3: Geometry
	(*OldGeometry)(nil),           // 4: OldGeometry
	nil,                           // 5: GetPointsResponse.EntrySourceEntry
	(*timestamppb.Timestamp)(nil), // 6: google.protobuf.Timestamp
	(*latlng.LatLng)(nil),         // 7: google.type.LatLng
}
var file_points_proto_depIdxs = []int32{
	6, // 0: GetPointsResponse.min_timestamp:type_name -> google.protobuf.Timestamp
	6, // 1: GetPointsResponse.max_timestamp:type_name -> google.protobuf.Timestamp
	5, // 2: GetPointsResponse.entry_source:type_name -> GetPointsResponse.EntrySourceEntry
	2, // 3: GetPointsResponse.point:type_name -> Point
	6, // 4: GetPointsRequest.from:type_name -> google.protobuf.Timestamp
	6, // 5: GetPointsRequest.to:type_name -> google.protobuf.Timestamp
	6, // 6: Point.entry_date:type_name -> google.protobuf.Timestamp
	3, // 7: Point.loc:type_name -> Geometry
	7, // 8: Geometry.coordinates:type_name -> google.type.LatLng
	9, // [9:9] is the sub-list for method output_type
	9, // [9:9] is the sub-list for method input_type
	9, // [9:9] is the sub-list for extension type_name
	9, // [9:9] is the sub-list for extension extendee
	0, // [0:9] is the sub-list for field type_name
}

func init() { file_points_proto_init() }
func file_points_proto_init() {
	if File_points_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_points_proto_rawDesc), len(file_points_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_points_proto_goTypes,
		DependencyIndexes: file_points_proto_depIdxs,
		MessageInfos:      file_points_proto_msgTypes,
	}.Build()
	File_points_proto = out.File
	file_points_proto_goTypes = nil
	file_points_proto_depIdxs = nil
}
