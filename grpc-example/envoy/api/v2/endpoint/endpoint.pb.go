// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: envoy/api/v2/endpoint/endpoint.proto

/*
	Package endpoint is a generated protocol buffer package.

	It is generated from these files:
		envoy/api/v2/endpoint/endpoint.proto
		envoy/api/v2/endpoint/load_report.proto

	It has these top-level messages:
		Endpoint
		LbEndpoint
		LocalityLbEndpoints
		UpstreamLocalityStats
		UpstreamEndpointStats
		EndpointLoadMetricStats
		ClusterStats
*/
package endpoint

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import envoy_api_v2_core1 "github.com/envoyproxy/go-control-plane/envoy/api/v2/core"
import envoy_api_v2_core "github.com/envoyproxy/go-control-plane/envoy/api/v2/core"
import envoy_api_v2_core2 "github.com/envoyproxy/go-control-plane/envoy/api/v2/core"
import google_protobuf2 "github.com/gogo/protobuf/types"
import _ "github.com/envoyproxy/protoc-gen-validate/validate"
import _ "github.com/gogo/protobuf/gogoproto"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Upstream host identifier.
type Endpoint struct {
	// The upstream host address.
	//
	// .. attention::
	//
	//   The form of host address depends on the given cluster type. For STATIC or EDS,
	//   it is expected to be a direct IP address (or something resolvable by the
	//   specified :ref:`resolver <envoy_api_field_core.SocketAddress.resolver_name>`
	//   in the Address). For LOGICAL or STRICT DNS, it is expected to be hostname,
	//   and will be resolved via DNS.
	Address *envoy_api_v2_core1.Address `protobuf:"bytes,1,opt,name=address" json:"address,omitempty"`
	// The optional health check configuration is used as configuration for the
	// health checker to contact the health checked host.
	//
	// .. attention::
	//
	//   This takes into effect only for upstream clusters with
	//   :ref:`active health checking <arch_overview_health_checking>` enabled.
	HealthCheckConfig *Endpoint_HealthCheckConfig `protobuf:"bytes,2,opt,name=health_check_config,json=healthCheckConfig" json:"health_check_config,omitempty"`
}

func (m *Endpoint) Reset()                    { *m = Endpoint{} }
func (m *Endpoint) String() string            { return proto.CompactTextString(m) }
func (*Endpoint) ProtoMessage()               {}
func (*Endpoint) Descriptor() ([]byte, []int) { return fileDescriptorEndpoint, []int{0} }

func (m *Endpoint) GetAddress() *envoy_api_v2_core1.Address {
	if m != nil {
		return m.Address
	}
	return nil
}

func (m *Endpoint) GetHealthCheckConfig() *Endpoint_HealthCheckConfig {
	if m != nil {
		return m.HealthCheckConfig
	}
	return nil
}

// The optional health check configuration.
type Endpoint_HealthCheckConfig struct {
	// Optional alternative health check port value.
	//
	// By default the health check address port of an upstream host is the same
	// as the host's serving address port. This provides an alternative health
	// check port. Setting this with a non-zero value allows an upstream host
	// to have different health check address port.
	PortValue uint32 `protobuf:"varint,1,opt,name=port_value,json=portValue,proto3" json:"port_value,omitempty"`
}

func (m *Endpoint_HealthCheckConfig) Reset()         { *m = Endpoint_HealthCheckConfig{} }
func (m *Endpoint_HealthCheckConfig) String() string { return proto.CompactTextString(m) }
func (*Endpoint_HealthCheckConfig) ProtoMessage()    {}
func (*Endpoint_HealthCheckConfig) Descriptor() ([]byte, []int) {
	return fileDescriptorEndpoint, []int{0, 0}
}

func (m *Endpoint_HealthCheckConfig) GetPortValue() uint32 {
	if m != nil {
		return m.PortValue
	}
	return 0
}

