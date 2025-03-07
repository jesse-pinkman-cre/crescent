// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: crescent/referral/v1beta1/referral.proto

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

type Referral struct {
	Id     uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Addr   string `protobuf:"bytes,2,opt,name=addr,proto3" json:"addr,omitempty"`
	Code   string `protobuf:"bytes,3,opt,name=code,proto3" json:"code,omitempty"`
	Parent string `protobuf:"bytes,4,opt,name=parent,proto3" json:"parent,omitempty"`
	Tier   uint64 `protobuf:"varint,5,opt,name=tier,proto3" json:"tier,omitempty"`
}

func (m *Referral) Reset()         { *m = Referral{} }
func (m *Referral) String() string { return proto.CompactTextString(m) }
func (*Referral) ProtoMessage()    {}
func (*Referral) Descriptor() ([]byte, []int) {
	return fileDescriptor_1ce64a9d1383f4a0, []int{0}
}
func (m *Referral) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Referral) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Referral.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Referral) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Referral.Merge(m, src)
}
func (m *Referral) XXX_Size() int {
	return m.Size()
}
func (m *Referral) XXX_DiscardUnknown() {
	xxx_messageInfo_Referral.DiscardUnknown(m)
}

var xxx_messageInfo_Referral proto.InternalMessageInfo

