// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.21.1
// source: proto/service/message/widget.proto

package message

import (
	widget "github.com/dtc03012/me/protobuf/proto/entity/widget"
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

type FetchDistrictWeatherRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Nx int32 `protobuf:"varint,1,opt,name=nx,proto3" json:"nx,omitempty"`
	Ny int32 `protobuf:"varint,2,opt,name=ny,proto3" json:"ny,omitempty"`
}

func (x *FetchDistrictWeatherRequest) Reset() {
	*x = FetchDistrictWeatherRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_service_message_widget_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FetchDistrictWeatherRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FetchDistrictWeatherRequest) ProtoMessage() {}

func (x *FetchDistrictWeatherRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_service_message_widget_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FetchDistrictWeatherRequest.ProtoReflect.Descriptor instead.
func (*FetchDistrictWeatherRequest) Descriptor() ([]byte, []int) {
	return file_proto_service_message_widget_proto_rawDescGZIP(), []int{0}
}

func (x *FetchDistrictWeatherRequest) GetNx() int32 {
	if x != nil {
		return x.Nx
	}
	return 0
}

func (x *FetchDistrictWeatherRequest) GetNy() int32 {
	if x != nil {
		return x.Ny
	}
	return 0
}

type FetchDistrictWeatherResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Temperature   *widget.Temperature     `protobuf:"bytes,1,opt,name=temperature,proto3" json:"temperature,omitempty"`
	Sky           []*widget.Sky           `protobuf:"bytes,2,rep,name=sky,proto3" json:"sky,omitempty"`
	Precipitation []*widget.Precipitation `protobuf:"bytes,3,rep,name=precipitation,proto3" json:"precipitation,omitempty"`
}

func (x *FetchDistrictWeatherResponse) Reset() {
	*x = FetchDistrictWeatherResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_service_message_widget_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FetchDistrictWeatherResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FetchDistrictWeatherResponse) ProtoMessage() {}

func (x *FetchDistrictWeatherResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_service_message_widget_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FetchDistrictWeatherResponse.ProtoReflect.Descriptor instead.
func (*FetchDistrictWeatherResponse) Descriptor() ([]byte, []int) {
	return file_proto_service_message_widget_proto_rawDescGZIP(), []int{1}
}

func (x *FetchDistrictWeatherResponse) GetTemperature() *widget.Temperature {
	if x != nil {
		return x.Temperature
	}
	return nil
}

func (x *FetchDistrictWeatherResponse) GetSky() []*widget.Sky {
	if x != nil {
		return x.Sky
	}
	return nil
}

func (x *FetchDistrictWeatherResponse) GetPrecipitation() []*widget.Precipitation {
	if x != nil {
		return x.Precipitation
	}
	return nil
}

var File_proto_service_message_widget_proto protoreflect.FileDescriptor

var file_proto_service_message_widget_proto_rawDesc = []byte{
	0x0a, 0x22, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2f, 0x77, 0x69, 0x64, 0x67, 0x65, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x76, 0x32, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x21, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2f, 0x77, 0x69, 0x64, 0x67, 0x65, 0x74, 0x2f, 0x77, 0x65,
	0x61, 0x74, 0x68, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x3d, 0x0a, 0x1b, 0x46,
	0x65, 0x74, 0x63, 0x68, 0x44, 0x69, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x57, 0x65, 0x61, 0x74,
	0x68, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x6e, 0x78,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x6e, 0x78, 0x12, 0x0e, 0x0a, 0x02, 0x6e, 0x79,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x6e, 0x79, 0x22, 0xcf, 0x01, 0x0a, 0x1c, 0x46,
	0x65, 0x74, 0x63, 0x68, 0x44, 0x69, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x57, 0x65, 0x61, 0x74,
	0x68, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3f, 0x0a, 0x0b, 0x74,
	0x65, 0x6d, 0x70, 0x65, 0x72, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1d, 0x2e, 0x76, 0x32, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x77, 0x69, 0x64,
	0x67, 0x65, 0x74, 0x2e, 0x74, 0x65, 0x6d, 0x70, 0x65, 0x72, 0x61, 0x74, 0x75, 0x72, 0x65, 0x52,
	0x0b, 0x74, 0x65, 0x6d, 0x70, 0x65, 0x72, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x27, 0x0a, 0x03,
	0x73, 0x6b, 0x79, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x76, 0x32, 0x2e, 0x65,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x77, 0x69, 0x64, 0x67, 0x65, 0x74, 0x2e, 0x73, 0x6b, 0x79,
	0x52, 0x03, 0x73, 0x6b, 0x79, 0x12, 0x45, 0x0a, 0x0d, 0x70, 0x72, 0x65, 0x63, 0x69, 0x70, 0x69,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x76,
	0x32, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x77, 0x69, 0x64, 0x67, 0x65, 0x74, 0x2e,
	0x70, 0x72, 0x65, 0x63, 0x69, 0x70, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0d, 0x70,
	0x72, 0x65, 0x63, 0x69, 0x70, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x37, 0x5a, 0x35,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x74, 0x63, 0x30, 0x33,
	0x30, 0x31, 0x32, 0x2f, 0x6d, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_service_message_widget_proto_rawDescOnce sync.Once
	file_proto_service_message_widget_proto_rawDescData = file_proto_service_message_widget_proto_rawDesc
)

func file_proto_service_message_widget_proto_rawDescGZIP() []byte {
	file_proto_service_message_widget_proto_rawDescOnce.Do(func() {
		file_proto_service_message_widget_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_service_message_widget_proto_rawDescData)
	})
	return file_proto_service_message_widget_proto_rawDescData
}

var file_proto_service_message_widget_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_service_message_widget_proto_goTypes = []interface{}{
	(*FetchDistrictWeatherRequest)(nil),  // 0: v2.service.message.FetchDistrictWeatherRequest
	(*FetchDistrictWeatherResponse)(nil), // 1: v2.service.message.FetchDistrictWeatherResponse
	(*widget.Temperature)(nil),           // 2: v2.entity.widget.temperature
	(*widget.Sky)(nil),                   // 3: v2.entity.widget.sky
	(*widget.Precipitation)(nil),         // 4: v2.entity.widget.precipitation
}
var file_proto_service_message_widget_proto_depIdxs = []int32{
	2, // 0: v2.service.message.FetchDistrictWeatherResponse.temperature:type_name -> v2.entity.widget.temperature
	3, // 1: v2.service.message.FetchDistrictWeatherResponse.sky:type_name -> v2.entity.widget.sky
	4, // 2: v2.service.message.FetchDistrictWeatherResponse.precipitation:type_name -> v2.entity.widget.precipitation
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_proto_service_message_widget_proto_init() }
func file_proto_service_message_widget_proto_init() {
	if File_proto_service_message_widget_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_service_message_widget_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FetchDistrictWeatherRequest); i {
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
		file_proto_service_message_widget_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FetchDistrictWeatherResponse); i {
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
			RawDescriptor: file_proto_service_message_widget_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_service_message_widget_proto_goTypes,
		DependencyIndexes: file_proto_service_message_widget_proto_depIdxs,
		MessageInfos:      file_proto_service_message_widget_proto_msgTypes,
	}.Build()
	File_proto_service_message_widget_proto = out.File
	file_proto_service_message_widget_proto_rawDesc = nil
	file_proto_service_message_widget_proto_goTypes = nil
	file_proto_service_message_widget_proto_depIdxs = nil
}
