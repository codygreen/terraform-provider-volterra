// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ves.io/schema/cloud_credentials/object.proto

package cloud_credentials

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	golang_proto "github.com/golang/protobuf/proto"
	schema "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema"
	_ "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/vesenv"
	io "io"
	math "math"
	math_bits "math/bits"
	reflect "reflect"
	strings "strings"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = golang_proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// Cloud Credentials object
//
// x-displayName: "Cloud Credentials"
// Cloud Credentials object
type Object struct {
	// metadata
	//
	// x-displayName: "Metadata"
	// Standard object's metadata
	Metadata *schema.ObjectMetaType `protobuf:"bytes,1,opt,name=metadata,proto3" json:"metadata,omitempty"`
	// system_metadata
	//
	// x-displayName: "System Metadata"
	// System generated object's metadata
	SystemMetadata *schema.SystemObjectMetaType `protobuf:"bytes,2,opt,name=system_metadata,json=systemMetadata,proto3" json:"system_metadata,omitempty"`
	// spec
	//
	// x-displayName: "Spec"
	// Specification of the cloud_credentials
	Spec *SpecType `protobuf:"bytes,3,opt,name=spec,proto3" json:"spec,omitempty"`
}

func (m *Object) Reset()      { *m = Object{} }
func (*Object) ProtoMessage() {}
func (*Object) Descriptor() ([]byte, []int) {
	return fileDescriptor_b3d3e0f702922184, []int{0}
}
func (m *Object) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Object) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *Object) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Object.Merge(m, src)
}
func (m *Object) XXX_Size() int {
	return m.Size()
}
func (m *Object) XXX_DiscardUnknown() {
	xxx_messageInfo_Object.DiscardUnknown(m)
}

var xxx_messageInfo_Object proto.InternalMessageInfo

func (m *Object) GetMetadata() *schema.ObjectMetaType {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *Object) GetSystemMetadata() *schema.SystemObjectMetaType {
	if m != nil {
		return m.SystemMetadata
	}
	return nil
}

func (m *Object) GetSpec() *SpecType {
	if m != nil {
		return m.Spec
	}
	return nil
}

// Specification for Cloud Credentials object
//
// x-displayName: "Specification"
// Shape of the cloud_credentials specification
type SpecType struct {
	// gc_spec
	//
	// x-displayName: "GC Spec"
	// desired state of cloud_credentials
	GcSpec *GlobalSpecType `protobuf:"bytes,2,opt,name=gc_spec,json=gcSpec,proto3" json:"gc_spec,omitempty"`
}

func (m *SpecType) Reset()      { *m = SpecType{} }
func (*SpecType) ProtoMessage() {}
func (*SpecType) Descriptor() ([]byte, []int) {
	return fileDescriptor_b3d3e0f702922184, []int{1}
}
func (m *SpecType) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SpecType) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *SpecType) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SpecType.Merge(m, src)
}
func (m *SpecType) XXX_Size() int {
	return m.Size()
}
func (m *SpecType) XXX_DiscardUnknown() {
	xxx_messageInfo_SpecType.DiscardUnknown(m)
}

var xxx_messageInfo_SpecType proto.InternalMessageInfo

func (m *SpecType) GetGcSpec() *GlobalSpecType {
	if m != nil {
		return m.GcSpec
	}
	return nil
}

// Cloud Credentials Status Object
//
// x-displayName: "Cloud Credentials Status"
// Most recently observed status of object
type StatusObject struct {
	// metadata
	//
	// x-displayName: "Metadata"
	// Standard status's metadata
	Metadata *schema.StatusMetaType `protobuf:"bytes,1,opt,name=metadata,proto3" json:"metadata,omitempty"`
	// object_refs
	//
	// x-displayName: "Config Object"
	// Object reference
	ObjectRefs []*schema.ObjectRefType `protobuf:"bytes,2,rep,name=object_refs,json=objectRefs,proto3" json:"object_refs,omitempty"`
	// conditions
	//
	// x-displayName: "Conditions"
	// Conditions
	Conditions []*schema.ConditionType `protobuf:"bytes,4,rep,name=conditions,proto3" json:"conditions,omitempty"`
}

