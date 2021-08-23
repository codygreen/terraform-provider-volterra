// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ves.io/schema/log/audit_log/types.proto

package audit_log

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema"
	io "io"
	math "math"
	math_bits "math/bits"
	reflect "reflect"
	strconv "strconv"
	strings "strings"
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

// KeyField
//
// x-displayName: "Key Field"
type KeyField int32

const (
	// x-displayName: "Authority"
	AUTHORITY KeyField = 0
	// x-displayName: "Destination Service"
	DST KeyField = 1
	// x-displayName: "Destination Instance"
	DST_INSTANCE KeyField = 2
	// x-displayName: "Destination Site"
	DST_SITE KeyField = 3
	// x-displayName: "Method"
	METHOD KeyField = 4
	// x-displayName: "Scheme"
	SCHEME KeyField = 5
	// x-displayName: "Request Path"
	REQ_PATH KeyField = 6
	// x-displayName: "Response Code"
	RSP_CODE KeyField = 7
	// x-displayName: "Source Service"
	SRC KeyField = 8
	// x-displayName: "Source Instance"
	SRC_INSTANCE KeyField = 9
	// x-displayName: "Source Site"
	SRC_SITE KeyField = 10
	// x-displayName: "Trasnsport"
	TRANSPORT KeyField = 11
	// x-displayName: "User"
	USER KeyField = 12
)

var KeyField_name = map[int32]string{
	0:  "AUTHORITY",
	1:  "DST",
	2:  "DST_INSTANCE",
	3:  "DST_SITE",
	4:  "METHOD",
	5:  "SCHEME",
	6:  "REQ_PATH",
	7:  "RSP_CODE",
	8:  "SRC",
	9:  "SRC_INSTANCE",
	10: "SRC_SITE",
	11: "TRANSPORT",
	12: "USER",
}

var KeyField_value = map[string]int32{
	"AUTHORITY":    0,
	"DST":          1,
	"DST_INSTANCE": 2,
	"DST_SITE":     3,
	"METHOD":       4,
	"SCHEME":       5,
	"REQ_PATH":     6,
	"RSP_CODE":     7,
	"SRC":          8,
	"SRC_INSTANCE": 9,
	"SRC_SITE":     10,
	"TRANSPORT":    11,
	"USER":         12,
}

func (KeyField) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_86c9591c7d7e7801, []int{0}
}

// Date Aggregation
//
// x-displayName: "Date Aggregation"
// Aggregate access logs based on timestamp in the log
type DateAggregation struct {
	// step
	//
	// x-displayName: "Step"
	// x-required
	// x-example: "5m"
	//
	// step is the resolution width, which determines the number of the data points [x-axis (time)] to be returned in the response.
	// The timestamps in the response will be t1=start_time, t2=t1+step, ... tn=tn-1+step, where tn <= end_time.
	// Format: [0-9][smhd], where s - seconds, m - minutes, h - hours, d - days
	Step string `protobuf:"bytes,1,opt,name=step,proto3" json:"step,omitempty"`
}

func (m *DateAggregation) Reset()      { *m = DateAggregation{} }
func (*DateAggregation) ProtoMessage() {}
func (*DateAggregation) Descriptor() ([]byte, []int) {
	return fileDescriptor_86c9591c7d7e7801, []int{0}
}
func (m *DateAggregation) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DateAggregation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *DateAggregation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DateAggregation.Merge(m, src)
}
func (m *DateAggregation) XXX_Size() int {
	return m.Size()
}
func (m *DateAggregation) XXX_DiscardUnknown() {
	xxx_messageInfo_DateAggregation.DiscardUnknown(m)
}

var xxx_messageInfo_DateAggregation proto.InternalMessageInfo

func (m *DateAggregation) GetStep() string {
	if m != nil {
		return m.Step
	}
	return ""
}

