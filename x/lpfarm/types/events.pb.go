// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: crescent/lpfarm/v1beta1/events.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
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

type EventCreatePrivatePlan struct {
	Creator            string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	PlanId             uint64 `protobuf:"varint,2,opt,name=plan_id,json=planId,proto3" json:"plan_id,omitempty"`
	FarmingPoolAddress string `protobuf:"bytes,3,opt,name=farming_pool_address,json=farmingPoolAddress,proto3" json:"farming_pool_address,omitempty"`
}

func (m *EventCreatePrivatePlan) Reset()         { *m = EventCreatePrivatePlan{} }
func (m *EventCreatePrivatePlan) String() string { return proto.CompactTextString(m) }
func (*EventCreatePrivatePlan) ProtoMessage()    {}
func (*EventCreatePrivatePlan) Descriptor() ([]byte, []int) {
	return fileDescriptor_d74bdb17e60e7c6f, []int{0}
}
func (m *EventCreatePrivatePlan) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventCreatePrivatePlan) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventCreatePrivatePlan.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventCreatePrivatePlan) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventCreatePrivatePlan.Merge(m, src)
}
func (m *EventCreatePrivatePlan) XXX_Size() int {
	return m.Size()
}
func (m *EventCreatePrivatePlan) XXX_DiscardUnknown() {
	xxx_messageInfo_EventCreatePrivatePlan.DiscardUnknown(m)
}

var xxx_messageInfo_EventCreatePrivatePlan proto.InternalMessageInfo

type EventFarm struct {
	Farmer           string                                   `protobuf:"bytes,1,opt,name=farmer,proto3" json:"farmer,omitempty"`
	Coin             types.Coin                               `protobuf:"bytes,2,opt,name=coin,proto3" json:"coin"`
	WithdrawnRewards github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,3,rep,name=withdrawn_rewards,json=withdrawnRewards,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"withdrawn_rewards"`
}

func (m *EventFarm) Reset()         { *m = EventFarm{} }
func (m *EventFarm) String() string { return proto.CompactTextString(m) }
func (*EventFarm) ProtoMessage()    {}
func (*EventFarm) Descriptor() ([]byte, []int) {
	return fileDescriptor_d74bdb17e60e7c6f, []int{1}
}
func (m *EventFarm) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventFarm) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventFarm.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventFarm) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventFarm.Merge(m, src)
}
func (m *EventFarm) XXX_Size() int {
	return m.Size()
}
func (m *EventFarm) XXX_DiscardUnknown() {
	xxx_messageInfo_EventFarm.DiscardUnknown(m)
}

var xxx_messageInfo_EventFarm proto.InternalMessageInfo

type EventUnfarm struct {
	Farmer           string                                   `protobuf:"bytes,1,opt,name=farmer,proto3" json:"farmer,omitempty"`
	Coin             types.Coin                               `protobuf:"bytes,2,opt,name=coin,proto3" json:"coin"`
	WithdrawnRewards github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,3,rep,name=withdrawn_rewards,json=withdrawnRewards,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"withdrawn_rewards"`
}

func (m *EventUnfarm) Reset()         { *m = EventUnfarm{} }
func (m *EventUnfarm) String() string { return proto.CompactTextString(m) }
func (*EventUnfarm) ProtoMessage()    {}
func (*EventUnfarm) Descriptor() ([]byte, []int) {
	return fileDescriptor_d74bdb17e60e7c6f, []int{2}
}
func (m *EventUnfarm) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventUnfarm) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventUnfarm.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventUnfarm) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventUnfarm.Merge(m, src)
}
func (m *EventUnfarm) XXX_Size() int {
	return m.Size()
}
func (m *EventUnfarm) XXX_DiscardUnknown() {
	xxx_messageInfo_EventUnfarm.DiscardUnknown(m)
}

var xxx_messageInfo_EventUnfarm proto.InternalMessageInfo

