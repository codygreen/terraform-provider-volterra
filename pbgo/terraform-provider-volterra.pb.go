// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ves.io/terraform-provider-volterra/terraform-provider-volterra.proto

package pbgo

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
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
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

func init() {
	proto.RegisterFile("ves.io/terraform-provider-volterra/terraform-provider-volterra.proto", fileDescriptor_6ae8cbeecec2517a)
}

var fileDescriptor_6ae8cbeecec2517a = []byte{
	// 159 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x72, 0x29, 0x4b, 0x2d, 0xd6,
	0xcb, 0xcc, 0xd7, 0x2f, 0x49, 0x2d, 0x2a, 0x4a, 0x4c, 0xcb, 0x2f, 0xca, 0xd5, 0x2d, 0x28, 0xca,
	0x2f, 0xcb, 0x4c, 0x49, 0x2d, 0xd2, 0x2d, 0xcb, 0xcf, 0x01, 0x8b, 0xe2, 0x93, 0xd3, 0x2b, 0x28,
	0xca, 0x2f, 0xc9, 0x17, 0x52, 0x82, 0x98, 0xa2, 0x07, 0x57, 0x19, 0x0f, 0x53, 0x19, 0x0f, 0x53,
	0xe9, 0x94, 0x71, 0xe1, 0xa1, 0x1c, 0xc3, 0x8d, 0x87, 0x72, 0x0c, 0x1f, 0x1e, 0xca, 0x31, 0x36,
	0x3c, 0x92, 0x63, 0x5c, 0xf1, 0x48, 0x8e, 0xf1, 0xc4, 0x23, 0x39, 0xc6, 0x0b, 0x8f, 0xe4, 0x18,
	0x1f, 0x3c, 0x92, 0x63, 0x7c, 0xf1, 0x48, 0x8e, 0xe1, 0xc3, 0x23, 0x39, 0xc6, 0x09, 0x8f, 0xe5,
	0x18, 0x2e, 0x3c, 0x96, 0x63, 0xb8, 0xf1, 0x58, 0x8e, 0x21, 0xca, 0x28, 0x3d, 0xbf, 0x20, 0x3b,
	0x5d, 0x0f, 0x6e, 0x5f, 0x69, 0x31, 0x5e, 0xa7, 0x16, 0x24, 0xa5, 0xe7, 0x27, 0xb1, 0x81, 0x1d,
	0x65, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x39, 0x00, 0xde, 0xb6, 0xdc, 0x00, 0x00, 0x00,
}