// Field Aggregation
//
// x-displayName: "Field Aggregation"
// Aggregate access logs based on the key fields in the log.
type FieldAggregation struct {
	// field
	//
	// x-displayName: "Field"
	// x-required
	//
	// Field name by which the logs should be aggregated.
	Field KeyField `protobuf:"varint,1,opt,name=field,proto3,enum=ves.io.schema.log.audit_log.KeyField" json:"field,omitempty"`
	// topk
	//
	// x-displayName: "TopK"
	//
	// Number of top field values to be returned in the response.
	// Optional: If not specified, top 5 values will be returned in the response.
	Topk uint32 `protobuf:"varint,2,opt,name=topk,proto3" json:"topk,omitempty"`
}

func (m *FieldAggregation) Reset()      { *m = FieldAggregation{} }
func (*FieldAggregation) ProtoMessage() {}
func (*FieldAggregation) Descriptor() ([]byte, []int) {
	return fileDescriptor_86c9591c7d7e7801, []int{1}
}
func (m *FieldAggregation) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *FieldAggregation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *FieldAggregation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FieldAggregation.Merge(m, src)
}
func (m *FieldAggregation) XXX_Size() int {
	return m.Size()
}
func (m *FieldAggregation) XXX_DiscardUnknown() {
	xxx_messageInfo_FieldAggregation.DiscardUnknown(m)
}

var xxx_messageInfo_FieldAggregation proto.InternalMessageInfo

func (m *FieldAggregation) GetField() KeyField {
	if m != nil {
		return m.Field
	}
	return AUTHORITY
}

func (m *FieldAggregation) GetTopk() uint32 {
	if m != nil {
		return m.Topk
	}
	return 0
}

// Aggregation Request
//
// x-displayName: "Aggregation Request"
// Aggregation request to provide analytics data over the log response
type AggregationRequest struct {
	// aggregation type
	//
	// x-displayName: "Aggregation Type"
	// Specify one of the aggregation types
	//
	// Types that are valid to be assigned to AggregationType:
	//	*AggregationRequest_DateAggregation
	//	*AggregationRequest_FieldAggregation
	AggregationType isAggregationRequest_AggregationType `protobuf_oneof:"aggregation_type"`
}

func (m *AggregationRequest) Reset()      { *m = AggregationRequest{} }
func (*AggregationRequest) ProtoMessage() {}
func (*AggregationRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_86c9591c7d7e7801, []int{2}
}
func (m *AggregationRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AggregationRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *AggregationRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AggregationRequest.Merge(m, src)
}
func (m *AggregationRequest) XXX_Size() int {
	return m.Size()
}
func (m *AggregationRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AggregationRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AggregationRequest proto.InternalMessageInfo

type isAggregationRequest_AggregationType interface {
	isAggregationRequest_AggregationType()
	Equal(interface{}) bool
	MarshalTo([]byte) (int, error)
	Size() int
}

type AggregationRequest_DateAggregation struct {
	DateAggregation *DateAggregation `protobuf:"bytes,1,opt,name=date_aggregation,json=dateAggregation,proto3,oneof" json:"date_aggregation,omitempty"`
}
type AggregationRequest_FieldAggregation struct {
	FieldAggregation *FieldAggregation `protobuf:"bytes,2,opt,name=field_aggregation,json=fieldAggregation,proto3,oneof" json:"field_aggregation,omitempty"`
}

func (*AggregationRequest_DateAggregation) isAggregationRequest_AggregationType()  {}
func (*AggregationRequest_FieldAggregation) isAggregationRequest_AggregationType() {}

func (m *AggregationRequest) GetAggregationType() isAggregationRequest_AggregationType {
	if m != nil {
		return m.AggregationType
	}
	return nil
}

func (m *AggregationRequest) GetDateAggregation() *DateAggregation {
	if x, ok := m.GetAggregationType().(*AggregationRequest_DateAggregation); ok {
		return x.DateAggregation
	}
	return nil
}

func (m *AggregationRequest) GetFieldAggregation() *FieldAggregation {
	if x, ok := m.GetAggregationType().(*AggregationRequest_FieldAggregation); ok {
		return x.FieldAggregation
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*AggregationRequest) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*AggregationRequest_DateAggregation)(nil),
		(*AggregationRequest_FieldAggregation)(nil),
	}
}

