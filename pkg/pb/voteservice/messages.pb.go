// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: voteservice/messages.proto

package voteservice

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

type VoteRequest_Value int32

const (
	VoteRequest_UP   VoteRequest_Value = 0
	VoteRequest_DOWN VoteRequest_Value = 1
)

// Enum value maps for VoteRequest_Value.
var (
	VoteRequest_Value_name = map[int32]string{
		0: "UP",
		1: "DOWN",
	}
	VoteRequest_Value_value = map[string]int32{
		"UP":   0,
		"DOWN": 1,
	}
)

func (x VoteRequest_Value) Enum() *VoteRequest_Value {
	p := new(VoteRequest_Value)
	*p = x
	return p
}

func (x VoteRequest_Value) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (VoteRequest_Value) Descriptor() protoreflect.EnumDescriptor {
	return file_voteservice_messages_proto_enumTypes[0].Descriptor()
}

func (VoteRequest_Value) Type() protoreflect.EnumType {
	return &file_voteservice_messages_proto_enumTypes[0]
}

func (x VoteRequest_Value) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use VoteRequest_Value.Descriptor instead.
func (VoteRequest_Value) EnumDescriptor() ([]byte, []int) {
	return file_voteservice_messages_proto_rawDescGZIP(), []int{4, 0}
}

type Cryptocurrency struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Symbol string `protobuf:"bytes,1,opt,name=symbol,proto3" json:"symbol,omitempty"`
	Name   string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Votes  int32  `protobuf:"varint,3,opt,name=votes,proto3" json:"votes,omitempty"`
}

func (x *Cryptocurrency) Reset() {
	*x = Cryptocurrency{}
	if protoimpl.UnsafeEnabled {
		mi := &file_voteservice_messages_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Cryptocurrency) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Cryptocurrency) ProtoMessage() {}

func (x *Cryptocurrency) ProtoReflect() protoreflect.Message {
	mi := &file_voteservice_messages_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Cryptocurrency.ProtoReflect.Descriptor instead.
func (*Cryptocurrency) Descriptor() ([]byte, []int) {
	return file_voteservice_messages_proto_rawDescGZIP(), []int{0}
}

func (x *Cryptocurrency) GetSymbol() string {
	if x != nil {
		return x.Symbol
	}
	return ""
}

func (x *Cryptocurrency) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Cryptocurrency) GetVotes() int32 {
	if x != nil {
		return x.Votes
	}
	return 0
}

type GetAllCryptoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cryptocurrencies []*Cryptocurrency `protobuf:"bytes,1,rep,name=cryptocurrencies,proto3" json:"cryptocurrencies,omitempty"`
}

func (x *GetAllCryptoResponse) Reset() {
	*x = GetAllCryptoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_voteservice_messages_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllCryptoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllCryptoResponse) ProtoMessage() {}

func (x *GetAllCryptoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_voteservice_messages_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllCryptoResponse.ProtoReflect.Descriptor instead.
func (*GetAllCryptoResponse) Descriptor() ([]byte, []int) {
	return file_voteservice_messages_proto_rawDescGZIP(), []int{1}
}

func (x *GetAllCryptoResponse) GetCryptocurrencies() []*Cryptocurrency {
	if x != nil {
		return x.Cryptocurrencies
	}
	return nil
}

type GetCryptoBySymbolStreamRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Symbol string `protobuf:"bytes,1,opt,name=symbol,proto3" json:"symbol,omitempty"`
}

func (x *GetCryptoBySymbolStreamRequest) Reset() {
	*x = GetCryptoBySymbolStreamRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_voteservice_messages_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCryptoBySymbolStreamRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCryptoBySymbolStreamRequest) ProtoMessage() {}

func (x *GetCryptoBySymbolStreamRequest) ProtoReflect() protoreflect.Message {
	mi := &file_voteservice_messages_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCryptoBySymbolStreamRequest.ProtoReflect.Descriptor instead.
func (*GetCryptoBySymbolStreamRequest) Descriptor() ([]byte, []int) {
	return file_voteservice_messages_proto_rawDescGZIP(), []int{2}
}

func (x *GetCryptoBySymbolStreamRequest) GetSymbol() string {
	if x != nil {
		return x.Symbol
	}
	return ""
}

type CryptoVotesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cryptocurrencies []*Cryptocurrency `protobuf:"bytes,1,rep,name=cryptocurrencies,proto3" json:"cryptocurrencies,omitempty"`
}

func (x *CryptoVotesResponse) Reset() {
	*x = CryptoVotesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_voteservice_messages_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CryptoVotesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CryptoVotesResponse) ProtoMessage() {}

