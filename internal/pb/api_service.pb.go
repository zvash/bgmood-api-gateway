// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.23.4
// source: api_service.proto

package pb

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

var File_api_service_proto protoreflect.FileDescriptor

var file_api_service_proto_rawDesc = []byte{
	0x0a, 0x11, 0x61, 0x70, 0x69, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x1a, 0x1c, 0x72, 0x70, 0x63, 0x5f, 0x61, 0x70, 0x69,
	0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1a, 0x72, 0x70, 0x63, 0x5f, 0x61, 0x70, 0x69, 0x5f, 0x62,
	0x75, 0x69, 0x6c, 0x64, 0x5f, 0x63, 0x69, 0x72, 0x63, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x18, 0x72, 0x70, 0x63, 0x5f, 0x61, 0x70, 0x69, 0x5f, 0x73, 0x68, 0x61, 0x72, 0x65,
	0x5f, 0x6d, 0x6f, 0x6f, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x26, 0x72, 0x70, 0x63,
	0x5f, 0x61, 0x70, 0x69, 0x5f, 0x67, 0x65, 0x74, 0x5f, 0x63, 0x69, 0x72, 0x63, 0x6c, 0x65, 0x5f,
	0x6a, 0x6f, 0x69, 0x6e, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x32, 0xb6, 0x02, 0x0a, 0x03, 0x41, 0x70, 0x70, 0x12, 0x48, 0x0a, 0x0d, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x18, 0x2e, 0x70,
	0x62, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x70, 0x62, 0x2e, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x28, 0x01, 0x12, 0x42, 0x0a, 0x0b, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x43, 0x69,
	0x72, 0x63, 0x6c, 0x65, 0x12, 0x16, 0x2e, 0x70, 0x62, 0x2e, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x43,
	0x69, 0x72, 0x63, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x70,
	0x62, 0x2e, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x43, 0x69, 0x72, 0x63, 0x6c, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x28, 0x01, 0x12, 0x3c, 0x0a, 0x09, 0x53, 0x68, 0x61,
	0x72, 0x65, 0x4d, 0x6f, 0x6f, 0x64, 0x12, 0x14, 0x2e, 0x70, 0x62, 0x2e, 0x53, 0x68, 0x61, 0x72,
	0x65, 0x4d, 0x6f, 0x6f, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x70,
	0x62, 0x2e, 0x53, 0x68, 0x61, 0x72, 0x65, 0x4d, 0x6f, 0x6f, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x28, 0x01, 0x12, 0x63, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x43, 0x69,
	0x72, 0x63, 0x6c, 0x65, 0x4a, 0x6f, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73,
	0x12, 0x20, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x69, 0x72, 0x63, 0x6c, 0x65, 0x4a,
	0x6f, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x26, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x69, 0x72, 0x63, 0x6c, 0x65, 0x4a, 0x6f,
	0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x57, 0x69, 0x74, 0x68, 0x55, 0x73,
	0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x28, 0x5a, 0x26,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x7a, 0x76, 0x61, 0x73, 0x68,
	0x2f, 0x62, 0x67, 0x6d, 0x6f, 0x6f, 0x64, 0x2d, 0x61, 0x70, 0x69, 0x2d, 0x67, 0x61, 0x74, 0x65,
	0x77, 0x61, 0x79, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_api_service_proto_goTypes = []interface{}{
	(*UpdateProfileRequest)(nil),               // 0: pb.UpdateProfileRequest
	(*BuildCircleRequest)(nil),                 // 1: pb.BuildCircleRequest
	(*ShareMoodRequest)(nil),                   // 2: pb.ShareMoodRequest
	(*GetCircleJoinRequestsRequest)(nil),       // 3: pb.GetCircleJoinRequestsRequest
	(*UpdateProfileResponse)(nil),              // 4: pb.UpdateProfileResponse
	(*BuildCircleResponse)(nil),                // 5: pb.BuildCircleResponse
	(*ShareMoodResponse)(nil),                  // 6: pb.ShareMoodResponse
	(*CircleJoinRequestsWithUserResponse)(nil), // 7: pb.CircleJoinRequestsWithUserResponse
}
var file_api_service_proto_depIdxs = []int32{
	0, // 0: pb.App.UpdateProfile:input_type -> pb.UpdateProfileRequest
	1, // 1: pb.App.BuildCircle:input_type -> pb.BuildCircleRequest
	2, // 2: pb.App.ShareMood:input_type -> pb.ShareMoodRequest
	3, // 3: pb.App.GetCircleJoinRequests:input_type -> pb.GetCircleJoinRequestsRequest
	4, // 4: pb.App.UpdateProfile:output_type -> pb.UpdateProfileResponse
	5, // 5: pb.App.BuildCircle:output_type -> pb.BuildCircleResponse
	6, // 6: pb.App.ShareMood:output_type -> pb.ShareMoodResponse
	7, // 7: pb.App.GetCircleJoinRequests:output_type -> pb.CircleJoinRequestsWithUserResponse
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_service_proto_init() }
func file_api_service_proto_init() {
	if File_api_service_proto != nil {
		return
	}
	file_rpc_api_update_profile_proto_init()
	file_rpc_api_build_circle_proto_init()
	file_rpc_api_share_mood_proto_init()
	file_rpc_api_get_circle_join_requests_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_service_proto_goTypes,
		DependencyIndexes: file_api_service_proto_depIdxs,
	}.Build()
	File_api_service_proto = out.File
	file_api_service_proto_rawDesc = nil
	file_api_service_proto_goTypes = nil
	file_api_service_proto_depIdxs = nil
}
