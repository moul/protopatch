// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: tests/message/message_renames.proto

package message

import (
	_ "github.com/alta/protopatch/patch/gopb"
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

type Frank struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Frank) Reset() {
	*x = Frank{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tests_message_message_renames_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Frank) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Frank) ProtoMessage() {}

func (x *Frank) ProtoReflect() protoreflect.Message {
	mi := &file_tests_message_message_renames_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Francis.ProtoReflect.Descriptor instead.
func (*Frank) Descriptor() ([]byte, []int) {
	return file_tests_message_message_renames_proto_rawDescGZIP(), []int{0}
}

type RenamedOneofMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Contents:
	//	*OriginalOneofMessage_Id
	//	*OriginalOneofMessage_Name
	Contents isRenamedOneofMessage_Contents `protobuf_oneof:"contents"`
}

func (x *RenamedOneofMessage) Reset() {
	*x = RenamedOneofMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tests_message_message_renames_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RenamedOneofMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RenamedOneofMessage) ProtoMessage() {}

func (x *RenamedOneofMessage) ProtoReflect() protoreflect.Message {
	mi := &file_tests_message_message_renames_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OriginalOneofMessage.ProtoReflect.Descriptor instead.
func (*RenamedOneofMessage) Descriptor() ([]byte, []int) {
	return file_tests_message_message_renames_proto_rawDescGZIP(), []int{1}
}

func (m *RenamedOneofMessage) GetContents() isRenamedOneofMessage_Contents {
	if m != nil {
		return m.Contents
	}
	return nil
}

func (x *RenamedOneofMessage) GetId() int32 {
	if x, ok := x.GetContents().(*RenamedOneofMessage_Id); ok {
		return x.Id
	}
	return 0
}

func (x *RenamedOneofMessage) GetName() string {
	if x, ok := x.GetContents().(*RenamedOneofMessage_Name); ok {
		return x.Name
	}
	return ""
}

type isRenamedOneofMessage_Contents interface {
	isRenamedOneofMessage_Contents()
}

type RenamedOneofMessage_Id struct {
	Id int32 `protobuf:"varint,1,opt,name=id,proto3,oneof"`
}

type RenamedOneofMessage_Name struct {
	Name string `protobuf:"bytes,2,opt,name=name,proto3,oneof"`
}

func (*RenamedOneofMessage_Id) isRenamedOneofMessage_Contents() {}

func (*RenamedOneofMessage_Name) isRenamedOneofMessage_Contents() {}

type RenamedOuterMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Inner *RenamedOuterMessage_InnerMessage `protobuf:"bytes,1,opt,name=inner,proto3" json:"inner,omitempty"`
}

func (x *RenamedOuterMessage) Reset() {
	*x = RenamedOuterMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tests_message_message_renames_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RenamedOuterMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RenamedOuterMessage) ProtoMessage() {}

func (x *RenamedOuterMessage) ProtoReflect() protoreflect.Message {
	mi := &file_tests_message_message_renames_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OriginalOuterMessage.ProtoReflect.Descriptor instead.
func (*RenamedOuterMessage) Descriptor() ([]byte, []int) {
	return file_tests_message_message_renames_proto_rawDescGZIP(), []int{2}
}

func (x *RenamedOuterMessage) GetInner() *RenamedOuterMessage_InnerMessage {
	if x != nil {
		return x.Inner
	}
	return nil
}

type OuterMessageWithRenamedInnerMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Inner *RenamedInnerMessage `protobuf:"bytes,1,opt,name=inner,proto3" json:"inner,omitempty"`
}

func (x *OuterMessageWithRenamedInnerMessage) Reset() {
	*x = OuterMessageWithRenamedInnerMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tests_message_message_renames_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OuterMessageWithRenamedInnerMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OuterMessageWithRenamedInnerMessage) ProtoMessage() {}

