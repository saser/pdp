// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.6
// source: aip/fieldbehavior/internal/testing/testing.proto

package testing_go_proto

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type Test struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Unspecified string  `protobuf:"bytes,1,opt,name=unspecified,proto3" json:"unspecified,omitempty"`
	OutputOnly  string  `protobuf:"bytes,2,opt,name=output_only,json=outputOnly,proto3" json:"output_only,omitempty"`
	Nested      *Nested `protobuf:"bytes,3,opt,name=nested,proto3" json:"nested,omitempty"`
}

func (x *Test) Reset() {
	*x = Test{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aip_fieldbehavior_internal_testing_testing_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Test) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Test) ProtoMessage() {}

func (x *Test) ProtoReflect() protoreflect.Message {
	mi := &file_aip_fieldbehavior_internal_testing_testing_proto_msgTypes[0]
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
	return file_aip_fieldbehavior_internal_testing_testing_proto_rawDescGZIP(), []int{0}
}

func (x *Test) GetUnspecified() string {
	if x != nil {
		return x.Unspecified
	}
	return ""
}

func (x *Test) GetOutputOnly() string {
	if x != nil {
		return x.OutputOnly
	}
	return ""
}

func (x *Test) GetNested() *Nested {
	if x != nil {
		return x.Nested
	}
	return nil
}

type Nested struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Unspecified string `protobuf:"bytes,1,opt,name=unspecified,proto3" json:"unspecified,omitempty"`
	OutputOnly  string `protobuf:"bytes,2,opt,name=output_only,json=outputOnly,proto3" json:"output_only,omitempty"`
}

func (x *Nested) Reset() {
	*x = Nested{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aip_fieldbehavior_internal_testing_testing_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Nested) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Nested) ProtoMessage() {}

func (x *Nested) ProtoReflect() protoreflect.Message {
	mi := &file_aip_fieldbehavior_internal_testing_testing_proto_msgTypes[1]
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
	return file_aip_fieldbehavior_internal_testing_testing_proto_rawDescGZIP(), []int{1}
}

func (x *Nested) GetUnspecified() string {
	if x != nil {
		return x.Unspecified
	}
	return ""
}

func (x *Nested) GetOutputOnly() string {
	if x != nil {
		return x.OutputOnly
	}
	return ""
}

var File_aip_fieldbehavior_internal_testing_testing_proto protoreflect.FileDescriptor

var file_aip_fieldbehavior_internal_testing_testing_proto_rawDesc = []byte{
	0x0a, 0x30, 0x61, 0x69, 0x70, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x62, 0x65, 0x68, 0x61, 0x76,
	0x69, 0x6f, 0x72, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x74, 0x65, 0x73,
	0x74, 0x69, 0x6e, 0x67, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x22, 0x61, 0x69, 0x70, 0x2e, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x62, 0x65, 0x68,
	0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x74,
	0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x92, 0x01, 0x0a, 0x04, 0x54, 0x65, 0x73, 0x74,
	0x12, 0x20, 0x0a, 0x0b, 0x75, 0x6e, 0x73, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x65, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x75, 0x6e, 0x73, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69,
	0x65, 0x64, 0x12, 0x24, 0x0a, 0x0b, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x5f, 0x6f, 0x6e, 0x6c,
	0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x0a, 0x6f, 0x75,
	0x74, 0x70, 0x75, 0x74, 0x4f, 0x6e, 0x6c, 0x79, 0x12, 0x42, 0x0a, 0x06, 0x6e, 0x65, 0x73, 0x74,
	0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x61, 0x69, 0x70, 0x2e, 0x66,
	0x69, 0x65, 0x6c, 0x64, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x69, 0x6e, 0x74,
	0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x4e, 0x65,
	0x73, 0x74, 0x65, 0x64, 0x52, 0x06, 0x6e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x22, 0x50, 0x0a, 0x06,
	0x4e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x75, 0x6e, 0x73, 0x70, 0x65, 0x63,
	0x69, 0x66, 0x69, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x75, 0x6e, 0x73,
	0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x65, 0x64, 0x12, 0x24, 0x0a, 0x0b, 0x6f, 0x75, 0x74, 0x70,
	0x75, 0x74, 0x5f, 0x6f, 0x6e, 0x6c, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0,
	0x41, 0x03, 0x52, 0x0a, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x4f, 0x6e, 0x6c, 0x79, 0x42, 0x4a,
	0x5a, 0x48, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x53, 0x61, 0x73,
	0x65, 0x72, 0x2f, 0x70, 0x64, 0x70, 0x2f, 0x61, 0x69, 0x70, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64,
	0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61,
	0x6c, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x69, 0x6e,
	0x67, 0x5f, 0x67, 0x6f, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_aip_fieldbehavior_internal_testing_testing_proto_rawDescOnce sync.Once
	file_aip_fieldbehavior_internal_testing_testing_proto_rawDescData = file_aip_fieldbehavior_internal_testing_testing_proto_rawDesc
)

func file_aip_fieldbehavior_internal_testing_testing_proto_rawDescGZIP() []byte {
	file_aip_fieldbehavior_internal_testing_testing_proto_rawDescOnce.Do(func() {
		file_aip_fieldbehavior_internal_testing_testing_proto_rawDescData = protoimpl.X.CompressGZIP(file_aip_fieldbehavior_internal_testing_testing_proto_rawDescData)
	})
	return file_aip_fieldbehavior_internal_testing_testing_proto_rawDescData
}

var file_aip_fieldbehavior_internal_testing_testing_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_aip_fieldbehavior_internal_testing_testing_proto_goTypes = []interface{}{
	(*Test)(nil),   // 0: aip.fieldbehavior.internal.testing.Test
	(*Nested)(nil), // 1: aip.fieldbehavior.internal.testing.Nested
}
var file_aip_fieldbehavior_internal_testing_testing_proto_depIdxs = []int32{
	1, // 0: aip.fieldbehavior.internal.testing.Test.nested:type_name -> aip.fieldbehavior.internal.testing.Nested
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_aip_fieldbehavior_internal_testing_testing_proto_init() }
func file_aip_fieldbehavior_internal_testing_testing_proto_init() {
	if File_aip_fieldbehavior_internal_testing_testing_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_aip_fieldbehavior_internal_testing_testing_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_aip_fieldbehavior_internal_testing_testing_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_aip_fieldbehavior_internal_testing_testing_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_aip_fieldbehavior_internal_testing_testing_proto_goTypes,
		DependencyIndexes: file_aip_fieldbehavior_internal_testing_testing_proto_depIdxs,
		MessageInfos:      file_aip_fieldbehavior_internal_testing_testing_proto_msgTypes,
	}.Build()
	File_aip_fieldbehavior_internal_testing_testing_proto = out.File
	file_aip_fieldbehavior_internal_testing_testing_proto_rawDesc = nil
	file_aip_fieldbehavior_internal_testing_testing_proto_goTypes = nil
	file_aip_fieldbehavior_internal_testing_testing_proto_depIdxs = nil
}
