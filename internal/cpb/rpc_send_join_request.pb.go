// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.23.4
// source: rpc_send_join_request.proto

package cpb

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

type SendJoinRequestRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CircleId string `protobuf:"bytes,1,opt,name=circle_id,json=circleId,proto3" json:"circle_id,omitempty"`
}

func (x *SendJoinRequestRequest) Reset() {
	*x = SendJoinRequestRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_send_join_request_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendJoinRequestRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendJoinRequestRequest) ProtoMessage() {}

func (x *SendJoinRequestRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_send_join_request_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendJoinRequestRequest.ProtoReflect.Descriptor instead.
func (*SendJoinRequestRequest) Descriptor() ([]byte, []int) {
	return file_rpc_send_join_request_proto_rawDescGZIP(), []int{0}
}

func (x *SendJoinRequestRequest) GetCircleId() string {
	if x != nil {
		return x.CircleId
	}
	return ""
}

type SendJoinRequestResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *SendJoinRequestResponse) Reset() {
	*x = SendJoinRequestResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_send_join_request_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendJoinRequestResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendJoinRequestResponse) ProtoMessage() {}

func (x *SendJoinRequestResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_send_join_request_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendJoinRequestResponse.ProtoReflect.Descriptor instead.
func (*SendJoinRequestResponse) Descriptor() ([]byte, []int) {
	return file_rpc_send_join_request_proto_rawDescGZIP(), []int{1}
}

func (x *SendJoinRequestResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_rpc_send_join_request_proto protoreflect.FileDescriptor

var file_rpc_send_join_request_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x72, 0x70, 0x63, 0x5f, 0x73, 0x65, 0x6e, 0x64, 0x5f, 0x6a, 0x6f, 0x69, 0x6e, 0x5f,
	0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x63,
	0x70, 0x62, 0x22, 0x35, 0x0a, 0x16, 0x53, 0x65, 0x6e, 0x64, 0x4a, 0x6f, 0x69, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09,
	0x63, 0x69, 0x72, 0x63, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x63, 0x69, 0x72, 0x63, 0x6c, 0x65, 0x49, 0x64, 0x22, 0x33, 0x0a, 0x17, 0x53, 0x65, 0x6e,
	0x64, 0x4a, 0x6f, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x42, 0x2d,
	0x5a, 0x2b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x7a, 0x76, 0x61,
	0x73, 0x68, 0x2f, 0x62, 0x67, 0x6d, 0x6f, 0x6f, 0x64, 0x2d, 0x63, 0x69, 0x72, 0x63, 0x6c, 0x65,
	0x73, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x63, 0x70, 0x62, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_rpc_send_join_request_proto_rawDescOnce sync.Once
	file_rpc_send_join_request_proto_rawDescData = file_rpc_send_join_request_proto_rawDesc
)

func file_rpc_send_join_request_proto_rawDescGZIP() []byte {
	file_rpc_send_join_request_proto_rawDescOnce.Do(func() {
		file_rpc_send_join_request_proto_rawDescData = protoimpl.X.CompressGZIP(file_rpc_send_join_request_proto_rawDescData)
	})
	return file_rpc_send_join_request_proto_rawDescData
}

var file_rpc_send_join_request_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_rpc_send_join_request_proto_goTypes = []interface{}{
	(*SendJoinRequestRequest)(nil),  // 0: cpb.SendJoinRequestRequest
	(*SendJoinRequestResponse)(nil), // 1: cpb.SendJoinRequestResponse
}
var file_rpc_send_join_request_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_rpc_send_join_request_proto_init() }
func file_rpc_send_join_request_proto_init() {
	if File_rpc_send_join_request_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_rpc_send_join_request_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendJoinRequestRequest); i {
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
		file_rpc_send_join_request_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendJoinRequestResponse); i {
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
			RawDescriptor: file_rpc_send_join_request_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_rpc_send_join_request_proto_goTypes,
		DependencyIndexes: file_rpc_send_join_request_proto_depIdxs,
		MessageInfos:      file_rpc_send_join_request_proto_msgTypes,
	}.Build()
	File_rpc_send_join_request_proto = out.File
	file_rpc_send_join_request_proto_rawDesc = nil
	file_rpc_send_join_request_proto_goTypes = nil
	file_rpc_send_join_request_proto_depIdxs = nil
}