func (m *StatusObject) Reset()      { *m = StatusObject{} }
func (*StatusObject) ProtoMessage() {}
func (*StatusObject) Descriptor() ([]byte, []int) {
	return fileDescriptor_b3d3e0f702922184, []int{2}
}
func (m *StatusObject) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *StatusObject) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *StatusObject) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StatusObject.Merge(m, src)
}
func (m *StatusObject) XXX_Size() int {
	return m.Size()
}
func (m *StatusObject) XXX_DiscardUnknown() {
	xxx_messageInfo_StatusObject.DiscardUnknown(m)
}

var xxx_messageInfo_StatusObject proto.InternalMessageInfo

func (m *StatusObject) GetMetadata() *schema.StatusMetaType {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *StatusObject) GetObjectRefs() []*schema.ObjectRefType {
	if m != nil {
		return m.ObjectRefs
	}
	return nil
}

func (m *StatusObject) GetConditions() []*schema.ConditionType {
	if m != nil {
		return m.Conditions
	}
	return nil
}

func init() {
	proto.RegisterType((*Object)(nil), "ves.io.schema.cloud_credentials.Object")
	golang_proto.RegisterType((*Object)(nil), "ves.io.schema.cloud_credentials.Object")
	proto.RegisterType((*SpecType)(nil), "ves.io.schema.cloud_credentials.SpecType")
	golang_proto.RegisterType((*SpecType)(nil), "ves.io.schema.cloud_credentials.SpecType")
	proto.RegisterType((*StatusObject)(nil), "ves.io.schema.cloud_credentials.StatusObject")
	golang_proto.RegisterType((*StatusObject)(nil), "ves.io.schema.cloud_credentials.StatusObject")
}

func init() {
	proto.RegisterFile("ves.io/schema/cloud_credentials/object.proto", fileDescriptor_b3d3e0f702922184)
}
func init() {
	golang_proto.RegisterFile("ves.io/schema/cloud_credentials/object.proto", fileDescriptor_b3d3e0f702922184)
}