func (x *OuterMessageWithRenamedInnerMessage) ProtoReflect() protoreflect.Message {
	mi := &file_tests_message_message_renames_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OuterMessageWithRenamedInnerMessage.ProtoReflect.Descriptor instead.
func (*OuterMessageWithRenamedInnerMessage) Descriptor() ([]byte, []int) {
	return file_tests_message_message_renames_proto_rawDescGZIP(), []int{3}
}

func (x *OuterMessageWithRenamedInnerMessage) GetInner() *RenamedInnerMessage {
	if x != nil {
		return x.Inner
	}
	return nil
}

type MessageWithRenamedField struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *MessageWithRenamedField) Reset() {
	*x = MessageWithRenamedField{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tests_message_message_renames_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MessageWithRenamedField) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageWithRenamedField) ProtoMessage() {}

func (x *MessageWithRenamedField) ProtoReflect() protoreflect.Message {
	mi := &file_tests_message_message_renames_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageWithRenamedField.ProtoReflect.Descriptor instead.
func (*MessageWithRenamedField) Descriptor() ([]byte, []int) {
	return file_tests_message_message_renames_proto_rawDescGZIP(), []int{4}
}

func (x *MessageWithRenamedField) GetID() int32 {
	if x != nil {
		return x.ID
	}
	return 0
}

type MessageWithEmbeddedField struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	*Embedded `protobuf:"bytes,5,opt,name=embedded_message,json=embeddedMessage,proto3" json:"embedded_message,omitempty"`
}

func (x *MessageWithEmbeddedField) Reset() {
	*x = MessageWithEmbeddedField{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tests_message_message_renames_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MessageWithEmbeddedField) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageWithEmbeddedField) ProtoMessage() {}

func (x *MessageWithEmbeddedField) ProtoReflect() protoreflect.Message {
	mi := &file_tests_message_message_renames_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageWithEmbeddedField.ProtoReflect.Descriptor instead.
func (*MessageWithEmbeddedField) Descriptor() ([]byte, []int) {
	return file_tests_message_message_renames_proto_rawDescGZIP(), []int{5}
}

func (x *MessageWithEmbeddedField) GetEmbedded() *Embedded {
	if x != nil {
		return x.Embedded
	}
	return nil
}

type Embedded struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *Embedded) Reset() {
	*x = Embedded{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tests_message_message_renames_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Embedded) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Embedded) ProtoMessage() {}

func (x *Embedded) ProtoReflect() protoreflect.Message {
	mi := &file_tests_message_message_renames_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Embedded.ProtoReflect.Descriptor instead.
func (*Embedded) Descriptor() ([]byte, []int) {
	return file_tests_message_message_renames_proto_rawDescGZIP(), []int{6}
}

func (x *Embedded) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type MessageWithOptionals struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OptionalString *string  `protobuf:"bytes,1,opt,name=string_field,json=stringField,proto3,oneof" json:"string_field,omitempty"`
	OptionalInt32  *int32   `protobuf:"varint,2,opt,name=int32_field,json=int32Field,proto3,oneof" json:"int32_field,omitempty"`
	OptionalInt64  *int64   `protobuf:"varint,3,opt,name=int64_field,json=int64Field,proto3,oneof" json:"int64_field,omitempty"`
	OptionalFloat  *float32 `protobuf:"fixed32,4,opt,name=float_field,json=floatField,proto3,oneof" json:"float_field,omitempty"`
	OptionalDouble *float64 `protobuf:"fixed64,5,opt,name=double_field,json=doubleField,proto3,oneof" json:"double_field,omitempty"`
	OptionalUInt32 *uint32  `protobuf:"varint,6,opt,name=uint32_field,json=uint32Field,proto3,oneof" json:"uint32_field,omitempty"`
	OptionalUInt64 *uint64  `protobuf:"varint,7,opt,name=uint64_field,json=uint64Field,proto3,oneof" json:"uint64_field,omitempty"`
	OptionalBool   *bool    `protobuf:"varint,8,opt,name=bool_field,json=boolField,proto3,oneof" json:"bool_field,omitempty"`
}

func (x *MessageWithOptionals) Reset() {
	*x = MessageWithOptionals{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tests_message_message_renames_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MessageWithOptionals) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageWithOptionals) ProtoMessage() {}