func (x *CryptoVotesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_voteservice_messages_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CryptoVotesResponse.ProtoReflect.Descriptor instead.
func (*CryptoVotesResponse) Descriptor() ([]byte, []int) {
	return file_voteservice_messages_proto_rawDescGZIP(), []int{3}
}

func (x *CryptoVotesResponse) GetCryptocurrencies() []*Cryptocurrency {
	if x != nil {
		return x.Cryptocurrencies
	}
	return nil
}

type VoteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Symbol   string            `protobuf:"bytes,1,opt,name=symbol,proto3" json:"symbol,omitempty"`
	Value    VoteRequest_Value `protobuf:"varint,2,opt,name=value,proto3,enum=voteservice.VoteRequest_Value" json:"value,omitempty"`
	Username string            `protobuf:"bytes,3,opt,name=username,proto3" json:"username,omitempty"`
}

func (x *VoteRequest) Reset() {
	*x = VoteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_voteservice_messages_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VoteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VoteRequest) ProtoMessage() {}

func (x *VoteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_voteservice_messages_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VoteRequest.ProtoReflect.Descriptor instead.
func (*VoteRequest) Descriptor() ([]byte, []int) {
	return file_voteservice_messages_proto_rawDescGZIP(), []int{4}
}

func (x *VoteRequest) GetSymbol() string {
	if x != nil {
		return x.Symbol
	}
	return ""
}

func (x *VoteRequest) GetValue() VoteRequest_Value {
	if x != nil {
		return x.Value
	}
	return VoteRequest_UP
}

func (x *VoteRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

type VoteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Response bool `protobuf:"varint,1,opt,name=response,proto3" json:"response,omitempty"`
}

func (x *VoteResponse) Reset() {
	*x = VoteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_voteservice_messages_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VoteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VoteResponse) ProtoMessage() {}

func (x *VoteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_voteservice_messages_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VoteResponse.ProtoReflect.Descriptor instead.
func (*VoteResponse) Descriptor() ([]byte, []int) {
	return file_voteservice_messages_proto_rawDescGZIP(), []int{5}
}

func (x *VoteResponse) GetResponse() bool {
	if x != nil {
		return x.Response
	}
	return false
}

type RemoveVoteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Symbol   string `protobuf:"bytes,1,opt,name=symbol,proto3" json:"symbol,omitempty"`
	Username string `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
}

func (x *RemoveVoteRequest) Reset() {
	*x = RemoveVoteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_voteservice_messages_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveVoteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveVoteRequest) ProtoMessage() {}

func (x *RemoveVoteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_voteservice_messages_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveVoteRequest.ProtoReflect.Descriptor instead.
func (*RemoveVoteRequest) Descriptor() ([]byte, []int) {
	return file_voteservice_messages_proto_rawDescGZIP(), []int{6}
}

func (x *RemoveVoteRequest) GetSymbol() string {
	if x != nil {
		return x.Symbol
	}
	return ""
}

func (x *RemoveVoteRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

type RemoveVoteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Response bool `protobuf:"varint,1,opt,name=response,proto3" json:"response,omitempty"`
}

func (x *RemoveVoteResponse) Reset() {
	*x = RemoveVoteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_voteservice_messages_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveVoteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveVoteResponse) ProtoMessage() {}

func (x *RemoveVoteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_voteservice_messages_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveVoteResponse.ProtoReflect.Descriptor instead.
func (*RemoveVoteResponse) Descriptor() ([]byte, []int) {
	return file_voteservice_messages_proto_rawDescGZIP(), []int{7}
}

func (x *RemoveVoteResponse) GetResponse() bool {
	if x != nil {
		return x.Response
	}
	return false
}

var File_voteservice_messages_proto protoreflect.FileDescriptor

var file_voteservice_messages_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x76, 0x6f, 0x74, 0x65, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x76, 0x6f,
	0x74, 0x65, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x22, 0x52, 0x0a, 0x0e, 0x43, 0x72, 0x79,
	0x70, 0x74, 0x6f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x73,
	0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x79, 0x6d,
	0x62, 0x6f, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x6f, 0x74, 0x65, 0x73,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x76, 0x6f, 0x74, 0x65, 0x73, 0x22, 0x5f, 0x0a,
	0x14, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x43, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x47, 0x0a, 0x10, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x63,
	0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x69, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x1b, 0x2e, 0x76, 0x6f, 0x74, 0x65, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x43, 0x72,
	0x79, 0x70, 0x74, 0x6f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x52, 0x10, 0x63, 0x72,
	0x79, 0x70, 0x74, 0x6f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x69, 0x65, 0x73, 0x22, 0x38,
	0x0a, 0x1e, 0x47, 0x65, 0x74, 0x43, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x42, 0x79, 0x53, 0x79, 0x6d,
	0x62, 0x6f, 0x6c, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x16, 0x0a, 0x06, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x22, 0x5e, 0x0a, 0x13, 0x43, 0x72, 0x79, 0x70,
	0x74, 0x6f, 0x56, 0x6f, 0x74, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x47, 0x0a, 0x10, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63,
	0x69, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x76, 0x6f, 0x74, 0x65,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x43, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x63, 0x75,
	0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x52, 0x10, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x63, 0x75,
	0x72, 0x72, 0x65, 0x6e, 0x63, 0x69, 0x65, 0x73, 0x22, 0x92, 0x01, 0x0a, 0x0b, 0x56, 0x6f, 0x74,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x79, 0x6d, 0x62,
	0x6f, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c,
	0x12, 0x34, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x1e, 0x2e, 0x76, 0x6f, 0x74, 0x65, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x56, 0x6f,
	0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61,
	0x6d, 0x65, 0x22, 0x19, 0x0a, 0x05, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x06, 0x0a, 0x02, 0x55,
	0x50, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x44, 0x4f, 0x57, 0x4e, 0x10, 0x01, 0x22, 0x2a, 0x0a,
	0x0c, 0x56, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a,
	0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x47, 0x0a, 0x11, 0x52, 0x65, 0x6d,
	0x6f, 0x76, 0x65, 0x56, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16,
	0x0a, 0x06, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61,
	0x6d, 0x65, 0x22, 0x30, 0x0a, 0x12, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x56, 0x6f, 0x74, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x72, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x42, 0x49, 0x5a, 0x47, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x72, 0x6f, 0x67, 0x65, 0x72, 0x69, 0x6f, 0x34, 0x31, 0x30, 0x2f, 0x63, 0x72,
	0x79, 0x70, 0x74, 0x6f, 0x76, 0x6f, 0x74, 0x65, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x62, 0x2f, 0x76, 0x6f, 0x74, 0x65, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x3b, 0x76, 0x6f, 0x74, 0x65, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_voteservice_messages_proto_rawDescOnce sync.Once
	file_voteservice_messages_proto_rawDescData = file_voteservice_messages_proto_rawDesc
)

func file_voteservice_messages_proto_rawDescGZIP() []byte {
	file_voteservice_messages_proto_rawDescOnce.Do(func() {
		file_voteservice_messages_proto_rawDescData = protoimpl.X.CompressGZIP(file_voteservice_messages_proto_rawDescData)
	})
	return file_voteservice_messages_proto_rawDescData
}

var file_voteservice_messages_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_voteservice_messages_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_voteservice_messages_proto_goTypes = []interface{}{
	(VoteRequest_Value)(0),                 // 0: voteservice.VoteRequest.Value
	(*Cryptocurrency)(nil),                 // 1: voteservice.Cryptocurrency
	(*GetAllCryptoResponse)(nil),           // 2: voteservice.GetAllCryptoResponse
	(*GetCryptoBySymbolStreamRequest)(nil), // 3: voteservice.GetCryptoBySymbolStreamRequest
	(*CryptoVotesResponse)(nil),            // 4: voteservice.CryptoVotesResponse
	(*VoteRequest)(nil),                    // 5: voteservice.VoteRequest
	(*VoteResponse)(nil),                   // 6: voteservice.VoteResponse
	(*RemoveVoteRequest)(nil),              // 7: voteservice.RemoveVoteRequest
	(*RemoveVoteResponse)(nil),             // 8: voteservice.RemoveVoteResponse
}
var file_voteservice_messages_proto_depIdxs = []int32{
	1, // 0: voteservice.GetAllCryptoResponse.cryptocurrencies:type_name -> voteservice.Cryptocurrency
	1, // 1: voteservice.CryptoVotesResponse.cryptocurrencies:type_name -> voteservice.Cryptocurrency
	0, // 2: voteservice.VoteRequest.value:type_name -> voteservice.VoteRequest.Value
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_voteservice_messages_proto_init() }
func file_voteservice_messages_proto_init() {
	if File_voteservice_messages_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_voteservice_messages_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Cryptocurrency); i {
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
		file_voteservice_messages_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllCryptoResponse); i {
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
		file_voteservice_messages_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCryptoBySymbolStreamRequest); i {
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
		file_voteservice_messages_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CryptoVotesResponse); i {
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
		file_voteservice_messages_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VoteRequest); i {
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
		file_voteservice_messages_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VoteResponse); i {
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
		file_voteservice_messages_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoveVoteRequest); i {
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
		file_voteservice_messages_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoveVoteResponse); i {
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
			RawDescriptor: file_voteservice_messages_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_voteservice_messages_proto_goTypes,
		DependencyIndexes: file_voteservice_messages_proto_depIdxs,
		EnumInfos:         file_voteservice_messages_proto_enumTypes,
		MessageInfos:      file_voteservice_messages_proto_msgTypes,
	}.Build()
	File_voteservice_messages_proto = out.File
	file_voteservice_messages_proto_rawDesc = nil
	file_voteservice_messages_proto_goTypes = nil
	file_voteservice_messages_proto_depIdxs = nil
}