var fileDescriptor_b3d3e0f702922184 = []byte{
	// 646 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x54, 0xcf, 0x4f, 0xd4, 0x4e,
	0x14, 0xdf, 0x61, 0xf9, 0xee, 0x97, 0x0c, 0x84, 0x2f, 0xdf, 0xd5, 0x43, 0x45, 0x1d, 0xc9, 0x7a,
	0xc1, 0x48, 0xa7, 0x0a, 0x5e, 0x24, 0xfe, 0x48, 0xf0, 0xa0, 0x07, 0x89, 0x09, 0x8b, 0x17, 0x2f,
	0x4d, 0x77, 0xfa, 0x5a, 0x46, 0xba, 0x9d, 0x66, 0x66, 0x5a, 0x58, 0x13, 0x13, 0xe3, 0x5f, 0x60,
	0xe2, 0xdd, 0x9b, 0x89, 0xf1, 0xaf, 0x10, 0x4f, 0x46, 0x2f, 0xc4, 0x13, 0xf1, 0x24, 0xdd, 0x8b,
	0x47, 0xfc, 0x0f, 0xcc, 0x4e, 0xb7, 0xc0, 0xc2, 0x12, 0x0c, 0xa7, 0xbe, 0x97, 0xf7, 0xf9, 0xf1,
	0xa6, 0xf3, 0xe6, 0xe1, 0xb9, 0x0c, 0x14, 0xe5, 0xc2, 0x51, 0x6c, 0x0d, 0xda, 0x9e, 0xc3, 0x22,
	0x91, 0xfa, 0x2e, 0x93, 0xe0, 0x43, 0xac, 0xb9, 0x17, 0x29, 0x47, 0xb4, 0x9e, 0x03, 0xd3, 0x34,
	0x91, 0x42, 0x8b, 0xfa, 0x95, 0x02, 0x4d, 0x0b, 0x34, 0x3d, 0x86, 0x9e, 0xb6, 0x43, 0xae, 0xd7,
	0xd2, 0x16, 0x65, 0xa2, 0xed, 0x84, 0x22, 0x14, 0x8e, 0xe1, 0xb5, 0xd2, 0xc0, 0x64, 0x26, 0x31,
	0x51, 0xa1, 0x37, 0x7d, 0xfd, 0x34, 0x77, 0xdd, 0x49, 0x40, 0xf5, 0xc1, 0x17, 0x07, 0xc1, 0x22,
	0xd1, 0x5c, 0xc4, 0x65, 0xf1, 0xc2, 0x60, 0xf1, 0x30, 0xaf, 0x31, 0x58, 0xca, 0x40, 0x41, 0x9c,
	0x0d, 0xd2, 0x1b, 0xef, 0xfe, 0xc1, 0xb5, 0x27, 0xe6, 0xa4, 0xf5, 0xdb, 0x78, 0xac, 0x0d, 0xda,
	0xf3, 0x3d, 0xed, 0x59, 0x68, 0x06, 0xcd, 0x8e, 0xcf, 0x5f, 0xa6, 0x83, 0xc7, 0x2e, 0x80, 0xcb,
	0xa0, 0xbd, 0xd5, 0x4e, 0x02, 0x2b, 0xfb, 0xf0, 0xfa, 0x63, 0xfc, 0x9f, 0xea, 0x28, 0x0d, 0x6d,
	0x77, 0x5f, 0x61, 0xc4, 0x28, 0x5c, 0x3d, 0xa2, 0xd0, 0x34, 0xa8, 0x23, 0x3a, 0x93, 0x05, 0x77,
	0xb9, 0x54, 0xbb, 0x8b, 0x47, 0x55, 0x02, 0xcc, 0xaa, 0x1a, 0x89, 0x6b, 0xf4, 0x94, 0x7f, 0x4f,
	0x9b, 0x09, 0x30, 0x23, 0x64, 0x68, 0x8b, 0xdf, 0xaa, 0xaf, 0xef, 0xd7, 0x0a, 0xc9, 0xcf, 0x5b,
	0xd6, 0xc8, 0x14, 0xfa, 0xba, 0x65, 0x4d, 0x66, 0xa0, 0x6c, 0x2e, 0xec, 0x10, 0x62, 0x90, 0x5e,
	0xf4, 0x63, 0xcb, 0x7a, 0x5f, 0xc5, 0x0c, 0xcf, 0xf5, 0x08, 0x34, 0x64, 0xae, 0xf9, 0x1a, 0x65,
	0xea, 0x6d, 0x28, 0x57, 0x01, 0x93, 0xa0, 0xdd, 0x75, 0xe8, 0xd0, 0x83, 0xb0, 0xbe, 0xd0, 0x17,
	0xc9, 0x38, 0x6c, 0xd8, 0x8a, 0x6b, 0xb0, 0x7b, 0x34, 0x3b, 0x64, 0xc5, 0xd7, 0x87, 0x24, 0x12,
	0x9d, 0x36, 0xc4, 0xda, 0x36, 0x4a, 0x76, 0xaf, 0x47, 0xbc, 0x86, 0x6f, 0x0e, 0x33, 0x79, 0x91,
	0x4a, 0x70, 0x93, 0x60, 0xd3, 0x65, 0x20, 0x35, 0x0f, 0x38, 0xf3, 0x34, 0xd0, 0xc4, 0x53, 0x6a,
	0x43, 0x48, 0xff, 0x6c, 0x4e, 0xeb, 0xf8, 0xd6, 0x89, 0x4e, 0x2c, 0xe2, 0x10, 0xeb, 0xfe, 0xb9,
	0xe8, 0x40, 0x76, 0x36, 0xb3, 0x10, 0xdf, 0x18, 0x62, 0x16, 0xb2, 0xc4, 0xdc, 0x8c, 0x1b, 0xf0,
	0x08, 0xe8, 0xc1, 0x1d, 0x99, 0xfc, 0x4c, 0x46, 0x8d, 0x55, 0x3c, 0x56, 0xde, 0x6f, 0xfd, 0x11,
	0xfe, 0xb7, 0xef, 0xd7, 0x1f, 0x2f, 0xe7, 0xd4, 0xd9, 0x78, 0x18, 0x89, 0x96, 0x17, 0xed, 0x4f,
	0x48, 0x2d, 0x64, 0xbd, 0xb8, 0xf1, 0x1b, 0xe1, 0x89, 0xa6, 0xf6, 0x74, 0xaa, 0xfe, 0x7a, 0xf8,
	0x0b, 0xf8, 0x90, 0xe1, 0x6f, 0xe2, 0xf1, 0x62, 0x57, 0xb8, 0x12, 0x02, 0x65, 0x8d, 0xcc, 0x54,
	0x67, 0xc7, 0xe7, 0x2f, 0x0d, 0x7d, 0x3a, 0x2b, 0x10, 0xf4, 0xc8, 0x4b, 0xe7, 0x3f, 0xbe, 0xfc,
	0xff, 0x58, 0xaf, 0x2b, 0x58, 0x94, 0x20, 0x55, 0xbf, 0x83, 0x31, 0x13, 0xb1, 0xcf, 0xcd, 0x5b,
	0xb5, 0x46, 0x87, 0x6a, 0x3e, 0x28, 0x01, 0xa6, 0xa1, 0x43, 0xf8, 0xc5, 0x73, 0xdf, 0xef, 0x4d,
	0xe1, 0x49, 0x3c, 0x51, 0xf6, 0x48, 0x53, 0xee, 0x2f, 0xbd, 0x45, 0xdb, 0xbb, 0xa4, 0xb2, 0xb3,
	0x4b, 0x2a, 0x7b, 0xbb, 0x04, 0xbd, 0xca, 0x09, 0xfa, 0x90, 0x13, 0xf4, 0x25, 0x27, 0x68, 0x3b,
	0x27, 0x68, 0x27, 0x27, 0xe8, 0x67, 0x4e, 0xd0, 0xaf, 0x9c, 0x54, 0xf6, 0x72, 0x82, 0xde, 0x74,
	0x49, 0xe5, 0x53, 0x97, 0xa0, 0xed, 0x2e, 0xa9, 0xec, 0x74, 0x49, 0xe5, 0xd9, 0xd3, 0x50, 0x24,
	0xeb, 0x21, 0xcd, 0x44, 0xa4, 0x41, 0x4a, 0x8f, 0xa6, 0xca, 0x31, 0x41, 0x20, 0x64, 0xdb, 0x4e,
	0xa4, 0xc8, 0xb8, 0x0f, 0xd2, 0x2e, 0xcb, 0x4e, 0xd2, 0x0a, 0x85, 0x03, 0x9b, 0xba, 0xbf, 0x7c,
	0x4e, 0x5a, 0x74, 0xad, 0x9a, 0xd9, 0x43, 0x0b, 0x7f, 0x02, 0x00, 0x00, 0xff, 0xff, 0x96, 0xef,
	0xa1, 0x5e, 0x90, 0x05, 0x00, 0x00,
}