func (x *MessageWithOptionals) ProtoReflect() protoreflect.Message {
	mi := &file_tests_message_message_renames_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageWithOptionals.ProtoReflect.Descriptor instead.
func (*MessageWithOptionals) Descriptor() ([]byte, []int) {
	return file_tests_message_message_renames_proto_rawDescGZIP(), []int{7}
}

func (x *MessageWithOptionals) GetOptionalString() string {
	if x != nil && x.OptionalString != nil {
		return *x.OptionalString
	}
	return ""
}

func (x *MessageWithOptionals) GetOptionalInt32() int32 {
	if x != nil && x.OptionalInt32 != nil {
		return *x.OptionalInt32
	}
	return 0
}

func (x *MessageWithOptionals) GetOptionalInt64() int64 {
	if x != nil && x.OptionalInt64 != nil {
		return *x.OptionalInt64
	}
	return 0
}

func (x *MessageWithOptionals) GetOptionalFloat() float32 {
	if x != nil && x.OptionalFloat != nil {
		return *x.OptionalFloat
	}
	return 0
}

func (x *MessageWithOptionals) GetOptionalDouble() float64 {
	if x != nil && x.OptionalDouble != nil {
		return *x.OptionalDouble
	}
	return 0
}

func (x *MessageWithOptionals) GetOptionalUInt32() uint32 {
	if x != nil && x.OptionalUInt32 != nil {
		return *x.OptionalUInt32
	}
	return 0
}

func (x *MessageWithOptionals) GetOptionalUInt64() uint64 {
	if x != nil && x.OptionalUInt64 != nil {
		return *x.OptionalUInt64
	}
	return 0
}

func (x *MessageWithOptionals) GetOptionalBool() bool {
	if x != nil && x.OptionalBool != nil {
		return *x.OptionalBool
	}
	return false
}

type RenamedOuterMessage_InnerMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *RenamedOuterMessage_InnerMessage) Reset() {
	*x = RenamedOuterMessage_InnerMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tests_message_message_renames_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RenamedOuterMessage_InnerMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RenamedOuterMessage_InnerMessage) ProtoMessage() {}

func (x *RenamedOuterMessage_InnerMessage) ProtoReflect() protoreflect.Message {
	mi := &file_tests_message_message_renames_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OriginalOuterMessage_InnerMessage.ProtoReflect.Descriptor instead.
func (*RenamedOuterMessage_InnerMessage) Descriptor() ([]byte, []int) {
	return file_tests_message_message_renames_proto_rawDescGZIP(), []int{2, 0}
}

type RenamedInnerMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *RenamedInnerMessage) Reset() {
	*x = RenamedInnerMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tests_message_message_renames_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RenamedInnerMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RenamedInnerMessage) ProtoMessage() {}

