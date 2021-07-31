// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.15.6
// source: aip/fieldmask/internal/testing/testing.proto

package testing_go_proto

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

// A message containing various types of fields that should be tested against.
type Test struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	S         string    `protobuf:"bytes,1,opt,name=s,proto3" json:"s,omitempty"`
	RepS      []string  `protobuf:"bytes,2,rep,name=rep_s,json=repS,proto3" json:"rep_s,omitempty"`
	Nested    *Nested   `protobuf:"bytes,3,opt,name=nested,proto3" json:"nested,omitempty"`
	RepNested []*Nested `protobuf:"bytes,4,rep,name=rep_nested,json=repNested,proto3" json:"rep_nested,omitempty"`
	// Types that are assignable to Oo:
	//	*Test_OoS
	//	*Test_OoNested
	Oo isTest_Oo `protobuf_oneof:"oo"`
}

func (x *Test) Reset() {
	*x = Test{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aip_fieldmask_internal_testing_testing_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Test) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Test) ProtoMessage() {}

func (x *Test) ProtoReflect() protoreflect.Message {
	mi := &file_aip_fieldmask_internal_testing_testing_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Test.ProtoReflect.Descriptor instead.
func (*Test) Descriptor() ([]byte, []int) {
	return file_aip_fieldmask_internal_testing_testing_proto_rawDescGZIP(), []int{0}
}

func (x *Test) GetS() string {
	if x != nil {
		return x.S
	}
	return ""
}

func (x *Test) GetRepS() []string {
	if x != nil {
		return x.RepS
	}
	return nil
}

func (x *Test) GetNested() *Nested {
	if x != nil {
		return x.Nested
	}
	return nil
}

func (x *Test) GetRepNested() []*Nested {
	if x != nil {
		return x.RepNested
	}
	return nil
}

func (m *Test) GetOo() isTest_Oo {
	if m != nil {
		return m.Oo
	}
	return nil
}

func (x *Test) GetOoS() string {
	if x, ok := x.GetOo().(*Test_OoS); ok {
		return x.OoS
	}
	return ""
}

func (x *Test) GetOoNested() *Nested {
	if x, ok := x.GetOo().(*Test_OoNested); ok {
		return x.OoNested
	}
	return nil
}

type isTest_Oo interface {
	isTest_Oo()
}

type Test_OoS struct {
	OoS string `protobuf:"bytes,5,opt,name=oo_s,json=ooS,proto3,oneof"`
}

type Test_OoNested struct {
	OoNested *Nested `protobuf:"bytes,6,opt,name=oo_nested,json=ooNested,proto3,oneof"`
}

func (*Test_OoS) isTest_Oo() {}

func (*Test_OoNested) isTest_Oo() {}

// A message that is intended to be used in a field in the Test message.
type Nested struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Foo string `protobuf:"bytes,1,opt,name=foo,proto3" json:"foo,omitempty"`
	Bar string `protobuf:"bytes,2,opt,name=bar,proto3" json:"bar,omitempty"`
}

func (x *Nested) Reset() {
	*x = Nested{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aip_fieldmask_internal_testing_testing_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Nested) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Nested) ProtoMessage() {}

