// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.21.1
// source: proto/service/service.proto

package service

import (
	message "github.com/dtc03012/me/protobuf/proto/service/message"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

var File_proto_service_service_proto protoreflect.FileDescriptor

var file_proto_service_service_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x76,
	0x32, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x22, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2f, 0x77,
	0x69, 0x64, 0x67, 0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x21, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x2f, 0x70, 0x6f, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x32, 0xc4, 0x13, 0x0a, 0x02, 0x6d, 0x65, 0x12, 0x84, 0x01, 0x0a, 0x0d, 0x46, 0x69, 0x6e, 0x64,
	0x41, 0x64, 0x6d, 0x69, 0x6e, 0x55, 0x55, 0x49, 0x44, 0x12, 0x28, 0x2e, 0x76, 0x32, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x46,
	0x69, 0x6e, 0x64, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x55, 0x55, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x29, 0x2e, 0x76, 0x32, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x41, 0x64, 0x6d,
	0x69, 0x6e, 0x55, 0x55, 0x49, 0x44, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1e,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x18, 0x22, 0x13, 0x2f, 0x76, 0x32, 0x2f, 0x66, 0x69, 0x6e, 0x64,
	0x2d, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2d, 0x75, 0x75, 0x69, 0x64, 0x3a, 0x01, 0x2a, 0x12, 0x8c,
	0x01, 0x0a, 0x0f, 0x49, 0x6e, 0x73, 0x65, 0x72, 0x74, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x55, 0x55,
	0x49, 0x44, 0x12, 0x2a, 0x2e, 0x76, 0x32, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x49, 0x6e, 0x73, 0x65, 0x72, 0x74, 0x41, 0x64,
	0x6d, 0x69, 0x6e, 0x55, 0x55, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2b,
	0x2e, 0x76, 0x32, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x2e, 0x49, 0x6e, 0x73, 0x65, 0x72, 0x74, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x55,
	0x55, 0x49, 0x44, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x20, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x1a, 0x22, 0x15, 0x2f, 0x76, 0x32, 0x2f, 0x69, 0x6e, 0x73, 0x65, 0x72, 0x74, 0x2d,
	0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2d, 0x75, 0x75, 0x69, 0x64, 0x3a, 0x01, 0x2a, 0x12, 0x77, 0x0a,
	0x0a, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x12, 0x25, 0x2e, 0x76, 0x32,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x26, 0x2e, 0x76, 0x32, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x41, 0x64, 0x6d,
	0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1a, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x14, 0x22, 0x0f, 0x2f, 0x76, 0x32, 0x2f, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x2d, 0x61, 0x64,
	0x6d, 0x69, 0x6e, 0x3a, 0x01, 0x2a, 0x12, 0xa7, 0x01, 0x0a, 0x14, 0x46, 0x65, 0x74, 0x63, 0x68,
	0x44, 0x69, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x57, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x12,
	0x2f, 0x2e, 0x76, 0x32, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x2e, 0x46, 0x65, 0x74, 0x63, 0x68, 0x44, 0x69, 0x73, 0x74, 0x72, 0x69,
	0x63, 0x74, 0x57, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x30, 0x2e, 0x76, 0x32, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x46, 0x65, 0x74, 0x63, 0x68, 0x44, 0x69, 0x73, 0x74, 0x72,
	0x69, 0x63, 0x74, 0x57, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x2c, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x26, 0x12, 0x24, 0x2f, 0x76, 0x32, 0x2f,
	0x66, 0x65, 0x74, 0x63, 0x68, 0x2d, 0x64, 0x69, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x2d, 0x77,
	0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x2f, 0x7b, 0x6e, 0x78, 0x7d, 0x2f, 0x7b, 0x6e, 0x79, 0x7d,
	0x12, 0x7d, 0x0a, 0x0a, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x50, 0x6f, 0x73, 0x74, 0x12, 0x25,
	0x2e, 0x76, 0x32, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x2e, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x76, 0x32, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x55, 0x70, 0x6c, 0x6f, 0x61,
	0x64, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x20, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x1a, 0x22, 0x15, 0x2f, 0x76, 0x32, 0x2f, 0x75, 0x70, 0x6c, 0x6f, 0x61,
	0x64, 0x2d, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2d, 0x70, 0x6f, 0x73, 0x74, 0x3a, 0x01, 0x2a, 0x12,
	0x87, 0x01, 0x0a, 0x0d, 0x46, 0x65, 0x74, 0x63, 0x68, 0x50, 0x6f, 0x73, 0x74, 0x4c, 0x69, 0x73,
	0x74, 0x12, 0x28, 0x2e, 0x76, 0x32, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x46, 0x65, 0x74, 0x63, 0x68, 0x50, 0x6f, 0x73, 0x74,
	0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x29, 0x2e, 0x76, 0x32,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x2e, 0x46, 0x65, 0x74, 0x63, 0x68, 0x50, 0x6f, 0x73, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x21, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1b, 0x12, 0x19,
	0x2f, 0x76, 0x32, 0x2f, 0x66, 0x65, 0x74, 0x63, 0x68, 0x2d, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2d,
	0x70, 0x6f, 0x73, 0x74, 0x2d, 0x6c, 0x69, 0x73, 0x74, 0x12, 0x76, 0x0a, 0x09, 0x46, 0x65, 0x74,
	0x63, 0x68, 0x50, 0x6f, 0x73, 0x74, 0x12, 0x24, 0x2e, 0x76, 0x32, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x46, 0x65, 0x74, 0x63,
	0x68, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x76,
	0x32, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x2e, 0x46, 0x65, 0x74, 0x63, 0x68, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x1c, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x16, 0x12, 0x14, 0x2f, 0x76, 0x32,
	0x2f, 0x66, 0x65, 0x74, 0x63, 0x68, 0x2d, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2d, 0x70, 0x6f, 0x73,
	0x74, 0x12, 0x7d, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x74, 0x12,
	0x25, 0x2e, 0x76, 0x32, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x76, 0x32, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x20,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1a, 0x22, 0x15, 0x2f, 0x76, 0x32, 0x2f, 0x64, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x2d, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2d, 0x70, 0x6f, 0x73, 0x74, 0x3a, 0x01, 0x2a,
	0x12, 0x7d, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x74, 0x12, 0x25,
	0x2e, 0x76, 0x32, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x76, 0x32, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x20, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x1a, 0x22, 0x15, 0x2f, 0x76, 0x32, 0x2f, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x2d, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2d, 0x70, 0x6f, 0x73, 0x74, 0x3a, 0x01, 0x2a, 0x12,
	0x94, 0x01, 0x0a, 0x11, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x50, 0x6f, 0x73, 0x74, 0x50, 0x61, 0x73,
	0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x2c, 0x2e, 0x76, 0x32, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b,
	0x50, 0x6f, 0x73, 0x74, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x2d, 0x2e, 0x76, 0x32, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x50, 0x6f,
	0x73, 0x74, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x22, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1c, 0x22, 0x17, 0x2f, 0x76, 0x32, 0x2f,
	0x63, 0x68, 0x65, 0x63, 0x6b, 0x2d, 0x70, 0x6f, 0x73, 0x74, 0x2d, 0x70, 0x61, 0x73, 0x73, 0x77,
	0x6f, 0x72, 0x64, 0x3a, 0x01, 0x2a, 0x12, 0x86, 0x01, 0x0a, 0x0d, 0x49, 0x6e, 0x63, 0x72, 0x65,
	0x6d, 0x65, 0x6e, 0x74, 0x56, 0x69, 0x65, 0x77, 0x12, 0x28, 0x2e, 0x76, 0x32, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x49, 0x6e,
	0x63, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x56, 0x69, 0x65, 0x77, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x29, 0x2e, 0x76, 0x32, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x49, 0x6e, 0x63, 0x72, 0x65, 0x6d, 0x65, 0x6e,
	0x74, 0x56, 0x69, 0x65, 0x77, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x20, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x1a, 0x1a, 0x18, 0x2f, 0x76, 0x32, 0x2f, 0x69, 0x6e, 0x63, 0x72, 0x65,
	0x6d, 0x65, 0x6e, 0x74, 0x2d, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2d, 0x76, 0x69, 0x65, 0x77, 0x12,
	0x85, 0x01, 0x0a, 0x0c, 0x4c, 0x65, 0x61, 0x76, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74,
	0x12, 0x27, 0x2e, 0x76, 0x32, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x4c, 0x65, 0x61, 0x76, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x65,
	0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e, 0x76, 0x32, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x4c,
	0x65, 0x61, 0x76, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x22, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1c, 0x22, 0x17, 0x2f, 0x76, 0x32,
	0x2f, 0x6c, 0x65, 0x61, 0x76, 0x65, 0x2d, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2d, 0x63, 0x6f, 0x6d,
	0x6d, 0x65, 0x6e, 0x74, 0x3a, 0x01, 0x2a, 0x12, 0x93, 0x01, 0x0a, 0x10, 0x46, 0x65, 0x74, 0x63,
	0x68, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x2b, 0x2e, 0x76,
	0x32, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x2e, 0x46, 0x65, 0x74, 0x63, 0x68, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x4c, 0x69,
	0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2c, 0x2e, 0x76, 0x32, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x46,
	0x65, 0x74, 0x63, 0x68, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x24, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1e, 0x12,
	0x1c, 0x2f, 0x76, 0x32, 0x2f, 0x66, 0x65, 0x74, 0x63, 0x68, 0x2d, 0x62, 0x6f, 0x61, 0x72, 0x64,
	0x2d, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x2d, 0x6c, 0x69, 0x73, 0x74, 0x12, 0x89, 0x01,
	0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x12,
	0x28, 0x2e, 0x76, 0x32, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x65,
	0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x29, 0x2e, 0x76, 0x32, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x23, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1d, 0x22, 0x18, 0x2f, 0x76,
	0x32, 0x2f, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x2d, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2d, 0x63,
	0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x3a, 0x01, 0x2a, 0x12, 0x86, 0x01, 0x0a, 0x0d, 0x49, 0x6e,
	0x63, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x4c, 0x69, 0x6b, 0x65, 0x12, 0x28, 0x2e, 0x76, 0x32,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x2e, 0x49, 0x6e, 0x63, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x4c, 0x69, 0x6b, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x29, 0x2e, 0x76, 0x32, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x49, 0x6e, 0x63, 0x72, 0x65,
	0x6d, 0x65, 0x6e, 0x74, 0x4c, 0x69, 0x6b, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x20, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1a, 0x1a, 0x18, 0x2f, 0x76, 0x32, 0x2f, 0x69, 0x6e,
	0x63, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2d, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2d, 0x6c, 0x69,
	0x6b, 0x65, 0x12, 0x86, 0x01, 0x0a, 0x0d, 0x44, 0x65, 0x63, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74,
	0x4c, 0x69, 0x6b, 0x65, 0x12, 0x28, 0x2e, 0x76, 0x32, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x44, 0x65, 0x63, 0x72, 0x65, 0x6d,
	0x65, 0x6e, 0x74, 0x4c, 0x69, 0x6b, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x29,
	0x2e, 0x76, 0x32, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x2e, 0x44, 0x65, 0x63, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x4c, 0x69, 0x6b,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x20, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x1a, 0x2a, 0x18, 0x2f, 0x76, 0x32, 0x2f, 0x64, 0x65, 0x63, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74,
	0x2d, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2d, 0x6c, 0x69, 0x6b, 0x65, 0x12, 0x8d, 0x01, 0x0a, 0x10,
	0x43, 0x68, 0x65, 0x63, 0x6b, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x50, 0x6f, 0x73, 0x74, 0x49, 0x64,
	0x12, 0x2b, 0x2e, 0x76, 0x32, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x56, 0x61, 0x6c, 0x69, 0x64,
	0x50, 0x6f, 0x73, 0x74, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2c, 0x2e,
	0x76, 0x32, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x50, 0x6f, 0x73,
	0x74, 0x49, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1e, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x18, 0x12, 0x16, 0x2f, 0x76, 0x32, 0x2f, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x2d, 0x76,
	0x61, 0x6c, 0x69, 0x64, 0x2d, 0x70, 0x6f, 0x73, 0x74, 0x49, 0x64, 0x12, 0x99, 0x01, 0x0a, 0x13,
	0x43, 0x68, 0x65, 0x63, 0x6b, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e,
	0x74, 0x49, 0x64, 0x12, 0x2e, 0x2e, 0x76, 0x32, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x56, 0x61,
	0x6c, 0x69, 0x64, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x2f, 0x2e, 0x76, 0x32, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x56, 0x61,
	0x6c, 0x69, 0x64, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x21, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1b, 0x12, 0x19, 0x2f, 0x76,
	0x32, 0x2f, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x2d, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x2d, 0x63, 0x6f,
	0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x42, 0x2f, 0x5a, 0x2d, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x74, 0x63, 0x30, 0x33, 0x30, 0x31, 0x32, 0x2f, 0x6d,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_proto_service_service_proto_goTypes = []interface{}{
	(*message.FindAdminUUIDRequest)(nil),         // 0: v2.service.message.FindAdminUUIDRequest
	(*message.InsertAdminUUIDRequest)(nil),       // 1: v2.service.message.InsertAdminUUIDRequest
	(*message.LoginAdminRequest)(nil),            // 2: v2.service.message.LoginAdminRequest
	(*message.FetchDistrictWeatherRequest)(nil),  // 3: v2.service.message.FetchDistrictWeatherRequest
	(*message.UploadPostRequest)(nil),            // 4: v2.service.message.UploadPostRequest
	(*message.FetchPostListRequest)(nil),         // 5: v2.service.message.FetchPostListRequest
	(*message.FetchPostRequest)(nil),             // 6: v2.service.message.FetchPostRequest
	(*message.DeletePostRequest)(nil),            // 7: v2.service.message.DeletePostRequest
	(*message.UpdatePostRequest)(nil),            // 8: v2.service.message.UpdatePostRequest
	(*message.CheckPostPasswordRequest)(nil),     // 9: v2.service.message.CheckPostPasswordRequest
	(*message.IncrementViewRequest)(nil),         // 10: v2.service.message.IncrementViewRequest
	(*message.LeaveCommentRequest)(nil),          // 11: v2.service.message.LeaveCommentRequest
	(*message.FetchCommentListRequest)(nil),      // 12: v2.service.message.FetchCommentListRequest
	(*message.DeleteCommentRequest)(nil),         // 13: v2.service.message.DeleteCommentRequest
	(*message.IncrementLikeRequest)(nil),         // 14: v2.service.message.IncrementLikeRequest
	(*message.DecrementLikeRequest)(nil),         // 15: v2.service.message.DecrementLikeRequest
	(*message.CheckValidPostIdRequest)(nil),      // 16: v2.service.message.CheckValidPostIdRequest
	(*message.CheckValidCommentIdRequest)(nil),   // 17: v2.service.message.CheckValidCommentIdRequest
	(*message.FindAdminUUIDResponse)(nil),        // 18: v2.service.message.FindAdminUUIDResponse
	(*message.InsertAdminUUIDResponse)(nil),      // 19: v2.service.message.InsertAdminUUIDResponse
	(*message.LoginAdminResponse)(nil),           // 20: v2.service.message.LoginAdminResponse
	(*message.FetchDistrictWeatherResponse)(nil), // 21: v2.service.message.FetchDistrictWeatherResponse
	(*message.UploadPostResponse)(nil),           // 22: v2.service.message.UploadPostResponse
	(*message.FetchPostListResponse)(nil),        // 23: v2.service.message.FetchPostListResponse
	(*message.FetchPostResponse)(nil),            // 24: v2.service.message.FetchPostResponse
	(*message.DeletePostResponse)(nil),           // 25: v2.service.message.DeletePostResponse
	(*message.UpdatePostResponse)(nil),           // 26: v2.service.message.UpdatePostResponse
	(*message.CheckPostPasswordResponse)(nil),    // 27: v2.service.message.CheckPostPasswordResponse
	(*message.IncrementViewResponse)(nil),        // 28: v2.service.message.IncrementViewResponse
	(*message.LeaveCommentResponse)(nil),         // 29: v2.service.message.LeaveCommentResponse
	(*message.FetchCommentListResponse)(nil),     // 30: v2.service.message.FetchCommentListResponse
	(*message.DeleteCommentResponse)(nil),        // 31: v2.service.message.DeleteCommentResponse
	(*message.IncrementLikeResponse)(nil),        // 32: v2.service.message.IncrementLikeResponse
	(*message.DecrementLikeResponse)(nil),        // 33: v2.service.message.DecrementLikeResponse
	(*message.CheckValidPostIdResponse)(nil),     // 34: v2.service.message.CheckValidPostIdResponse
	(*message.CheckValidCommentIdResponse)(nil),  // 35: v2.service.message.CheckValidCommentIdResponse
}
var file_proto_service_service_proto_depIdxs = []int32{
	0,  // 0: v2.service.me.FindAdminUUID:input_type -> v2.service.message.FindAdminUUIDRequest
	1,  // 1: v2.service.me.InsertAdminUUID:input_type -> v2.service.message.InsertAdminUUIDRequest
	2,  // 2: v2.service.me.LoginAdmin:input_type -> v2.service.message.LoginAdminRequest
	3,  // 3: v2.service.me.FetchDistrictWeather:input_type -> v2.service.message.FetchDistrictWeatherRequest
	4,  // 4: v2.service.me.UploadPost:input_type -> v2.service.message.UploadPostRequest
	5,  // 5: v2.service.me.FetchPostList:input_type -> v2.service.message.FetchPostListRequest
	6,  // 6: v2.service.me.FetchPost:input_type -> v2.service.message.FetchPostRequest
	7,  // 7: v2.service.me.DeletePost:input_type -> v2.service.message.DeletePostRequest
	8,  // 8: v2.service.me.UpdatePost:input_type -> v2.service.message.UpdatePostRequest
	9,  // 9: v2.service.me.CheckPostPassword:input_type -> v2.service.message.CheckPostPasswordRequest
	10, // 10: v2.service.me.IncrementView:input_type -> v2.service.message.IncrementViewRequest
	11, // 11: v2.service.me.LeaveComment:input_type -> v2.service.message.LeaveCommentRequest
	12, // 12: v2.service.me.FetchCommentList:input_type -> v2.service.message.FetchCommentListRequest
	13, // 13: v2.service.me.DeleteComment:input_type -> v2.service.message.DeleteCommentRequest
	14, // 14: v2.service.me.IncrementLike:input_type -> v2.service.message.IncrementLikeRequest
	15, // 15: v2.service.me.DecrementLike:input_type -> v2.service.message.DecrementLikeRequest
	16, // 16: v2.service.me.CheckValidPostId:input_type -> v2.service.message.CheckValidPostIdRequest
	17, // 17: v2.service.me.CheckValidCommentId:input_type -> v2.service.message.CheckValidCommentIdRequest
	18, // 18: v2.service.me.FindAdminUUID:output_type -> v2.service.message.FindAdminUUIDResponse
	19, // 19: v2.service.me.InsertAdminUUID:output_type -> v2.service.message.InsertAdminUUIDResponse
	20, // 20: v2.service.me.LoginAdmin:output_type -> v2.service.message.LoginAdminResponse
	21, // 21: v2.service.me.FetchDistrictWeather:output_type -> v2.service.message.FetchDistrictWeatherResponse
	22, // 22: v2.service.me.UploadPost:output_type -> v2.service.message.UploadPostResponse
	23, // 23: v2.service.me.FetchPostList:output_type -> v2.service.message.FetchPostListResponse
	24, // 24: v2.service.me.FetchPost:output_type -> v2.service.message.FetchPostResponse
	25, // 25: v2.service.me.DeletePost:output_type -> v2.service.message.DeletePostResponse
	26, // 26: v2.service.me.UpdatePost:output_type -> v2.service.message.UpdatePostResponse
	27, // 27: v2.service.me.CheckPostPassword:output_type -> v2.service.message.CheckPostPasswordResponse
	28, // 28: v2.service.me.IncrementView:output_type -> v2.service.message.IncrementViewResponse
	29, // 29: v2.service.me.LeaveComment:output_type -> v2.service.message.LeaveCommentResponse
	30, // 30: v2.service.me.FetchCommentList:output_type -> v2.service.message.FetchCommentListResponse
	31, // 31: v2.service.me.DeleteComment:output_type -> v2.service.message.DeleteCommentResponse
	32, // 32: v2.service.me.IncrementLike:output_type -> v2.service.message.IncrementLikeResponse
	33, // 33: v2.service.me.DecrementLike:output_type -> v2.service.message.DecrementLikeResponse
	34, // 34: v2.service.me.CheckValidPostId:output_type -> v2.service.message.CheckValidPostIdResponse
	35, // 35: v2.service.me.CheckValidCommentId:output_type -> v2.service.message.CheckValidCommentIdResponse
	18, // [18:36] is the sub-list for method output_type
	0,  // [0:18] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_proto_service_service_proto_init() }
func file_proto_service_service_proto_init() {
	if File_proto_service_service_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_service_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_service_service_proto_goTypes,
		DependencyIndexes: file_proto_service_service_proto_depIdxs,
	}.Build()
	File_proto_service_service_proto = out.File
	file_proto_service_service_proto_rawDesc = nil
	file_proto_service_service_proto_goTypes = nil
	file_proto_service_service_proto_depIdxs = nil
}
