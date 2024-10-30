// goe_service.proto

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.28.3
// source: goe_service.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_goe_service_proto protoreflect.FileDescriptor

var file_goe_service_proto_rawDesc = []byte{
	0x0a, 0x11, 0x67, 0x6f, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x0b, 0x73, 0x74, 0x61, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x0f, 0x70, 0x6f, 0x6c, 0x79, 0x6c, 0x69, 0x6e, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x0c, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32,
	0xb2, 0x01, 0x0a, 0x0a, 0x47, 0x4f, 0x45, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x2f,
	0x0a, 0x08, 0x47, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x73, 0x12, 0x10, 0x2e, 0x47, 0x65, 0x74,
	0x53, 0x74, 0x61, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x47,
	0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x3d, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x6c, 0x79, 0x6c, 0x69, 0x6e, 0x65, 0x73, 0x12,
	0x14, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x6c, 0x79, 0x6c, 0x69, 0x6e, 0x65, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x6c, 0x79, 0x6c,
	0x69, 0x6e, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x30, 0x01, 0x12, 0x34,
	0x0a, 0x09, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x12, 0x11, 0x2e, 0x47, 0x65,
	0x74, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12,
	0x2e, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x30, 0x01, 0x42, 0x15, 0x5a, 0x13, 0x65, 0x72, 0x6e, 0x69, 0x65, 0x2e, 0x6f, 0x72,
	0x67, 0x2f, 0x67, 0x6f, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var file_goe_service_proto_goTypes = []any{
	(*GetStatsRequest)(nil),      // 0: GetStatsRequest
	(*GetPolylinesRequest)(nil),  // 1: GetPolylinesRequest
	(*GetPointsRequest)(nil),     // 2: GetPointsRequest
	(*GetStatsResponse)(nil),     // 3: GetStatsResponse
	(*GetPolylinesResponse)(nil), // 4: GetPolylinesResponse
	(*GetPointsResponse)(nil),    // 5: GetPointsResponse
}
var file_goe_service_proto_depIdxs = []int32{
	0, // 0: GOEService.GetStats:input_type -> GetStatsRequest
	1, // 1: GOEService.GetPolylines:input_type -> GetPolylinesRequest
	2, // 2: GOEService.GetPoints:input_type -> GetPointsRequest
	3, // 3: GOEService.GetStats:output_type -> GetStatsResponse
	4, // 4: GOEService.GetPolylines:output_type -> GetPolylinesResponse
	5, // 5: GOEService.GetPoints:output_type -> GetPointsResponse
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_goe_service_proto_init() }
func file_goe_service_proto_init() {
	if File_goe_service_proto != nil {
		return
	}
	file_stats_proto_init()
	file_polylines_proto_init()
	file_points_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_goe_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_goe_service_proto_goTypes,
		DependencyIndexes: file_goe_service_proto_depIdxs,
	}.Build()
	File_goe_service_proto = out.File
	file_goe_service_proto_rawDesc = nil
	file_goe_service_proto_goTypes = nil
	file_goe_service_proto_depIdxs = nil
}
