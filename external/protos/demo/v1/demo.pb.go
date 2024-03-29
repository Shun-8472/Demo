// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        (unknown)
// source: demo/v1/demo.proto

package demov1

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

type Demo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Version          string `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
	LatestUpdateTime string `protobuf:"bytes,2,opt,name=latest_update_time,json=latestUpdateTime,proto3" json:"latest_update_time,omitempty"`
}

func (x *Demo) Reset() {
	*x = Demo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_demo_v1_demo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Demo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Demo) ProtoMessage() {}

func (x *Demo) ProtoReflect() protoreflect.Message {
	mi := &file_demo_v1_demo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Demo.ProtoReflect.Descriptor instead.
func (*Demo) Descriptor() ([]byte, []int) {
	return file_demo_v1_demo_proto_rawDescGZIP(), []int{0}
}

func (x *Demo) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *Demo) GetLatestUpdateTime() string {
	if x != nil {
		return x.LatestUpdateTime
	}
	return ""
}

type WhatIsDemoInfoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DemoName string `protobuf:"bytes,1,opt,name=demo_name,json=demoName,proto3" json:"demo_name,omitempty"`
}

func (x *WhatIsDemoInfoRequest) Reset() {
	*x = WhatIsDemoInfoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_demo_v1_demo_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WhatIsDemoInfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WhatIsDemoInfoRequest) ProtoMessage() {}

func (x *WhatIsDemoInfoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_demo_v1_demo_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WhatIsDemoInfoRequest.ProtoReflect.Descriptor instead.
func (*WhatIsDemoInfoRequest) Descriptor() ([]byte, []int) {
	return file_demo_v1_demo_proto_rawDescGZIP(), []int{1}
}

func (x *WhatIsDemoInfoRequest) GetDemoName() string {
	if x != nil {
		return x.DemoName
	}
	return ""
}

type WhatIsDemoInfoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Version          string `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
	LatestUpdateTime string `protobuf:"bytes,2,opt,name=latest_update_time,json=latestUpdateTime,proto3" json:"latest_update_time,omitempty"`
}

func (x *WhatIsDemoInfoResponse) Reset() {
	*x = WhatIsDemoInfoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_demo_v1_demo_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WhatIsDemoInfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WhatIsDemoInfoResponse) ProtoMessage() {}

func (x *WhatIsDemoInfoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_demo_v1_demo_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WhatIsDemoInfoResponse.ProtoReflect.Descriptor instead.
func (*WhatIsDemoInfoResponse) Descriptor() ([]byte, []int) {
	return file_demo_v1_demo_proto_rawDescGZIP(), []int{2}
}

func (x *WhatIsDemoInfoResponse) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *WhatIsDemoInfoResponse) GetLatestUpdateTime() string {
	if x != nil {
		return x.LatestUpdateTime
	}
	return ""
}

var File_demo_v1_demo_proto protoreflect.FileDescriptor