// An Endpoint that Envoy can route traffic to.
type LbEndpoint struct {
	// Upstream host identifier
	Endpoint *Endpoint `protobuf:"bytes,1,opt,name=endpoint" json:"endpoint,omitempty"`
	// Optional health status when known and supplied by EDS server.
	HealthStatus envoy_api_v2_core2.HealthStatus `protobuf:"varint,2,opt,name=health_status,json=healthStatus,proto3,enum=envoy.api.v2.core.HealthStatus" json:"health_status,omitempty"`
	// The endpoint metadata specifies values that may be used by the load
	// balancer to select endpoints in a cluster for a given request. The filter
	// name should be specified as *envoy.lb*. An example boolean key-value pair
	// is *canary*, providing the optional canary status of the upstream host.
	// This may be matched against in a route's
	// :ref:`RouteAction <envoy_api_msg_route.RouteAction>` metadata_match field
	// to subset the endpoints considered in cluster load balancing.
	Metadata *envoy_api_v2_core.Metadata `protobuf:"bytes,3,opt,name=metadata" json:"metadata,omitempty"`
	// The optional load balancing weight of the upstream host, in the range 1 -
	// 128. Envoy uses the load balancing weight in some of the built in load
	// balancers. The load balancing weight for an endpoint is divided by the sum
	// of the weights of all endpoints in the endpoint's locality to produce a
	// percentage of traffic for the endpoint. This percentage is then further
	// weighted by the endpoint's locality's load balancing weight from
	// LocalityLbEndpoints. If unspecified, each host is presumed to have equal
	// weight in a locality.
	//
	// .. attention::
	//
	//   The limit of 128 is somewhat arbitrary, but is applied due to performance
	//   concerns with the current implementation and can be removed when
	//   `this issue <https://github.com/envoyproxy/envoy/issues/1285>`_ is fixed.
	LoadBalancingWeight *google_protobuf2.UInt32Value `protobuf:"bytes,4,opt,name=load_balancing_weight,json=loadBalancingWeight" json:"load_balancing_weight,omitempty"`
}

func (m *LbEndpoint) Reset()                    { *m = LbEndpoint{} }
func (m *LbEndpoint) String() string            { return proto.CompactTextString(m) }
func (*LbEndpoint) ProtoMessage()               {}
func (*LbEndpoint) Descriptor() ([]byte, []int) { return fileDescriptorEndpoint, []int{1} }

func (m *LbEndpoint) GetEndpoint() *Endpoint {
	if m != nil {
		return m.Endpoint
	}
	return nil
}

func (m *LbEndpoint) GetHealthStatus() envoy_api_v2_core2.HealthStatus {
	if m != nil {
		return m.HealthStatus
	}
	return envoy_api_v2_core2.HealthStatus_UNKNOWN
}

func (m *LbEndpoint) GetMetadata() *envoy_api_v2_core.Metadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *LbEndpoint) GetLoadBalancingWeight() *google_protobuf2.UInt32Value {
	if m != nil {
		return m.LoadBalancingWeight
	}
	return nil
}

