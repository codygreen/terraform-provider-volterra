// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ves.io/schema/errors.proto

/*
	Package schema is a generated protocol buffer package.

	It is generated from these files:
		ves.io/schema/errors.proto
		ves.io/schema/net.proto
		ves.io/schema/options.proto
		ves.io/schema/pvt_types.proto
		ves.io/schema/types.proto
		ves.io/schema/validate.proto
		ves.io/schema/view_options.proto

	It has these top-level messages:
		ErrorType
		MacAddressType
		Ipv4AddressType
		Ipv4SubnetType
		Ipv6AddressType
		Ipv6SubnetType
		IpAddressType
		IpSubnetType
		PrefixListType
		Via
		On
		Dependencies
		Key
		Keys
		MetricDef
		DaemonTLSParamsType
		DaemonTlsCertificateType
		DaemonTlsParametersType
		UseragentType
		ServiceParameters
		OperMetaType
		DaemonEnvironmentType
		StatusServerParamsType
		SyncServerParamsType
		Empty
		ObjectRefType
		LabelSelectorType
		LabelMatcherType
		ConditionType
		StatusType
		InitializerType
		InitializersType
		StatusMetaType
		ObjectMetaType
		ListMetaType
		ObjectGetMetaType
		ObjectCreateMetaType
		ObjectReplaceMetaType
		MessageMetaType
		ViewRefType
		KubeRefType
		SystemObjectMetaType
		SystemObjectGetMetaType
		AuthnTypeBasicAuth
		AuthnTypeHeaders
		AuthnTypeQueryParams
		BlindfoldSecretInfoType
		VaultSecretInfoType
		ClearSecretInfoType
		WingmanSecretInfoType
		SecretType
		NetworkRefType
		SiteRefType
		IpPrefixSetRefType
		VSiteRefType
		PolicerRefType
		ProtocolPolicerRefType
		NetworkSiteRefSelector
		SiteVirtualSiteRefSelector
		HeaderManipulationOptionType
		TlsValidationParamsType
		TlsCertificateType
		TlsParamsType
		UpstreamTlsParamsType
		DownstreamTlsParamsType
		DomainType
		L4DestType
		TlsInterceptionRule
		TlsInterceptionPolicy
		TlsInterceptionType
		FractionalPercent
		BufferConfigType
		CorsPolicy
		PathMatcherType
		HeaderMatcherType
		QueryParameterMatcherType
		RouteMatch
		WafRefType
		WafRulesRefType
		WafType
		AppRoleAuthInfoType
		VaultAuthInfoType
		RestAuthInfoType
		HostAccessInfoType
		VaultAccessInfoType
		VaultSecretType
		VolterraSecretType
		PortValueType
		VirtualNetworkReferenceType
		VirtualNetworkSelectorType
		RetryBackOff
		RetryPolicyType
		MetricValue
		NextHopType
		StaticRouteType
		ForwardProxyConfigType
		HostIdentifier
		InterfaceIdentifier
		InterfaceOrNetwork
		RouteTarget2ByteAsn
		RouteTarget4ByteAsn
		RouteTargetIPv4Addr
		RouteTarget
		FieldRules
		FloatRules
		DoubleRules
		Int32Rules
		Int64Rules
		UInt32Rules
		UInt64Rules
		SInt32Rules
		SInt64Rules
		Fixed32Rules
		Fixed64Rules
		SFixed32Rules
		SFixed64Rules
		BoolRules
		StringRules
		BytesRules
		EnumRules
		MessageRules
		RepeatedRules
		MapRules
		AnyRules
		DurationRules
		TimestampRules
		Choices
		ChoiceItem
		ChoiceItemList
		LabelKeyClassList
		HiddenConditions
		FieldViewOptions
		MapOptions
		RepeatedOptions
		Tile
		Tiles
*/
package schema

import proto "github.com/gogo/protobuf/proto"
import golang_proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"
import google_protobuf1 "github.com/gogo/protobuf/types"