type Revenue struct {
	ReferralId uint64                                   `protobuf:"varint,1,opt,name=referral_id,json=referralId,proto3" json:"referral_id,omitempty"`
	Coins      github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,2,rep,name=coins,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"coins"`
}

func (m *Revenue) Reset()         { *m = Revenue{} }
func (m *Revenue) String() string { return proto.CompactTextString(m) }
func (*Revenue) ProtoMessage()    {}
func (*Revenue) Descriptor() ([]byte, []int) {
	return fileDescriptor_1ce64a9d1383f4a0, []int{1}
}
func (m *Revenue) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Revenue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Revenue.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Revenue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Revenue.Merge(m, src)
}
func (m *Revenue) XXX_Size() int {
	return m.Size()
}
func (m *Revenue) XXX_DiscardUnknown() {
	xxx_messageInfo_Revenue.DiscardUnknown(m)
}

var xxx_messageInfo_Revenue proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Referral)(nil), "crescent.referral.v1beta1.Referral")
	proto.RegisterType((*Revenue)(nil), "crescent.referral.v1beta1.Revenue")
}

func init() {
	proto.RegisterFile("crescent/referral/v1beta1/referral.proto", fileDescriptor_1ce64a9d1383f4a0)
}

var fileDescriptor_1ce64a9d1383f4a0 = []byte{
	// 330 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x51, 0xbb, 0x4e, 0xc3, 0x30,
	0x14, 0x4d, 0xd2, 0x07, 0xe0, 0x4a, 0x0c, 0x11, 0x42, 0x69, 0x07, 0xb7, 0xea, 0x94, 0xa5, 0x31,
	0x05, 0x89, 0x0f, 0x28, 0x13, 0x6b, 0xc4, 0xc4, 0x82, 0x1c, 0xfb, 0x52, 0x4c, 0xa9, 0x1d, 0xd9,
	0x6e, 0x05, 0x1f, 0x81, 0xc4, 0x77, 0xf0, 0x25, 0x1d, 0x3b, 0x32, 0xf1, 0x68, 0x7f, 0x04, 0xc5,
	0x79, 0xc0, 0xe4, 0xe3, 0x73, 0xcf, 0x3d, 0xd7, 0xc7, 0x17, 0xc5, 0x4c, 0x83, 0x61, 0x20, 0x2d,
	0xd1, 0x70, 0x0f, 0x5a, 0xd3, 0x27, 0xb2, 0x9e, 0x66, 0x60, 0xe9, 0xb4, 0x21, 0x92, 0x5c, 0x2b,
	0xab, 0xc2, 0x7e, 0xad, 0x4c, 0x9a, 0x42, 0xa5, 0x1c, 0x9c, 0xcc, 0xd5, 0x5c, 0x39, 0x15, 0x29,
	0x50, 0xd9, 0x30, 0xc0, 0x4c, 0x99, 0xa5, 0x32, 0x24, 0xa3, 0x06, 0x1a, 0x53, 0xa6, 0x84, 0x2c,
	0xeb, 0x63, 0x89, 0x0e, 0xd3, 0xca, 0x29, 0x3c, 0x46, 0x81, 0xe0, 0x91, 0x3f, 0xf2, 0xe3, 0x76,
	0x1a, 0x08, 0x1e, 0x86, 0xa8, 0x4d, 0x39, 0xd7, 0x51, 0x30, 0xf2, 0xe3, 0xa3, 0xd4, 0xe1, 0x82,
	0x63, 0x8a, 0x43, 0xd4, 0x2a, 0xb9, 0x02, 0x87, 0xa7, 0xa8, 0x9b, 0x53, 0x0d, 0xd2, 0x46, 0x6d,
	0xc7, 0x56, 0xb7, 0x42, 0x6b, 0x05, 0xe8, 0xa8, 0xe3, 0x1c, 0x1d, 0x1e, 0xbf, 0xfa, 0xe8, 0x20,
	0x85, 0x35, 0xc8, 0x15, 0x84, 0x43, 0xd4, 0xab, 0x53, 0xdc, 0x35, 0x83, 0x51, 0x4d, 0x5d, 0xf3,
	0x90, 0xa2, 0x4e, 0xf1, 0x54, 0x13, 0x05, 0xa3, 0x56, 0xdc, 0x3b, 0xef, 0x27, 0x65, 0x98, 0xa4,
	0x08, 0x53, 0xe7, 0x4e, 0xae, 0x94, 0x90, 0xb3, 0xb3, 0xcd, 0xe7, 0xd0, 0x7b, 0xff, 0x1a, 0xc6,
	0x73, 0x61, 0x1f, 0x56, 0x59, 0xc2, 0xd4, 0x92, 0x54, 0xc9, 0xcb, 0x63, 0x62, 0xf8, 0x82, 0xd8,
	0x97, 0x1c, 0x8c, 0x6b, 0x30, 0x69, 0xe9, 0x3c, 0xbb, 0xd9, 0xfc, 0x60, 0x6f, 0xb3, 0xc3, 0xfe,
	0x76, 0x87, 0xfd, 0xef, 0x1d, 0xf6, 0xdf, 0xf6, 0xd8, 0xdb, 0xee, 0xb1, 0xf7, 0xb1, 0xc7, 0xde,
	0xed, 0xe5, 0x3f, 0xbb, 0x47, 0x30, 0x06, 0x26, 0xb9, 0x90, 0x8b, 0x25, 0x95, 0x13, 0xa6, 0x81,
	0x34, 0x5b, 0x7b, 0xfe, 0xdb, 0x9b, 0x1b, 0x91, 0x75, 0xdd, 0xe7, 0x5e, 0xfc, 0x06, 0x00, 0x00,
	0xff, 0xff, 0xfe, 0x4b, 0xaa, 0xaf, 0xd9, 0x01, 0x00, 0x00,
}

func (m *Referral) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Referral) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Referral) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Tier != 0 {
		i = encodeVarintReferral(dAtA, i, uint64(m.Tier))
		i--
		dAtA[i] = 0x28
	}
	if len(m.Parent) > 0 {
		i -= len(m.Parent)
		copy(dAtA[i:], m.Parent)
		i = encodeVarintReferral(dAtA, i, uint64(len(m.Parent)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Code) > 0 {
		i -= len(m.Code)
		copy(dAtA[i:], m.Code)
		i = encodeVarintReferral(dAtA, i, uint64(len(m.Code)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Addr) > 0 {
		i -= len(m.Addr)
		copy(dAtA[i:], m.Addr)
		i = encodeVarintReferral(dAtA, i, uint64(len(m.Addr)))
		i--
		dAtA[i] = 0x12
	}
	if m.Id != 0 {
		i = encodeVarintReferral(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *Revenue) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Revenue) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Revenue) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Coins) > 0 {
		for iNdEx := len(m.Coins) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Coins[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintReferral(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if m.ReferralId != 0 {
		i = encodeVarintReferral(dAtA, i, uint64(m.ReferralId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintReferral(dAtA []byte, offset int, v uint64) int {
	offset -= sovReferral(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Referral) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovReferral(uint64(m.Id))
	}
	l = len(m.Addr)
	if l > 0 {
		n += 1 + l + sovReferral(uint64(l))
	}
	l = len(m.Code)
	if l > 0 {
		n += 1 + l + sovReferral(uint64(l))
	}
	l = len(m.Parent)
	if l > 0 {
		n += 1 + l + sovReferral(uint64(l))
	}
	if m.Tier != 0 {
		n += 1 + sovReferral(uint64(m.Tier))
	}
	return n
}

func (m *Revenue) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ReferralId != 0 {
		n += 1 + sovReferral(uint64(m.ReferralId))
	}
	if len(m.Coins) > 0 {
		for _, e := range m.Coins {
			l = e.Size()
			n += 1 + l + sovReferral(uint64(l))
		}
	}
	return n
}

func sovReferral(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozReferral(x uint64) (n int) {
	return sovReferral(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Referral) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowReferral
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
			return fmt.Errorf("proto: Referral: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Referral: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowReferral
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Addr", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowReferral
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
				return ErrInvalidLengthReferral
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthReferral
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Addr = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Code", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowReferral
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
				return ErrInvalidLengthReferral
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthReferral
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Code = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Parent", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowReferral
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
				return ErrInvalidLengthReferral
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthReferral
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Parent = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Tier", wireType)
			}
			m.Tier = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowReferral
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Tier |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipReferral(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthReferral
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
func (m *Revenue) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowReferral
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
			return fmt.Errorf("proto: Revenue: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Revenue: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ReferralId", wireType)
			}
			m.ReferralId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowReferral
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ReferralId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Coins", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowReferral
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
				return ErrInvalidLengthReferral
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthReferral
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Coins = append(m.Coins, types.Coin{})
			if err := m.Coins[len(m.Coins)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipReferral(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthReferral
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
func skipReferral(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowReferral
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
					return 0, ErrIntOverflowReferral
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
					return 0, ErrIntOverflowReferral
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
				return 0, ErrInvalidLengthReferral
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupReferral
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthReferral
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthReferral        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowReferral          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupReferral = fmt.Errorf("proto: unexpected end of group")
)
