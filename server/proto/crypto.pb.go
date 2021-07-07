// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: crypto.proto

package __

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type Crypto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name     string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Upvote   int32  `protobuf:"varint,3,opt,name=upvote,proto3" json:"upvote,omitempty"`
	Downvote int32  `protobuf:"varint,4,opt,name=downvote,proto3" json:"downvote,omitempty"`
}

func (x *Crypto) Reset() {
	*x = Crypto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_crypto_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Crypto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Crypto) ProtoMessage() {}

func (x *Crypto) ProtoReflect() protoreflect.Message {
	mi := &file_crypto_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Crypto.ProtoReflect.Descriptor instead.
func (*Crypto) Descriptor() ([]byte, []int) {
	return file_crypto_proto_rawDescGZIP(), []int{0}
}

func (x *Crypto) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Crypto) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Crypto) GetUpvote() int32 {
	if x != nil {
		return x.Upvote
	}
	return 0
}

func (x *Crypto) GetDownvote() int32 {
	if x != nil {
		return x.Downvote
	}
	return 0
}

type InsertCryptoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Crypto *Crypto `protobuf:"bytes,1,opt,name=crypto,proto3" json:"crypto,omitempty"`
}

func (x *InsertCryptoRequest) Reset() {
	*x = InsertCryptoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_crypto_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InsertCryptoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InsertCryptoRequest) ProtoMessage() {}

func (x *InsertCryptoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_crypto_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InsertCryptoRequest.ProtoReflect.Descriptor instead.
func (*InsertCryptoRequest) Descriptor() ([]byte, []int) {
	return file_crypto_proto_rawDescGZIP(), []int{1}
}

func (x *InsertCryptoRequest) GetCrypto() *Crypto {
	if x != nil {
		return x.Crypto
	}
	return nil
}

type UpdateCryptoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Upvote   int32 `protobuf:"varint,1,opt,name=upvote,proto3" json:"upvote,omitempty"`
	Downvote int32 `protobuf:"varint,2,opt,name=downvote,proto3" json:"downvote,omitempty"`
}

func (x *UpdateCryptoRequest) Reset() {
	*x = UpdateCryptoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_crypto_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateCryptoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateCryptoRequest) ProtoMessage() {}

func (x *UpdateCryptoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_crypto_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateCryptoRequest.ProtoReflect.Descriptor instead.
func (*UpdateCryptoRequest) Descriptor() ([]byte, []int) {
	return file_crypto_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateCryptoRequest) GetUpvote() int32 {
	if x != nil {
		return x.Upvote
	}
	return 0
}

func (x *UpdateCryptoRequest) GetDownvote() int32 {
	if x != nil {
		return x.Downvote
	}
	return 0
}

type UpdateCryptoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Crypto *Crypto `protobuf:"bytes,1,opt,name=crypto,proto3" json:"crypto,omitempty"`
}

func (x *UpdateCryptoResponse) Reset() {
	*x = UpdateCryptoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_crypto_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateCryptoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateCryptoResponse) ProtoMessage() {}

func (x *UpdateCryptoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_crypto_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateCryptoResponse.ProtoReflect.Descriptor instead.
func (*UpdateCryptoResponse) Descriptor() ([]byte, []int) {
	return file_crypto_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateCryptoResponse) GetCrypto() *Crypto {
	if x != nil {
		return x.Crypto
	}
	return nil
}

type CryptoID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *CryptoID) Reset() {
	*x = CryptoID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_crypto_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CryptoID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CryptoID) ProtoMessage() {}

func (x *CryptoID) ProtoReflect() protoreflect.Message {
	mi := &file_crypto_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CryptoID.ProtoReflect.Descriptor instead.
func (*CryptoID) Descriptor() ([]byte, []int) {
	return file_crypto_proto_rawDescGZIP(), []int{4}
}

func (x *CryptoID) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type ListCryptoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListCryptoRequest) Reset() {
	*x = ListCryptoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_crypto_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListCryptoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListCryptoRequest) ProtoMessage() {}

func (x *ListCryptoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_crypto_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListCryptoRequest.ProtoReflect.Descriptor instead.
func (*ListCryptoRequest) Descriptor() ([]byte, []int) {
	return file_crypto_proto_rawDescGZIP(), []int{5}
}

type ListCryptoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Crypto []*Crypto `protobuf:"bytes,1,rep,name=crypto,proto3" json:"crypto,omitempty"`
}

func (x *ListCryptoResponse) Reset() {
	*x = ListCryptoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_crypto_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListCryptoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListCryptoResponse) ProtoMessage() {}