var file_demo_v1_demo_proto_rawDesc = []byte{
	0x0a, 0x12, 0x64, 0x65, 0x6d, 0x6f, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x65, 0x6d, 0x6f, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x64, 0x65, 0x6d, 0x6f, 0x2e, 0x76, 0x31, 0x22, 0x4e, 0x0a,
	0x04, 0x44, 0x65, 0x6d, 0x6f, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12,
	0x2c, 0x0a, 0x12, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x6c, 0x61, 0x74,
	0x65, 0x73, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x34, 0x0a,
	0x15, 0x57, 0x68, 0x61, 0x74, 0x49, 0x73, 0x44, 0x65, 0x6d, 0x6f, 0x49, 0x6e, 0x66, 0x6f, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x64, 0x65, 0x6d, 0x6f, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x65, 0x6d, 0x6f, 0x4e,
	0x61, 0x6d, 0x65, 0x22, 0x60, 0x0a, 0x16, 0x57, 0x68, 0x61, 0x74, 0x49, 0x73, 0x44, 0x65, 0x6d,
	0x6f, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x2c, 0x0a, 0x12, 0x6c, 0x61, 0x74, 0x65, 0x73,
	0x74, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x10, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x54, 0x69, 0x6d, 0x65, 0x32, 0x62, 0x0a, 0x0b, 0x44, 0x65, 0x6d, 0x6f, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x53, 0x0a, 0x0e, 0x57, 0x68, 0x61, 0x74, 0x49, 0x73, 0x44, 0x65,
	0x6d, 0x6f, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1e, 0x2e, 0x64, 0x65, 0x6d, 0x6f, 0x2e, 0x76, 0x31,
	0x2e, 0x57, 0x68, 0x61, 0x74, 0x49, 0x73, 0x44, 0x65, 0x6d, 0x6f, 0x49, 0x6e, 0x66, 0x6f, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x64, 0x65, 0x6d, 0x6f, 0x2e, 0x76, 0x31,
	0x2e, 0x57, 0x68, 0x61, 0x74, 0x49, 0x73, 0x44, 0x65, 0x6d, 0x6f, 0x49, 0x6e, 0x66, 0x6f, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x7a, 0x0a, 0x0b, 0x63, 0x6f, 0x6d,
	0x2e, 0x64, 0x65, 0x6d, 0x6f, 0x2e, 0x76, 0x31, 0x42, 0x09, 0x44, 0x65, 0x6d, 0x6f, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x23, 0x64, 0x65, 0x6d, 0x6f, 0x2f, 0x65, 0x78, 0x74, 0x65,
	0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x64, 0x65, 0x6d, 0x6f,
	0x2f, 0x76, 0x31, 0x3b, 0x64, 0x65, 0x6d, 0x6f, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x44, 0x58, 0x58,
	0xaa, 0x02, 0x07, 0x44, 0x65, 0x6d, 0x6f, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x07, 0x44, 0x65, 0x6d,
	0x6f, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x13, 0x44, 0x65, 0x6d, 0x6f, 0x5c, 0x56, 0x31, 0x5c, 0x47,
	0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x08, 0x44, 0x65, 0x6d,
	0x6f, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_demo_v1_demo_proto_rawDescOnce sync.Once
	file_demo_v1_demo_proto_rawDescData = file_demo_v1_demo_proto_rawDesc
)

func file_demo_v1_demo_proto_rawDescGZIP() []byte {
	file_demo_v1_demo_proto_rawDescOnce.Do(func() {
		file_demo_v1_demo_proto_rawDescData = protoimpl.X.CompressGZIP(file_demo_v1_demo_proto_rawDescData)
	})
	return file_demo_v1_demo_proto_rawDescData
}

var file_demo_v1_demo_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_demo_v1_demo_proto_goTypes = []interface{}{
	(*Demo)(nil),                   // 0: demo.v1.Demo
	(*WhatIsDemoInfoRequest)(nil),  // 1: demo.v1.WhatIsDemoInfoRequest
	(*WhatIsDemoInfoResponse)(nil), // 2: demo.v1.WhatIsDemoInfoResponse
}
var file_demo_v1_demo_proto_depIdxs = []int32{
	1, // 0: demo.v1.DemoService.WhatIsDemoInfo:input_type -> demo.v1.WhatIsDemoInfoRequest
	2, // 1: demo.v1.DemoService.WhatIsDemoInfo:output_type -> demo.v1.WhatIsDemoInfoResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_demo_v1_demo_proto_init() }
func file_demo_v1_demo_proto_init() {
	if File_demo_v1_demo_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_demo_v1_demo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Demo); i {
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
		file_demo_v1_demo_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WhatIsDemoInfoRequest); i {
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
		file_demo_v1_demo_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WhatIsDemoInfoResponse); i {
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
			RawDescriptor: file_demo_v1_demo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_demo_v1_demo_proto_goTypes,
		DependencyIndexes: file_demo_v1_demo_proto_depIdxs,
		MessageInfos:      file_demo_v1_demo_proto_msgTypes,
	}.Build()
	File_demo_v1_demo_proto = out.File
	file_demo_v1_demo_proto_rawDesc = nil
	file_demo_v1_demo_proto_goTypes = nil
	file_demo_v1_demo_proto_depIdxs = nil
}
