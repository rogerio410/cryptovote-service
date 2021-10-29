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

type Cryptocurrency struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Symbol string `protobuf:"bytes,1,opt,name=symbol,proto3" json:"symbol,omitempty"`
	Name   string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
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

type GetAllCriptoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetAllCriptoRequest) Reset() {
	*x = GetAllCriptoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_voteservice_messages_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllCriptoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllCriptoRequest) ProtoMessage() {}

func (x *GetAllCriptoRequest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use GetAllCriptoRequest.ProtoReflect.Descriptor instead.
func (*GetAllCriptoRequest) Descriptor() ([]byte, []int) {
	return file_voteservice_messages_proto_rawDescGZIP(), []int{1}
}

type GetAllCriptoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cryptocurrencies []*Cryptocurrency `protobuf:"bytes,1,rep,name=cryptocurrencies,proto3" json:"cryptocurrencies,omitempty"`
}

func (x *GetAllCriptoResponse) Reset() {
	*x = GetAllCriptoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_voteservice_messages_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllCriptoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllCriptoResponse) ProtoMessage() {}

func (x *GetAllCriptoResponse) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use GetAllCriptoResponse.ProtoReflect.Descriptor instead.
func (*GetAllCriptoResponse) Descriptor() ([]byte, []int) {
	return file_voteservice_messages_proto_rawDescGZIP(), []int{2}
}

func (x *GetAllCriptoResponse) GetCryptocurrencies() []*Cryptocurrency {
	if x != nil {
		return x.Cryptocurrencies
	}
	return nil
}

type VoteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Symbol string `protobuf:"bytes,1,opt,name=symbol,proto3" json:"symbol,omitempty"`
}

func (x *VoteRequest) Reset() {
	*x = VoteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_voteservice_messages_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VoteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VoteRequest) ProtoMessage() {}

func (x *VoteRequest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use VoteRequest.ProtoReflect.Descriptor instead.
func (*VoteRequest) Descriptor() ([]byte, []int) {
	return file_voteservice_messages_proto_rawDescGZIP(), []int{3}
}

func (x *VoteRequest) GetSymbol() string {
	if x != nil {
		return x.Symbol
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
		mi := &file_voteservice_messages_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VoteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VoteResponse) ProtoMessage() {}

func (x *VoteResponse) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use VoteResponse.ProtoReflect.Descriptor instead.
func (*VoteResponse) Descriptor() ([]byte, []int) {
	return file_voteservice_messages_proto_rawDescGZIP(), []int{4}
}

func (x *VoteResponse) GetResponse() bool {
	if x != nil {
		return x.Response
	}
	return false
}

var File_voteservice_messages_proto protoreflect.FileDescriptor

var file_voteservice_messages_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x76, 0x6f, 0x74, 0x65, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x76, 0x6f,
	0x74, 0x65, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x22, 0x3c, 0x0a, 0x0e, 0x43, 0x72, 0x79,
	0x70, 0x74, 0x6f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x73,
	0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x79, 0x6d,
	0x62, 0x6f, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x15, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x41, 0x6c,
	0x6c, 0x43, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x5f,
	0x0a, 0x14, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x43, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x47, 0x0a, 0x10, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f,
	0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x69, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x1b, 0x2e, 0x76, 0x6f, 0x74, 0x65, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x43,
	0x72, 0x79, 0x70, 0x74, 0x6f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x52, 0x10, 0x63,
	0x72, 0x79, 0x70, 0x74, 0x6f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x69, 0x65, 0x73, 0x22,
	0x25, 0x0a, 0x0b, 0x56, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16,
	0x0a, 0x06, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x22, 0x2a, 0x0a, 0x0c, 0x56, 0x6f, 0x74, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x42, 0x49, 0x5a, 0x47, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x72, 0x6f, 0x67, 0x65, 0x72, 0x69, 0x6f, 0x34, 0x31, 0x30, 0x2f, 0x63, 0x72, 0x79, 0x70,
	0x74, 0x6f, 0x76, 0x6f, 0x74, 0x65, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x70,
	0x6b, 0x67, 0x2f, 0x70, 0x62, 0x2f, 0x76, 0x6f, 0x74, 0x65, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x3b, 0x76, 0x6f, 0x74, 0x65, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
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

var file_voteservice_messages_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_voteservice_messages_proto_goTypes = []interface{}{
	(*Cryptocurrency)(nil),       // 0: voteservice.Cryptocurrency
	(*GetAllCriptoRequest)(nil),  // 1: voteservice.GetAllCriptoRequest
	(*GetAllCriptoResponse)(nil), // 2: voteservice.GetAllCriptoResponse
	(*VoteRequest)(nil),          // 3: voteservice.VoteRequest
	(*VoteResponse)(nil),         // 4: voteservice.VoteResponse
}
var file_voteservice_messages_proto_depIdxs = []int32{
	0, // 0: voteservice.GetAllCriptoResponse.cryptocurrencies:type_name -> voteservice.Cryptocurrency
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
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
			switch v := v.(*GetAllCriptoRequest); i {
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
			switch v := v.(*GetAllCriptoResponse); i {
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
		file_voteservice_messages_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
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
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_voteservice_messages_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_voteservice_messages_proto_goTypes,
		DependencyIndexes: file_voteservice_messages_proto_depIdxs,
		MessageInfos:      file_voteservice_messages_proto_msgTypes,
	}.Build()
	File_voteservice_messages_proto = out.File
	file_voteservice_messages_proto_rawDesc = nil
	file_voteservice_messages_proto_goTypes = nil
	file_voteservice_messages_proto_depIdxs = nil
}