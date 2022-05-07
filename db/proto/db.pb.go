// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: db/proto/db.proto

package proto

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

type SetEntryPrepareRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IdempotencyKey int64             `protobuf:"varint,1,opt,name=idempotencyKey,proto3" json:"idempotencyKey,omitempty"`
	Key            string            `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	Updates        map[string]string `protobuf:"bytes,3,rep,name=updates,proto3" json:"updates,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *SetEntryPrepareRequest) Reset() {
	*x = SetEntryPrepareRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_db_proto_db_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetEntryPrepareRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetEntryPrepareRequest) ProtoMessage() {}

func (x *SetEntryPrepareRequest) ProtoReflect() protoreflect.Message {
	mi := &file_db_proto_db_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetEntryPrepareRequest.ProtoReflect.Descriptor instead.
func (*SetEntryPrepareRequest) Descriptor() ([]byte, []int) {
	return file_db_proto_db_proto_rawDescGZIP(), []int{0}
}

func (x *SetEntryPrepareRequest) GetIdempotencyKey() int64 {
	if x != nil {
		return x.IdempotencyKey
	}
	return 0
}

func (x *SetEntryPrepareRequest) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *SetEntryPrepareRequest) GetUpdates() map[string]string {
	if x != nil {
		return x.Updates
	}
	return nil
}

type SetEntryPrepareResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SetEntryPrepareResponse) Reset() {
	*x = SetEntryPrepareResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_db_proto_db_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetEntryPrepareResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetEntryPrepareResponse) ProtoMessage() {}

func (x *SetEntryPrepareResponse) ProtoReflect() protoreflect.Message {
	mi := &file_db_proto_db_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetEntryPrepareResponse.ProtoReflect.Descriptor instead.
func (*SetEntryPrepareResponse) Descriptor() ([]byte, []int) {
	return file_db_proto_db_proto_rawDescGZIP(), []int{1}
}

type SetEntryCommitRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IdempotencyKey int64             `protobuf:"varint,1,opt,name=idempotencyKey,proto3" json:"idempotencyKey,omitempty"`
	Key            string            `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	Updates        map[string]string `protobuf:"bytes,3,rep,name=updates,proto3" json:"updates,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *SetEntryCommitRequest) Reset() {
	*x = SetEntryCommitRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_db_proto_db_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetEntryCommitRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetEntryCommitRequest) ProtoMessage() {}

func (x *SetEntryCommitRequest) ProtoReflect() protoreflect.Message {
	mi := &file_db_proto_db_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetEntryCommitRequest.ProtoReflect.Descriptor instead.
func (*SetEntryCommitRequest) Descriptor() ([]byte, []int) {
	return file_db_proto_db_proto_rawDescGZIP(), []int{2}
}

func (x *SetEntryCommitRequest) GetIdempotencyKey() int64 {
	if x != nil {
		return x.IdempotencyKey
	}
	return 0
}

func (x *SetEntryCommitRequest) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *SetEntryCommitRequest) GetUpdates() map[string]string {
	if x != nil {
		return x.Updates
	}
	return nil
}

type SetEntryCommitResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SetEntryCommitResponse) Reset() {
	*x = SetEntryCommitResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_db_proto_db_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetEntryCommitResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetEntryCommitResponse) ProtoMessage() {}

func (x *SetEntryCommitResponse) ProtoReflect() protoreflect.Message {
	mi := &file_db_proto_db_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetEntryCommitResponse.ProtoReflect.Descriptor instead.
func (*SetEntryCommitResponse) Descriptor() ([]byte, []int) {
	return file_db_proto_db_proto_rawDescGZIP(), []int{3}
}

type SetEntryAbortRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IdempotencyKey int64 `protobuf:"varint,1,opt,name=idempotencyKey,proto3" json:"idempotencyKey,omitempty"`
}

func (x *SetEntryAbortRequest) Reset() {
	*x = SetEntryAbortRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_db_proto_db_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetEntryAbortRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetEntryAbortRequest) ProtoMessage() {}

func (x *SetEntryAbortRequest) ProtoReflect() protoreflect.Message {
	mi := &file_db_proto_db_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetEntryAbortRequest.ProtoReflect.Descriptor instead.
func (*SetEntryAbortRequest) Descriptor() ([]byte, []int) {
	return file_db_proto_db_proto_rawDescGZIP(), []int{4}
}

func (x *SetEntryAbortRequest) GetIdempotencyKey() int64 {
	if x != nil {
		return x.IdempotencyKey
	}
	return 0
}

type SetEntryAbortResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SetEntryAbortResponse) Reset() {
	*x = SetEntryAbortResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_db_proto_db_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetEntryAbortResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetEntryAbortResponse) ProtoMessage() {}

