// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: envoy/config/filter/http/squash/v2/squash.proto

/*
	Package v2 is a generated protocol buffer package.

	It is generated from these files:
		envoy/config/filter/http/squash/v2/squash.proto

	It has these top-level messages:
		Squash
*/
package v2

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/types"
import google_protobuf1 "github.com/gogo/protobuf/types"
import _ "github.com/envoyproxy/protoc-gen-validate/validate"
import _ "github.com/gogo/protobuf/gogoproto"

import time "time"

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

// [#proto-status: experimental]
type Squash struct {
	// The name of the cluster that hosts the Squash server.
	Cluster string `protobuf:"bytes,1,opt,name=cluster,proto3" json:"cluster,omitempty"`
	// When the filter requests the Squash server to create a DebugAttachment, it will use this
	// structure as template for the body of the request. It can contain reference to environment
	// variables in the form of '{{ ENV_VAR_NAME }}'. These can be used to provide the Squash server
	// with more information to find the process to attach the debugger to. For example, in a
	// Istio/k8s environment, this will contain information on the pod:
	//
	// .. code-block:: json
	//
	//  {
	//    "spec": {
	//      "attachment": {
	//        "pod": "{{ POD_NAME }}",
	//        "namespace": "{{ POD_NAMESPACE }}"
	//      },
	//      "match_request": true
	//    }
	//  }
	//
	// (where POD_NAME, POD_NAMESPACE are configured in the pod via the Downward API)
	AttachmentTemplate *google_protobuf1.Struct `protobuf:"bytes,2,opt,name=attachment_template,json=attachmentTemplate" json:"attachment_template,omitempty"`
	// The timeout for individual requests sent to the Squash cluster. Defaults to 1 second.
	RequestTimeout *time.Duration `protobuf:"bytes,3,opt,name=request_timeout,json=requestTimeout,stdduration" json:"request_timeout,omitempty"`
	// The total timeout Squash will delay a request and wait for it to be attached. Defaults to 60
	// seconds.
	AttachmentTimeout *time.Duration `protobuf:"bytes,4,opt,name=attachment_timeout,json=attachmentTimeout,stdduration" json:"attachment_timeout,omitempty"`
	// Amount of time to poll for the status of the attachment object in the Squash server
	// (to check if has been attached). Defaults to 1 second.
	AttachmentPollPeriod *time.Duration `protobuf:"bytes,5,opt,name=attachment_poll_period,json=attachmentPollPeriod,stdduration" json:"attachment_poll_period,omitempty"`
}

func (m *Squash) Reset()                    { *m = Squash{} }
func (m *Squash) String() string            { return proto.CompactTextString(m) }
func (*Squash) ProtoMessage()               {}
func (*Squash) Descriptor() ([]byte, []int) { return fileDescriptorSquash, []int{0} }

func (m *Squash) GetCluster() string {
	if m != nil {
		return m.Cluster
	}
	return ""
}

func (m *Squash) GetAttachmentTemplate() *google_protobuf1.Struct {
	if m != nil {
		return m.AttachmentTemplate
	}
	return nil
}

func (m *Squash) GetRequestTimeout() *time.Duration {
	if m != nil {
		return m.RequestTimeout
	}
	return nil
}

func (m *Squash) GetAttachmentTimeout() *time.Duration {
	if m != nil {
		return m.AttachmentTimeout
	}
	return nil
}

func (m *Squash) GetAttachmentPollPeriod() *time.Duration {
	if m != nil {
		return m.AttachmentPollPeriod
	}
	return nil
}