func (x *ListCryptoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_crypto_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListCryptoResponse.ProtoReflect.Descriptor instead.
func (*ListCryptoResponse) Descriptor() ([]byte, []int) {
	return file_crypto_proto_rawDescGZIP(), []int{6}
}

func (x *ListCryptoResponse) GetCrypto() []*Crypto {
	if x != nil {
		return x.Crypto
	}
	return nil
}

var File_crypto_proto protoreflect.FileDescriptor

var file_crypto_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06,
	0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x22, 0x60, 0x0a, 0x06, 0x43, 0x72, 0x79, 0x70, 0x74, 0x6f,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x70, 0x76, 0x6f, 0x74, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x75, 0x70, 0x76, 0x6f, 0x74, 0x65, 0x12, 0x1a, 0x0a, 0x08,
	0x64, 0x6f, 0x77, 0x6e, 0x76, 0x6f, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08,
	0x64, 0x6f, 0x77, 0x6e, 0x76, 0x6f, 0x74, 0x65, 0x22, 0x3d, 0x0a, 0x13, 0x49, 0x6e, 0x73, 0x65,
	0x72, 0x74, 0x43, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x26, 0x0a, 0x06, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0e, 0x2e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x2e, 0x43, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x52,
	0x06, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x22, 0x49, 0x0a, 0x13, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x43, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16,
	0x0a, 0x06, 0x75, 0x70, 0x76, 0x6f, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06,
	0x75, 0x70, 0x76, 0x6f, 0x74, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x6f, 0x77, 0x6e, 0x76, 0x6f,
	0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x64, 0x6f, 0x77, 0x6e, 0x76, 0x6f,
	0x74, 0x65, 0x22, 0x3e, 0x0a, 0x14, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x72, 0x79, 0x70,
	0x74, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x26, 0x0a, 0x06, 0x63, 0x72,
	0x79, 0x70, 0x74, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x63, 0x72, 0x79,
	0x70, 0x74, 0x6f, 0x2e, 0x43, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x52, 0x06, 0x63, 0x72, 0x79, 0x70,
	0x74, 0x6f, 0x22, 0x1a, 0x0a, 0x08, 0x43, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x49, 0x44, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x22, 0x13,
	0x0a, 0x11, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x22, 0x3c, 0x0a, 0x12, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x72, 0x79, 0x70, 0x74,
	0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x26, 0x0a, 0x06, 0x63, 0x72, 0x79,
	0x70, 0x74, 0x6f, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x63, 0x72, 0x79, 0x70,
	0x74, 0x6f, 0x2e, 0x43, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x52, 0x06, 0x63, 0x72, 0x79, 0x70, 0x74,
	0x6f, 0x32, 0xc2, 0x02, 0x0a, 0x0d, 0x43, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x3d, 0x0a, 0x0c, 0x49, 0x6e, 0x73, 0x65, 0x72, 0x74, 0x43, 0x72, 0x79,
	0x70, 0x74, 0x6f, 0x12, 0x1b, 0x2e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x2e, 0x49, 0x6e, 0x73,
	0x65, 0x72, 0x74, 0x43, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x10, 0x2e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x2e, 0x43, 0x72, 0x79, 0x70, 0x74, 0x6f,
	0x49, 0x44, 0x12, 0x2e, 0x0a, 0x0a, 0x52, 0x65, 0x61, 0x64, 0x43, 0x72, 0x79, 0x70, 0x74, 0x6f,
	0x12, 0x10, 0x2e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x2e, 0x43, 0x72, 0x79, 0x70, 0x74, 0x6f,
	0x49, 0x44, 0x1a, 0x0e, 0x2e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x2e, 0x43, 0x72, 0x79, 0x70,
	0x74, 0x6f, 0x12, 0x49, 0x0a, 0x0c, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x72, 0x79, 0x70,
	0x74, 0x6f, 0x12, 0x1b, 0x2e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x2e, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x43, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1c, 0x2e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43,
	0x72, 0x79, 0x70, 0x74, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x32, 0x0a,
	0x0c, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x12, 0x10, 0x2e,
	0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x2e, 0x43, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x49, 0x44, 0x1a,
	0x10, 0x2e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x2e, 0x43, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x49,
	0x44, 0x12, 0x43, 0x0a, 0x0a, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x12,
	0x19, 0x2e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x72, 0x79,
	0x70, 0x74, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x63, 0x72, 0x79,
	0x70, 0x74, 0x6f, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x04, 0x5a, 0x02, 0x2e, 0x2f, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_crypto_proto_rawDescOnce sync.Once
	file_crypto_proto_rawDescData = file_crypto_proto_rawDesc
)

