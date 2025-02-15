// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        (unknown)
// source: tests/middleware/protovalidate/v1/person.proto

package protovalidatepb

import (
	_ "buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
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

type Person struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    uint64       `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Email string       `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	Name  string       `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Home  *Coordinates `protobuf:"bytes,4,opt,name=home,proto3" json:"home,omitempty"`
}

func (x *Person) Reset() {
	*x = Person{}
	mi := &file_tests_middleware_protovalidate_v1_person_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Person) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Person) ProtoMessage() {}

func (x *Person) ProtoReflect() protoreflect.Message {
	mi := &file_tests_middleware_protovalidate_v1_person_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Person.ProtoReflect.Descriptor instead.
func (*Person) Descriptor() ([]byte, []int) {
	return file_tests_middleware_protovalidate_v1_person_proto_rawDescGZIP(), []int{0}
}

func (x *Person) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Person) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *Person) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Person) GetHome() *Coordinates {
	if x != nil {
		return x.Home
	}
	return nil
}

type Coordinates struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Lat float64 `protobuf:"fixed64,1,opt,name=lat,proto3" json:"lat,omitempty"`
	Lng float64 `protobuf:"fixed64,2,opt,name=lng,proto3" json:"lng,omitempty"`
}

func (x *Coordinates) Reset() {
	*x = Coordinates{}
	mi := &file_tests_middleware_protovalidate_v1_person_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Coordinates) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Coordinates) ProtoMessage() {}

func (x *Coordinates) ProtoReflect() protoreflect.Message {
	mi := &file_tests_middleware_protovalidate_v1_person_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Coordinates.ProtoReflect.Descriptor instead.
func (*Coordinates) Descriptor() ([]byte, []int) {
	return file_tests_middleware_protovalidate_v1_person_proto_rawDescGZIP(), []int{1}
}

func (x *Coordinates) GetLat() float64 {
	if x != nil {
		return x.Lat
	}
	return 0
}

func (x *Coordinates) GetLng() float64 {
	if x != nil {
		return x.Lng
	}
	return 0
}

var File_tests_middleware_protovalidate_v1_person_proto protoreflect.FileDescriptor

var file_tests_middleware_protovalidate_v1_person_proto_rawDesc = []byte{
	0x0a, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x73, 0x2f, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61,
	0x72, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65,
	0x2f, 0x76, 0x31, 0x2f, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x21, 0x74, 0x65, 0x73, 0x74, 0x73, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61,
	0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65,
	0x2e, 0x76, 0x31, 0x1a, 0x1b, 0x62, 0x75, 0x66, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xc3, 0x01, 0x0a, 0x06, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x42, 0x08, 0xba, 0x48, 0x05, 0x32, 0x03, 0x20, 0xe7,
	0x07, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1d, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xba, 0x48, 0x04, 0x72, 0x02, 0x60, 0x01, 0x52, 0x05, 0x65,
	0x6d, 0x61, 0x69, 0x6c, 0x12, 0x3c, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x28, 0xba, 0x48, 0x25, 0x72, 0x23, 0x28, 0x80, 0x02, 0x32, 0x1e, 0x5e, 0x5b,
	0x5b, 0x3a, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x3a, 0x5d, 0x5d, 0x2b, 0x28, 0x20, 0x5b, 0x5b, 0x3a,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x3a, 0x5d, 0x5d, 0x2b, 0x29, 0x2a, 0x24, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x42, 0x0a, 0x04, 0x68, 0x6f, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x2e, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x73, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77,
	0x61, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x73,
	0x52, 0x04, 0x68, 0x6f, 0x6d, 0x65, 0x22, 0x63, 0x0a, 0x0b, 0x43, 0x6f, 0x6f, 0x72, 0x64, 0x69,
	0x6e, 0x61, 0x74, 0x65, 0x73, 0x12, 0x29, 0x0a, 0x03, 0x6c, 0x61, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x01, 0x42, 0x17, 0xba, 0x48, 0x14, 0x12, 0x12, 0x19, 0x00, 0x00, 0x00, 0x00, 0x00, 0x80,
	0x56, 0x40, 0x29, 0x00, 0x00, 0x00, 0x00, 0x00, 0x80, 0x56, 0xc0, 0x52, 0x03, 0x6c, 0x61, 0x74,
	0x12, 0x29, 0x0a, 0x03, 0x6c, 0x6e, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x42, 0x17, 0xba,
	0x48, 0x14, 0x12, 0x12, 0x19, 0x00, 0x00, 0x00, 0x00, 0x00, 0x80, 0x66, 0x40, 0x29, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x80, 0x66, 0xc0, 0x52, 0x03, 0x6c, 0x6e, 0x67, 0x42, 0x33, 0x5a, 0x31, 0x74,
	0x65, 0x73, 0x74, 0x73, 0x2f, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x31,
	0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x70, 0x62,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_tests_middleware_protovalidate_v1_person_proto_rawDescOnce sync.Once
	file_tests_middleware_protovalidate_v1_person_proto_rawDescData = file_tests_middleware_protovalidate_v1_person_proto_rawDesc
)

func file_tests_middleware_protovalidate_v1_person_proto_rawDescGZIP() []byte {
	file_tests_middleware_protovalidate_v1_person_proto_rawDescOnce.Do(func() {
		file_tests_middleware_protovalidate_v1_person_proto_rawDescData = protoimpl.X.CompressGZIP(file_tests_middleware_protovalidate_v1_person_proto_rawDescData)
	})
	return file_tests_middleware_protovalidate_v1_person_proto_rawDescData
}

var file_tests_middleware_protovalidate_v1_person_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_tests_middleware_protovalidate_v1_person_proto_goTypes = []any{
	(*Person)(nil),      // 0: tests.middleware.protovalidate.v1.Person
	(*Coordinates)(nil), // 1: tests.middleware.protovalidate.v1.Coordinates
}
var file_tests_middleware_protovalidate_v1_person_proto_depIdxs = []int32{
	1, // 0: tests.middleware.protovalidate.v1.Person.home:type_name -> tests.middleware.protovalidate.v1.Coordinates
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_tests_middleware_protovalidate_v1_person_proto_init() }
func file_tests_middleware_protovalidate_v1_person_proto_init() {
	if File_tests_middleware_protovalidate_v1_person_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_tests_middleware_protovalidate_v1_person_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_tests_middleware_protovalidate_v1_person_proto_goTypes,
		DependencyIndexes: file_tests_middleware_protovalidate_v1_person_proto_depIdxs,
		MessageInfos:      file_tests_middleware_protovalidate_v1_person_proto_msgTypes,
	}.Build()
	File_tests_middleware_protovalidate_v1_person_proto = out.File
	file_tests_middleware_protovalidate_v1_person_proto_rawDesc = nil
	file_tests_middleware_protovalidate_v1_person_proto_goTypes = nil
	file_tests_middleware_protovalidate_v1_person_proto_depIdxs = nil
}