func init() {
	proto.RegisterType((*Squash)(nil), "envoy.config.filter.http.squash.v2.Squash")
}
func (m *Squash) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Squash) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Cluster) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintSquash(dAtA, i, uint64(len(m.Cluster)))
		i += copy(dAtA[i:], m.Cluster)
	}
	if m.AttachmentTemplate != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintSquash(dAtA, i, uint64(m.AttachmentTemplate.Size()))
		n1, err := m.AttachmentTemplate.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if m.RequestTimeout != nil {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintSquash(dAtA, i, uint64(types.SizeOfStdDuration(*m.RequestTimeout)))
		n2, err := types.StdDurationMarshalTo(*m.RequestTimeout, dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	if m.AttachmentTimeout != nil {
		dAtA[i] = 0x22
		i++
		i = encodeVarintSquash(dAtA, i, uint64(types.SizeOfStdDuration(*m.AttachmentTimeout)))
		n3, err := types.StdDurationMarshalTo(*m.AttachmentTimeout, dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n3
	}
	if m.AttachmentPollPeriod != nil {
		dAtA[i] = 0x2a
		i++
		i = encodeVarintSquash(dAtA, i, uint64(types.SizeOfStdDuration(*m.AttachmentPollPeriod)))
		n4, err := types.StdDurationMarshalTo(*m.AttachmentPollPeriod, dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n4
	}
	return i, nil
}

func encodeVarintSquash(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *Squash) Size() (n int) {
	var l int
	_ = l
	l = len(m.Cluster)
	if l > 0 {
		n += 1 + l + sovSquash(uint64(l))
	}
	if m.AttachmentTemplate != nil {
		l = m.AttachmentTemplate.Size()
		n += 1 + l + sovSquash(uint64(l))
	}
	if m.RequestTimeout != nil {
		l = types.SizeOfStdDuration(*m.RequestTimeout)
		n += 1 + l + sovSquash(uint64(l))
	}
	if m.AttachmentTimeout != nil {
		l = types.SizeOfStdDuration(*m.AttachmentTimeout)
		n += 1 + l + sovSquash(uint64(l))
	}
	if m.AttachmentPollPeriod != nil {
		l = types.SizeOfStdDuration(*m.AttachmentPollPeriod)
		n += 1 + l + sovSquash(uint64(l))
	}
	return n
}

func sovSquash(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozSquash(x uint64) (n int) {
	return sovSquash(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Squash) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSquash
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
			return fmt.Errorf("proto: Squash: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Squash: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Cluster", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSquash
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
				return ErrInvalidLengthSquash
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Cluster = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AttachmentTemplate", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSquash
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
				return ErrInvalidLengthSquash
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.AttachmentTemplate == nil {
				m.AttachmentTemplate = &google_protobuf1.Struct{}
			}
			if err := m.AttachmentTemplate.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RequestTimeout", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSquash
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
				return ErrInvalidLengthSquash
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.RequestTimeout == nil {
				m.RequestTimeout = new(time.Duration)
			}
			if err := types.StdDurationUnmarshal(m.RequestTimeout, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AttachmentTimeout", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSquash
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
				return ErrInvalidLengthSquash
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.AttachmentTimeout == nil {
				m.AttachmentTimeout = new(time.Duration)
			}
			if err := types.StdDurationUnmarshal(m.AttachmentTimeout, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AttachmentPollPeriod", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSquash
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
				return ErrInvalidLengthSquash
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.AttachmentPollPeriod == nil {
				m.AttachmentPollPeriod = new(time.Duration)
			}
			if err := types.StdDurationUnmarshal(m.AttachmentPollPeriod, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSquash(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthSquash
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
func skipSquash(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowSquash
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
					return 0, ErrIntOverflowSquash
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
					return 0, ErrIntOverflowSquash
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
				return 0, ErrInvalidLengthSquash
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowSquash
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
				next, err := skipSquash(dAtA[start:])
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
	ErrInvalidLengthSquash = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowSquash   = fmt.Errorf("proto: integer overflow")
)

func init() {
	proto.RegisterFile("envoy/config/filter/http/squash/v2/squash.proto", fileDescriptorSquash)
}

var fileDescriptorSquash = []byte{
	// 343 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x90, 0xc1, 0x4e, 0x22, 0x31,
	0x1c, 0xc6, 0xb7, 0x03, 0xcb, 0x86, 0x6e, 0xb2, 0x1b, 0x2b, 0x91, 0x91, 0x98, 0x91, 0xe0, 0x85,
	0x53, 0x9b, 0x8c, 0x6f, 0x40, 0x3c, 0x70, 0x32, 0x04, 0xf0, 0xe2, 0x85, 0x94, 0xa1, 0x0c, 0x93,
	0x94, 0xe9, 0xd0, 0xf9, 0x77, 0x12, 0xdf, 0x84, 0x67, 0xf1, 0xe4, 0xd1, 0xa3, 0x6f, 0xa0, 0xc1,
	0x8b, 0xbe, 0x85, 0xa1, 0x53, 0x02, 0xd1, 0x0b, 0xb7, 0x2f, 0xfd, 0xfa, 0xfb, 0xf5, 0x4b, 0x31,
	0x13, 0x69, 0xa1, 0x1e, 0x58, 0xa4, 0xd2, 0x79, 0x12, 0xb3, 0x79, 0x22, 0x41, 0x68, 0xb6, 0x00,
	0xc8, 0x58, 0xbe, 0x32, 0x3c, 0x5f, 0xb0, 0x22, 0x74, 0x89, 0x66, 0x5a, 0x81, 0x22, 0x1d, 0x0b,
	0xd0, 0x12, 0xa0, 0x25, 0x40, 0xb7, 0x00, 0x75, 0xd7, 0x8a, 0xb0, 0x15, 0xc4, 0x4a, 0xc5, 0x52,
	0x30, 0x4b, 0x4c, 0xcd, 0x9c, 0xcd, 0x8c, 0xe6, 0x90, 0xa8, 0xb4, 0x74, 0xb4, 0x2e, 0xbe, 0xf7,
	0x39, 0x68, 0x13, 0x81, 0x6b, 0x9b, 0x05, 0x97, 0xc9, 0x8c, 0x83, 0x60, 0xbb, 0xe0, 0x8a, 0x46,
	0xac, 0x62, 0x65, 0x23, 0xdb, 0xa6, 0xf2, 0xb4, 0xf3, 0xe1, 0xe1, 0xda, 0xc8, 0x3e, 0x4d, 0xae,
	0xf0, 0x9f, 0x48, 0x9a, 0x1c, 0x84, 0xf6, 0x51, 0x1b, 0x75, 0xeb, 0xbd, 0xfa, 0xe3, 0xe7, 0x53,
	0xa5, 0xaa, 0xbd, 0x36, 0x1a, 0xee, 0x1a, 0xd2, 0xc7, 0xa7, 0x1c, 0x80, 0x47, 0x8b, 0xa5, 0x48,
	0x61, 0x02, 0x62, 0x99, 0x49, 0x0e, 0xc2, 0xf7, 0xda, 0xa8, 0xfb, 0x37, 0x6c, 0xd2, 0x72, 0x1a,
	0xdd, 0x4d, 0xa3, 0x23, 0x3b, 0x6d, 0x48, 0xf6, 0xcc, 0xd8, 0x21, 0xa4, 0x8f, 0xff, 0x6b, 0xb1,
	0x32, 0x22, 0x87, 0x09, 0x24, 0x4b, 0xa1, 0x0c, 0xf8, 0x15, 0x6b, 0x39, 0xff, 0x61, 0xb9, 0x71,
	0x1f, 0xd0, 0xab, 0xae, 0x5f, 0x2f, 0xd1, 0xf0, 0x9f, 0xe3, 0xc6, 0x25, 0x46, 0x6e, 0x31, 0x39,
	0xdc, 0xe4, 0x64, 0xd5, 0xe3, 0x64, 0x27, 0x07, 0xd3, 0x9c, 0xef, 0x0e, 0x9f, 0x1d, 0xf8, 0x32,
	0x25, 0xe5, 0x24, 0x13, 0x3a, 0x51, 0x33, 0xff, 0xf7, 0x71, 0xce, 0xc6, 0x1e, 0x1f, 0x28, 0x29,
	0x07, 0x16, 0xee, 0x35, 0x9e, 0x37, 0x01, 0x7a, 0xd9, 0x04, 0xe8, 0x6d, 0x13, 0xa0, 0xf5, 0x7b,
	0xf0, 0xeb, 0xde, 0x2b, 0xc2, 0x69, 0xcd, 0x4a, 0xae, 0xbf, 0x02, 0x00, 0x00, 0xff, 0xff, 0x89,
	0x64, 0x2d, 0x74, 0x4b, 0x02, 0x00, 0x00,
}
