// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.23.4
// source: tests/plugin/validate.proto

package plugin

import (
	_ "github.com/alta/protopatch/patch/gopb"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
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

type InterfaceStatus int32

const (
	StatusUnknown InterfaceStatus = 0
	StatusUp      InterfaceStatus = 1
	StatusDown    InterfaceStatus = 2
)

// Enum value maps for Interface_Status.
var (
	InterfaceStatus_name = map[int32]string{
		0: "UNKNOWN",
		1: "UP",
		2: "DOWN",
	}
	InterfaceStatus_value = map[string]int32{
		"UNKNOWN": 0,
		"UP":      1,
		"DOWN":    2,
	}
)

func (x InterfaceStatus) Enum() *InterfaceStatus {
	p := new(InterfaceStatus)
	*p = x
	return p
}

func (x InterfaceStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (InterfaceStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_tests_plugin_validate_proto_enumTypes[0].Descriptor()
}

func (InterfaceStatus) Type() protoreflect.EnumType {
	return &file_tests_plugin_validate_proto_enumTypes[0]
}

func (x InterfaceStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Interface_Status.Descriptor instead.
func (InterfaceStatus) EnumDescriptor() ([]byte, []int) {
	return file_tests_plugin_validate_proto_rawDescGZIP(), []int{0, 0}
}

type Interface struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name      string          `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Status    InterfaceStatus `protobuf:"varint,2,opt,name=status,proto3,enum=tests.plugin.Interface_Status" json:"status,omitempty"`
	Addresses []*IPAddress    `protobuf:"bytes,3,rep,name=addresses,proto3" json:"addresses,omitempty"`
}

func (x *Interface) Reset() {
	*x = Interface{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tests_plugin_validate_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Interface) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Interface) ProtoMessage() {}

func (x *Interface) ProtoReflect() protoreflect.Message {
	mi := &file_tests_plugin_validate_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Interface.ProtoReflect.Descriptor instead.
func (*Interface) Descriptor() ([]byte, []int) {
	return file_tests_plugin_validate_proto_rawDescGZIP(), []int{0}
}

func (x *Interface) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Interface) GetStatus() InterfaceStatus {
	if x != nil {
		return x.Status
	}
	return StatusUnknown
}

func (x *Interface) GetAddresses() []*IPAddress {
	if x != nil {
		return x.Addresses
	}
	return nil
}

type IPAddress struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Address:
	//
	//	*IPAddress_Ipv4
	//	*IPAddress_Ipv6
	Address isIPAddress_Address `protobuf_oneof:"Address"`
}

func (x *IPAddress) Reset() {
	*x = IPAddress{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tests_plugin_validate_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IPAddress) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IPAddress) ProtoMessage() {}

func (x *IPAddress) ProtoReflect() protoreflect.Message {
	mi := &file_tests_plugin_validate_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IPAddress.ProtoReflect.Descriptor instead.
func (*IPAddress) Descriptor() ([]byte, []int) {
	return file_tests_plugin_validate_proto_rawDescGZIP(), []int{1}
}

func (m *IPAddress) GetAddress() isIPAddress_Address {
	if m != nil {
		return m.Address
	}
	return nil
}

func (x *IPAddress) GetIPV4() string {
	if x, ok := x.GetAddress().(*IPAddress_IPV4); ok {
		return x.IPV4
	}
	return ""
}

func (x *IPAddress) GetIPV6() string {
	if x, ok := x.GetAddress().(*IPAddress_IPV6); ok {
		return x.IPV6
	}
	return ""
}

type isIPAddress_Address interface {
	isIPAddress_Address()
}

type IPAddress_IPV4 struct {
	IPV4 string `protobuf:"bytes,3,opt,name=ipv4,proto3,oneof"`
}

type IPAddress_IPV6 struct {
	IPV6 string `protobuf:"bytes,4,opt,name=ipv6,proto3,oneof"`
}

func (*IPAddress_IPV4) isIPAddress_Address() {}

func (*IPAddress_IPV6) isIPAddress_Address() {}

type InterfaceWithCustomTypes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name      Name        `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Addresses IPAddresses `protobuf:"bytes,3,rep,name=addresses,proto3" json:"addresses,omitempty"`
	Aliases   Names       `protobuf:"bytes,4,rep,name=aliases,proto3" json:"aliases,omitempty"`
}

func (x *InterfaceWithCustomTypes) Reset() {
	*x = InterfaceWithCustomTypes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tests_plugin_validate_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InterfaceWithCustomTypes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InterfaceWithCustomTypes) ProtoMessage() {}

func (x *InterfaceWithCustomTypes) ProtoReflect() protoreflect.Message {
	mi := &file_tests_plugin_validate_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InterfaceWithCustomTypes.ProtoReflect.Descriptor instead.
func (*InterfaceWithCustomTypes) Descriptor() ([]byte, []int) {
	return file_tests_plugin_validate_proto_rawDescGZIP(), []int{2}
}

func (x *InterfaceWithCustomTypes) GetName() Name {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *InterfaceWithCustomTypes) GetAddresses() IPAddresses {
	if x != nil {
		return x.Addresses
	}
	return nil
}

func (x *InterfaceWithCustomTypes) GetAliases() Names {
	if x != nil {
		return x.Aliases
	}
	return nil
}

var File_tests_plugin_validate_proto protoreflect.FileDescriptor

var file_tests_plugin_validate_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x74, 0x65, 0x73, 0x74, 0x73, 0x2f, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2f, 0x76,
	0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x74,
	0x65, 0x73, 0x74, 0x73, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x1a, 0x0e, 0x70, 0x61, 0x74,
	0x63, 0x68, 0x2f, 0x67, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xad, 0x02, 0x0a, 0x09, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61,
	0x63, 0x65, 0x12, 0x2e, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x1a, 0xfa, 0x42, 0x17, 0x72, 0x15, 0x10, 0x02, 0x18, 0x0a, 0x32, 0x0f, 0x5b, 0x30, 0x2d,
	0x39, 0x61, 0x2d, 0x7a, 0x41, 0x2d, 0x5a, 0x2e, 0x2d, 0x5f, 0x5d, 0x2a, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x42, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x1e, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x73, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69,
	0x6e, 0x2e, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x2e, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x42, 0x0a, 0xfa, 0x42, 0x07, 0x82, 0x01, 0x04, 0x10, 0x01, 0x20, 0x00, 0x52, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x35, 0x0a, 0x09, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x74, 0x65, 0x73, 0x74,
	0x73, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e, 0x49, 0x50, 0x41, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x52, 0x09, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x65, 0x73, 0x22, 0x75, 0x0a,
	0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x20, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f,
	0x57, 0x4e, 0x10, 0x00, 0x1a, 0x13, 0xca, 0xb5, 0x03, 0x0f, 0x0a, 0x0d, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x55, 0x6e, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x12, 0x16, 0x0a, 0x02, 0x55, 0x50, 0x10,
	0x01, 0x1a, 0x0e, 0xca, 0xb5, 0x03, 0x0a, 0x0a, 0x08, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x55,
	0x70, 0x12, 0x1a, 0x0a, 0x04, 0x44, 0x4f, 0x57, 0x4e, 0x10, 0x02, 0x1a, 0x10, 0xca, 0xb5, 0x03,
	0x0c, 0x0a, 0x0a, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x44, 0x6f, 0x77, 0x6e, 0x1a, 0x15, 0xca,
	0xb5, 0x03, 0x11, 0x0a, 0x0f, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x22, 0x69, 0x0a, 0x09, 0x49, 0x50, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x12, 0x27, 0x0a, 0x04, 0x69, 0x70, 0x76, 0x34, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x11, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x78, 0x01, 0xca, 0xb5, 0x03, 0x06, 0x0a, 0x04, 0x49, 0x50,
	0x56, 0x34, 0x48, 0x00, 0x52, 0x04, 0x69, 0x70, 0x76, 0x34, 0x12, 0x28, 0x0a, 0x04, 0x69, 0x70,
	0x76, 0x36, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x42, 0x12, 0xfa, 0x42, 0x05, 0x72, 0x03, 0x80,
	0x01, 0x01, 0xca, 0xb5, 0x03, 0x06, 0x0a, 0x04, 0x49, 0x50, 0x56, 0x36, 0x48, 0x00, 0x52, 0x04,
	0x69, 0x70, 0x76, 0x36, 0x42, 0x09, 0x0a, 0x07, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x22,
	0xe6, 0x01, 0x0a, 0x18, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x57, 0x69, 0x74,
	0x68, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x54, 0x79, 0x70, 0x65, 0x73, 0x12, 0x38, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x24, 0xfa, 0x42, 0x17, 0x72,
	0x15, 0x10, 0x02, 0x18, 0x0a, 0x32, 0x0f, 0x5b, 0x30, 0x2d, 0x39, 0x61, 0x2d, 0x7a, 0x41, 0x2d,
	0x5a, 0x2e, 0x2d, 0x5f, 0x5d, 0x2a, 0xca, 0xb5, 0x03, 0x06, 0x1a, 0x04, 0x4e, 0x61, 0x6d, 0x65,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x48, 0x0a, 0x09, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x74, 0x65, 0x73, 0x74,
	0x73, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e, 0x49, 0x50, 0x41, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x42, 0x11, 0xca, 0xb5, 0x03, 0x0d, 0x1a, 0x0b, 0x49, 0x50, 0x41, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x65, 0x73, 0x52, 0x09, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x65, 0x73,
	0x12, 0x46, 0x0a, 0x07, 0x61, 0x6c, 0x69, 0x61, 0x73, 0x65, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28,
	0x09, 0x42, 0x2c, 0xfa, 0x42, 0x1e, 0x92, 0x01, 0x1b, 0x10, 0x0a, 0x22, 0x17, 0x72, 0x15, 0x10,
	0x02, 0x18, 0x0a, 0x32, 0x0f, 0x5b, 0x30, 0x2d, 0x39, 0x61, 0x2d, 0x7a, 0x41, 0x2d, 0x5a, 0x2e,
	0x2d, 0x5f, 0x5d, 0x2a, 0xca, 0xb5, 0x03, 0x07, 0x1a, 0x05, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x52,
	0x07, 0x61, 0x6c, 0x69, 0x61, 0x73, 0x65, 0x73, 0x42, 0x29, 0x5a, 0x27, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x6c, 0x74, 0x61, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x70, 0x61, 0x74, 0x63, 0x68, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x73, 0x2f, 0x70, 0x6c, 0x75,
	0x67, 0x69, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_tests_plugin_validate_proto_rawDescOnce sync.Once
	file_tests_plugin_validate_proto_rawDescData = file_tests_plugin_validate_proto_rawDesc
)

func file_tests_plugin_validate_proto_rawDescGZIP() []byte {
	file_tests_plugin_validate_proto_rawDescOnce.Do(func() {
		file_tests_plugin_validate_proto_rawDescData = protoimpl.X.CompressGZIP(file_tests_plugin_validate_proto_rawDescData)
	})
	return file_tests_plugin_validate_proto_rawDescData
}

var file_tests_plugin_validate_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_tests_plugin_validate_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_tests_plugin_validate_proto_goTypes = []interface{}{
	(InterfaceStatus)(0),             // 0: tests.plugin.Interface.Status
	(*Interface)(nil),                // 1: tests.plugin.Interface
	(*IPAddress)(nil),                // 2: tests.plugin.IPAddress
	(*InterfaceWithCustomTypes)(nil), // 3: tests.plugin.InterfaceWithCustomTypes
}
var file_tests_plugin_validate_proto_depIdxs = []int32{
	0, // 0: tests.plugin.Interface.status:type_name -> tests.plugin.Interface.Status
	2, // 1: tests.plugin.Interface.addresses:type_name -> tests.plugin.IPAddress
	2, // 2: tests.plugin.InterfaceWithCustomTypes.addresses:type_name -> tests.plugin.IPAddress
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_tests_plugin_validate_proto_init() }
func file_tests_plugin_validate_proto_init() {
	if File_tests_plugin_validate_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_tests_plugin_validate_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Interface); i {
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
		file_tests_plugin_validate_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IPAddress); i {
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
		file_tests_plugin_validate_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InterfaceWithCustomTypes); i {
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
	file_tests_plugin_validate_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*IPAddress_IPV4)(nil),
		(*IPAddress_IPV6)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_tests_plugin_validate_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_tests_plugin_validate_proto_goTypes,
		DependencyIndexes: file_tests_plugin_validate_proto_depIdxs,
		EnumInfos:         file_tests_plugin_validate_proto_enumTypes,
		MessageInfos:      file_tests_plugin_validate_proto_msgTypes,
	}.Build()
	File_tests_plugin_validate_proto = out.File
	file_tests_plugin_validate_proto_rawDesc = nil
	file_tests_plugin_validate_proto_goTypes = nil
	file_tests_plugin_validate_proto_depIdxs = nil
}