type EventHarvest struct {
	Farmer           string                                   `protobuf:"bytes,1,opt,name=farmer,proto3" json:"farmer,omitempty"`
	Denom            string                                   `protobuf:"bytes,2,opt,name=denom,proto3" json:"denom,omitempty"`
	WithdrawnRewards github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,3,rep,name=withdrawn_rewards,json=withdrawnRewards,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"withdrawn_rewards"`
}

func (m *EventHarvest) Reset()         { *m = EventHarvest{} }
func (m *EventHarvest) String() string { return proto.CompactTextString(m) }
func (*EventHarvest) ProtoMessage()    {}
func (*EventHarvest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d74bdb17e60e7c6f, []int{3}
}
func (m *EventHarvest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventHarvest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventHarvest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventHarvest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventHarvest.Merge(m, src)
}
func (m *EventHarvest) XXX_Size() int {
	return m.Size()
}
func (m *EventHarvest) XXX_DiscardUnknown() {
	xxx_messageInfo_EventHarvest.DiscardUnknown(m)
}

var xxx_messageInfo_EventHarvest proto.InternalMessageInfo

type EventTerminatePlan struct {
	PlanId uint64 `protobuf:"varint,1,opt,name=plan_id,json=planId,proto3" json:"plan_id,omitempty"`
}

func (m *EventTerminatePlan) Reset()         { *m = EventTerminatePlan{} }
func (m *EventTerminatePlan) String() string { return proto.CompactTextString(m) }
func (*EventTerminatePlan) ProtoMessage()    {}
func (*EventTerminatePlan) Descriptor() ([]byte, []int) {
	return fileDescriptor_d74bdb17e60e7c6f, []int{4}
}
func (m *EventTerminatePlan) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventTerminatePlan) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventTerminatePlan.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventTerminatePlan) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventTerminatePlan.Merge(m, src)
}
func (m *EventTerminatePlan) XXX_Size() int {
	return m.Size()
}
func (m *EventTerminatePlan) XXX_DiscardUnknown() {
	xxx_messageInfo_EventTerminatePlan.DiscardUnknown(m)
}

var xxx_messageInfo_EventTerminatePlan proto.InternalMessageInfo

func init() {
	proto.RegisterType((*EventCreatePrivatePlan)(nil), "crescent.lpfarm.v1beta1.EventCreatePrivatePlan")
	proto.RegisterType((*EventFarm)(nil), "crescent.lpfarm.v1beta1.EventFarm")
	proto.RegisterType((*EventUnfarm)(nil), "crescent.lpfarm.v1beta1.EventUnfarm")
	proto.RegisterType((*EventHarvest)(nil), "crescent.lpfarm.v1beta1.EventHarvest")
	proto.RegisterType((*EventTerminatePlan)(nil), "crescent.lpfarm.v1beta1.EventTerminatePlan")
}

func init() {
	proto.RegisterFile("crescent/lpfarm/v1beta1/events.proto", fileDescriptor_d74bdb17e60e7c6f)
}

var fileDescriptor_d74bdb17e60e7c6f = []byte{
	// 443 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xdc, 0x53, 0x4d, 0x6e, 0x13, 0x31,
	0x14, 0x1e, 0xd3, 0x90, 0x2a, 0x2e, 0x0b, 0xb0, 0xa2, 0x36, 0x74, 0xe1, 0x56, 0x11, 0x8b, 0x6c,
	0x32, 0x6e, 0x29, 0x17, 0xa0, 0x15, 0x08, 0x76, 0xd5, 0x08, 0x36, 0x6c, 0x22, 0x67, 0xfc, 0x9a,
	0x9a, 0xce, 0xd8, 0x23, 0xdb, 0x4c, 0xca, 0x82, 0x3b, 0x70, 0x0e, 0x0e, 0xc0, 0x19, 0xb2, 0x8c,
	0x10, 0x0b, 0x56, 0xfc, 0x24, 0x17, 0x41, 0xb6, 0x67, 0xa2, 0x6c, 0x60, 0x8b, 0xc4, 0xca, 0xf3,
	0xe6, 0x7d, 0x7e, 0xdf, 0xf7, 0x59, 0xdf, 0xc3, 0x8f, 0x72, 0x03, 0x36, 0x07, 0xe5, 0x58, 0x51,
	0x5d, 0x71, 0x53, 0xb2, 0xfa, 0x74, 0x0a, 0x8e, 0x9f, 0x32, 0xa8, 0x41, 0x39, 0x9b, 0x56, 0x46,
	0x3b, 0x4d, 0x0e, 0x5a, 0x54, 0x1a, 0x51, 0x69, 0x83, 0x3a, 0xec, 0xcf, 0xf4, 0x4c, 0x07, 0x0c,
	0xf3, 0x5f, 0x11, 0x7e, 0x48, 0x73, 0x6d, 0x4b, 0x6d, 0xd9, 0x94, 0x5b, 0xd8, 0x0c, 0xcc, 0xb5,
	0x54, 0xb1, 0x3f, 0xfc, 0x80, 0xf7, 0x9f, 0xf9, 0xf1, 0x17, 0x06, 0xb8, 0x83, 0x4b, 0x23, 0x6b,
	0x7f, 0x14, 0x5c, 0x91, 0x01, 0xde, 0xcd, 0xfd, 0x4f, 0x6d, 0x06, 0xe8, 0x18, 0x8d, 0x7a, 0x59,
	0x5b, 0x92, 0x03, 0xbc, 0x5b, 0x15, 0x5c, 0x4d, 0xa4, 0x18, 0xdc, 0x39, 0x46, 0xa3, 0x4e, 0xd6,
	0xf5, 0xe5, 0x4b, 0x41, 0x4e, 0x70, 0xdf, 0x4b, 0x92, 0x6a, 0x36, 0xa9, 0xb4, 0x2e, 0x26, 0x5c,
	0x08, 0x03, 0xd6, 0x0e, 0x76, 0xc2, 0x7d, 0xd2, 0xf4, 0x2e, 0xb5, 0x2e, 0x9e, 0xc6, 0xce, 0xf0,
	0x0b, 0xc2, 0xbd, 0xc0, 0xff, 0x9c, 0x9b, 0x92, 0xec, 0xe3, 0xae, 0xc7, 0x40, 0xcb, 0xd8, 0x54,
	0xe4, 0x0c, 0x77, 0xbc, 0xe4, 0xc0, 0xb6, 0xf7, 0xf8, 0x61, 0x1a, 0x3d, 0xa5, 0xde, 0x53, 0x6b,
	0x3f, 0xbd, 0xd0, 0x52, 0x9d, 0x77, 0x16, 0xdf, 0x8f, 0x92, 0x2c, 0x80, 0xc9, 0x2d, 0x7e, 0x30,
	0x97, 0xee, 0x5a, 0x18, 0x3e, 0x57, 0x13, 0x03, 0x73, 0x6e, 0x84, 0x57, 0xb2, 0xf3, 0xf7, 0x09,
	0x27, 0x7e, 0xc2, 0xa7, 0x1f, 0x47, 0xa3, 0x99, 0x74, 0xd7, 0xef, 0xa6, 0x69, 0xae, 0x4b, 0xd6,
	0x3c, 0x61, 0x3c, 0xc6, 0x56, 0xdc, 0x30, 0xf7, 0xbe, 0x02, 0x1b, 0x2e, 0xd8, 0xec, 0xfe, 0x86,
	0x25, 0x8b, 0x24, 0xc3, 0xaf, 0x08, 0xef, 0x05, 0x53, 0xaf, 0xd5, 0xd5, 0x7f, 0x64, 0xeb, 0x33,
	0xc2, 0xf7, 0x82, 0xad, 0x17, 0xdc, 0xd4, 0x60, 0xdd, 0x1f, 0x7d, 0xf5, 0xf1, 0x5d, 0x01, 0x4a,
	0x97, 0xc1, 0x58, 0x2f, 0x8b, 0xc5, 0x3f, 0x14, 0x3e, 0xc6, 0x24, 0xe8, 0x7e, 0x05, 0x3e, 0x7f,
	0x6d, 0xbe, 0xb7, 0x52, 0x8c, 0xb6, 0x53, 0x7c, 0x9e, 0x2d, 0x7e, 0xd1, 0x64, 0xb1, 0xa2, 0x68,
	0xb9, 0xa2, 0xe8, 0xe7, 0x8a, 0xa2, 0x8f, 0x6b, 0x9a, 0x2c, 0xd7, 0x34, 0xf9, 0xb6, 0xa6, 0xc9,
	0x9b, 0x27, 0x5b, 0x42, 0xde, 0x82, 0xb5, 0x30, 0xae, 0xa4, 0xba, 0x29, 0xb9, 0x1a, 0xe7, 0x06,
	0xd8, 0x66, 0x85, 0x6f, 0xdb, 0x25, 0x0e, 0xd2, 0xa6, 0xdd, 0xb0, 0x6d, 0x67, 0xbf, 0x03, 0x00,
	0x00, 0xff, 0xff, 0x9f, 0x99, 0x14, 0x92, 0xe4, 0x03, 0x00, 0x00,
}

func (m *EventCreatePrivatePlan) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventCreatePrivatePlan) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventCreatePrivatePlan) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.FarmingPoolAddress) > 0 {
		i -= len(m.FarmingPoolAddress)
		copy(dAtA[i:], m.FarmingPoolAddress)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.FarmingPoolAddress)))
		i--
		dAtA[i] = 0x1a
	}
	if m.PlanId != 0 {
		i = encodeVarintEvents(dAtA, i, uint64(m.PlanId))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *EventFarm) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventFarm) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventFarm) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.WithdrawnRewards) > 0 {
		for iNdEx := len(m.WithdrawnRewards) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.WithdrawnRewards[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintEvents(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	{
		size, err := m.Coin.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintEvents(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.Farmer) > 0 {
		i -= len(m.Farmer)
		copy(dAtA[i:], m.Farmer)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Farmer)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *EventUnfarm) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventUnfarm) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventUnfarm) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.WithdrawnRewards) > 0 {
		for iNdEx := len(m.WithdrawnRewards) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.WithdrawnRewards[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintEvents(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	{
		size, err := m.Coin.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintEvents(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.Farmer) > 0 {
		i -= len(m.Farmer)
		copy(dAtA[i:], m.Farmer)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Farmer)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *EventHarvest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventHarvest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventHarvest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.WithdrawnRewards) > 0 {
		for iNdEx := len(m.WithdrawnRewards) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.WithdrawnRewards[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintEvents(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.Denom) > 0 {
		i -= len(m.Denom)
		copy(dAtA[i:], m.Denom)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Denom)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Farmer) > 0 {
		i -= len(m.Farmer)
		copy(dAtA[i:], m.Farmer)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Farmer)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *EventTerminatePlan) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventTerminatePlan) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventTerminatePlan) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.PlanId != 0 {
		i = encodeVarintEvents(dAtA, i, uint64(m.PlanId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintEvents(dAtA []byte, offset int, v uint64) int {
	offset -= sovEvents(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *EventCreatePrivatePlan) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	if m.PlanId != 0 {
		n += 1 + sovEvents(uint64(m.PlanId))
	}
	l = len(m.FarmingPoolAddress)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	return n
}

func (m *EventFarm) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Farmer)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	l = m.Coin.Size()
	n += 1 + l + sovEvents(uint64(l))
	if len(m.WithdrawnRewards) > 0 {
		for _, e := range m.WithdrawnRewards {
			l = e.Size()
			n += 1 + l + sovEvents(uint64(l))
		}
	}
	return n
}

func (m *EventUnfarm) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Farmer)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	l = m.Coin.Size()
	n += 1 + l + sovEvents(uint64(l))
	if len(m.WithdrawnRewards) > 0 {
		for _, e := range m.WithdrawnRewards {
			l = e.Size()
			n += 1 + l + sovEvents(uint64(l))
		}
	}
	return n
}

func (m *EventHarvest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Farmer)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	l = len(m.Denom)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	if len(m.WithdrawnRewards) > 0 {
		for _, e := range m.WithdrawnRewards {
			l = e.Size()
			n += 1 + l + sovEvents(uint64(l))
		}
	}
	return n
}

func (m *EventTerminatePlan) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.PlanId != 0 {
		n += 1 + sovEvents(uint64(m.PlanId))
	}
	return n
}

func sovEvents(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozEvents(x uint64) (n int) {
	return sovEvents(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *EventCreatePrivatePlan) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvents
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
			return fmt.Errorf("proto: EventCreatePrivatePlan: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventCreatePrivatePlan: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PlanId", wireType)
			}
			m.PlanId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PlanId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FarmingPoolAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.FarmingPoolAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvents(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvents
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
func (m *EventFarm) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvents
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
			return fmt.Errorf("proto: EventFarm: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventFarm: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Farmer", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Farmer = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Coin", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Coin.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field WithdrawnRewards", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.WithdrawnRewards = append(m.WithdrawnRewards, types.Coin{})
			if err := m.WithdrawnRewards[len(m.WithdrawnRewards)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvents(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvents
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
func (m *EventUnfarm) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvents
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
			return fmt.Errorf("proto: EventUnfarm: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventUnfarm: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Farmer", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Farmer = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Coin", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Coin.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field WithdrawnRewards", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.WithdrawnRewards = append(m.WithdrawnRewards, types.Coin{})
			if err := m.WithdrawnRewards[len(m.WithdrawnRewards)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvents(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvents
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
func (m *EventHarvest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvents
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
			return fmt.Errorf("proto: EventHarvest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventHarvest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Farmer", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Farmer = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Denom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Denom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field WithdrawnRewards", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.WithdrawnRewards = append(m.WithdrawnRewards, types.Coin{})
			if err := m.WithdrawnRewards[len(m.WithdrawnRewards)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvents(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvents
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
func (m *EventTerminatePlan) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvents
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
			return fmt.Errorf("proto: EventTerminatePlan: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventTerminatePlan: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PlanId", wireType)
			}
			m.PlanId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PlanId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipEvents(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvents
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
func skipEvents(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowEvents
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
					return 0, ErrIntOverflowEvents
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
					return 0, ErrIntOverflowEvents
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
				return 0, ErrInvalidLengthEvents
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupEvents
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthEvents
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthEvents        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowEvents          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupEvents = fmt.Errorf("proto: unexpected end of group")
)