func file_crypto_proto_rawDescGZIP() []byte {
	file_crypto_proto_rawDescOnce.Do(func() {
		file_crypto_proto_rawDescData = protoimpl.X.CompressGZIP(file_crypto_proto_rawDescData)
	})
	return file_crypto_proto_rawDescData
}

var file_crypto_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_crypto_proto_goTypes = []interface{}{
	(*Crypto)(nil),               // 0: crypto.Crypto
	(*InsertCryptoRequest)(nil),  // 1: crypto.InsertCryptoRequest
	(*UpdateCryptoRequest)(nil),  // 2: crypto.UpdateCryptoRequest
	(*UpdateCryptoResponse)(nil), // 3: crypto.UpdateCryptoResponse
	(*CryptoID)(nil),             // 4: crypto.CryptoID
	(*ListCryptoRequest)(nil),    // 5: crypto.ListCryptoRequest
	(*ListCryptoResponse)(nil),   // 6: crypto.ListCryptoResponse
}
var file_crypto_proto_depIdxs = []int32{
	0, // 0: crypto.InsertCryptoRequest.crypto:type_name -> crypto.Crypto
	0, // 1: crypto.UpdateCryptoResponse.crypto:type_name -> crypto.Crypto
	0, // 2: crypto.ListCryptoResponse.crypto:type_name -> crypto.Crypto
	1, // 3: crypto.CryptoService.InsertCrypto:input_type -> crypto.InsertCryptoRequest
	4, // 4: crypto.CryptoService.ReadCrypto:input_type -> crypto.CryptoID
	2, // 5: crypto.CryptoService.UpdateCrypto:input_type -> crypto.UpdateCryptoRequest
	4, // 6: crypto.CryptoService.DeleteCrypto:input_type -> crypto.CryptoID
	5, // 7: crypto.CryptoService.ListCrypto:input_type -> crypto.ListCryptoRequest
	4, // 8: crypto.CryptoService.InsertCrypto:output_type -> crypto.CryptoID
	0, // 9: crypto.CryptoService.ReadCrypto:output_type -> crypto.Crypto
	3, // 10: crypto.CryptoService.UpdateCrypto:output_type -> crypto.UpdateCryptoResponse
	4, // 11: crypto.CryptoService.DeleteCrypto:output_type -> crypto.CryptoID
	6, // 12: crypto.CryptoService.ListCrypto:output_type -> crypto.ListCryptoResponse
	8, // [8:13] is the sub-list for method output_type
	3, // [3:8] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_crypto_proto_init() }
func file_crypto_proto_init() {
	if File_crypto_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_crypto_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Crypto); i {
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
		file_crypto_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InsertCryptoRequest); i {
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
		file_crypto_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateCryptoRequest); i {
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
		file_crypto_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateCryptoResponse); i {
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
		file_crypto_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CryptoID); i {
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
		file_crypto_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListCryptoRequest); i {
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
		file_crypto_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListCryptoResponse); i {
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
			RawDescriptor: file_crypto_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_crypto_proto_goTypes,
		DependencyIndexes: file_crypto_proto_depIdxs,
		MessageInfos:      file_crypto_proto_msgTypes,
	}.Build()
	File_crypto_proto = out.File
	file_crypto_proto_rawDesc = nil
	file_crypto_proto_goTypes = nil
	file_crypto_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// CryptoServiceClient is the client API for CryptoService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CryptoServiceClient interface {
	InsertCrypto(ctx context.Context, in *InsertCryptoRequest, opts ...grpc.CallOption) (*CryptoID, error)
	ReadCrypto(ctx context.Context, in *CryptoID, opts ...grpc.CallOption) (*Crypto, error)
	UpdateCrypto(ctx context.Context, in *UpdateCryptoRequest, opts ...grpc.CallOption) (*UpdateCryptoResponse, error)
	DeleteCrypto(ctx context.Context, in *CryptoID, opts ...grpc.CallOption) (*CryptoID, error)
	ListCrypto(ctx context.Context, in *ListCryptoRequest, opts ...grpc.CallOption) (*ListCryptoResponse, error)
}

type cryptoServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCryptoServiceClient(cc grpc.ClientConnInterface) CryptoServiceClient {
	return &cryptoServiceClient{cc}
}

