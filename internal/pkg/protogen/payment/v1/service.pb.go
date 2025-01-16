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
// source: stocklet/payment/v1/service.proto

package payment_v1

import (
	_ "buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
	v1 "github.com/hexolan/stocklet/internal/pkg/protogen/common/v1"
	v11 "github.com/hexolan/stocklet/internal/pkg/protogen/events/v1"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	_ "google.golang.org/genproto/googleapis/api/visibility"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ViewTransactionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TransactionId string `protobuf:"bytes,1,opt,name=transaction_id,json=transactionId,proto3" json:"transaction_id,omitempty"`
}

func (x *ViewTransactionRequest) Reset() {
	*x = ViewTransactionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stocklet_payment_v1_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ViewTransactionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ViewTransactionRequest) ProtoMessage() {}

func (x *ViewTransactionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_stocklet_payment_v1_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ViewTransactionRequest.ProtoReflect.Descriptor instead.
func (*ViewTransactionRequest) Descriptor() ([]byte, []int) {
	return file_stocklet_payment_v1_service_proto_rawDescGZIP(), []int{0}
}

func (x *ViewTransactionRequest) GetTransactionId() string {
	if x != nil {
		return x.TransactionId
	}
	return ""
}

type ViewTransactionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Transaction *Transaction `protobuf:"bytes,1,opt,name=transaction,proto3" json:"transaction,omitempty"`
}

func (x *ViewTransactionResponse) Reset() {
	*x = ViewTransactionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stocklet_payment_v1_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ViewTransactionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ViewTransactionResponse) ProtoMessage() {}

