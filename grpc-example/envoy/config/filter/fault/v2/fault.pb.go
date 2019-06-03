// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: envoy/config/filter/fault/v2/fault.proto

/*
	Package v2 is a generated protocol buffer package.

	It is generated from these files:
		envoy/config/filter/fault/v2/fault.proto

	It has these top-level messages:
		FaultDelay
*/
package v2

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import envoy_type "github.com/envoyproxy/go-control-plane/envoy/type"
import _ "github.com/gogo/protobuf/types"
import _ "github.com/envoyproxy/protoc-gen-validate/validate"
import _ "github.com/gogo/protobuf/gogoproto"

import time "time"
import github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"

import types "github.com/gogo/protobuf/types"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type FaultDelay_FaultDelayType int32

const (
	// Fixed delay (step function).
	FaultDelay_FIXED FaultDelay_FaultDelayType = 0
)

var FaultDelay_FaultDelayType_name = map[int32]string{
	0: "FIXED",
}
var FaultDelay_FaultDelayType_value = map[string]int32{
	"FIXED": 0,
}

func (x FaultDelay_FaultDelayType) String() string {
	return proto.EnumName(FaultDelay_FaultDelayType_name, int32(x))
}
func (FaultDelay_FaultDelayType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptorFault, []int{0, 0}
}

// Delay specification is used to inject latency into the
// HTTP/gRPC/Mongo/Redis operation or delay proxying of TCP connections.
type FaultDelay struct {
	// Delay type to use (fixed|exponential|..). Currently, only fixed delay (step function) is
	// supported.
	Type FaultDelay_FaultDelayType `protobuf:"varint,1,opt,name=type,proto3,enum=envoy.config.filter.fault.v2.FaultDelay_FaultDelayType" json:"type,omitempty"`
	// Types that are valid to be assigned to FaultDelaySecifier:
	//	*FaultDelay_FixedDelay
	FaultDelaySecifier isFaultDelay_FaultDelaySecifier `protobuf_oneof:"fault_delay_secifier"`
	// The percentage of operations/connection requests on which the delay will be injected.
	Percentage *envoy_type.FractionalPercent `protobuf:"bytes,4,opt,name=percentage" json:"percentage,omitempty"`
}

func (m *FaultDelay) Reset()                    { *m = FaultDelay{} }
func (m *FaultDelay) String() string            { return proto.CompactTextString(m) }
func (*FaultDelay) ProtoMessage()               {}
func (*FaultDelay) Descriptor() ([]byte, []int) { return fileDescriptorFault, []int{0} }

type isFaultDelay_FaultDelaySecifier interface {
	isFaultDelay_FaultDelaySecifier()
	MarshalTo([]byte) (int, error)
	Size() int
}

type FaultDelay_FixedDelay struct {
	FixedDelay *time.Duration `protobuf:"bytes,3,opt,name=fixed_delay,json=fixedDelay,oneof,stdduration"`
}

func (*FaultDelay_FixedDelay) isFaultDelay_FaultDelaySecifier() {}

func (m *FaultDelay) GetFaultDelaySecifier() isFaultDelay_FaultDelaySecifier {
	if m != nil {
		return m.FaultDelaySecifier
	}
	return nil
}

func (m *FaultDelay) GetType() FaultDelay_FaultDelayType {
	if m != nil {
		return m.Type
	}
	return FaultDelay_FIXED
}

func (m *FaultDelay) GetFixedDelay() *time.Duration {
	if x, ok := m.GetFaultDelaySecifier().(*FaultDelay_FixedDelay); ok {
		return x.FixedDelay
	}
	return nil
}

func (m *FaultDelay) GetPercentage() *envoy_type.FractionalPercent {
	if m != nil {
		return m.Percentage
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*FaultDelay) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _FaultDelay_OneofMarshaler, _FaultDelay_OneofUnmarshaler, _FaultDelay_OneofSizer, []interface{}{
		(*FaultDelay_FixedDelay)(nil),
	}
}

func _FaultDelay_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*FaultDelay)
	// fault_delay_secifier
	switch x := m.FaultDelaySecifier.(type) {
	case *FaultDelay_FixedDelay:
		_ = b.EncodeVarint(3<<3 | proto.WireBytes)
		dAtA, err := github_com_gogo_protobuf_types.StdDurationMarshal(*x.FixedDelay)
		if err != nil {
			return err
		}
		if err := b.EncodeRawBytes(dAtA); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("FaultDelay.FaultDelaySecifier has unexpected type %T", x)
	}
	return nil
}