func (c *cryptoServiceClient) InsertCrypto(ctx context.Context, in *InsertCryptoRequest, opts ...grpc.CallOption) (*CryptoID, error) {
	out := new(CryptoID)
	err := c.cc.Invoke(ctx, "/crypto.CryptoService/InsertCrypto", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cryptoServiceClient) ReadCrypto(ctx context.Context, in *CryptoID, opts ...grpc.CallOption) (*Crypto, error) {
	out := new(Crypto)
	err := c.cc.Invoke(ctx, "/crypto.CryptoService/ReadCrypto", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cryptoServiceClient) UpdateCrypto(ctx context.Context, in *UpdateCryptoRequest, opts ...grpc.CallOption) (*UpdateCryptoResponse, error) {
	out := new(UpdateCryptoResponse)
	err := c.cc.Invoke(ctx, "/crypto.CryptoService/UpdateCrypto", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cryptoServiceClient) DeleteCrypto(ctx context.Context, in *CryptoID, opts ...grpc.CallOption) (*CryptoID, error) {
	out := new(CryptoID)
	err := c.cc.Invoke(ctx, "/crypto.CryptoService/DeleteCrypto", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cryptoServiceClient) ListCrypto(ctx context.Context, in *ListCryptoRequest, opts ...grpc.CallOption) (*ListCryptoResponse, error) {
	out := new(ListCryptoResponse)
	err := c.cc.Invoke(ctx, "/crypto.CryptoService/ListCrypto", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CryptoServiceServer is the server API for CryptoService service.
type CryptoServiceServer interface {
	InsertCrypto(context.Context, *InsertCryptoRequest) (*CryptoID, error)
	ReadCrypto(context.Context, *CryptoID) (*Crypto, error)
	UpdateCrypto(context.Context, *UpdateCryptoRequest) (*UpdateCryptoResponse, error)
	DeleteCrypto(context.Context, *CryptoID) (*CryptoID, error)
	ListCrypto(context.Context, *ListCryptoRequest) (*ListCryptoResponse, error)
}

// UnimplementedCryptoServiceServer can be embedded to have forward compatible implementations.
type UnimplementedCryptoServiceServer struct {
}

func (*UnimplementedCryptoServiceServer) InsertCrypto(context.Context, *InsertCryptoRequest) (*CryptoID, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InsertCrypto not implemented")
}
func (*UnimplementedCryptoServiceServer) ReadCrypto(context.Context, *CryptoID) (*Crypto, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadCrypto not implemented")
}
func (*UnimplementedCryptoServiceServer) UpdateCrypto(context.Context, *UpdateCryptoRequest) (*UpdateCryptoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCrypto not implemented")
}
func (*UnimplementedCryptoServiceServer) DeleteCrypto(context.Context, *CryptoID) (*CryptoID, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCrypto not implemented")
}
func (*UnimplementedCryptoServiceServer) ListCrypto(context.Context, *ListCryptoRequest) (*ListCryptoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListCrypto not implemented")
}

func RegisterCryptoServiceServer(s *grpc.Server, srv CryptoServiceServer) {
	s.RegisterService(&_CryptoService_serviceDesc, srv)
}

func _CryptoService_InsertCrypto_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InsertCryptoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CryptoServiceServer).InsertCrypto(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/crypto.CryptoService/InsertCrypto",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CryptoServiceServer).InsertCrypto(ctx, req.(*InsertCryptoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CryptoService_ReadCrypto_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CryptoID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CryptoServiceServer).ReadCrypto(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/crypto.CryptoService/ReadCrypto",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CryptoServiceServer).ReadCrypto(ctx, req.(*CryptoID))
	}
	return interceptor(ctx, in, info, handler)
}

func _CryptoService_UpdateCrypto_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateCryptoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CryptoServiceServer).UpdateCrypto(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/crypto.CryptoService/UpdateCrypto",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CryptoServiceServer).UpdateCrypto(ctx, req.(*UpdateCryptoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CryptoService_DeleteCrypto_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CryptoID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CryptoServiceServer).DeleteCrypto(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/crypto.CryptoService/DeleteCrypto",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CryptoServiceServer).DeleteCrypto(ctx, req.(*CryptoID))
	}
	return interceptor(ctx, in, info, handler)
}

func _CryptoService_ListCrypto_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListCryptoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CryptoServiceServer).ListCrypto(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/crypto.CryptoService/ListCrypto",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CryptoServiceServer).ListCrypto(ctx, req.(*ListCryptoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CryptoService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "crypto.CryptoService",
	HandlerType: (*CryptoServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "InsertCrypto",
			Handler:    _CryptoService_InsertCrypto_Handler,
		},
		{
			MethodName: "ReadCrypto",
			Handler:    _CryptoService_ReadCrypto_Handler,
		},
		{
			MethodName: "UpdateCrypto",
			Handler:    _CryptoService_UpdateCrypto_Handler,
		},
		{
			MethodName: "DeleteCrypto",
			Handler:    _CryptoService_DeleteCrypto_Handler,
		},
		{
			MethodName: "ListCrypto",
			Handler:    _CryptoService_ListCrypto_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "crypto.proto",
}