func init() {
	proto.RegisterEnum("ves.io.schema.log.audit_log.KeyField", KeyField_name, KeyField_value)
	proto.RegisterType((*DateAggregation)(nil), "ves.io.schema.log.audit_log.DateAggregation")
	proto.RegisterType((*FieldAggregation)(nil), "ves.io.schema.log.audit_log.FieldAggregation")
	proto.RegisterType((*AggregationRequest)(nil), "ves.io.schema.log.audit_log.AggregationRequest")
}

func init() {
	proto.RegisterFile("ves.io/schema/log/audit_log/types.proto", fileDescriptor_86c9591c7d7e7801)
}

var fileDescriptor_86c9591c7d7e7801 = []byte{
	// 571 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x93, 0xb1, 0x4f, 0xdb, 0x4e,
	0x14, 0xc7, 0x7d, 0x89, 0x01, 0xe7, 0x80, 0x1f, 0xf7, 0xbb, 0x09, 0x01, 0x3d, 0x21, 0x24, 0x54,
	0x84, 0x1a, 0x5b, 0xa2, 0x63, 0xa7, 0x90, 0xb8, 0x32, 0xad, 0x20, 0x70, 0x36, 0x03, 0x55, 0x25,
	0xcb, 0xc1, 0x17, 0x63, 0x11, 0x38, 0xd7, 0xbe, 0xa4, 0x65, 0xeb, 0xde, 0xa5, 0x7f, 0x46, 0xff,
	0x84, 0xaa, 0x2c, 0x8c, 0x1d, 0x33, 0x66, 0x6c, 0x9c, 0xa5, 0xdd, 0xe8, 0x7f, 0x50, 0xdd, 0x25,
	0x94, 0x24, 0xaa, 0xb2, 0xbd, 0xf7, 0xbe, 0xdf, 0xf7, 0x79, 0x7e, 0xa7, 0x67, 0xf8, 0xb4, 0xc3,
	0x32, 0x33, 0xe6, 0x56, 0x76, 0x7e, 0xc1, 0xae, 0x02, 0xab, 0xc5, 0x23, 0x2b, 0x68, 0x87, 0xb1,
	0xf0, 0x65, 0x24, 0x6e, 0x12, 0x96, 0x99, 0x49, 0xca, 0x05, 0xc7, 0xeb, 0x43, 0xa3, 0x39, 0x34,
	0x9a, 0x2d, 0x1e, 0x99, 0x7f, 0x8d, 0x6b, 0xe5, 0x28, 0x16, 0x17, 0xed, 0x86, 0x79, 0xce, 0xaf,
	0xac, 0x88, 0x47, 0xdc, 0x52, 0x3d, 0x8d, 0x76, 0x53, 0x65, 0x2a, 0x51, 0xd1, 0x90, 0xb5, 0xb6,
	0x3e, 0x39, 0x94, 0x27, 0x22, 0xe6, 0xd7, 0xa3, 0x41, 0x6b, 0x1b, 0x93, 0x62, 0x27, 0x68, 0xc5,
	0x61, 0x20, 0xd8, 0x48, 0xdd, 0x9c, 0x52, 0x63, 0xf6, 0xde, 0x9f, 0xe8, 0xdf, 0xda, 0x86, 0x2b,
	0xb5, 0x40, 0xb0, 0x4a, 0x14, 0xa5, 0x2c, 0x0a, 0xa4, 0x82, 0x31, 0xd4, 0x33, 0xc1, 0x92, 0x55,
	0xb0, 0x09, 0x76, 0x4a, 0x54, 0xc5, 0x5b, 0xd7, 0x10, 0xbd, 0x8c, 0x59, 0x2b, 0x1c, 0xf7, 0xbd,
	0x80, 0x73, 0x4d, 0x59, 0x53, 0xc6, 0xff, 0xf6, 0xb6, 0xcd, 0x19, 0x3b, 0x9b, 0xaf, 0xd9, 0x8d,
	0x02, 0xd0, 0x61, 0x0f, 0x7e, 0x02, 0x75, 0xc1, 0x93, 0xcb, 0xd5, 0xc2, 0x26, 0xd8, 0x59, 0xde,
	0x2f, 0x7d, 0xfb, 0x75, 0x57, 0xd4, 0x77, 0x0b, 0xab, 0x21, 0x55, 0xe5, 0xad, 0xdf, 0x00, 0xe2,
	0xb1, 0x59, 0x94, 0xbd, 0x6b, 0xb3, 0x4c, 0xe0, 0x33, 0x88, 0xe4, 0x76, 0x7e, 0xf0, 0x28, 0xa9,
	0xe9, 0x8b, 0x7b, 0xcf, 0x66, 0x4e, 0x9f, 0x5a, 0xd1, 0xd1, 0xe8, 0x4a, 0x38, 0xb5, 0xf5, 0x5b,
	0xf8, 0xbf, 0xfa, 0xb2, 0x09, 0x76, 0x41, 0xb1, 0xcb, 0x33, 0xd9, 0xd3, 0xef, 0xe2, 0x68, 0x14,
	0x35, 0xa7, 0x6a, 0xfb, 0x1b, 0x10, 0x8d, 0x71, 0x7d, 0x79, 0x2a, 0xd8, 0xb8, 0xbb, 0x05, 0xa0,
	0x7b, 0x0b, 0x8a, 0xaf, 0x74, 0xa3, 0x88, 0xf4, 0xdd, 0xaf, 0x00, 0x1a, 0x0f, 0xcf, 0x84, 0x97,
	0x61, 0xa9, 0x72, 0xea, 0x39, 0x75, 0x7a, 0xe0, 0x9d, 0x21, 0x0d, 0x2f, 0xc0, 0x62, 0xcd, 0xf5,
	0x10, 0xc0, 0x08, 0x2e, 0xd5, 0x5c, 0xcf, 0x3f, 0x38, 0x72, 0xbd, 0xca, 0x51, 0xd5, 0x46, 0x05,
	0xbc, 0x04, 0x0d, 0x59, 0x71, 0x0f, 0x3c, 0x1b, 0x15, 0x31, 0x84, 0xf3, 0x87, 0xb6, 0xe7, 0xd4,
	0x6b, 0x48, 0x97, 0xb1, 0x5b, 0x75, 0xec, 0x43, 0x1b, 0xcd, 0x49, 0x17, 0xb5, 0x4f, 0xfc, 0xe3,
	0x8a, 0xe7, 0xa0, 0x79, 0x95, 0xb9, 0xc7, 0x7e, 0xb5, 0x5e, 0xb3, 0xd1, 0x82, 0x84, 0xbb, 0xb4,
	0x8a, 0x0c, 0x09, 0x77, 0x69, 0xf5, 0x11, 0x5e, 0x92, 0x46, 0x59, 0x51, 0x70, 0x28, 0x3f, 0xca,
	0xa3, 0x95, 0x23, 0xf7, 0xb8, 0x4e, 0x3d, 0xb4, 0x88, 0x0d, 0xa8, 0x9f, 0xba, 0x36, 0x45, 0x4b,
	0xfb, 0x9f, 0x40, 0xb7, 0x4f, 0xb4, 0x5e, 0x9f, 0x68, 0xf7, 0x7d, 0x02, 0x3e, 0xe6, 0x04, 0x7c,
	0xc9, 0x09, 0xf8, 0x9e, 0x13, 0xd0, 0xcd, 0x09, 0xe8, 0xe5, 0x04, 0xfc, 0xc8, 0x09, 0xf8, 0x99,
	0x13, 0xed, 0x3e, 0x27, 0xe0, 0xf3, 0x80, 0x68, 0xdd, 0x01, 0xd1, 0x7a, 0x03, 0xa2, 0xbd, 0x39,
	0x89, 0x78, 0x72, 0x19, 0x99, 0x1d, 0xde, 0x12, 0x2c, 0x4d, 0x03, 0xb3, 0x9d, 0x59, 0x2a, 0x68,
	0xf2, 0xf4, 0xaa, 0x9c, 0xa4, 0xbc, 0x13, 0x87, 0x2c, 0x2d, 0x3f, 0xc8, 0x56, 0xd2, 0x88, 0xb8,
	0xc5, 0x3e, 0x88, 0xd1, 0x45, 0xff, 0xeb, 0x47, 0x6c, 0xcc, 0xab, 0xd3, 0x7e, 0xfe, 0x27, 0x00,
	0x00, 0xff, 0xff, 0x6e, 0xba, 0x4c, 0xf6, 0xae, 0x03, 0x00, 0x00,
}

