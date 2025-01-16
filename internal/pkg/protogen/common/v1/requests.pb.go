// Copyright (C) 2024 Declan Teevan
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: stocklet/common/v1/requests.proto

package common_v1

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

type ServiceInfoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ServiceInfoRequest) Reset() {
	*x = ServiceInfoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stocklet_common_v1_requests_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServiceInfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServiceInfoRequest) ProtoMessage() {}

func (x *ServiceInfoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_stocklet_common_v1_requests_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServiceInfoRequest.ProtoReflect.Descriptor instead.
func (*ServiceInfoRequest) Descriptor() ([]byte, []int) {
	return file_stocklet_common_v1_requests_proto_rawDescGZIP(), []int{0}
}

type ServiceInfoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name          string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Source        string `protobuf:"bytes,2,opt,name=source,proto3" json:"source,omitempty"`
	SourceLicense string `protobuf:"bytes,3,opt,name=source_license,json=sourceLicense,proto3" json:"source_license,omitempty"`
}

func (x *ServiceInfoResponse) Reset() {
	*x = ServiceInfoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stocklet_common_v1_requests_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServiceInfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServiceInfoResponse) ProtoMessage() {}

func (x *ServiceInfoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_stocklet_common_v1_requests_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServiceInfoResponse.ProtoReflect.Descriptor instead.
func (*ServiceInfoResponse) Descriptor() ([]byte, []int) {
	return file_stocklet_common_v1_requests_proto_rawDescGZIP(), []int{1}
}

func (x *ServiceInfoResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ServiceInfoResponse) GetSource() string {
	if x != nil {
		return x.Source
	}
	return ""
}

func (x *ServiceInfoResponse) GetSourceLicense() string {
	if x != nil {
		return x.SourceLicense
	}
	return ""
}

var File_stocklet_common_v1_requests_proto protoreflect.FileDescriptor

var file_stocklet_common_v1_requests_proto_rawDesc = []byte{
	0x0a, 0x21, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x6c, 0x65, 0x74, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x12, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x6c, 0x65, 0x74, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x22, 0x14, 0x0a, 0x12, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x68, 0x0a,
	0x13, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x12, 0x25, 0x0a, 0x0e, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x6c, 0x69, 0x63, 0x65, 0x6e,
	0x73, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x4c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x42, 0x47, 0x5a, 0x45, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x68, 0x65, 0x78, 0x6f, 0x6c, 0x61, 0x6e, 0x2f, 0x73, 0x74,
	0x6f, 0x63, 0x6b, 0x6c, 0x65, 0x74, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f,
	0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x67, 0x65, 0x6e, 0x2f, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x3b, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x5f, 0x76, 0x31,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_stocklet_common_v1_requests_proto_rawDescOnce sync.Once
	file_stocklet_common_v1_requests_proto_rawDescData = file_stocklet_common_v1_requests_proto_rawDesc
)

func file_stocklet_common_v1_requests_proto_rawDescGZIP() []byte {
	file_stocklet_common_v1_requests_proto_rawDescOnce.Do(func() {
		file_stocklet_common_v1_requests_proto_rawDescData = protoimpl.X.CompressGZIP(file_stocklet_common_v1_requests_proto_rawDescData)
	})
	return file_stocklet_common_v1_requests_proto_rawDescData
}

var file_stocklet_common_v1_requests_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_stocklet_common_v1_requests_proto_goTypes = []interface{}{
	(*ServiceInfoRequest)(nil),  // 0: stocklet.common.v1.ServiceInfoRequest
	(*ServiceInfoResponse)(nil), // 1: stocklet.common.v1.ServiceInfoResponse
}
var file_stocklet_common_v1_requests_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_stocklet_common_v1_requests_proto_init() }
func file_stocklet_common_v1_requests_proto_init() {
	if File_stocklet_common_v1_requests_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_stocklet_common_v1_requests_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServiceInfoRequest); i {
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
		file_stocklet_common_v1_requests_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServiceInfoResponse); i {
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
			RawDescriptor: file_stocklet_common_v1_requests_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_stocklet_common_v1_requests_proto_goTypes,
		DependencyIndexes: file_stocklet_common_v1_requests_proto_depIdxs,
		MessageInfos:      file_stocklet_common_v1_requests_proto_msgTypes,
	}.Build()
	File_stocklet_common_v1_requests_proto = out.File
	file_stocklet_common_v1_requests_proto_rawDesc = nil
	file_stocklet_common_v1_requests_proto_goTypes = nil
	file_stocklet_common_v1_requests_proto_depIdxs = nil
}