import strconv "strconv"

import strings "strings"
import reflect "reflect"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = golang_proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// ErrorCode
//
// x-displayName: "Error Code"
// Union of all possible error-codes from system
type ErrorCode int32

const (
	// No error
	EOK ErrorCode = 0
	// Permissions error
	EPERMS ErrorCode = 1
	// Input is not correct
	EBADINPUT ErrorCode = 2
	// Not found
	ENOTFOUND ErrorCode = 3
	// Already exists
	EEXISTS ErrorCode = 4
	// Unknown/catchall error
	EUNKNOWN ErrorCode = 5
	// Error in serializing/de-serializing
	ESERIALIZE ErrorCode = 6
	// Server error
	EINTERNAL ErrorCode = 7
)

var ErrorCode_name = map[int32]string{
	0: "EOK",
	1: "EPERMS",
	2: "EBADINPUT",
	3: "ENOTFOUND",
	4: "EEXISTS",
	5: "EUNKNOWN",
	6: "ESERIALIZE",
	7: "EINTERNAL",
}
var ErrorCode_value = map[string]int32{
	"EOK":        0,
	"EPERMS":     1,
	"EBADINPUT":  2,
	"ENOTFOUND":  3,
	"EEXISTS":    4,
	"EUNKNOWN":   5,
	"ESERIALIZE": 6,
	"EINTERNAL":  7,
}

func (ErrorCode) EnumDescriptor() ([]byte, []int) { return fileDescriptorErrors, []int{0} }

// ErrorType
//
// x-displayName: "Error Type"
// Information about a error in API operation
type ErrorType struct {
	// code
	//
	// x-displayName: "Code"
	// A simple general code by category
	Code ErrorCode `protobuf:"varint,1,opt,name=code,proto3,enum=ves.io.schema.ErrorCode" json:"code,omitempty"`
	// message
	//
	// x-displayName: "Message"
	// x-example: "value"
	// A human readable string of the error
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	// error_obj
	//
	// x-displayName: "Error Object"
	// A structured error object for machine parsing
	ErrorObj *google_protobuf1.Any `protobuf:"bytes,3,opt,name=error_obj,json=errorObj" json:"error_obj,omitempty"`
}

func (m *ErrorType) Reset()                    { *m = ErrorType{} }
func (*ErrorType) ProtoMessage()               {}
func (*ErrorType) Descriptor() ([]byte, []int) { return fileDescriptorErrors, []int{0} }

func (m *ErrorType) GetCode() ErrorCode {
	if m != nil {
		return m.Code
	}
	return EOK
}

func (m *ErrorType) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *ErrorType) GetErrorObj() *google_protobuf1.Any {
	if m != nil {
		return m.ErrorObj
	}
	return nil
}

func init() {
	proto.RegisterType((*ErrorType)(nil), "ves.io.schema.ErrorType")
	golang_proto.RegisterType((*ErrorType)(nil), "ves.io.schema.ErrorType")
	proto.RegisterEnum("ves.io.schema.ErrorCode", ErrorCode_name, ErrorCode_value)
	golang_proto.RegisterEnum("ves.io.schema.ErrorCode", ErrorCode_name, ErrorCode_value)
}
func (x ErrorCode) String() string {
	s, ok := ErrorCode_name[int32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}
func (this *ErrorType) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ErrorType)
	if !ok {
		that2, ok := that.(ErrorType)
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
	if this.Code != that1.Code {
		return false
	}
	if this.Message != that1.Message {
		return false
	}
	if !this.ErrorObj.Equal(that1.ErrorObj) {
		return false
	}
	return true
}
func (this *ErrorType) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 7)
	s = append(s, "&schema.ErrorType{")
	s = append(s, "Code: "+fmt.Sprintf("%#v", this.Code)+",\n")
	s = append(s, "Message: "+fmt.Sprintf("%#v", this.Message)+",\n")
	if this.ErrorObj != nil {
		s = append(s, "ErrorObj: "+fmt.Sprintf("%#v", this.ErrorObj)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringErrors(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func (m *ErrorType) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ErrorType) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Code != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintErrors(dAtA, i, uint64(m.Code))
	}
	if len(m.Message) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintErrors(dAtA, i, uint64(len(m.Message)))
		i += copy(dAtA[i:], m.Message)
	}
	if m.ErrorObj != nil {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintErrors(dAtA, i, uint64(m.ErrorObj.Size()))
		n1, err := m.ErrorObj.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	return i, nil
}

