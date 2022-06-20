// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.21.1
// source: proto/handler/message/admin.proto

package message

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

type CheckAdminRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Password string `protobuf:"bytes,1,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *CheckAdminRequest) Reset() {
	*x = CheckAdminRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_service_message_admin_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckAdminRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckAdminRequest) ProtoMessage() {}

func (x *CheckAdminRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_service_message_admin_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckAdminRequest.ProtoReflect.Descriptor instead.
func (*CheckAdminRequest) Descriptor() ([]byte, []int) {
	return file_proto_service_message_admin_proto_rawDescGZIP(), []int{0}
}

func (x *CheckAdminRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type CheckAdminResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsAdmin bool `protobuf:"varint,1,opt,name=is_admin,json=isAdmin,proto3" json:"is_admin,omitempty"`
}

func (x *CheckAdminResponse) Reset() {
	*x = CheckAdminResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_service_message_admin_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckAdminResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckAdminResponse) ProtoMessage() {}

func (x *CheckAdminResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_service_message_admin_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckAdminResponse.ProtoReflect.Descriptor instead.
func (*CheckAdminResponse) Descriptor() ([]byte, []int) {
	return file_proto_service_message_admin_proto_rawDescGZIP(), []int{1}
}

func (x *CheckAdminResponse) GetIsAdmin() bool {
	if x != nil {
		return x.IsAdmin
	}
	return false
}

var File_proto_service_message_admin_proto protoreflect.FileDescriptor

var file_proto_service_message_admin_proto_rawDesc = []byte{
	0x0a, 0x21, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x12, 0x76, 0x32, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x2f, 0x0a, 0x11, 0x43, 0x68, 0x65, 0x63, 0x6b,
	0x41, 0x64, 0x6d, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08,
	0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x2f, 0x0a, 0x12, 0x43, 0x68, 0x65, 0x63,
	0x6b, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x19,
	0x0a, 0x08, 0x69, 0x73, 0x5f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x07, 0x69, 0x73, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x42, 0x37, 0x5a, 0x35, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x74, 0x63, 0x30, 0x33, 0x30, 0x31, 0x32,
	0x2f, 0x6d, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_service_message_admin_proto_rawDescOnce sync.Once
	file_proto_service_message_admin_proto_rawDescData = file_proto_service_message_admin_proto_rawDesc
)

func file_proto_service_message_admin_proto_rawDescGZIP() []byte {
	file_proto_service_message_admin_proto_rawDescOnce.Do(func() {
		file_proto_service_message_admin_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_service_message_admin_proto_rawDescData)
	})
	return file_proto_service_message_admin_proto_rawDescData
}

var file_proto_service_message_admin_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_service_message_admin_proto_goTypes = []interface{}{
	(*CheckAdminRequest)(nil),  // 0: v2.handler.message.CheckAdminRequest
	(*CheckAdminResponse)(nil), // 1: v2.handler.message.CheckAdminResponse
}
var file_proto_service_message_admin_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_service_message_admin_proto_init() }
func file_proto_service_message_admin_proto_init() {
	if File_proto_service_message_admin_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_service_message_admin_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckAdminRequest); i {
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
		file_proto_service_message_admin_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckAdminResponse); i {
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
			RawDescriptor: file_proto_service_message_admin_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_service_message_admin_proto_goTypes,
		DependencyIndexes: file_proto_service_message_admin_proto_depIdxs,
		MessageInfos:      file_proto_service_message_admin_proto_msgTypes,
	}.Build()
	File_proto_service_message_admin_proto = out.File
	file_proto_service_message_admin_proto_rawDesc = nil
	file_proto_service_message_admin_proto_goTypes = nil
	file_proto_service_message_admin_proto_depIdxs = nil
}