func _FaultDelay_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*FaultDelay)
	switch tag {
	case 3: // fault_delay_secifier.fixed_delay
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeRawBytes(true)
		if err != nil {
			return true, err
		}
		c := new(time.Duration)
		if err2 := github_com_gogo_protobuf_types.StdDurationUnmarshal(c, x); err2 != nil {
			return true, err
		}
		m.FaultDelaySecifier = &FaultDelay_FixedDelay{c}
		return true, err
	default:
		return false, nil
	}
}

func _FaultDelay_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*FaultDelay)
	// fault_delay_secifier
	switch x := m.FaultDelaySecifier.(type) {
	case *FaultDelay_FixedDelay:
		s := github_com_gogo_protobuf_types.SizeOfStdDuration(*x.FixedDelay)
		n += proto.SizeVarint(3<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

func init() {
	proto.RegisterType((*FaultDelay)(nil), "envoy.config.filter.fault.v2.FaultDelay")
	proto.RegisterEnum("envoy.config.filter.fault.v2.FaultDelay_FaultDelayType", FaultDelay_FaultDelayType_name, FaultDelay_FaultDelayType_value)
}
func (m *FaultDelay) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *FaultDelay) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Type != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintFault(dAtA, i, uint64(m.Type))
	}
	if m.FaultDelaySecifier != nil {
		nn1, err := m.FaultDelaySecifier.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += nn1
	}
	if m.Percentage != nil {
		dAtA[i] = 0x22
		i++
		i = encodeVarintFault(dAtA, i, uint64(m.Percentage.Size()))
		n2, err := m.Percentage.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	return i, nil
}