func (x *ViewTransactionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_stocklet_payment_v1_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ViewTransactionResponse.ProtoReflect.Descriptor instead.
func (*ViewTransactionResponse) Descriptor() ([]byte, []int) {
	return file_stocklet_payment_v1_service_proto_rawDescGZIP(), []int{1}
}

func (x *ViewTransactionResponse) GetTransaction() *Transaction {
	if x != nil {
		return x.Transaction
	}
	return nil
}

type ViewBalanceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CustomerId string `protobuf:"bytes,1,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
}

func (x *ViewBalanceRequest) Reset() {
	*x = ViewBalanceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stocklet_payment_v1_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ViewBalanceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ViewBalanceRequest) ProtoMessage() {}

func (x *ViewBalanceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_stocklet_payment_v1_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ViewBalanceRequest.ProtoReflect.Descriptor instead.
func (*ViewBalanceRequest) Descriptor() ([]byte, []int) {
	return file_stocklet_payment_v1_service_proto_rawDescGZIP(), []int{2}
}

func (x *ViewBalanceRequest) GetCustomerId() string {
	if x != nil {
		return x.CustomerId
	}
	return ""
}

type ViewBalanceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Balance *CustomerBalance `protobuf:"bytes,1,opt,name=balance,proto3" json:"balance,omitempty"`
}

func (x *ViewBalanceResponse) Reset() {
	*x = ViewBalanceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stocklet_payment_v1_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ViewBalanceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ViewBalanceResponse) ProtoMessage() {}

func (x *ViewBalanceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_stocklet_payment_v1_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ViewBalanceResponse.ProtoReflect.Descriptor instead.
func (*ViewBalanceResponse) Descriptor() ([]byte, []int) {
	return file_stocklet_payment_v1_service_proto_rawDescGZIP(), []int{3}
}

func (x *ViewBalanceResponse) GetBalance() *CustomerBalance {
	if x != nil {
		return x.Balance
	}
	return nil
}

var File_stocklet_payment_v1_service_proto protoreflect.FileDescriptor

var file_stocklet_payment_v1_service_proto_rawDesc = []byte{
	0x0a, 0x21, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x6c, 0x65, 0x74, 0x2f, 0x70, 0x61, 0x79, 0x6d, 0x65,
	0x6e, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x13, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x6c, 0x65, 0x74, 0x2e, 0x70, 0x61,
	0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x1a, 0x1b, 0x62, 0x75, 0x66, 0x2f, 0x76, 0x61,
	0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x76, 0x69, 0x73, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x21, 0x73,
	0x74, 0x6f, 0x63, 0x6b, 0x6c, 0x65, 0x74, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76,
	0x31, 0x2f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x21, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x6c, 0x65, 0x74, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74,
	0x73, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x68, 0x69, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x6c, 0x65, 0x74, 0x2f, 0x65, 0x76,
	0x65, 0x6e, 0x74, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1f, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x6c, 0x65, 0x74, 0x2f, 0x70, 0x61, 0x79,
	0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x48, 0x0a, 0x16, 0x56, 0x69, 0x65, 0x77, 0x54, 0x72, 0x61, 0x6e, 0x73,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2e, 0x0a,
	0x0e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xba, 0x48, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x0d,
	0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x22, 0x5d, 0x0a,
	0x17, 0x56, 0x69, 0x65, 0x77, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x42, 0x0a, 0x0b, 0x74, 0x72, 0x61, 0x6e,
	0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e,
	0x73, 0x74, 0x6f, 0x63, 0x6b, 0x6c, 0x65, 0x74, 0x2e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74,
	0x2e, 0x76, 0x31, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x0b, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x3e, 0x0a, 0x12,
	0x56, 0x69, 0x65, 0x77, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x28, 0x0a, 0x0b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xba, 0x48, 0x04, 0x72, 0x02, 0x10, 0x01,
	0x52, 0x0a, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x49, 0x64, 0x22, 0x55, 0x0a, 0x13,
	0x56, 0x69, 0x65, 0x77, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x3e, 0x0a, 0x07, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x6c, 0x65, 0x74, 0x2e,
	0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f,
	0x6d, 0x65, 0x72, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x07, 0x62, 0x61, 0x6c, 0x61,
	0x6e, 0x63, 0x65, 0x32, 0x8b, 0x06, 0x0a, 0x0e, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x7b, 0x0a, 0x0b, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x26, 0x2e, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x6c, 0x65, 0x74,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e,
	0x73, 0x74, 0x6f, 0x63, 0x6b, 0x6c, 0x65, 0x74, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x76, 0x31, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x15, 0x12, 0x13,
	0x2f, 0x76, 0x31, 0x2f, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x9e, 0x01, 0x0a, 0x0f, 0x56, 0x69, 0x65, 0x77, 0x54, 0x72, 0x61, 0x6e,
	0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2b, 0x2e, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x6c,
	0x65, 0x74, 0x2e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x56, 0x69,
	0x65, 0x77, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x2c, 0x2e, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x6c, 0x65, 0x74, 0x2e,
	0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x56, 0x69, 0x65, 0x77, 0x54,
	0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x30, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x2a, 0x12, 0x28, 0x2f, 0x76, 0x31, 0x2f,
	0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x2f, 0x7b, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x69, 0x64, 0x7d, 0x12, 0x8b, 0x01, 0x0a, 0x0b, 0x56, 0x69, 0x65, 0x77, 0x42, 0x61, 0x6c,
	0x61, 0x6e, 0x63, 0x65, 0x12, 0x27, 0x2e, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x6c, 0x65, 0x74, 0x2e,
	0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x56, 0x69, 0x65, 0x77, 0x42,
	0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e,
	0x73, 0x74, 0x6f, 0x63, 0x6b, 0x6c, 0x65, 0x74, 0x2e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74,
	0x2e, 0x76, 0x31, 0x2e, 0x56, 0x69, 0x65, 0x77, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x29, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x23, 0x12,
	0x21, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x62, 0x61, 0x6c,
	0x61, 0x6e, 0x63, 0x65, 0x2f, 0x7b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x7d, 0x12, 0x69, 0x0a, 0x17, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x55, 0x73, 0x65,
	0x72, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x24, 0x2e,
	0x73, 0x74, 0x6f, 0x63, 0x6b, 0x6c, 0x65, 0x74, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x10, 0xfa, 0xd2, 0xe4,
	0x93, 0x02, 0x0a, 0x12, 0x08, 0x49, 0x4e, 0x54, 0x45, 0x52, 0x4e, 0x41, 0x4c, 0x12, 0x69, 0x0a,
	0x17, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x55, 0x73, 0x65, 0x72, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x64, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x24, 0x2e, 0x73, 0x74, 0x6f, 0x63, 0x6b,
	0x6c, 0x65, 0x74, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73,
	0x65, 0x72, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x1a, 0x16,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x10, 0xfa, 0xd2, 0xe4, 0x93, 0x02, 0x0a, 0x12, 0x08,
	0x49, 0x4e, 0x54, 0x45, 0x52, 0x4e, 0x41, 0x4c, 0x12, 0x77, 0x0a, 0x1e, 0x50, 0x72, 0x6f, 0x63,
	0x65, 0x73, 0x73, 0x53, 0x68, 0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x41, 0x6c, 0x6c, 0x6f, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x2b, 0x2e, 0x73, 0x74, 0x6f,
	0x63, 0x6b, 0x6c, 0x65, 0x74, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e,
	0x53, 0x68, 0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x41, 0x6c, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22,
	0x10, 0xfa, 0xd2, 0xe4, 0x93, 0x02, 0x0a, 0x12, 0x08, 0x49, 0x4e, 0x54, 0x45, 0x52, 0x4e, 0x41,
	0x4c, 0x42, 0x49, 0x5a, 0x47, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x68, 0x65, 0x78, 0x6f, 0x6c, 0x61, 0x6e, 0x2f, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x6c, 0x65, 0x74,
	0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x67, 0x65, 0x6e, 0x2f, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x76,
	0x31, 0x3b, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_stocklet_payment_v1_service_proto_rawDescOnce sync.Once
	file_stocklet_payment_v1_service_proto_rawDescData = file_stocklet_payment_v1_service_proto_rawDesc
)

func file_stocklet_payment_v1_service_proto_rawDescGZIP() []byte {
	file_stocklet_payment_v1_service_proto_rawDescOnce.Do(func() {
		file_stocklet_payment_v1_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_stocklet_payment_v1_service_proto_rawDescData)
	})
	return file_stocklet_payment_v1_service_proto_rawDescData
}

var file_stocklet_payment_v1_service_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_stocklet_payment_v1_service_proto_goTypes = []interface{}{
	(*ViewTransactionRequest)(nil),      // 0: stocklet.payment.v1.ViewTransactionRequest
	(*ViewTransactionResponse)(nil),     // 1: stocklet.payment.v1.ViewTransactionResponse
	(*ViewBalanceRequest)(nil),          // 2: stocklet.payment.v1.ViewBalanceRequest
	(*ViewBalanceResponse)(nil),         // 3: stocklet.payment.v1.ViewBalanceResponse
	(*Transaction)(nil),                 // 4: stocklet.payment.v1.Transaction
	(*CustomerBalance)(nil),             // 5: stocklet.payment.v1.CustomerBalance
	(*v1.ServiceInfoRequest)(nil),       // 6: stocklet.common.v1.ServiceInfoRequest
	(*v11.UserCreatedEvent)(nil),        // 7: stocklet.events.v1.UserCreatedEvent
	(*v11.UserDeletedEvent)(nil),        // 8: stocklet.events.v1.UserDeletedEvent
	(*v11.ShipmentAllocationEvent)(nil), // 9: stocklet.events.v1.ShipmentAllocationEvent
	(*v1.ServiceInfoResponse)(nil),      // 10: stocklet.common.v1.ServiceInfoResponse
	(*emptypb.Empty)(nil),               // 11: google.protobuf.Empty
}
var file_stocklet_payment_v1_service_proto_depIdxs = []int32{
	4,  // 0: stocklet.payment.v1.ViewTransactionResponse.transaction:type_name -> stocklet.payment.v1.Transaction
	5,  // 1: stocklet.payment.v1.ViewBalanceResponse.balance:type_name -> stocklet.payment.v1.CustomerBalance
	6,  // 2: stocklet.payment.v1.PaymentService.ServiceInfo:input_type -> stocklet.common.v1.ServiceInfoRequest
	0,  // 3: stocklet.payment.v1.PaymentService.ViewTransaction:input_type -> stocklet.payment.v1.ViewTransactionRequest
	2,  // 4: stocklet.payment.v1.PaymentService.ViewBalance:input_type -> stocklet.payment.v1.ViewBalanceRequest
	7,  // 5: stocklet.payment.v1.PaymentService.ProcessUserCreatedEvent:input_type -> stocklet.events.v1.UserCreatedEvent
	8,  // 6: stocklet.payment.v1.PaymentService.ProcessUserDeletedEvent:input_type -> stocklet.events.v1.UserDeletedEvent
	9,  // 7: stocklet.payment.v1.PaymentService.ProcessShipmentAllocationEvent:input_type -> stocklet.events.v1.ShipmentAllocationEvent
	10, // 8: stocklet.payment.v1.PaymentService.ServiceInfo:output_type -> stocklet.common.v1.ServiceInfoResponse
	1,  // 9: stocklet.payment.v1.PaymentService.ViewTransaction:output_type -> stocklet.payment.v1.ViewTransactionResponse
	3,  // 10: stocklet.payment.v1.PaymentService.ViewBalance:output_type -> stocklet.payment.v1.ViewBalanceResponse
	11, // 11: stocklet.payment.v1.PaymentService.ProcessUserCreatedEvent:output_type -> google.protobuf.Empty
	11, // 12: stocklet.payment.v1.PaymentService.ProcessUserDeletedEvent:output_type -> google.protobuf.Empty
	11, // 13: stocklet.payment.v1.PaymentService.ProcessShipmentAllocationEvent:output_type -> google.protobuf.Empty
	8,  // [8:14] is the sub-list for method output_type
	2,  // [2:8] is the sub-list for method input_type
	2,  // [2:2] is the sub-list for extension type_name
	2,  // [2:2] is the sub-list for extension extendee
	0,  // [0:2] is the sub-list for field type_name
}

func init() { file_stocklet_payment_v1_service_proto_init() }
func file_stocklet_payment_v1_service_proto_init() {
	if File_stocklet_payment_v1_service_proto != nil {
		return
	}
	file_stocklet_payment_v1_types_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_stocklet_payment_v1_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ViewTransactionRequest); i {
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
		file_stocklet_payment_v1_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ViewTransactionResponse); i {
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
		file_stocklet_payment_v1_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ViewBalanceRequest); i {
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
		file_stocklet_payment_v1_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ViewBalanceResponse); i {
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
			RawDescriptor: file_stocklet_payment_v1_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_stocklet_payment_v1_service_proto_goTypes,
		DependencyIndexes: file_stocklet_payment_v1_service_proto_depIdxs,
		MessageInfos:      file_stocklet_payment_v1_service_proto_msgTypes,
	}.Build()
	File_stocklet_payment_v1_service_proto = out.File
	file_stocklet_payment_v1_service_proto_rawDesc = nil
	file_stocklet_payment_v1_service_proto_goTypes = nil
	file_stocklet_payment_v1_service_proto_depIdxs = nil
}
