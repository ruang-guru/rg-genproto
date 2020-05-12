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

type Variant int32

const (
	Variant_ID                 Variant = 0
	Variant_HyphenatedID       Variant = 1
	Variant_Email              Variant = 2
	Variant_MacAddress         Variant = 3
	Variant_DomainName         Variant = 4
	Variant_UserName           Variant = 5
	Variant_URL                Variant = 6
	Variant_IPV4               Variant = 7
	Variant_IPV6               Variant = 8
	Variant_Password           Variant = 9
	Variant_Latitude           Variant = 10
	Variant_Longitude          Variant = 11
	Variant_CreditCardNumber   Variant = 12
	Variant_CreditCardType     Variant = 13
	Variant_PhoneNumber        Variant = 14
	Variant_TollFreeNumber     Variant = 15
	Variant_E164PhoneNumber    Variant = 16
	Variant_TitleMale          Variant = 17
	Variant_TitleFemale        Variant = 18
	Variant_FirstName          Variant = 19
	Variant_FirstNameMale      Variant = 20
	Variant_FirstNameFemale    Variant = 21
	Variant_LastName           Variant = 22
	Variant_Name               Variant = 23
	Variant_UnixTime           Variant = 24
	Variant_Date               Variant = 25
	Variant_Time               Variant = 26
	Variant_MonthName          Variant = 27
	Variant_Year               Variant = 28
	Variant_DayOfWeek          Variant = 29
	Variant_DayOfMonth         Variant = 30
	Variant_Timestamp          Variant = 31
	Variant_Century            Variant = 32
	Variant_Timezoe            Variant = 33
	Variant_TimePeriod         Variant = 34
	Variant_Word               Variant = 35
	Variant_Sentence           Variant = 36
	Variant_Paragraph          Variant = 37
	Variant_Currency           Variant = 38
	Variant_Amount             Variant = 39
	Variant_AmountWithCurrency Variant = 40
	Variant_Page               Variant = 41
	Variant_PageSize           Variant = 42
	Variant_PageTotal          Variant = 43
	Variant_Number             Variant = 44
)

var Variant_name = map[int32]string{
	0:  "ID",
	1:  "HyphenatedID",
	2:  "Email",
	3:  "MacAddress",
	4:  "DomainName",
	5:  "UserName",
	6:  "URL",
	7:  "IPV4",
	8:  "IPV6",
	9:  "Password",
	10: "Latitude",
	11: "Longitude",
	12: "CreditCardNumber",
	13: "CreditCardType",
	14: "PhoneNumber",
	15: "TollFreeNumber",
	16: "E164PhoneNumber",
	17: "TitleMale",
	18: "TitleFemale",
	19: "FirstName",
	20: "FirstNameMale",
	21: "FirstNameFemale",
	22: "LastName",
	23: "Name",
	24: "UnixTime",
	25: "Date",
	26: "Time",
	27: "MonthName",
	28: "Year",
	29: "DayOfWeek",
	30: "DayOfMonth",
	31: "Timestamp",
	32: "Century",
	33: "Timezoe",
	34: "TimePeriod",
	35: "Word",
	36: "Sentence",
	37: "Paragraph",
	38: "Currency",
	39: "Amount",
	40: "AmountWithCurrency",
	41: "Page",
	42: "PageSize",
	43: "PageTotal",
	44: "Number",
}

var Variant_value = map[string]int32{
	"ID":                 0,
	"HyphenatedID":       1,
	"Email":              2,
	"MacAddress":         3,
	"DomainName":         4,
	"UserName":           5,
	"URL":                6,
	"IPV4":               7,
	"IPV6":               8,
	"Password":           9,
	"Latitude":           10,
	"Longitude":          11,
	"CreditCardNumber":   12,
	"CreditCardType":     13,
	"PhoneNumber":        14,
	"TollFreeNumber":     15,
	"E164PhoneNumber":    16,
	"TitleMale":          17,
	"TitleFemale":        18,
	"FirstName":          19,
	"FirstNameMale":      20,
	"FirstNameFemale":    21,
	"LastName":           22,
	"Name":               23,
	"UnixTime":           24,
	"Date":               25,
	"Time":               26,
	"MonthName":          27,
	"Year":               28,
	"DayOfWeek":          29,
	"DayOfMonth":         30,
	"Timestamp":          31,
	"Century":            32,
	"Timezoe":            33,
	"TimePeriod":         34,
	"Word":               35,
	"Sentence":           36,
	"Paragraph":          37,
	"Currency":           38,
	"Amount":             39,
	"AmountWithCurrency": 40,
	"Page":               41,
	"PageSize":           42,
	"PageTotal":          43,
	"Number":             44,
}

func (x Variant) String() string {
	return proto.EnumName(Variant_name, int32(x))
}

