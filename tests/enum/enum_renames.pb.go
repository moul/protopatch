// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: tests/enum/enum_renames.proto

// clang-format off

package enum

import (
	_ "github.com/alta/protopatch/patch/gopb"
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type Flavour int32

const (
	Flavour_INVALID Flavour = 0
	Flavour_SWEET   Flavour = 1
	Flavour_SALTY   Flavour = 2
	Flavour_SOUR    Flavour = 3
	Flavour_BITTER  Flavour = 4
)

// Enum value maps for Flavor.
var (
	Flavour_name = map[int32]string{
		0: "INVALID",
		1: "SWEET",
		2: "SALTY",
		3: "SOUR",
		4: "BITTER",
	}
	Flavour_value = map[string]int32{
		"INVALID": 0,
		"SWEET":   1,
		"SALTY":   2,
		"SOUR":    3,
		"BITTER":  4,
	}
)

func (x Flavour) Enum() *Flavour {
	p := new(Flavour)
	*p = x
	return p
}

func (x Flavour) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Flavour) Descriptor() protoreflect.EnumDescriptor {
	return file_tests_enum_enum_renames_proto_enumTypes[0].Descriptor()
}

func (Flavour) Type() protoreflect.EnumType {
	return &file_tests_enum_enum_renames_proto_enumTypes[0]
}

func (x Flavour) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Flavor.Descriptor instead.
func (Flavour) EnumDescriptor() ([]byte, []int) {
	return file_tests_enum_enum_renames_proto_rawDescGZIP(), []int{0}
}

type Level int32

const (
	LevelSimple   Level = 0
	Level_COMPLEX Level = 1
)

// Enum value maps for Level.
var (
	Level_name = map[int32]string{
		0: "SIMPLE",
		1: "COMPLEX",
	}
	Level_value = map[string]int32{
		"SIMPLE":  0,
		"COMPLEX": 1,
	}
)

func (x Level) Enum() *Level {
	p := new(Level)
	*p = x
	return p
}

func (x Level) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Level) Descriptor() protoreflect.EnumDescriptor {
	return file_tests_enum_enum_renames_proto_enumTypes[1].Descriptor()
}

func (Level) Type() protoreflect.EnumType {
	return &file_tests_enum_enum_renames_proto_enumTypes[1]
}

func (x Level) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Level.Descriptor instead.
func (Level) EnumDescriptor() ([]byte, []int) {
	return file_tests_enum_enum_renames_proto_rawDescGZIP(), []int{1}
}

type RenamedNested int32

const (
	RenamedValueInvalid RenamedNested = 0
	RenamedValueA       RenamedNested = 1
	RenamedValueB       RenamedNested = 2
	RenamedValueC       RenamedNested = 3
)

// Enum value maps for Wrapper_Original.
var (
	RenamedNested_name = map[int32]string{
		0: "INVALID",
		1: "A",
		2: "B",
		3: "C",
	}
	RenamedNested_value = map[string]int32{
		"INVALID": 0,
		"A":       1,
		"B":       2,
		"C":       3,
	}
)

func (x RenamedNested) Enum() *RenamedNested {
	p := new(RenamedNested)
	*p = x
	return p
}

func (x RenamedNested) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RenamedNested) Descriptor() protoreflect.EnumDescriptor {
	return file_tests_enum_enum_renames_proto_enumTypes[2].Descriptor()
}

func (RenamedNested) Type() protoreflect.EnumType {
	return &file_tests_enum_enum_renames_proto_enumTypes[2]
}

func (x RenamedNested) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Wrapper_Original.Descriptor instead.
func (RenamedNested) EnumDescriptor() ([]byte, []int) {
	return file_tests_enum_enum_renames_proto_rawDescGZIP(), []int{0, 0}
}

type Holiday_Route int32

const (
	Holiday_INVALID Holiday_Route = 0
	Holiday_FAST    Holiday_Route = 1
	Holiday_SLOW    Holiday_Route = 2
	Holiday_SCENIC  Holiday_Route = 3
)

// Enum value maps for Vacation_Route.
var (
	Holiday_Route_name = map[int32]string{
		0: "INVALID",
		1: "FAST",
		2: "SLOW",
		3: "SCENIC",
	}
	Holiday_Route_value = map[string]int32{
		"INVALID": 0,
		"FAST":    1,
		"SLOW":    2,
		"SCENIC":  3,
	}
)

func (x Holiday_Route) Enum() *Holiday_Route {
	p := new(Holiday_Route)
	*p = x
	return p
}

func (x Holiday_Route) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Holiday_Route) Descriptor() protoreflect.EnumDescriptor {
	return file_tests_enum_enum_renames_proto_enumTypes[3].Descriptor()
}

func (Holiday_Route) Type() protoreflect.EnumType {
	return &file_tests_enum_enum_renames_proto_enumTypes[3]
}