func encodeVarintErrors(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *ErrorType) Size() (n int) {
	var l int
	_ = l
	if m.Code != 0 {
		n += 1 + sovErrors(uint64(m.Code))
	}
	l = len(m.Message)
	if l > 0 {
		n += 1 + l + sovErrors(uint64(l))
	}
	if m.ErrorObj != nil {
		l = m.ErrorObj.Size()
		n += 1 + l + sovErrors(uint64(l))
	}
	return n
}

func sovErrors(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozErrors(x uint64) (n int) {
	return sovErrors(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *ErrorType) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&ErrorType{`,
		`Code:` + fmt.Sprintf("%v", this.Code) + `,`,
		`Message:` + fmt.Sprintf("%v", this.Message) + `,`,
		`ErrorObj:` + strings.Replace(fmt.Sprintf("%v", this.ErrorObj), "Any", "google_protobuf1.Any", 1) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringErrors(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *ErrorType) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowErrors
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ErrorType: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ErrorType: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Code", wireType)
			}
			m.Code = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowErrors
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Code |= (ErrorCode(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Message", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowErrors
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthErrors
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Message = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ErrorObj", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowErrors
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthErrors
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.ErrorObj == nil {
				m.ErrorObj = &google_protobuf1.Any{}
			}
			if err := m.ErrorObj.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipErrors(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthErrors
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
func skipErrors(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowErrors
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
					return 0, ErrIntOverflowErrors
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowErrors
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
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthErrors
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowErrors
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipErrors(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthErrors = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowErrors   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("ves.io/schema/errors.proto", fileDescriptorErrors) }
func init() { golang_proto.RegisterFile("ves.io/schema/errors.proto", fileDescriptorErrors) }

var fileDescriptorErrors = []byte{
	// 406 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x91, 0xc1, 0x6e, 0xd3, 0x30,
	0x1c, 0xc6, 0xe3, 0x75, 0xb4, 0xab, 0xc7, 0x26, 0xcb, 0xe2, 0x10, 0x7a, 0xb0, 0x2a, 0x4e, 0x15,
	0xa2, 0x8e, 0x18, 0x4f, 0x90, 0x6d, 0x46, 0x8a, 0x36, 0x9c, 0x29, 0x49, 0x05, 0xda, 0x05, 0x25,
	0xad, 0xe7, 0x75, 0x2c, 0xfb, 0x47, 0x4e, 0x52, 0xad, 0x37, 0x4e, 0x9c, 0x79, 0x0c, 0x1e, 0x83,
	0x23, 0xc7, 0x1d, 0x39, 0xd2, 0xec, 0xc2, 0x71, 0x8f, 0x80, 0xe6, 0xb4, 0x42, 0x3b, 0xf9, 0xff,
	0xe9, 0xfb, 0xfd, 0xfd, 0x7d, 0xb2, 0xf1, 0x60, 0xa1, 0x4a, 0x3e, 0x07, 0xaf, 0x9c, 0x5e, 0xaa,
	0x3c, 0xf5, 0x94, 0x31, 0x60, 0x4a, 0x5e, 0x18, 0xa8, 0x80, 0xee, 0xb5, 0x1e, 0x6f, 0xbd, 0xc1,
	0x58, 0xcf, 0xab, 0xcb, 0x3a, 0xe3, 0x53, 0xc8, 0x3d, 0x0d, 0x1a, 0x3c, 0x4b, 0x65, 0xf5, 0x85,
	0x55, 0x56, 0xd8, 0xa9, 0xdd, 0x1e, 0xbc, 0xd4, 0x00, 0xfa, 0x5a, 0xfd, 0xa7, 0xd2, 0x9b, 0x65,
	0x6b, 0xbd, 0xfa, 0x86, 0x70, 0x5f, 0x3c, 0x26, 0x25, 0xcb, 0x42, 0xd1, 0x37, 0x78, 0x7b, 0x0a,
	0x33, 0xe5, 0xa2, 0x21, 0x1a, 0xed, 0x1f, 0xb8, 0xfc, 0x49, 0x2a, 0xb7, 0xdc, 0x11, 0xcc, 0x54,
	0x64, 0x29, 0xea, 0xe2, 0x5e, 0xae, 0xca, 0x32, 0xd5, 0xca, 0xdd, 0x1a, 0xa2, 0x51, 0x3f, 0xda,
	0x48, 0xfa, 0x16, 0xf7, 0x6d, 0xfd, 0xcf, 0x90, 0x5d, 0xb9, 0x9d, 0x21, 0x1a, 0xed, 0x1e, 0xbc,
	0xe0, 0x6d, 0x09, 0xbe, 0x29, 0xc1, 0xfd, 0x9b, 0x65, 0xb4, 0x63, 0xb1, 0x30, 0xbb, 0x7a, 0x7d,
	0xbb, 0xee, 0xf1, 0x78, 0x3f, 0xed, 0xe1, 0x8e, 0x08, 0x4f, 0x88, 0x43, 0x31, 0xee, 0x8a, 0x33,
	0x11, 0x7d, 0x88, 0x09, 0xa2, 0x7b, 0xb8, 0x2f, 0x0e, 0xfd, 0xe3, 0x40, 0x9e, 0x4d, 0x12, 0xb2,
	0x65, 0xa5, 0x0c, 0x93, 0xf7, 0xe1, 0x44, 0x1e, 0x93, 0x0e, 0xdd, 0xc5, 0x3d, 0x21, 0x3e, 0x05,
	0x71, 0x12, 0x93, 0x6d, 0xfa, 0x1c, 0xef, 0x88, 0x89, 0x3c, 0x91, 0xe1, 0x47, 0x49, 0x9e, 0xd1,
	0x7d, 0x8c, 0x45, 0x2c, 0xa2, 0xc0, 0x3f, 0x0d, 0xce, 0x05, 0xe9, 0xda, 0xcd, 0x40, 0x26, 0x22,
	0x92, 0xfe, 0x29, 0xe9, 0x1d, 0xd6, 0x77, 0x2b, 0xe6, 0xfc, 0x5e, 0x31, 0xe7, 0x61, 0xc5, 0xd0,
	0xd7, 0x86, 0xa1, 0x1f, 0x0d, 0x43, 0xbf, 0x1a, 0x86, 0xee, 0x1a, 0x86, 0xfe, 0x34, 0x0c, 0xfd,
	0x6d, 0x98, 0xf3, 0xd0, 0x30, 0xf4, 0xfd, 0x9e, 0x39, 0x3f, 0xef, 0x19, 0x3a, 0x3f, 0xd2, 0x50,
	0x7c, 0xd1, 0x7c, 0x01, 0xd7, 0x95, 0x32, 0x26, 0xe5, 0x75, 0xe9, 0xd9, 0xe1, 0x02, 0x4c, 0x3e,
	0x2e, 0x0c, 0x2c, 0xe6, 0x33, 0x65, 0xc6, 0x1b, 0xdb, 0x2b, 0x32, 0x0d, 0x9e, 0xba, 0xad, 0xd6,
	0x1f, 0xdb, 0x1e, 0x59, 0xd7, 0x3e, 0xc4, 0xbb, 0x7f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x9d, 0x37,
	0xbc, 0xf4, 0xf7, 0x01, 0x00, 0x00,
}
