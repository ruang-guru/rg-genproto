// Code generated by protoc-gen-go. DO NOT EDIT.
// source: extension/validation/validation.proto

package validation

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	descriptor "github.com/golang/protobuf/protoc-gen-go/descriptor"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

var E_Required = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FieldOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         820001,
	Name:          "validation.required",
	Tag:           "varint,820001,opt,name=required",
	Filename:      "extension/validation/validation.proto",
}

var E_Editable = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FieldOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         820002,
	Name:          "validation.editable",
	Tag:           "varint,820002,opt,name=editable",
	Filename:      "extension/validation/validation.proto",
}

var E_Hidden = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FieldOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         820003,
	Name:          "validation.hidden",
	Tag:           "varint,820003,opt,name=hidden",
	Filename:      "extension/validation/validation.proto",
}

var E_Nullable = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FieldOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         820004,
	Name:          "validation.nullable",
	Tag:           "varint,820004,opt,name=nullable",
	Filename:      "extension/validation/validation.proto",
}

var E_Datatype = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FieldOptions)(nil),
	ExtensionType: (*string)(nil),
	Field:         820005,
	Name:          "validation.datatype",
	Tag:           "bytes,820005,opt,name=datatype",
	Filename:      "extension/validation/validation.proto",
}

var E_Variant = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FieldOptions)(nil),
	ExtensionType: (*string)(nil),
	Field:         820006,
	Name:          "validation.variant",
	Tag:           "bytes,820006,opt,name=variant",
	Filename:      "extension/validation/validation.proto",
}

func init() {
	proto.RegisterExtension(E_Required)
	proto.RegisterExtension(E_Editable)
	proto.RegisterExtension(E_Hidden)
	proto.RegisterExtension(E_Nullable)
	proto.RegisterExtension(E_Datatype)
	proto.RegisterExtension(E_Variant)
}

func init() {
	proto.RegisterFile("extension/validation/validation.proto", fileDescriptor_3379825bd37f6cc2)
}

var fileDescriptor_3379825bd37f6cc2 = []byte{
	// 248 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0xd1, 0xbd, 0x4a, 0x03, 0x41,
	0x10, 0xc0, 0x71, 0x6c, 0x62, 0xdc, 0x32, 0x95, 0x08, 0x42, 0x1a, 0xc1, 0x26, 0xbb, 0xa0, 0x85,
	0x78, 0xa9, 0x54, 0xb0, 0x15, 0x52, 0xda, 0xed, 0x65, 0xc6, 0xcd, 0xc0, 0xba, 0x7b, 0xce, 0xcd,
	0x06, 0x7d, 0x01, 0xdf, 0xc3, 0xcf, 0xe7, 0x94, 0xbb, 0xdc, 0x26, 0x08, 0xc2, 0xda, 0x0d, 0x77,
	0xff, 0xdf, 0x4c, 0xb1, 0xea, 0x04, 0x9f, 0x05, 0x43, 0x4b, 0x31, 0x98, 0xb5, 0xf5, 0x04, 0x56,
	0x7e, 0x8f, 0xba, 0xe1, 0x28, 0x71, 0xa2, 0x76, 0x5f, 0x8e, 0xa6, 0x2e, 0x46, 0xe7, 0xd1, 0xf4,
	0x7f, 0xea, 0xf4, 0x60, 0x00, 0xdb, 0x25, 0x53, 0x23, 0x91, 0x37, 0x75, 0x35, 0x57, 0x63, 0xc6,
	0xa7, 0x44, 0x8c, 0x30, 0x39, 0xd6, 0x9b, 0x5c, 0xe7, 0x5c, 0xdf, 0x12, 0x7a, 0xb8, 0x6b, 0xba,
	0x5d, 0xed, 0xe1, 0xdb, 0xeb, 0xd9, 0x74, 0xef, 0x74, 0xbc, 0xd8, 0x82, 0x0e, 0x23, 0x90, 0xd8,
	0xda, 0x63, 0x09, 0xbf, 0x67, 0x9c, 0x41, 0x75, 0xa1, 0x46, 0x2b, 0x02, 0xc0, 0x50, 0xa2, 0x1f,
	0x03, 0x1d, 0xf2, 0xee, 0x6a, 0x48, 0xde, 0xff, 0xe7, 0xea, 0x67, 0xbe, 0x9a, 0x41, 0x87, 0xc1,
	0x8a, 0x95, 0x97, 0xa6, 0x88, 0xbf, 0x7a, 0x7c, 0xb0, 0xd8, 0x82, 0xea, 0x52, 0xed, 0xaf, 0x2d,
	0x93, 0x0d, 0x52, 0xb2, 0xdf, 0x83, 0xcd, 0xfd, 0xf5, 0xcd, 0xfd, 0x95, 0x23, 0x59, 0xa5, 0x5a,
	0x2f, 0xe3, 0xa3, 0xe1, 0x64, 0x83, 0x9b, 0xb9, 0xc4, 0xc9, 0xb0, 0x9b, 0x39, 0x0c, 0xfd, 0x16,
	0xf3, 0xd7, 0x03, 0xcf, 0x77, 0x63, 0x3d, 0xea, 0xb3, 0xf3, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff,
	0xe2, 0x18, 0x4f, 0xae, 0x0a, 0x02, 0x00, 0x00,
}