func (x Holiday_Route) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Vacation_Route.Descriptor instead.
func (Holiday_Route) EnumDescriptor() ([]byte, []int) {
	return file_tests_enum_enum_renames_proto_rawDescGZIP(), []int{1, 0}
}

type Wrapper struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Inner RenamedNested `protobuf:"varint,1,opt,name=inner,proto3,enum=tests.enum.Wrapper_Original" json:"inner,omitempty"`
}

func (x *Wrapper) Reset() {
	*x = Wrapper{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tests_enum_enum_renames_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Wrapper) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Wrapper) ProtoMessage() {}

func (x *Wrapper) ProtoReflect() protoreflect.Message {
	mi := &file_tests_enum_enum_renames_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Wrapper.ProtoReflect.Descriptor instead.
func (*Wrapper) Descriptor() ([]byte, []int) {
	return file_tests_enum_enum_renames_proto_rawDescGZIP(), []int{0}
}

func (x *Wrapper) GetInner() RenamedNested {
	if x != nil {
		return x.Inner
	}
	return RenamedValueInvalid
}

type Holiday struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Inner Holiday_Route `protobuf:"varint,1,opt,name=inner,proto3,enum=tests.enum.Vacation_Route" json:"inner,omitempty"`
}

func (x *Holiday) Reset() {
	*x = Holiday{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tests_enum_enum_renames_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Holiday) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Holiday) ProtoMessage() {}

func (x *Holiday) ProtoReflect() protoreflect.Message {
	mi := &file_tests_enum_enum_renames_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Vacation.ProtoReflect.Descriptor instead.
func (*Holiday) Descriptor() ([]byte, []int) {
	return file_tests_enum_enum_renames_proto_rawDescGZIP(), []int{1}
}

func (x *Holiday) GetInner() Holiday_Route {
	if x != nil {
		return x.Inner
	}
	return Holiday_INVALID
}

var File_tests_enum_enum_renames_proto protoreflect.FileDescriptor

var file_tests_enum_enum_renames_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x74, 0x65, 0x73, 0x74, 0x73, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x2f, 0x65, 0x6e, 0x75,
	0x6d, 0x5f, 0x72, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0a, 0x74, 0x65, 0x73, 0x74, 0x73, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x1a, 0x0e, 0x70, 0x61, 0x74,
	0x63, 0x68, 0x2f, 0x67, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xdb, 0x01, 0x0a, 0x07,
	0x57, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x12, 0x32, 0x0a, 0x05, 0x69, 0x6e, 0x6e, 0x65, 0x72,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1c, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x73, 0x2e, 0x65,
	0x6e, 0x75, 0x6d, 0x2e, 0x57, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x2e, 0x4f, 0x72, 0x69, 0x67,
	0x69, 0x6e, 0x61, 0x6c, 0x52, 0x05, 0x69, 0x6e, 0x6e, 0x65, 0x72, 0x22, 0x9b, 0x01, 0x0a, 0x08,
	0x4f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x12, 0x26, 0x0a, 0x07, 0x49, 0x4e, 0x56, 0x41,
	0x4c, 0x49, 0x44, 0x10, 0x00, 0x1a, 0x19, 0xca, 0xb5, 0x03, 0x15, 0x0a, 0x13, 0x52, 0x65, 0x6e,
	0x61, 0x6d, 0x65, 0x64, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x49, 0x6e, 0x76, 0x61, 0x6c, 0x69, 0x64,
	0x12, 0x1a, 0x0a, 0x01, 0x41, 0x10, 0x01, 0x1a, 0x13, 0xca, 0xb5, 0x03, 0x0f, 0x0a, 0x0d, 0x52,
	0x65, 0x6e, 0x61, 0x6d, 0x65, 0x64, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x41, 0x12, 0x1a, 0x0a, 0x01,
	0x42, 0x10, 0x02, 0x1a, 0x13, 0xca, 0xb5, 0x03, 0x0f, 0x0a, 0x0d, 0x52, 0x65, 0x6e, 0x61, 0x6d,
	0x65, 0x64, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x12, 0x1a, 0x0a, 0x01, 0x43, 0x10, 0x03, 0x1a,
	0x13, 0xca, 0xb5, 0x03, 0x0f, 0x0a, 0x0d, 0x52, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x64, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x43, 0x1a, 0x13, 0xca, 0xb5, 0x03, 0x0f, 0x0a, 0x0d, 0x52, 0x65, 0x6e, 0x61,
	0x6d, 0x65, 0x64, 0x4e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x22, 0x81, 0x01, 0x0a, 0x08, 0x56, 0x61,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x30, 0x0a, 0x05, 0x69, 0x6e, 0x6e, 0x65, 0x72, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1a, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x73, 0x2e, 0x65, 0x6e,
	0x75, 0x6d, 0x2e, 0x56, 0x61, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x52, 0x6f, 0x75, 0x74,
	0x65, 0x52, 0x05, 0x69, 0x6e, 0x6e, 0x65, 0x72, 0x22, 0x34, 0x0a, 0x05, 0x52, 0x6f, 0x75, 0x74,
	0x65, 0x12, 0x0b, 0x0a, 0x07, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x10, 0x00, 0x12, 0x08,
	0x0a, 0x04, 0x46, 0x41, 0x53, 0x54, 0x10, 0x01, 0x12, 0x08, 0x0a, 0x04, 0x53, 0x4c, 0x4f, 0x57,
	0x10, 0x02, 0x12, 0x0a, 0x0a, 0x06, 0x53, 0x43, 0x45, 0x4e, 0x49, 0x43, 0x10, 0x03, 0x3a, 0x0d,
	0xca, 0xb5, 0x03, 0x09, 0x0a, 0x07, 0x48, 0x6f, 0x6c, 0x69, 0x64, 0x61, 0x79, 0x2a, 0x50, 0x0a,
	0x06, 0x46, 0x6c, 0x61, 0x76, 0x6f, 0x72, 0x12, 0x0b, 0x0a, 0x07, 0x49, 0x4e, 0x56, 0x41, 0x4c,
	0x49, 0x44, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x53, 0x57, 0x45, 0x45, 0x54, 0x10, 0x01, 0x12,
	0x09, 0x0a, 0x05, 0x53, 0x41, 0x4c, 0x54, 0x59, 0x10, 0x02, 0x12, 0x08, 0x0a, 0x04, 0x53, 0x4f,
	0x55, 0x52, 0x10, 0x03, 0x12, 0x0a, 0x0a, 0x06, 0x42, 0x49, 0x54, 0x54, 0x45, 0x52, 0x10, 0x04,
	0x1a, 0x0d, 0xca, 0xb5, 0x03, 0x09, 0x0a, 0x07, 0x46, 0x6c, 0x61, 0x76, 0x6f, 0x75, 0x72, 0x2a,
	0x33, 0x0a, 0x05, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x1d, 0x0a, 0x06, 0x53, 0x49, 0x4d, 0x50,
	0x4c, 0x45, 0x10, 0x00, 0x1a, 0x11, 0xca, 0xb5, 0x03, 0x0d, 0x0a, 0x0b, 0x4c, 0x65, 0x76, 0x65,
	0x6c, 0x53, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x43, 0x4f, 0x4d, 0x50, 0x4c,
	0x45, 0x58, 0x10, 0x01, 0x42, 0x27, 0x5a, 0x25, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x61, 0x6c, 0x74, 0x61, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x70, 0x61, 0x74,
	0x63, 0x68, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x73, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_tests_enum_enum_renames_proto_rawDescOnce sync.Once
	file_tests_enum_enum_renames_proto_rawDescData = file_tests_enum_enum_renames_proto_rawDesc
)