func (x KeyField) String() string {
	s, ok := KeyField_name[int32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}
func (this *DateAggregation) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*DateAggregation)
	if !ok {
		that2, ok := that.(DateAggregation)
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
	if this.Step != that1.Step {
		return false
	}
	return true
}
func (this *FieldAggregation) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*FieldAggregation)
	if !ok {
		that2, ok := that.(FieldAggregation)
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
	if this.Field != that1.Field {
		return false
	}
	if this.Topk != that1.Topk {
		return false
	}
	return true
}
func (this *AggregationRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*AggregationRequest)
	if !ok {
		that2, ok := that.(AggregationRequest)
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
	if that1.AggregationType == nil {
		if this.AggregationType != nil {
			return false
		}
	} else if this.AggregationType == nil {
		return false
	} else if !this.AggregationType.Equal(that1.AggregationType) {
		return false
	}
	return true
}
func (this *AggregationRequest_DateAggregation) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*AggregationRequest_DateAggregation)
	if !ok {
		that2, ok := that.(AggregationRequest_DateAggregation)
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
	if !this.DateAggregation.Equal(that1.DateAggregation) {
		return false
	}
	return true
}
func (this *AggregationRequest_FieldAggregation) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*AggregationRequest_FieldAggregation)
	if !ok {
		that2, ok := that.(AggregationRequest_FieldAggregation)
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
	if !this.FieldAggregation.Equal(that1.FieldAggregation) {
		return false
	}
	return true
}
func (this *DateAggregation) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 5)
	s = append(s, "&audit_log.DateAggregation{")
	s = append(s, "Step: "+fmt.Sprintf("%#v", this.Step)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *FieldAggregation) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 6)
	s = append(s, "&audit_log.FieldAggregation{")
	s = append(s, "Field: "+fmt.Sprintf("%#v", this.Field)+",\n")
	s = append(s, "Topk: "+fmt.Sprintf("%#v", this.Topk)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *AggregationRequest) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 6)
	s = append(s, "&audit_log.AggregationRequest{")
	if this.AggregationType != nil {
		s = append(s, "AggregationType: "+fmt.Sprintf("%#v", this.AggregationType)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *AggregationRequest_DateAggregation) GoString() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&audit_log.AggregationRequest_DateAggregation{` +
		`DateAggregation:` + fmt.Sprintf("%#v", this.DateAggregation) + `}`}, ", ")
	return s
}
func (this *AggregationRequest_FieldAggregation) GoString() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&audit_log.AggregationRequest_FieldAggregation{` +
		`FieldAggregation:` + fmt.Sprintf("%#v", this.FieldAggregation) + `}`}, ", ")
	return s
}
func valueToGoStringTypes(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func (m *DateAggregation) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DateAggregation) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DateAggregation) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Step) > 0 {
		i -= len(m.Step)
		copy(dAtA[i:], m.Step)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Step)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *FieldAggregation) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *FieldAggregation) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *FieldAggregation) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Topk != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.Topk))
		i--
		dAtA[i] = 0x10
	}
	if m.Field != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.Field))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *AggregationRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AggregationRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AggregationRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.AggregationType != nil {
		{
			size := m.AggregationType.Size()
			i -= size
			if _, err := m.AggregationType.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
		}
	}
	return len(dAtA) - i, nil
}