func (x *Nested) ProtoReflect() protoreflect.Message {
	mi := &file_aip_fieldmask_internal_testing_testing_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Nested.ProtoReflect.Descriptor instead.
func (*Nested) Descriptor() ([]byte, []int) {
	return file_aip_fieldmask_internal_testing_testing_proto_rawDescGZIP(), []int{1}
}

func (x *Nested) GetFoo() string {
	if x != nil {
		return x.Foo
	}
	return ""
}

func (x *Nested) GetBar() string {
	if x != nil {
		return x.Bar
	}
	return ""
}

var File_aip_fieldmask_internal_testing_testing_proto protoreflect.FileDescriptor

var file_aip_fieldmask_internal_testing_testing_proto_rawDesc = []byte{
	0x0a, 0x2c, 0x61, 0x69, 0x70, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x6d, 0x61, 0x73, 0x6b, 0x2f,
	0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67,
	0x2f, 0x74, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1e,
	0x61, 0x69, 0x70, 0x2e, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x6d, 0x61, 0x73, 0x6b, 0x2e, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x22, 0x92,
	0x02, 0x0a, 0x04, 0x54, 0x65, 0x73, 0x74, 0x12, 0x0c, 0x0a, 0x01, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x01, 0x73, 0x12, 0x13, 0x0a, 0x05, 0x72, 0x65, 0x70, 0x5f, 0x73, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x72, 0x65, 0x70, 0x53, 0x12, 0x3e, 0x0a, 0x06, 0x6e, 0x65,
	0x73, 0x74, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x61, 0x69, 0x70,
	0x2e, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x6d, 0x61, 0x73, 0x6b, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72,
	0x6e, 0x61, 0x6c, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x4e, 0x65, 0x73, 0x74,
	0x65, 0x64, 0x52, 0x06, 0x6e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x12, 0x45, 0x0a, 0x0a, 0x72, 0x65,
	0x70, 0x5f, 0x6e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x26,
	0x2e, 0x61, 0x69, 0x70, 0x2e, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x6d, 0x61, 0x73, 0x6b, 0x2e, 0x69,
	0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x2e,
	0x4e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x52, 0x09, 0x72, 0x65, 0x70, 0x4e, 0x65, 0x73, 0x74, 0x65,
	0x64, 0x12, 0x13, 0x0a, 0x04, 0x6f, 0x6f, 0x5f, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x48,
	0x00, 0x52, 0x03, 0x6f, 0x6f, 0x53, 0x12, 0x45, 0x0a, 0x09, 0x6f, 0x6f, 0x5f, 0x6e, 0x65, 0x73,
	0x74, 0x65, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x61, 0x69, 0x70, 0x2e,
	0x66, 0x69, 0x65, 0x6c, 0x64, 0x6d, 0x61, 0x73, 0x6b, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e,
	0x61, 0x6c, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x4e, 0x65, 0x73, 0x74, 0x65,
	0x64, 0x48, 0x00, 0x52, 0x08, 0x6f, 0x6f, 0x4e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x42, 0x04, 0x0a,
	0x02, 0x6f, 0x6f, 0x22, 0x2c, 0x0a, 0x06, 0x4e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x12, 0x10, 0x0a,
	0x03, 0x66, 0x6f, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x66, 0x6f, 0x6f, 0x12,
	0x10, 0x0a, 0x03, 0x62, 0x61, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x62, 0x61,
	0x72, 0x42, 0x46, 0x5a, 0x44, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x53, 0x61, 0x73, 0x65, 0x72, 0x2f, 0x70, 0x64, 0x70, 0x2f, 0x61, 0x69, 0x70, 0x2f, 0x66, 0x69,
	0x65, 0x6c, 0x64, 0x6d, 0x61, 0x73, 0x6b, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c,
	0x2f, 0x74, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67,
	0x5f, 0x67, 0x6f, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_aip_fieldmask_internal_testing_testing_proto_rawDescOnce sync.Once
	file_aip_fieldmask_internal_testing_testing_proto_rawDescData = file_aip_fieldmask_internal_testing_testing_proto_rawDesc
)

func file_aip_fieldmask_internal_testing_testing_proto_rawDescGZIP() []byte {
	file_aip_fieldmask_internal_testing_testing_proto_rawDescOnce.Do(func() {
		file_aip_fieldmask_internal_testing_testing_proto_rawDescData = protoimpl.X.CompressGZIP(file_aip_fieldmask_internal_testing_testing_proto_rawDescData)
	})
	return file_aip_fieldmask_internal_testing_testing_proto_rawDescData
}

var file_aip_fieldmask_internal_testing_testing_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_aip_fieldmask_internal_testing_testing_proto_goTypes = []interface{}{
	(*Test)(nil),   // 0: aip.fieldmask.internal.testing.Test
	(*Nested)(nil), // 1: aip.fieldmask.internal.testing.Nested
}
var file_aip_fieldmask_internal_testing_testing_proto_depIdxs = []int32{
	1, // 0: aip.fieldmask.internal.testing.Test.nested:type_name -> aip.fieldmask.internal.testing.Nested
	1, // 1: aip.fieldmask.internal.testing.Test.rep_nested:type_name -> aip.fieldmask.internal.testing.Nested
	1, // 2: aip.fieldmask.internal.testing.Test.oo_nested:type_name -> aip.fieldmask.internal.testing.Nested
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_aip_fieldmask_internal_testing_testing_proto_init() }
func file_aip_fieldmask_internal_testing_testing_proto_init() {
	if File_aip_fieldmask_internal_testing_testing_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_aip_fieldmask_internal_testing_testing_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Test); i {
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
		file_aip_fieldmask_internal_testing_testing_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Nested); i {
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
	file_aip_fieldmask_internal_testing_testing_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*Test_OoS)(nil),
		(*Test_OoNested)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_aip_fieldmask_internal_testing_testing_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_aip_fieldmask_internal_testing_testing_proto_goTypes,
		DependencyIndexes: file_aip_fieldmask_internal_testing_testing_proto_depIdxs,
		MessageInfos:      file_aip_fieldmask_internal_testing_testing_proto_msgTypes,
	}.Build()
	File_aip_fieldmask_internal_testing_testing_proto = out.File
	file_aip_fieldmask_internal_testing_testing_proto_rawDesc = nil
	file_aip_fieldmask_internal_testing_testing_proto_goTypes = nil
	file_aip_fieldmask_internal_testing_testing_proto_depIdxs = nil
}