func file_tests_enum_enum_renames_proto_rawDescGZIP() []byte {
	file_tests_enum_enum_renames_proto_rawDescOnce.Do(func() {
		file_tests_enum_enum_renames_proto_rawDescData = protoimpl.X.CompressGZIP(file_tests_enum_enum_renames_proto_rawDescData)
	})
	return file_tests_enum_enum_renames_proto_rawDescData
}

var file_tests_enum_enum_renames_proto_enumTypes = make([]protoimpl.EnumInfo, 4)
var file_tests_enum_enum_renames_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_tests_enum_enum_renames_proto_goTypes = []interface{}{
	(Flavour)(0),       // 0: tests.enum.Flavor
	(Level)(0),         // 1: tests.enum.Level
	(RenamedNested)(0), // 2: tests.enum.Wrapper.Original
	(Holiday_Route)(0), // 3: tests.enum.Vacation.Route
	(*Wrapper)(nil),    // 4: tests.enum.Wrapper
	(*Holiday)(nil),    // 5: tests.enum.Vacation
}
var file_tests_enum_enum_renames_proto_depIdxs = []int32{
	2, // 0: tests.enum.Wrapper.inner:type_name -> tests.enum.Wrapper.Original
	3, // 1: tests.enum.Vacation.inner:type_name -> tests.enum.Vacation.Route
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_tests_enum_enum_renames_proto_init() }
func file_tests_enum_enum_renames_proto_init() {
	if File_tests_enum_enum_renames_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_tests_enum_enum_renames_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Wrapper); i {
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
		file_tests_enum_enum_renames_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Holiday); i {
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
			RawDescriptor: file_tests_enum_enum_renames_proto_rawDesc,
			NumEnums:      4,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_tests_enum_enum_renames_proto_goTypes,
		DependencyIndexes: file_tests_enum_enum_renames_proto_depIdxs,
		EnumInfos:         file_tests_enum_enum_renames_proto_enumTypes,
		MessageInfos:      file_tests_enum_enum_renames_proto_msgTypes,
	}.Build()
	File_tests_enum_enum_renames_proto = out.File
	file_tests_enum_enum_renames_proto_rawDesc = nil
	file_tests_enum_enum_renames_proto_goTypes = nil
	file_tests_enum_enum_renames_proto_depIdxs = nil
}