func (m *FaultDelay_FixedDelay) MarshalTo(dAtA []byte) (int, error) {
	i := 0
	if m.FixedDelay != nil {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintFault(dAtA, i, uint64(types.SizeOfStdDuration(*m.FixedDelay)))
		n3, err := types.StdDurationMarshalTo(*m.FixedDelay, dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n3
	}
	return i, nil
}
func encodeVarintFault(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *FaultDelay) Size() (n int) {
	var l int
	_ = l
	if m.Type != 0 {
		n += 1 + sovFault(uint64(m.Type))
	}
	if m.FaultDelaySecifier != nil {
		n += m.FaultDelaySecifier.Size()
	}
	if m.Percentage != nil {
		l = m.Percentage.Size()
		n += 1 + l + sovFault(uint64(l))
	}
	return n
}

func (m *FaultDelay_FixedDelay) Size() (n int) {
	var l int
	_ = l
	if m.FixedDelay != nil {
		l = types.SizeOfStdDuration(*m.FixedDelay)
		n += 1 + l + sovFault(uint64(l))
	}
	return n
}

func sovFault(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozFault(x uint64) (n int) {
	return sovFault(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *FaultDelay) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowFault
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
			return fmt.Errorf("proto: FaultDelay: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: FaultDelay: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			m.Type = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFault
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Type |= (FaultDelay_FaultDelayType(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FixedDelay", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFault
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
				return ErrInvalidLengthFault
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := new(time.Duration)
			if err := types.StdDurationUnmarshal(v, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.FaultDelaySecifier = &FaultDelay_FixedDelay{v}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Percentage", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFault
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
				return ErrInvalidLengthFault
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Percentage == nil {
				m.Percentage = &envoy_type.FractionalPercent{}
			}
			if err := m.Percentage.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipFault(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthFault
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
func skipFault(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowFault
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
					return 0, ErrIntOverflowFault
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
					return 0, ErrIntOverflowFault
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
				return 0, ErrInvalidLengthFault
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowFault
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
				next, err := skipFault(dAtA[start:])
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
	ErrInvalidLengthFault = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowFault   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("envoy/config/filter/fault/v2/fault.proto", fileDescriptorFault) }

var fileDescriptorFault = []byte{
	// 354 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x50, 0xbf, 0x4a, 0xc3, 0x40,
	0x18, 0xcf, 0x25, 0xad, 0xe8, 0x15, 0x4a, 0x09, 0x05, 0x63, 0xb5, 0xb1, 0x74, 0x2a, 0x0e, 0x77,
	0x10, 0x07, 0x27, 0x97, 0x50, 0x8b, 0x8a, 0x83, 0x14, 0x41, 0x71, 0x29, 0xd7, 0xe4, 0x4b, 0x38,
	0x08, 0xb9, 0x10, 0xd3, 0x60, 0x56, 0x9f, 0xa2, 0xcf, 0xe0, 0xec, 0x20, 0x4e, 0x1d, 0x1d, 0x7d,
	0x03, 0xa5, 0x4e, 0x7d, 0x0b, 0xc9, 0x5d, 0x8a, 0x75, 0x71, 0xfb, 0x25, 0xdf, 0xef, 0xef, 0xe1,
	0x01, 0xc4, 0xb9, 0x28, 0xa8, 0x27, 0xe2, 0x80, 0x87, 0x34, 0xe0, 0x51, 0x06, 0x29, 0x0d, 0xd8,
	0x2c, 0xca, 0x68, 0xee, 0x28, 0x40, 0x92, 0x54, 0x64, 0xc2, 0x3c, 0x90, 0x4c, 0xa2, 0x98, 0x44,
	0x31, 0x89, 0x22, 0xe4, 0x4e, 0xc7, 0x52, 0x3e, 0x59, 0x91, 0x00, 0x4d, 0x20, 0xf5, 0x20, 0xae,
	0x74, 0x1d, 0x3b, 0x14, 0x22, 0x8c, 0x80, 0xca, 0xaf, 0xe9, 0x2c, 0xa0, 0xfe, 0x2c, 0x65, 0x19,
	0x17, 0x71, 0x75, 0xdf, 0xcd, 0x59, 0xc4, 0x7d, 0x96, 0x01, 0x5d, 0x83, 0xea, 0xd0, 0x0e, 0x45,
	0x28, 0x24, 0xa4, 0x25, 0x52, 0x7f, 0xfb, 0x2f, 0x3a, 0xc6, 0xa3, 0x32, 0x75, 0x08, 0x11, 0x2b,
	0xcc, 0x5b, 0x5c, 0x2b, 0x33, 0x2d, 0xd4, 0x43, 0x83, 0xa6, 0x73, 0x42, 0xfe, 0x2b, 0x49, 0x7e,
	0x75, 0x1b, 0xf0, 0xa6, 0x48, 0xc0, 0xc5, 0x6f, 0xab, 0x85, 0x51, 0x7f, 0x42, 0x7a, 0x0b, 0x8d,
	0xa5, 0xa1, 0x79, 0x85, 0x1b, 0x01, 0x7f, 0x04, 0x7f, 0xe2, 0x97, 0x24, 0xcb, 0xe8, 0xa1, 0x41,
	0xc3, 0xd9, 0x23, 0x6a, 0x0c, 0x59, 0x8f, 0x21, 0xc3, 0x6a, 0x8c, 0xdb, 0x9c, 0x7f, 0x1e, 0x22,
	0xe9, 0xf2, 0x8c, 0xf4, 0x23, 0xed, 0x5c, 0x1b, 0x63, 0xa9, 0x57, 0x35, 0x4f, 0x31, 0xae, 0x5e,
	0x85, 0x85, 0x60, 0xd5, 0xa4, 0x59, 0xb7, 0x2a, 0x5b, 0xc6, 0x91, 0x51, 0xca, 0xbc, 0xd2, 0x87,
	0x45, 0xd7, 0x8a, 0x37, 0xde, 0x10, 0xf4, 0xf7, 0x71, 0xf3, 0x6f, 0x61, 0x73, 0x07, 0xd7, 0x47,
	0x17, 0x77, 0x67, 0xc3, 0x96, 0xe6, 0x76, 0x71, 0x5b, 0x2e, 0x54, 0x4d, 0x27, 0x0f, 0xe0, 0xf1,
	0x80, 0x43, 0x6a, 0xd6, 0x5f, 0x57, 0x0b, 0x03, 0x5d, 0xd6, 0xb6, 0xf5, 0x96, 0xe1, 0xb6, 0xdf,
	0x97, 0x36, 0xfa, 0x58, 0xda, 0xe8, 0x6b, 0x69, 0xa3, 0xf9, 0xb7, 0xad, 0xdd, 0xeb, 0xb9, 0x33,
	0xdd, 0x92, 0x3b, 0x8e, 0x7f, 0x02, 0x00, 0x00, 0xff, 0xff, 0x38, 0xe4, 0x19, 0x3f, 0x06, 0x02,
	0x00, 0x00,
}
