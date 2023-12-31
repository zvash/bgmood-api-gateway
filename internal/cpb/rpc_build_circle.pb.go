// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.23.4
// source: rpc_build_circle.proto

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

type CreateCircleRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title       string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Avatar      string `protobuf:"bytes,2,opt,name=avatar,proto3" json:"avatar,omitempty"`
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	CircleType  string `protobuf:"bytes,4,opt,name=circle_type,json=circleType,proto3" json:"circle_type,omitempty"`
	IsPrivate   bool   `protobuf:"varint,5,opt,name=is_private,json=isPrivate,proto3" json:"is_private,omitempty"`
}

func (x *CreateCircleRequest) Reset() {
	*x = CreateCircleRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_build_circle_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateCircleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCircleRequest) ProtoMessage() {}

func (x *CreateCircleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_build_circle_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCircleRequest.ProtoReflect.Descriptor instead.
func (*CreateCircleRequest) Descriptor() ([]byte, []int) {
	return file_rpc_build_circle_proto_rawDescGZIP(), []int{0}
}

func (x *CreateCircleRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *CreateCircleRequest) GetAvatar() string {
	if x != nil {
		return x.Avatar
	}
	return ""
}

func (x *CreateCircleRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *CreateCircleRequest) GetCircleType() string {
	if x != nil {
		return x.CircleType
	}
	return ""
}

func (x *CreateCircleRequest) GetIsPrivate() bool {
	if x != nil {
		return x.IsPrivate
	}
	return false
}

type CreateCircleResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Circle *Circle `protobuf:"bytes,1,opt,name=circle,proto3" json:"circle,omitempty"`
}

func (x *CreateCircleResponse) Reset() {
	*x = CreateCircleResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_build_circle_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateCircleResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCircleResponse) ProtoMessage() {}

func (x *CreateCircleResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_build_circle_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCircleResponse.ProtoReflect.Descriptor instead.
func (*CreateCircleResponse) Descriptor() ([]byte, []int) {
	return file_rpc_build_circle_proto_rawDescGZIP(), []int{1}
}

func (x *CreateCircleResponse) GetCircle() *Circle {
	if x != nil {
		return x.Circle
	}
	return nil
}

var File_rpc_build_circle_proto protoreflect.FileDescriptor

var file_rpc_build_circle_proto_rawDesc = []byte{
	0x0a, 0x16, 0x72, 0x70, 0x63, 0x5f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x5f, 0x63, 0x69, 0x72, 0x63,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x63, 0x70, 0x62, 0x1a, 0x0c, 0x63,
	0x69, 0x72, 0x63, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa5, 0x01, 0x0a, 0x13,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x69, 0x72, 0x63, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x76, 0x61,
	0x74, 0x61, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x76, 0x61, 0x74, 0x61,
	0x72, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x69, 0x72, 0x63, 0x6c, 0x65, 0x5f, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x69, 0x72, 0x63, 0x6c, 0x65,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x73, 0x5f, 0x70, 0x72, 0x69, 0x76, 0x61,
	0x74, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x69, 0x73, 0x50, 0x72, 0x69, 0x76,
	0x61, 0x74, 0x65, 0x22, 0x3b, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x69, 0x72,
	0x63, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x23, 0x0a, 0x06, 0x63,
	0x69, 0x72, 0x63, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x63, 0x70,
	0x62, 0x2e, 0x43, 0x69, 0x72, 0x63, 0x6c, 0x65, 0x52, 0x06, 0x63, 0x69, 0x72, 0x63, 0x6c, 0x65,
	0x42, 0x2d, 0x5a, 0x2b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x7a,
	0x76, 0x61, 0x73, 0x68, 0x2f, 0x62, 0x67, 0x6d, 0x6f, 0x6f, 0x64, 0x2d, 0x63, 0x69, 0x72, 0x63,
	0x6c, 0x65, 0x73, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x63, 0x70, 0x62, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_rpc_build_circle_proto_rawDescOnce sync.Once
	file_rpc_build_circle_proto_rawDescData = file_rpc_build_circle_proto_rawDesc
)

func file_rpc_build_circle_proto_rawDescGZIP() []byte {
	file_rpc_build_circle_proto_rawDescOnce.Do(func() {
		file_rpc_build_circle_proto_rawDescData = protoimpl.X.CompressGZIP(file_rpc_build_circle_proto_rawDescData)
	})
	return file_rpc_build_circle_proto_rawDescData
}

var file_rpc_build_circle_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_rpc_build_circle_proto_goTypes = []interface{}{
	(*CreateCircleRequest)(nil),  // 0: cpb.CreateCircleRequest
	(*CreateCircleResponse)(nil), // 1: cpb.CreateCircleResponse
	(*Circle)(nil),               // 2: cpb.Circle
}
var file_rpc_build_circle_proto_depIdxs = []int32{
	2, // 0: cpb.CreateCircleResponse.circle:type_name -> cpb.Circle
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_rpc_build_circle_proto_init() }
func file_rpc_build_circle_proto_init() {
	if File_rpc_build_circle_proto != nil {
		return
	}
	file_circle_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_rpc_build_circle_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateCircleRequest); i {
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
		file_rpc_build_circle_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateCircleResponse); i {
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
			RawDescriptor: file_rpc_build_circle_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_rpc_build_circle_proto_goTypes,
		DependencyIndexes: file_rpc_build_circle_proto_depIdxs,
		MessageInfos:      file_rpc_build_circle_proto_msgTypes,
	}.Build()
	File_rpc_build_circle_proto = out.File
	file_rpc_build_circle_proto_rawDesc = nil
	file_rpc_build_circle_proto_goTypes = nil
	file_rpc_build_circle_proto_depIdxs = nil
}