func (this *Object) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Object)
	if !ok {
		that2, ok := that.(Object)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.Metadata.Equal(that1.Metadata) {
		return false
	}
	if !this.SystemMetadata.Equal(that1.SystemMetadata) {
		return false
	}
	if !this.Spec.Equal(that1.Spec) {
		return false
	}
	return true
}
func (this *SpecType) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*SpecType)
	if !ok {
		that2, ok := that.(SpecType)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.GcSpec.Equal(that1.GcSpec) {
		return false
	}
	return true
}
func (this *StatusObject) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*StatusObject)
	if !ok {
		that2, ok := that.(StatusObject)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.Metadata.Equal(that1.Metadata) {
		return false
	}
	if len(this.ObjectRefs) != len(that1.ObjectRefs) {
		return false
	}
	for i := range this.ObjectRefs {
		if !this.ObjectRefs[i].Equal(that1.ObjectRefs[i]) {
			return false
		}
	}
	if len(this.Conditions) != len(that1.Conditions) {
		return false
	}
	for i := range this.Conditions {
		if !this.Conditions[i].Equal(that1.Conditions[i]) {
			return false
		}
	}
	return true
}
func (this *Object) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 7)
	s = append(s, "&cloud_credentials.Object{")
	if this.Metadata != nil {
		s = append(s, "Metadata: "+fmt.Sprintf("%#v", this.Metadata)+",\n")
	}
	if this.SystemMetadata != nil {
		s = append(s, "SystemMetadata: "+fmt.Sprintf("%#v", this.SystemMetadata)+",\n")
	}
	if this.Spec != nil {
		s = append(s, "Spec: "+fmt.Sprintf("%#v", this.Spec)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *SpecType) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 5)
	s = append(s, "&cloud_credentials.SpecType{")
	if this.GcSpec != nil {
		s = append(s, "GcSpec: "+fmt.Sprintf("%#v", this.GcSpec)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *StatusObject) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 7)
	s = append(s, "&cloud_credentials.StatusObject{")
	if this.Metadata != nil {
		s = append(s, "Metadata: "+fmt.Sprintf("%#v", this.Metadata)+",\n")
	}
	if this.ObjectRefs != nil {
		s = append(s, "ObjectRefs: "+fmt.Sprintf("%#v", this.ObjectRefs)+",\n")
	}
	if this.Conditions != nil {
		s = append(s, "Conditions: "+fmt.Sprintf("%#v", this.Conditions)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringObject(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func (m *Object) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Object) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Object) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Spec != nil {
		{
			size, err := m.Spec.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintObject(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if m.SystemMetadata != nil {
		{
			size, err := m.SystemMetadata.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintObject(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if m.Metadata != nil {
		{
			size, err := m.Metadata.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintObject(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *SpecType) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SpecType) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SpecType) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.GcSpec != nil {
		{
			size, err := m.GcSpec.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintObject(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	return len(dAtA) - i, nil
}

func (m *StatusObject) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *StatusObject) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *StatusObject) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Conditions) > 0 {
		for iNdEx := len(m.Conditions) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Conditions[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintObject(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.ObjectRefs) > 0 {
		for iNdEx := len(m.ObjectRefs) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ObjectRefs[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintObject(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if m.Metadata != nil {
		{
			size, err := m.Metadata.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintObject(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintObject(dAtA []byte, offset int, v uint64) int {
	offset -= sovObject(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Object) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Metadata != nil {
		l = m.Metadata.Size()
		n += 1 + l + sovObject(uint64(l))
	}
	if m.SystemMetadata != nil {
		l = m.SystemMetadata.Size()
		n += 1 + l + sovObject(uint64(l))
	}
	if m.Spec != nil {
		l = m.Spec.Size()
		n += 1 + l + sovObject(uint64(l))
	}
	return n
}

func (m *SpecType) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.GcSpec != nil {
		l = m.GcSpec.Size()
		n += 1 + l + sovObject(uint64(l))
	}
	return n
}

func (m *StatusObject) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Metadata != nil {
		l = m.Metadata.Size()
		n += 1 + l + sovObject(uint64(l))
	}
	if len(m.ObjectRefs) > 0 {
		for _, e := range m.ObjectRefs {
			l = e.Size()
			n += 1 + l + sovObject(uint64(l))
		}
	}
	if len(m.Conditions) > 0 {
		for _, e := range m.Conditions {
			l = e.Size()
			n += 1 + l + sovObject(uint64(l))
		}
	}
	return n
}

func sovObject(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozObject(x uint64) (n int) {
	return sovObject(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *Object) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&Object{`,
		`Metadata:` + strings.Replace(fmt.Sprintf("%v", this.Metadata), "ObjectMetaType", "schema.ObjectMetaType", 1) + `,`,
		`SystemMetadata:` + strings.Replace(fmt.Sprintf("%v", this.SystemMetadata), "SystemObjectMetaType", "schema.SystemObjectMetaType", 1) + `,`,
		`Spec:` + strings.Replace(this.Spec.String(), "SpecType", "SpecType", 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *SpecType) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&SpecType{`,
		`GcSpec:` + strings.Replace(fmt.Sprintf("%v", this.GcSpec), "GlobalSpecType", "GlobalSpecType", 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *StatusObject) String() string {
	if this == nil {
		return "nil"
	}
	repeatedStringForObjectRefs := "[]*ObjectRefType{"
	for _, f := range this.ObjectRefs {
		repeatedStringForObjectRefs += strings.Replace(fmt.Sprintf("%v", f), "ObjectRefType", "schema.ObjectRefType", 1) + ","
	}
	repeatedStringForObjectRefs += "}"
	repeatedStringForConditions := "[]*ConditionType{"
	for _, f := range this.Conditions {
		repeatedStringForConditions += strings.Replace(fmt.Sprintf("%v", f), "ConditionType", "schema.ConditionType", 1) + ","
	}
	repeatedStringForConditions += "}"
	s := strings.Join([]string{`&StatusObject{`,
		`Metadata:` + strings.Replace(fmt.Sprintf("%v", this.Metadata), "StatusMetaType", "schema.StatusMetaType", 1) + `,`,
		`ObjectRefs:` + repeatedStringForObjectRefs + `,`,
		`Conditions:` + repeatedStringForConditions + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringObject(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *Object) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowObject
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Object: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Object: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Metadata", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowObject
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthObject
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthObject
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Metadata == nil {
				m.Metadata = &schema.ObjectMetaType{}
			}
			if err := m.Metadata.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SystemMetadata", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowObject
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthObject
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthObject
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.SystemMetadata == nil {
				m.SystemMetadata = &schema.SystemObjectMetaType{}
			}
			if err := m.SystemMetadata.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Spec", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowObject
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthObject
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthObject
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Spec == nil {
				m.Spec = &SpecType{}
			}
			if err := m.Spec.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipObject(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthObject
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthObject
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *SpecType) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowObject
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: SpecType: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SpecType: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GcSpec", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowObject
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthObject
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthObject
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.GcSpec == nil {
				m.GcSpec = &GlobalSpecType{}
			}
			if err := m.GcSpec.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipObject(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthObject
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthObject
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *StatusObject) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowObject
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: StatusObject: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: StatusObject: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Metadata", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowObject
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthObject
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthObject
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Metadata == nil {
				m.Metadata = &schema.StatusMetaType{}
			}
			if err := m.Metadata.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ObjectRefs", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowObject
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthObject
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthObject
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ObjectRefs = append(m.ObjectRefs, &schema.ObjectRefType{})
			if err := m.ObjectRefs[len(m.ObjectRefs)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Conditions", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowObject
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthObject
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthObject
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Conditions = append(m.Conditions, &schema.ConditionType{})
			if err := m.Conditions[len(m.Conditions)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipObject(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthObject
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthObject
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipObject(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowObject
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowObject
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowObject
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthObject
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupObject
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthObject
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthObject        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowObject          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupObject = fmt.Errorf("proto: unexpected end of group")
)