func (x *SetEntryAbortResponse) ProtoReflect() protoreflect.Message {
	mi := &file_db_proto_db_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetEntryAbortResponse.ProtoReflect.Descriptor instead.
func (*SetEntryAbortResponse) Descriptor() ([]byte, []int) {
	return file_db_proto_db_proto_rawDescGZIP(), []int{5}
}

type GetEntryByIdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetEntryByIdRequest) Reset() {
	*x = GetEntryByIdRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_db_proto_db_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetEntryByIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetEntryByIdRequest) ProtoMessage() {}

func (x *GetEntryByIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_db_proto_db_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetEntryByIdRequest.ProtoReflect.Descriptor instead.
func (*GetEntryByIdRequest) Descriptor() ([]byte, []int) {
	return file_db_proto_db_proto_rawDescGZIP(), []int{6}
}

func (x *GetEntryByIdRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetEntryByIdResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Entry string `protobuf:"bytes,1,opt,name=entry,proto3" json:"entry,omitempty"`
}

func (x *GetEntryByIdResponse) Reset() {
	*x = GetEntryByIdResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_db_proto_db_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetEntryByIdResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetEntryByIdResponse) ProtoMessage() {}

func (x *GetEntryByIdResponse) ProtoReflect() protoreflect.Message {
	mi := &file_db_proto_db_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetEntryByIdResponse.ProtoReflect.Descriptor instead.
func (*GetEntryByIdResponse) Descriptor() ([]byte, []int) {
	return file_db_proto_db_proto_rawDescGZIP(), []int{7}
}

func (x *GetEntryByIdResponse) GetEntry() string {
	if x != nil {
		return x.Entry
	}
	return ""
}

// For retrieving the id of a row using an indexed field.
type GetEntryByIndexedFieldRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Field string `protobuf:"bytes,1,opt,name=field,proto3" json:"field,omitempty"`
	Key   string `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
}

func (x *GetEntryByIndexedFieldRequest) Reset() {
	*x = GetEntryByIndexedFieldRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_db_proto_db_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetEntryByIndexedFieldRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetEntryByIndexedFieldRequest) ProtoMessage() {}

func (x *GetEntryByIndexedFieldRequest) ProtoReflect() protoreflect.Message {
	mi := &file_db_proto_db_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetEntryByIndexedFieldRequest.ProtoReflect.Descriptor instead.
func (*GetEntryByIndexedFieldRequest) Descriptor() ([]byte, []int) {
	return file_db_proto_db_proto_rawDescGZIP(), []int{8}
}

func (x *GetEntryByIndexedFieldRequest) GetField() string {
	if x != nil {
		return x.Field
	}
	return ""
}

func (x *GetEntryByIndexedFieldRequest) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

type GetEntryByIndexedFieldResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Entry string `protobuf:"bytes,1,opt,name=entry,proto3" json:"entry,omitempty"`
}

func (x *GetEntryByIndexedFieldResponse) Reset() {
	*x = GetEntryByIndexedFieldResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_db_proto_db_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetEntryByIndexedFieldResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetEntryByIndexedFieldResponse) ProtoMessage() {}

func (x *GetEntryByIndexedFieldResponse) ProtoReflect() protoreflect.Message {
	mi := &file_db_proto_db_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetEntryByIndexedFieldResponse.ProtoReflect.Descriptor instead.
func (*GetEntryByIndexedFieldResponse) Descriptor() ([]byte, []int) {
	return file_db_proto_db_proto_rawDescGZIP(), []int{9}
}

func (x *GetEntryByIndexedFieldResponse) GetEntry() string {
	if x != nil {
		return x.Entry
	}
	return ""
}

var File_db_proto_db_proto protoreflect.FileDescriptor

var file_db_proto_db_proto_rawDesc = []byte{
	0x0a, 0x11, 0x64, 0x62, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x64, 0x62, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x02, 0x64, 0x62, 0x22, 0xd1, 0x01, 0x0a, 0x16, 0x53, 0x65, 0x74, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x50, 0x72, 0x65, 0x70, 0x61, 0x72, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x26, 0x0a, 0x0e, 0x69, 0x64, 0x65, 0x6d, 0x70, 0x6f, 0x74, 0x65, 0x6e, 0x63,
	0x79, 0x4b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0e, 0x69, 0x64, 0x65, 0x6d,
	0x70, 0x6f, 0x74, 0x65, 0x6e, 0x63, 0x79, 0x4b, 0x65, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65,
	0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x41, 0x0a, 0x07,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x27, 0x2e,
	0x64, 0x62, 0x2e, 0x53, 0x65, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x50, 0x72, 0x65, 0x70, 0x61,
	0x72, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x73, 0x1a,
	0x3a, 0x0a, 0x0c, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12,
	0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65,
	0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x19, 0x0a, 0x17, 0x53,
	0x65, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x50, 0x72, 0x65, 0x70, 0x61, 0x72, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0xcf, 0x01, 0x0a, 0x15, 0x53, 0x65, 0x74, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x26, 0x0a, 0x0e, 0x69, 0x64, 0x65, 0x6d, 0x70, 0x6f, 0x74, 0x65, 0x6e, 0x63, 0x79, 0x4b,
	0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0e, 0x69, 0x64, 0x65, 0x6d, 0x70, 0x6f,
	0x74, 0x65, 0x6e, 0x63, 0x79, 0x4b, 0x65, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x40, 0x0a, 0x07, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x64, 0x62,
	0x2e, 0x53, 0x65, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x73, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x52, 0x07, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x73, 0x1a, 0x3a, 0x0a, 0x0c,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x18, 0x0a, 0x16, 0x53, 0x65, 0x74, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x3e, 0x0a, 0x14, 0x53, 0x65, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x41, 0x62,
	0x6f, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x26, 0x0a, 0x0e, 0x69, 0x64,
	0x65, 0x6d, 0x70, 0x6f, 0x74, 0x65, 0x6e, 0x63, 0x79, 0x4b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x0e, 0x69, 0x64, 0x65, 0x6d, 0x70, 0x6f, 0x74, 0x65, 0x6e, 0x63, 0x79, 0x4b,
	0x65, 0x79, 0x22, 0x17, 0x0a, 0x15, 0x53, 0x65, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x41, 0x62,
	0x6f, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x25, 0x0a, 0x13, 0x47,
	0x65, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x22, 0x2c, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x42, 0x79,
	0x49, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6e,
	0x74, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6e, 0x74, 0x72, 0x79,
	0x22, 0x47, 0x0a, 0x1d, 0x47, 0x65, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x42, 0x79, 0x49, 0x6e,
	0x64, 0x65, 0x78, 0x65, 0x64, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x14, 0x0a, 0x05, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x22, 0x36, 0x0a, 0x1e, 0x47, 0x65, 0x74,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x42, 0x79, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x65, 0x64, 0x46, 0x69,
	0x65, 0x6c, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65,
	0x6e, 0x74, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6e, 0x74, 0x72,
	0x79, 0x32, 0xff, 0x02, 0x0a, 0x05, 0x4a, 0x4b, 0x5a, 0x44, 0x42, 0x12, 0x4a, 0x0a, 0x0f, 0x53,
	0x65, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x50, 0x72, 0x65, 0x70, 0x61, 0x72, 0x65, 0x12, 0x1a,
	0x2e, 0x64, 0x62, 0x2e, 0x53, 0x65, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x50, 0x72, 0x65, 0x70,
	0x61, 0x72, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x64, 0x62, 0x2e,
	0x53, 0x65, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x50, 0x72, 0x65, 0x70, 0x61, 0x72, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x47, 0x0a, 0x0e, 0x53, 0x65, 0x74, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x12, 0x19, 0x2e, 0x64, 0x62, 0x2e, 0x53,
	0x65, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x64, 0x62, 0x2e, 0x53, 0x65, 0x74, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x44, 0x0a, 0x0d, 0x53, 0x65, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x41, 0x62, 0x6f, 0x72,
	0x74, 0x12, 0x18, 0x2e, 0x64, 0x62, 0x2e, 0x53, 0x65, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x41,
	0x62, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x64, 0x62,
	0x2e, 0x53, 0x65, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x41, 0x62, 0x6f, 0x72, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x41, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x42, 0x79, 0x49, 0x64, 0x12, 0x17, 0x2e, 0x64, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x18, 0x2e, 0x64, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x42, 0x79, 0x49,
	0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x58, 0x0a, 0x0f, 0x47, 0x65, 0x74,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x42, 0x79, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x21, 0x2e, 0x64,
	0x62, 0x2e, 0x47, 0x65, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x42, 0x79, 0x49, 0x6e, 0x64, 0x65,
	0x78, 0x65, 0x64, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x22, 0x2e, 0x64, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x42, 0x79, 0x49,
	0x6e, 0x64, 0x65, 0x78, 0x65, 0x64, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x42, 0x10, 0x5a, 0x0e, 0x4a, 0x4b, 0x5a, 0x44, 0x42, 0x2f, 0x64, 0x62, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_db_proto_db_proto_rawDescOnce sync.Once
	file_db_proto_db_proto_rawDescData = file_db_proto_db_proto_rawDesc
)

func file_db_proto_db_proto_rawDescGZIP() []byte {
	file_db_proto_db_proto_rawDescOnce.Do(func() {
		file_db_proto_db_proto_rawDescData = protoimpl.X.CompressGZIP(file_db_proto_db_proto_rawDescData)
	})
	return file_db_proto_db_proto_rawDescData
}

var file_db_proto_db_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_db_proto_db_proto_goTypes = []interface{}{
	(*SetEntryPrepareRequest)(nil),         // 0: db.SetEntryPrepareRequest
	(*SetEntryPrepareResponse)(nil),        // 1: db.SetEntryPrepareResponse
	(*SetEntryCommitRequest)(nil),          // 2: db.SetEntryCommitRequest
	(*SetEntryCommitResponse)(nil),         // 3: db.SetEntryCommitResponse
	(*SetEntryAbortRequest)(nil),           // 4: db.SetEntryAbortRequest
	(*SetEntryAbortResponse)(nil),          // 5: db.SetEntryAbortResponse
	(*GetEntryByIdRequest)(nil),            // 6: db.GetEntryByIdRequest
	(*GetEntryByIdResponse)(nil),           // 7: db.GetEntryByIdResponse
	(*GetEntryByIndexedFieldRequest)(nil),  // 8: db.GetEntryByIndexedFieldRequest
	(*GetEntryByIndexedFieldResponse)(nil), // 9: db.GetEntryByIndexedFieldResponse
	nil,                                    // 10: db.SetEntryPrepareRequest.UpdatesEntry
	nil,                                    // 11: db.SetEntryCommitRequest.UpdatesEntry
}
var file_db_proto_db_proto_depIdxs = []int32{
	10, // 0: db.SetEntryPrepareRequest.updates:type_name -> db.SetEntryPrepareRequest.UpdatesEntry
	11, // 1: db.SetEntryCommitRequest.updates:type_name -> db.SetEntryCommitRequest.UpdatesEntry
	0,  // 2: db.JKZDB.SetEntryPrepare:input_type -> db.SetEntryPrepareRequest
	2,  // 3: db.JKZDB.SetEntryCommit:input_type -> db.SetEntryCommitRequest
	4,  // 4: db.JKZDB.SetEntryAbort:input_type -> db.SetEntryAbortRequest
	6,  // 5: db.JKZDB.GetEntryById:input_type -> db.GetEntryByIdRequest
	8,  // 6: db.JKZDB.GetEntryByField:input_type -> db.GetEntryByIndexedFieldRequest
	1,  // 7: db.JKZDB.SetEntryPrepare:output_type -> db.SetEntryPrepareResponse
	3,  // 8: db.JKZDB.SetEntryCommit:output_type -> db.SetEntryCommitResponse
	5,  // 9: db.JKZDB.SetEntryAbort:output_type -> db.SetEntryAbortResponse
	7,  // 10: db.JKZDB.GetEntryById:output_type -> db.GetEntryByIdResponse
	9,  // 11: db.JKZDB.GetEntryByField:output_type -> db.GetEntryByIndexedFieldResponse
	7,  // [7:12] is the sub-list for method output_type
	2,  // [2:7] is the sub-list for method input_type
	2,  // [2:2] is the sub-list for extension type_name
	2,  // [2:2] is the sub-list for extension extendee
	0,  // [0:2] is the sub-list for field type_name
}

func init() { file_db_proto_db_proto_init() }
func file_db_proto_db_proto_init() {
	if File_db_proto_db_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_db_proto_db_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetEntryPrepareRequest); i {
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
		file_db_proto_db_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetEntryPrepareResponse); i {
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
		file_db_proto_db_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetEntryCommitRequest); i {
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
		file_db_proto_db_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetEntryCommitResponse); i {
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
		file_db_proto_db_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetEntryAbortRequest); i {
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
		file_db_proto_db_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetEntryAbortResponse); i {
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
		file_db_proto_db_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetEntryByIdRequest); i {
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
		file_db_proto_db_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetEntryByIdResponse); i {
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
		file_db_proto_db_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetEntryByIndexedFieldRequest); i {
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
		file_db_proto_db_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetEntryByIndexedFieldResponse); i {
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
			RawDescriptor: file_db_proto_db_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_db_proto_db_proto_goTypes,
		DependencyIndexes: file_db_proto_db_proto_depIdxs,
		MessageInfos:      file_db_proto_db_proto_msgTypes,
	}.Build()
	File_db_proto_db_proto = out.File
	file_db_proto_db_proto_rawDesc = nil
	file_db_proto_db_proto_goTypes = nil
	file_db_proto_db_proto_depIdxs = nil
}