func (x *RenamedInnerMessage) ProtoReflect() protoreflect.Message {
	mi := &file_tests_message_message_renames_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OuterMessageWithRenamedInnerMessage_InnerMessage.ProtoReflect.Descriptor instead.
func (*RenamedInnerMessage) Descriptor() ([]byte, []int) {
	return file_tests_message_message_renames_proto_rawDescGZIP(), []int{3, 0}
}

var File_tests_message_message_renames_proto protoreflect.FileDescriptor

var file_tests_message_message_renames_proto_rawDesc = []byte{
	0x0a, 0x23, 0x74, 0x65, 0x73, 0x74, 0x73, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2f,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x72, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x74, 0x65, 0x73, 0x74, 0x73, 0x2e, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x1a, 0x0e, 0x70, 0x61, 0x74, 0x63, 0x68, 0x2f, 0x67, 0x6f, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x16, 0x0a, 0x07, 0x46, 0x72, 0x61, 0x6e, 0x63, 0x69, 0x73, 0x3a,
	0x0b, 0xca, 0xb5, 0x03, 0x07, 0x0a, 0x05, 0x46, 0x72, 0x61, 0x6e, 0x6b, 0x22, 0x65, 0x0a, 0x14,
	0x4f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x4f, 0x6e, 0x65, 0x6f, 0x66, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x12, 0x10, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x48, 0x00, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x3a, 0x19, 0xca, 0xb5,
	0x03, 0x15, 0x0a, 0x13, 0x52, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x64, 0x4f, 0x6e, 0x65, 0x6f, 0x66,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x42, 0x0a, 0x0a, 0x08, 0x63, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x73, 0x22, 0x89, 0x01, 0x0a, 0x14, 0x4f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c,
	0x4f, 0x75, 0x74, 0x65, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x46, 0x0a, 0x05,
	0x69, 0x6e, 0x6e, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x30, 0x2e, 0x74, 0x65,
	0x73, 0x74, 0x73, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x4f, 0x72, 0x69, 0x67,
	0x69, 0x6e, 0x61, 0x6c, 0x4f, 0x75, 0x74, 0x65, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x2e, 0x49, 0x6e, 0x6e, 0x65, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x05, 0x69,
	0x6e, 0x6e, 0x65, 0x72, 0x1a, 0x0e, 0x0a, 0x0c, 0x49, 0x6e, 0x6e, 0x65, 0x72, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x3a, 0x19, 0xca, 0xb5, 0x03, 0x15, 0x0a, 0x13, 0x52, 0x65, 0x6e, 0x61,
	0x6d, 0x65, 0x64, 0x4f, 0x75, 0x74, 0x65, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22,
	0xa7, 0x01, 0x0a, 0x23, 0x4f, 0x75, 0x74, 0x65, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x57, 0x69, 0x74, 0x68, 0x52, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x64, 0x49, 0x6e, 0x6e, 0x65, 0x72,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x55, 0x0a, 0x05, 0x69, 0x6e, 0x6e, 0x65, 0x72,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x3f, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x73, 0x2e, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x4f, 0x75, 0x74, 0x65, 0x72, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x57, 0x69, 0x74, 0x68, 0x52, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x64, 0x49, 0x6e,
	0x6e, 0x65, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x49, 0x6e, 0x6e, 0x65, 0x72,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x05, 0x69, 0x6e, 0x6e, 0x65, 0x72, 0x1a, 0x29,
	0x0a, 0x0c, 0x49, 0x6e, 0x6e, 0x65, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x3a, 0x19,
	0xca, 0xb5, 0x03, 0x15, 0x0a, 0x13, 0x52, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x64, 0x49, 0x6e, 0x6e,
	0x65, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x33, 0x0a, 0x17, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x57, 0x69, 0x74, 0x68, 0x52, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x64, 0x46,
	0x69, 0x65, 0x6c, 0x64, 0x12, 0x18, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x42, 0x08, 0xca, 0xb5, 0x03, 0x04, 0x0a, 0x02, 0x49, 0x44, 0x52, 0x02, 0x69, 0x64, 0x22, 0x66,
	0x0a, 0x18, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x57, 0x69, 0x74, 0x68, 0x45, 0x6d, 0x62,
	0x65, 0x64, 0x64, 0x65, 0x64, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x4a, 0x0a, 0x10, 0x65, 0x6d,
	0x62, 0x65, 0x64, 0x64, 0x65, 0x64, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x73, 0x2e, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x2e, 0x45, 0x6d, 0x62, 0x65, 0x64, 0x64, 0x65, 0x64, 0x42, 0x06, 0xca,
	0xb5, 0x03, 0x02, 0x10, 0x01, 0x52, 0x0f, 0x65, 0x6d, 0x62, 0x65, 0x64, 0x64, 0x65, 0x64, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x24, 0x0a, 0x08, 0x45, 0x6d, 0x62, 0x65, 0x64, 0x64,
	0x65, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0xfa, 0x04, 0x0a,
	0x14, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x57, 0x69, 0x74, 0x68, 0x4f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x61, 0x6c, 0x73, 0x12, 0x3c, 0x0a, 0x0c, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x5f,
	0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x14, 0xca, 0xb5, 0x03,
	0x10, 0x0a, 0x0e, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x53, 0x74, 0x72, 0x69, 0x6e,
	0x67, 0x48, 0x00, 0x52, 0x0b, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x46, 0x69, 0x65, 0x6c, 0x64,
	0x88, 0x01, 0x01, 0x12, 0x39, 0x0a, 0x0b, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x5f, 0x66, 0x69, 0x65,
	0x6c, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x42, 0x13, 0xca, 0xb5, 0x03, 0x0f, 0x0a, 0x0d,
	0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x48, 0x01, 0x52,
	0x0a, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x88, 0x01, 0x01, 0x12, 0x39,
	0x0a, 0x0b, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x03, 0x42, 0x13, 0xca, 0xb5, 0x03, 0x0f, 0x0a, 0x0d, 0x4f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x61, 0x6c, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x48, 0x02, 0x52, 0x0a, 0x69, 0x6e, 0x74, 0x36,
	0x34, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x88, 0x01, 0x01, 0x12, 0x39, 0x0a, 0x0b, 0x66, 0x6c, 0x6f,
	0x61, 0x74, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x02, 0x42, 0x13,
	0xca, 0xb5, 0x03, 0x0f, 0x0a, 0x0d, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x46, 0x6c,
	0x6f, 0x61, 0x74, 0x48, 0x03, 0x52, 0x0a, 0x66, 0x6c, 0x6f, 0x61, 0x74, 0x46, 0x69, 0x65, 0x6c,
	0x64, 0x88, 0x01, 0x01, 0x12, 0x3c, 0x0a, 0x0c, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x5f, 0x66,
	0x69, 0x65, 0x6c, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x01, 0x42, 0x14, 0xca, 0xb5, 0x03, 0x10,
	0x0a, 0x0e, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x44, 0x6f, 0x75, 0x62, 0x6c, 0x65,
	0x48, 0x04, 0x52, 0x0b, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x88,
	0x01, 0x01, 0x12, 0x3c, 0x0a, 0x0c, 0x75, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x5f, 0x66, 0x69, 0x65,
	0x6c, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x42, 0x14, 0xca, 0xb5, 0x03, 0x10, 0x0a, 0x0e,
	0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x55, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x48, 0x05,
	0x52, 0x0b, 0x75, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x88, 0x01, 0x01,
	0x12, 0x3c, 0x0a, 0x0c, 0x75, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x04, 0x42, 0x14, 0xca, 0xb5, 0x03, 0x10, 0x0a, 0x0e, 0x4f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x55, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x48, 0x06, 0x52, 0x0b,
	0x75, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x88, 0x01, 0x01, 0x12, 0x36,
	0x0a, 0x0a, 0x62, 0x6f, 0x6f, 0x6c, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x08, 0x42, 0x12, 0xca, 0xb5, 0x03, 0x0e, 0x0a, 0x0c, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x61, 0x6c, 0x42, 0x6f, 0x6f, 0x6c, 0x48, 0x07, 0x52, 0x09, 0x62, 0x6f, 0x6f, 0x6c, 0x46, 0x69,
	0x65, 0x6c, 0x64, 0x88, 0x01, 0x01, 0x42, 0x0f, 0x0a, 0x0d, 0x5f, 0x73, 0x74, 0x72, 0x69, 0x6e,
	0x67, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x69, 0x6e, 0x74, 0x33,
	0x32, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x69, 0x6e, 0x74, 0x36,
	0x34, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x66, 0x6c, 0x6f, 0x61,
	0x74, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x42, 0x0f, 0x0a, 0x0d, 0x5f, 0x64, 0x6f, 0x75, 0x62,
	0x6c, 0x65, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x42, 0x0f, 0x0a, 0x0d, 0x5f, 0x75, 0x69, 0x6e,
	0x74, 0x33, 0x32, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x42, 0x0f, 0x0a, 0x0d, 0x5f, 0x75, 0x69,
	0x6e, 0x74, 0x36, 0x34, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x62,
	0x6f, 0x6f, 0x6c, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x42, 0x2a, 0x5a, 0x28, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x6c, 0x74, 0x61, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x70, 0x61, 0x74, 0x63, 0x68, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x73, 0x2f, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_tests_message_message_renames_proto_rawDescOnce sync.Once
	file_tests_message_message_renames_proto_rawDescData = file_tests_message_message_renames_proto_rawDesc
)

func file_tests_message_message_renames_proto_rawDescGZIP() []byte {
	file_tests_message_message_renames_proto_rawDescOnce.Do(func() {
		file_tests_message_message_renames_proto_rawDescData = protoimpl.X.CompressGZIP(file_tests_message_message_renames_proto_rawDescData)
	})
	return file_tests_message_message_renames_proto_rawDescData
}

var file_tests_message_message_renames_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_tests_message_message_renames_proto_goTypes = []interface{}{
	(*Frank)(nil),                               // 0: tests.message.Francis
	(*RenamedOneofMessage)(nil),                 // 1: tests.message.OriginalOneofMessage
	(*RenamedOuterMessage)(nil),                 // 2: tests.message.OriginalOuterMessage
	(*OuterMessageWithRenamedInnerMessage)(nil), // 3: tests.message.OuterMessageWithRenamedInnerMessage
	(*MessageWithRenamedField)(nil),             // 4: tests.message.MessageWithRenamedField
	(*MessageWithEmbeddedField)(nil),            // 5: tests.message.MessageWithEmbeddedField
	(*Embedded)(nil),                            // 6: tests.message.Embedded
	(*MessageWithOptionals)(nil),                // 7: tests.message.MessageWithOptionals
	(*RenamedOuterMessage_InnerMessage)(nil),    // 8: tests.message.OriginalOuterMessage.InnerMessage
	(*RenamedInnerMessage)(nil),                 // 9: tests.message.OuterMessageWithRenamedInnerMessage.InnerMessage
}
var file_tests_message_message_renames_proto_depIdxs = []int32{
	8, // 0: tests.message.OriginalOuterMessage.inner:type_name -> tests.message.OriginalOuterMessage.InnerMessage
	9, // 1: tests.message.OuterMessageWithRenamedInnerMessage.inner:type_name -> tests.message.OuterMessageWithRenamedInnerMessage.InnerMessage
	6, // 2: tests.message.MessageWithEmbeddedField.embedded_message:type_name -> tests.message.Embedded
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_tests_message_message_renames_proto_init() }
func file_tests_message_message_renames_proto_init() {
	if File_tests_message_message_renames_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_tests_message_message_renames_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Frank); i {
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
		file_tests_message_message_renames_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RenamedOneofMessage); i {
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
		file_tests_message_message_renames_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RenamedOuterMessage); i {
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
		file_tests_message_message_renames_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OuterMessageWithRenamedInnerMessage); i {
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
		file_tests_message_message_renames_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MessageWithRenamedField); i {
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
		file_tests_message_message_renames_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MessageWithEmbeddedField); i {
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
		file_tests_message_message_renames_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Embedded); i {
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
		file_tests_message_message_renames_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MessageWithOptionals); i {
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
		file_tests_message_message_renames_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RenamedOuterMessage_InnerMessage); i {
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
		file_tests_message_message_renames_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RenamedInnerMessage); i {
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
	file_tests_message_message_renames_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*RenamedOneofMessage_Id)(nil),
		(*RenamedOneofMessage_Name)(nil),
	}
	file_tests_message_message_renames_proto_msgTypes[7].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_tests_message_message_renames_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_tests_message_message_renames_proto_goTypes,
		DependencyIndexes: file_tests_message_message_renames_proto_depIdxs,
		MessageInfos:      file_tests_message_message_renames_proto_msgTypes,
	}.Build()
	File_tests_message_message_renames_proto = out.File
	file_tests_message_message_renames_proto_rawDesc = nil
	file_tests_message_message_renames_proto_goTypes = nil
	file_tests_message_message_renames_proto_depIdxs = nil
}