// A group of endpoints belonging to a Locality.
// One can have multiple LocalityLbEndpoints for a locality, but this is
// generally only done if the different groups need to have different load
// balancing weights or different priorities.
type LocalityLbEndpoints struct {
	// Identifies location of where the upstream hosts run.
	Locality *envoy_api_v2_core.Locality `protobuf:"bytes,1,opt,name=locality" json:"locality,omitempty"`
	// The group of endpoints belonging to the locality specified.
	LbEndpoints []LbEndpoint `protobuf:"bytes,2,rep,name=lb_endpoints,json=lbEndpoints" json:"lb_endpoints"`
	// Optional: Per priority/region/zone/sub_zone weight - range 1-128. The load
	// balancing weight for a locality is divided by the sum of the weights of all
	// localities  at the same priority level to produce the effective percentage
	// of traffic for the locality.
	//
	// Locality weights are only considered when :ref:`locality weighted load
	// balancing <arch_overview_load_balancing_locality_weighted_lb>` is
	// configured. These weights are ignored otherwise. If no weights are
	// specificed when locality weighted load balancing is enabled, the cluster is
	// assumed to have a weight of 1.
	//
	// .. attention::
	//
	//   The limit of 128 is somewhat arbitrary, but is applied due to performance
	//   concerns with the current implementation and can be removed when
	//   `this issue <https://github.com/envoyproxy/envoy/issues/1285>`_ is fixed.
	LoadBalancingWeight *google_protobuf2.UInt32Value `protobuf:"bytes,3,opt,name=load_balancing_weight,json=loadBalancingWeight" json:"load_balancing_weight,omitempty"`
	// Optional: the priority for this LocalityLbEndpoints. If unspecified this will
	// default to the highest priority (0).
	//
	// Under usual circumstances, Envoy will only select endpoints for the highest
	// priority (0). In the event all endpoints for a particular priority are
	// unavailable/unhealthy, Envoy will fail over to selecting endpoints for the
	// next highest priority group.
	//
	// Priorities should range from 0 (highest) to N (lowest) without skipping.
	Priority uint32 `protobuf:"varint,5,opt,name=priority,proto3" json:"priority,omitempty"`
}

func (m *LocalityLbEndpoints) Reset()                    { *m = LocalityLbEndpoints{} }
func (m *LocalityLbEndpoints) String() string            { return proto.CompactTextString(m) }
func (*LocalityLbEndpoints) ProtoMessage()               {}
func (*LocalityLbEndpoints) Descriptor() ([]byte, []int) { return fileDescriptorEndpoint, []int{2} }

func (m *LocalityLbEndpoints) GetLocality() *envoy_api_v2_core.Locality {
	if m != nil {
		return m.Locality
	}
	return nil
}

func (m *LocalityLbEndpoints) GetLbEndpoints() []LbEndpoint {
	if m != nil {
		return m.LbEndpoints
	}
	return nil
}

func (m *LocalityLbEndpoints) GetLoadBalancingWeight() *google_protobuf2.UInt32Value {
	if m != nil {
		return m.LoadBalancingWeight
	}
	return nil
}

func (m *LocalityLbEndpoints) GetPriority() uint32 {
	if m != nil {
		return m.Priority
	}
	return 0
}

func init() {
	proto.RegisterType((*Endpoint)(nil), "envoy.api.v2.endpoint.Endpoint")
	proto.RegisterType((*Endpoint_HealthCheckConfig)(nil), "envoy.api.v2.endpoint.Endpoint.HealthCheckConfig")
	proto.RegisterType((*LbEndpoint)(nil), "envoy.api.v2.endpoint.LbEndpoint")
	proto.RegisterType((*LocalityLbEndpoints)(nil), "envoy.api.v2.endpoint.LocalityLbEndpoints")
}
func (this *Endpoint) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Endpoint)
	if !ok {
		that2, ok := that.(Endpoint)
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
	if !this.Address.Equal(that1.Address) {
		return false
	}
	if !this.HealthCheckConfig.Equal(that1.HealthCheckConfig) {
		return false
	}
	return true
}
func (this *Endpoint_HealthCheckConfig) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Endpoint_HealthCheckConfig)
	if !ok {
		that2, ok := that.(Endpoint_HealthCheckConfig)
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
	if this.PortValue != that1.PortValue {
		return false
	}
	return true
}
func (this *LbEndpoint) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*LbEndpoint)
	if !ok {
		that2, ok := that.(LbEndpoint)
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
	if !this.Endpoint.Equal(that1.Endpoint) {
		return false
	}
	if this.HealthStatus != that1.HealthStatus {
		return false
	}
	if !this.Metadata.Equal(that1.Metadata) {
		return false
	}
	if !this.LoadBalancingWeight.Equal(that1.LoadBalancingWeight) {
		return false
	}
	return true
}
func (this *LocalityLbEndpoints) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*LocalityLbEndpoints)
	if !ok {
		that2, ok := that.(LocalityLbEndpoints)
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
	if !this.Locality.Equal(that1.Locality) {
		return false
	}
	if len(this.LbEndpoints) != len(that1.LbEndpoints) {
		return false
	}
	for i := range this.LbEndpoints {
		if !this.LbEndpoints[i].Equal(&that1.LbEndpoints[i]) {
			return false
		}
	}
	if !this.LoadBalancingWeight.Equal(that1.LoadBalancingWeight) {
		return false
	}
	if this.Priority != that1.Priority {
		return false
	}
	return true
}
func (m *Endpoint) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Endpoint) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Address != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintEndpoint(dAtA, i, uint64(m.Address.Size()))
		n1, err := m.Address.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if m.HealthCheckConfig != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintEndpoint(dAtA, i, uint64(m.HealthCheckConfig.Size()))
		n2, err := m.HealthCheckConfig.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	return i, nil
}