func (m *AggregationRequest_DateAggregation) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AggregationRequest_DateAggregation) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.DateAggregation != nil {
		{
			size, err := m.DateAggregation.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTypes(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}
func (m *AggregationRequest_FieldAggregation) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AggregationRequest_FieldAggregation) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.FieldAggregation != nil {
		{
			size, err := m.FieldAggregation.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTypes(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	return len(dAtA) - i, nil
}
func encodeVarintTypes(dAtA []byte, offset int, v uint64) int {
	offset -= sovTypes(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *DateAggregation) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Step)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	return n
}

func (m *FieldAggregation) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Field != 0 {
		n += 1 + sovTypes(uint64(m.Field))
	}
	if m.Topk != 0 {
		n += 1 + sovTypes(uint64(m.Topk))
	}
	return n
}

func (m *AggregationRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.AggregationType != nil {
		n += m.AggregationType.Size()
	}
	return n
}

func (m *AggregationRequest_DateAggregation) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.DateAggregation != nil {
		l = m.DateAggregation.Size()
		n += 1 + l + sovTypes(uint64(l))
	}
	return n
}
func (m *AggregationRequest_FieldAggregation) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.FieldAggregation != nil {
		l = m.FieldAggregation.Size()
		n += 1 + l + sovTypes(uint64(l))
	}
	return n
}