func (Variant) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_3379825bd37f6cc2, []int{0}
}

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
	ExtensionType: (*Variant)(nil),
	Field:         820006,
	Name:          "validation.variant",
	Tag:           "varint,820006,opt,name=variant,enum=validation.Variant",
	Filename:      "extension/validation/validation.proto",
}

func init() {
	proto.RegisterEnum("validation.Variant", Variant_name, Variant_value)
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
	// 663 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x94, 0xdb, 0x4e, 0x1b, 0x49,
	0x10, 0x86, 0x97, 0x93, 0x0f, 0x0d, 0x98, 0xa2, 0x61, 0x59, 0x96, 0x5d, 0x76, 0xd9, 0x4d, 0x48,
	0x08, 0x09, 0xb6, 0x42, 0x10, 0x91, 0xe0, 0x8a, 0xd8, 0x58, 0x41, 0xe2, 0x60, 0x19, 0x03, 0x4a,
	0xee, 0xda, 0x9e, 0x62, 0xdc, 0xca, 0x4c, 0xf7, 0xa4, 0xa7, 0x87, 0x60, 0x1e, 0x20, 0x8f, 0x90,
	0xfb, 0x1c, 0x9f, 0x33, 0xaa, 0x9e, 0xb1, 0x21, 0x52, 0x24, 0xe7, 0xae, 0xfe, 0xaa, 0xff, 0xab,
	0x72, 0xbb, 0x6b, 0x9a, 0xad, 0xe2, 0xb5, 0x45, 0x15, 0x4b, 0xad, 0x2a, 0x57, 0x22, 0x90, 0x9e,
	0xb0, 0x3f, 0x86, 0xe5, 0xc8, 0x68, 0xab, 0x39, 0xbb, 0xcd, 0x2c, 0xad, 0xf8, 0x5a, 0xfb, 0x01,
	0x56, 0x5c, 0xa5, 0x9d, 0x5c, 0x56, 0x3c, 0x8c, 0x3b, 0x46, 0x46, 0x56, 0x9b, 0xd4, 0xbd, 0xfe,
	0x61, 0x82, 0xe5, 0xcf, 0x85, 0x91, 0x42, 0x59, 0x9e, 0x63, 0xa3, 0x07, 0x35, 0xf8, 0x8d, 0x03,
	0x9b, 0x7a, 0xd9, 0x8b, 0xba, 0xa8, 0x84, 0x45, 0xef, 0xa0, 0x06, 0x23, 0xbc, 0xc8, 0x26, 0xf6,
	0x43, 0x21, 0x03, 0x18, 0xe5, 0x25, 0xc6, 0x8e, 0x44, 0x67, 0xcf, 0xf3, 0x0c, 0xc6, 0x31, 0x8c,
	0x91, 0xae, 0xe9, 0x50, 0x48, 0x75, 0x2c, 0x42, 0x84, 0x71, 0x3e, 0xc5, 0x0a, 0x67, 0x31, 0x1a,
	0xa7, 0x26, 0x78, 0x9e, 0x8d, 0x9d, 0x35, 0x0f, 0x21, 0xc7, 0x0b, 0x6c, 0xfc, 0xa0, 0x71, 0xbe,
	0x05, 0xf9, 0x2c, 0xda, 0x86, 0x02, 0x59, 0x1b, 0x22, 0x8e, 0xdf, 0x69, 0xe3, 0x41, 0x91, 0xd4,
	0xa1, 0xb0, 0xd2, 0x26, 0x1e, 0x02, 0xe3, 0xd3, 0xac, 0x78, 0xa8, 0x95, 0x9f, 0xca, 0x49, 0x3e,
	0xcf, 0xa0, 0x6a, 0xd0, 0x93, 0xb6, 0x2a, 0x8c, 0x77, 0x9c, 0x84, 0x6d, 0x34, 0x30, 0xc5, 0x39,
	0x2b, 0xdd, 0x66, 0x5b, 0xbd, 0x08, 0x61, 0x9a, 0xcf, 0xb0, 0xc9, 0x46, 0x57, 0x2b, 0xcc, 0x4c,
	0x25, 0x32, 0xb5, 0x74, 0x10, 0xd4, 0x0d, 0xf6, 0x73, 0x33, 0x7c, 0x8e, 0xcd, 0xec, 0x3f, 0xdd,
	0xde, 0xba, 0x6b, 0x04, 0x1a, 0xd9, 0x92, 0x36, 0xc0, 0x23, 0x11, 0x20, 0xcc, 0x52, 0x23, 0x27,
	0xeb, 0x18, 0x52, 0x82, 0x53, 0xbd, 0x2e, 0x4d, 0x6c, 0xdd, 0xd1, 0xe6, 0xf8, 0x2c, 0x9b, 0x1e,
	0x48, 0x87, 0xcc, 0x53, 0xdb, 0x41, 0x2a, 0xc3, 0x7e, 0x4f, 0xcf, 0x95, 0x51, 0x0b, 0x74, 0x7a,
	0x17, 0xfd, 0xe1, 0xfe, 0x28, 0x25, 0xaf, 0x5b, 0x32, 0x44, 0x58, 0xa4, 0x7c, 0x4d, 0x58, 0x84,
	0x3f, 0x29, 0x72, 0xb9, 0x25, 0x1a, 0x78, 0xa4, 0x95, 0xed, 0x3a, 0xe0, 0x2f, 0x2a, 0xbc, 0x42,
	0x61, 0xe0, 0x6f, 0x2a, 0xd4, 0x44, 0xef, 0xe4, 0xf2, 0x02, 0xf1, 0x0d, 0x2c, 0xbb, 0x2b, 0x20,
	0xe9, 0xcc, 0xf0, 0x4f, 0x7a, 0x90, 0x10, 0x63, 0x2b, 0xc2, 0x08, 0xfe, 0xe5, 0x93, 0x2c, 0x5f,
	0x45, 0x65, 0x13, 0xd3, 0x83, 0x15, 0x12, 0x54, 0xbb, 0xd1, 0x08, 0xff, 0x11, 0x48, 0xa2, 0x81,
	0x46, 0x6a, 0x0f, 0xfe, 0xa7, 0x09, 0x17, 0x74, 0x19, 0xf7, 0xe8, 0xc7, 0x9d, 0xa2, 0xb2, 0xa8,
	0x3a, 0x08, 0xf7, 0xa9, 0x61, 0x43, 0x18, 0xe1, 0x1b, 0x11, 0x75, 0x61, 0x95, 0x8a, 0xd5, 0xc4,
	0x18, 0x54, 0x9d, 0x1e, 0x3c, 0xe0, 0x8c, 0xe5, 0xf6, 0x42, 0x9d, 0x28, 0x0b, 0x0f, 0xf9, 0x02,
	0xe3, 0x69, 0x7c, 0x21, 0x6d, 0x77, 0xe0, 0x59, 0xa3, 0xc6, 0x0d, 0xe1, 0x23, 0x3c, 0x4a, 0xef,
	0xdc, 0xc7, 0x53, 0x79, 0x83, 0xb0, 0x9e, 0x36, 0xf6, 0xb1, 0xa5, 0xad, 0x08, 0xe0, 0x31, 0xb5,
	0xca, 0x6e, 0xe3, 0xc9, 0xce, 0x2e, 0x2b, 0x18, 0x7c, 0x9b, 0x48, 0x83, 0x1e, 0x5f, 0x2e, 0xa7,
	0x7b, 0x5c, 0xee, 0xef, 0x71, 0xb9, 0x2e, 0x31, 0xf0, 0x4e, 0x22, 0x5a, 0xf2, 0x78, 0xf1, 0xe3,
	0xfb, 0xcd, 0x95, 0x91, 0xb5, 0x42, 0x73, 0x00, 0x10, 0x4c, 0x6b, 0x21, 0xda, 0x01, 0x0e, 0x83,
	0x3f, 0xf5, 0xe1, 0x3e, 0xb0, 0xf3, 0x9c, 0xe5, 0xba, 0xd2, 0xf3, 0x50, 0x0d, 0x43, 0x3f, 0x67,
	0x68, 0x66, 0xa7, 0xa9, 0x2a, 0x09, 0x82, 0x5f, 0x99, 0xfa, 0xa5, 0x3f, 0xb5, 0x0f, 0x10, 0xec,
	0x09, 0x2b, 0x6c, 0x2f, 0x1a, 0x0a, 0x7f, 0x75, 0x70, 0xb1, 0x39, 0x00, 0x76, 0x8e, 0x59, 0xfe,
	0x2a, 0xfb, 0x88, 0x87, 0xb0, 0xdf, 0x1c, 0x5b, 0xda, 0x9c, 0x2b, 0xdf, 0x79, 0x38, 0xb2, 0x07,
	0xa0, 0xd9, 0x6f, 0xf2, 0xa2, 0xfa, 0x7a, 0xcf, 0x97, 0xb6, 0x9b, 0xb4, 0xcb, 0x1d, 0x1d, 0x56,
	0x4c, 0x22, 0x94, 0xbf, 0xe1, 0x27, 0x26, 0xa9, 0x18, 0x7f, 0xc3, 0x47, 0xe5, 0x5a, 0x57, 0x7e,
	0xf6, 0x1c, 0xed, 0xde, 0x86, 0xed, 0x9c, 0xb3, 0x3d, 0xfb, 0x1e, 0x00, 0x00, 0xff, 0xff, 0x45,
	0x5f, 0x3f, 0xed, 0xb8, 0x04, 0x00, 0x00,
}