func (m *Endpoint_HealthCheckConfig) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Endpoint_HealthCheckConfig) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.PortValue != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintEndpoint(dAtA, i, uint64(m.PortValue))
	}
	return i, nil
}

func (m *LbEndpoint) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *LbEndpoint) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Endpoint != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintEndpoint(dAtA, i, uint64(m.Endpoint.Size()))
		n3, err := m.Endpoint.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n3
	}
	if m.HealthStatus != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintEndpoint(dAtA, i, uint64(m.HealthStatus))
	}
	if m.Metadata != nil {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintEndpoint(dAtA, i, uint64(m.Metadata.Size()))
		n4, err := m.Metadata.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n4
	}
	if m.LoadBalancingWeight != nil {
		dAtA[i] = 0x22
		i++
		i = encodeVarintEndpoint(dAtA, i, uint64(m.LoadBalancingWeight.Size()))
		n5, err := m.LoadBalancingWeight.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n5
	}
	return i, nil
}

func (m *LocalityLbEndpoints) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *LocalityLbEndpoints) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Locality != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintEndpoint(dAtA, i, uint64(m.Locality.Size()))
		n6, err := m.Locality.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n6
	}
	if len(m.LbEndpoints) > 0 {
		for _, msg := range m.LbEndpoints {
			dAtA[i] = 0x12
			i++
			i = encodeVarintEndpoint(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if m.LoadBalancingWeight != nil {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintEndpoint(dAtA, i, uint64(m.LoadBalancingWeight.Size()))
		n7, err := m.LoadBalancingWeight.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n7
	}
	if m.Priority != 0 {
		dAtA[i] = 0x28
		i++
		i = encodeVarintEndpoint(dAtA, i, uint64(m.Priority))
	}
	return i, nil
}

func encodeVarintEndpoint(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *Endpoint) Size() (n int) {
	var l int
	_ = l
	if m.Address != nil {
		l = m.Address.Size()
		n += 1 + l + sovEndpoint(uint64(l))
	}
	if m.HealthCheckConfig != nil {
		l = m.HealthCheckConfig.Size()
		n += 1 + l + sovEndpoint(uint64(l))
	}
	return n
}

func (m *Endpoint_HealthCheckConfig) Size() (n int) {
	var l int
	_ = l
	if m.PortValue != 0 {
		n += 1 + sovEndpoint(uint64(m.PortValue))
	}
	return n
}

func (m *LbEndpoint) Size() (n int) {
	var l int
	_ = l
	if m.Endpoint != nil {
		l = m.Endpoint.Size()
		n += 1 + l + sovEndpoint(uint64(l))
	}
	if m.HealthStatus != 0 {
		n += 1 + sovEndpoint(uint64(m.HealthStatus))
	}
	if m.Metadata != nil {
		l = m.Metadata.Size()
		n += 1 + l + sovEndpoint(uint64(l))
	}
	if m.LoadBalancingWeight != nil {
		l = m.LoadBalancingWeight.Size()
		n += 1 + l + sovEndpoint(uint64(l))
	}
	return n
}

func (m *LocalityLbEndpoints) Size() (n int) {
	var l int
	_ = l
	if m.Locality != nil {
		l = m.Locality.Size()
		n += 1 + l + sovEndpoint(uint64(l))
	}
	if len(m.LbEndpoints) > 0 {
		for _, e := range m.LbEndpoints {
			l = e.Size()
			n += 1 + l + sovEndpoint(uint64(l))
		}
	}
	if m.LoadBalancingWeight != nil {
		l = m.LoadBalancingWeight.Size()
		n += 1 + l + sovEndpoint(uint64(l))
	}
	if m.Priority != 0 {
		n += 1 + sovEndpoint(uint64(m.Priority))
	}
	return n
}

func sovEndpoint(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozEndpoint(x uint64) (n int) {
	return sovEndpoint(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Endpoint) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEndpoint
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
			return fmt.Errorf("proto: Endpoint: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Endpoint: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEndpoint
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
				return ErrInvalidLengthEndpoint
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Address == nil {
				m.Address = &envoy_api_v2_core1.Address{}
			}
			if err := m.Address.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field HealthCheckConfig", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEndpoint
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
				return ErrInvalidLengthEndpoint
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.HealthCheckConfig == nil {
				m.HealthCheckConfig = &Endpoint_HealthCheckConfig{}
			}
			if err := m.HealthCheckConfig.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEndpoint(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthEndpoint
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
func (m *Endpoint_HealthCheckConfig) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEndpoint
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
			return fmt.Errorf("proto: HealthCheckConfig: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: HealthCheckConfig: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PortValue", wireType)
			}
			m.PortValue = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEndpoint
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PortValue |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipEndpoint(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthEndpoint
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
func (m *LbEndpoint) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEndpoint
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
			return fmt.Errorf("proto: LbEndpoint: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: LbEndpoint: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Endpoint", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEndpoint
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
				return ErrInvalidLengthEndpoint
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Endpoint == nil {
				m.Endpoint = &Endpoint{}
			}
			if err := m.Endpoint.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field HealthStatus", wireType)
			}
			m.HealthStatus = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEndpoint
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.HealthStatus |= (envoy_api_v2_core2.HealthStatus(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Metadata", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEndpoint
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
				return ErrInvalidLengthEndpoint
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Metadata == nil {
				m.Metadata = &envoy_api_v2_core.Metadata{}
			}
			if err := m.Metadata.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LoadBalancingWeight", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEndpoint
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
				return ErrInvalidLengthEndpoint
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.LoadBalancingWeight == nil {
				m.LoadBalancingWeight = &google_protobuf2.UInt32Value{}
			}
			if err := m.LoadBalancingWeight.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEndpoint(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthEndpoint
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
func (m *LocalityLbEndpoints) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEndpoint
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
			return fmt.Errorf("proto: LocalityLbEndpoints: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: LocalityLbEndpoints: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Locality", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEndpoint
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
				return ErrInvalidLengthEndpoint
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Locality == nil {
				m.Locality = &envoy_api_v2_core.Locality{}
			}
			if err := m.Locality.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LbEndpoints", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEndpoint
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
				return ErrInvalidLengthEndpoint
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.LbEndpoints = append(m.LbEndpoints, LbEndpoint{})
			if err := m.LbEndpoints[len(m.LbEndpoints)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LoadBalancingWeight", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEndpoint
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
				return ErrInvalidLengthEndpoint
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.LoadBalancingWeight == nil {
				m.LoadBalancingWeight = &google_protobuf2.UInt32Value{}
			}
			if err := m.LoadBalancingWeight.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Priority", wireType)
			}
			m.Priority = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEndpoint
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Priority |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipEndpoint(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthEndpoint
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
func skipEndpoint(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowEndpoint
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
					return 0, ErrIntOverflowEndpoint
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
					return 0, ErrIntOverflowEndpoint
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
				return 0, ErrInvalidLengthEndpoint
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowEndpoint
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
				next, err := skipEndpoint(dAtA[start:])
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
	ErrInvalidLengthEndpoint = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowEndpoint   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("envoy/api/v2/endpoint/endpoint.proto", fileDescriptorEndpoint) }

var fileDescriptorEndpoint = []byte{
	// 538 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x52, 0x41, 0x6f, 0xd3, 0x3e,
	0x1c, 0x9d, 0x9b, 0x76, 0xeb, 0xdf, 0xed, 0xfe, 0xd2, 0x52, 0x26, 0xa2, 0x32, 0xa5, 0x50, 0x4d,
	0xa8, 0xea, 0xc1, 0x11, 0x1d, 0x12, 0x07, 0x0e, 0x88, 0x0e, 0x24, 0x40, 0xe3, 0x62, 0x04, 0x48,
	0x1c, 0x88, 0x9c, 0xc4, 0x4b, 0x22, 0x4c, 0x1c, 0xa5, 0x6e, 0xa7, 0xdd, 0xf6, 0x31, 0x38, 0x73,
	0xda, 0x67, 0xe0, 0xc4, 0x71, 0x47, 0x3e, 0x01, 0x42, 0x85, 0x0b, 0x9f, 0x62, 0xc8, 0x8e, 0x9d,
	0x75, 0x6b, 0x27, 0x2e, 0xdc, 0x6c, 0xff, 0xde, 0x7b, 0x7e, 0xef, 0xe9, 0x07, 0x77, 0x69, 0x36,
	0xe3, 0xc7, 0x1e, 0xc9, 0x53, 0x6f, 0x36, 0xf2, 0x68, 0x16, 0xe5, 0x3c, 0xcd, 0x44, 0x75, 0x40,
	0x79, 0xc1, 0x05, 0xb7, 0xb7, 0x15, 0x0a, 0x91, 0x3c, 0x45, 0xb3, 0x11, 0x32, 0xc3, 0x6e, 0xef,
	0x12, 0x39, 0xe4, 0x05, 0xf5, 0x48, 0x14, 0x15, 0x74, 0x32, 0x29, 0x79, 0xdd, 0x9d, 0x65, 0x40,
	0x40, 0x26, 0x54, 0x4f, 0x77, 0x97, 0xa7, 0x09, 0x25, 0x4c, 0x24, 0x7e, 0x98, 0xd0, 0xf0, 0x83,
	0x46, 0xb9, 0x31, 0xe7, 0x31, 0xa3, 0x9e, 0xba, 0x05, 0xd3, 0x43, 0xef, 0xa8, 0x20, 0x79, 0x4e,
	0x0b, 0xf3, 0xc7, 0xcd, 0x19, 0x61, 0x69, 0x44, 0x04, 0xf5, 0xcc, 0x41, 0x0f, 0x6e, 0xc4, 0x3c,
	0xe6, 0xea, 0xe8, 0xc9, 0x53, 0xf9, 0xda, 0xff, 0x05, 0x60, 0xf3, 0xa9, 0x0e, 0x60, 0xdf, 0x87,
	0x1b, 0xda, 0xb0, 0x03, 0x6e, 0x83, 0x41, 0x6b, 0xd4, 0x45, 0x97, 0x92, 0x4a, 0x4f, 0xe8, 0x71,
	0x89, 0xc0, 0x06, 0x6a, 0x13, 0xd8, 0x59, 0xf4, 0xe9, 0x87, 0x3c, 0x3b, 0x4c, 0x63, 0xa7, 0xa6,
	0x14, 0xee, 0xa1, 0x95, 0x5d, 0x21, 0xf3, 0x27, 0x7a, 0xa6, 0xa8, 0xfb, 0x92, 0xb9, 0xaf, 0x88,
	0x78, 0x2b, 0xb9, 0xfa, 0xd4, 0x7d, 0x04, 0xb7, 0x96, 0x70, 0xf6, 0x10, 0xc2, 0x9c, 0x17, 0xc2,
	0x9f, 0x11, 0x36, 0xa5, 0xca, 0xf0, 0xe6, 0xb8, 0xf5, 0xe5, 0xf7, 0x57, 0x6b, 0x7d, 0x58, 0x77,
	0xce, 0xcf, 0x2d, 0xfc, 0x9f, 0x1c, 0xbf, 0x91, 0xd3, 0xfe, 0x69, 0x0d, 0xc2, 0x83, 0xa0, 0x0a,
	0xfa, 0x10, 0x36, 0x8d, 0x13, 0x9d, 0xb4, 0xf7, 0x17, 0x9f, 0xb8, 0x22, 0xd8, 0x4f, 0xe0, 0xa6,
	0xce, 0x3b, 0x11, 0x44, 0x4c, 0x27, 0x2a, 0xe9, 0xff, 0x57, 0x15, 0x54, 0x57, 0xa5, 0xe9, 0x57,
	0x0a, 0x86, 0xdb, 0xc9, 0xc2, 0xcd, 0x7e, 0x00, 0x9b, 0x1f, 0xa9, 0x20, 0x11, 0x11, 0xc4, 0xb1,
	0x94, 0x85, 0x5b, 0x2b, 0x04, 0x5e, 0x6a, 0x08, 0xae, 0xc0, 0xf6, 0x7b, 0xb8, 0xcd, 0x38, 0x89,
	0xfc, 0x80, 0x30, 0x92, 0x85, 0x69, 0x16, 0xfb, 0x47, 0x34, 0x8d, 0x13, 0xe1, 0xd4, 0x95, 0xca,
	0x0e, 0x2a, 0x17, 0x04, 0x99, 0x05, 0x41, 0xaf, 0x9f, 0x67, 0x62, 0x6f, 0xa4, 0x7a, 0x18, 0xb7,
	0x65, 0x3f, 0x1b, 0xc3, 0x86, 0x73, 0x02, 0x06, 0x00, 0x77, 0xa4, 0xd0, 0xd8, 0xe8, 0xbc, 0x55,
	0x32, 0xfd, 0xcf, 0x35, 0xd8, 0x39, 0xe0, 0x21, 0x61, 0xa9, 0x38, 0xbe, 0xa8, 0x4c, 0x19, 0x66,
	0xfa, 0x59, 0x77, 0xb6, 0xca, 0xb0, 0x61, 0xe2, 0x0a, 0x6c, 0xbf, 0x80, 0x6d, 0x16, 0xf8, 0xa6,
	0x3e, 0x59, 0x97, 0x35, 0x68, 0x8d, 0xee, 0x5c, 0x53, 0xf8, 0xc5, 0x97, 0xe3, 0xfa, 0xd9, 0xf7,
	0xde, 0x1a, 0x6e, 0xb1, 0x05, 0x13, 0xd7, 0x86, 0xb7, 0xfe, 0x49, 0x78, 0xfb, 0x2e, 0x6c, 0xe6,
	0x45, 0xca, 0x0b, 0x19, 0xb2, 0xa1, 0x36, 0x0a, 0x4a, 0x52, 0x63, 0x68, 0x39, 0x27, 0x00, 0x57,
	0xb3, 0xb1, 0x7b, 0x3a, 0x77, 0xc1, 0xd9, 0xdc, 0x05, 0xdf, 0xe6, 0x2e, 0xf8, 0x31, 0x77, 0xc1,
	0xa7, 0x9f, 0xee, 0xda, 0xbb, 0x6a, 0x47, 0x82, 0x75, 0x65, 0x60, 0xef, 0x4f, 0x00, 0x00, 0x00,
	0xff, 0xff, 0xcc, 0x94, 0x7b, 0x1c, 0x50, 0x04, 0x00, 0x00,
}