func sovTypes(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTypes(x uint64) (n int) {
	return sovTypes(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *DateAggregation) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&DateAggregation{`,
		`Step:` + fmt.Sprintf("%v", this.Step) + `,`,
		`}`,
	}, "")
	return s
}
func (this *FieldAggregation) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&FieldAggregation{`,
		`Field:` + fmt.Sprintf("%v", this.Field) + `,`,
		`Topk:` + fmt.Sprintf("%v", this.Topk) + `,`,
		`}`,
	}, "")
	return s
}
func (this *AggregationRequest) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&AggregationRequest{`,
		`AggregationType:` + fmt.Sprintf("%v", this.AggregationType) + `,`,
		`}`,
	}, "")
	return s
}
func (this *AggregationRequest_DateAggregation) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&AggregationRequest_DateAggregation{`,
		`DateAggregation:` + strings.Replace(fmt.Sprintf("%v", this.DateAggregation), "DateAggregation", "DateAggregation", 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *AggregationRequest_FieldAggregation) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&AggregationRequest_FieldAggregation{`,
		`FieldAggregation:` + strings.Replace(fmt.Sprintf("%v", this.FieldAggregation), "FieldAggregation", "FieldAggregation", 1) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringTypes(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *DateAggregation) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
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
			return fmt.Errorf("proto: DateAggregation: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DateAggregation: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Step", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Step = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTypes
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
func (m *FieldAggregation) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
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
			return fmt.Errorf("proto: FieldAggregation: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: FieldAggregation: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Field", wireType)
			}
			m.Field = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Field |= KeyField(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Topk", wireType)
			}
			m.Topk = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Topk |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTypes
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
func (m *AggregationRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
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
			return fmt.Errorf("proto: AggregationRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AggregationRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DateAggregation", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &DateAggregation{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.AggregationType = &AggregationRequest_DateAggregation{v}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FieldAggregation", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &FieldAggregation{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.AggregationType = &AggregationRequest_FieldAggregation{v}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTypes
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
func skipTypes(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTypes
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
					return 0, ErrIntOverflowTypes
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
					return 0, ErrIntOverflowTypes
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
				return 0, ErrInvalidLengthTypes
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTypes
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTypes
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTypes        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTypes          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTypes = fmt.Errorf("proto: unexpected end of group")
)
