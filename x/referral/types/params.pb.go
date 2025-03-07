// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: crescent/referral/v1beta1/params.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
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

type Params struct {
	// uint32 max_parent_depth = 1;
	MaxParentDepth           github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,1,opt,name=max_parent_depth,json=maxParentDepth,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"max_parent_depth"`
	SwapFeeRate              github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,2,opt,name=swap_fee_rate,json=swapFeeRate,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"swap_fee_rate"`
	SwapFeeReferralRate      github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,3,opt,name=swap_fee_referral_rate,json=swapFeeReferralRate,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"swap_fee_referral_rate"`
	SwapFeeParentRate        github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,4,opt,name=swap_fee_parent_rate,json=swapFeeParentRate,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"swap_fee_parent_rate"`
	MinNumChildForTierSilver github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,5,opt,name=min_num_child_for_tier_silver,json=minNumChildForTierSilver,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"min_num_child_for_tier_silver"`
	FeeCollectorAddress      string                                 `protobuf:"bytes,6,opt,name=fee_collector_address,json=feeCollectorAddress,proto3" json:"fee_collector_address,omitempty"`
}

func (m *Params) Reset()      { *m = Params{} }
func (*Params) ProtoMessage() {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_96228ce6359f5682, []int{0}
}
func (m *Params) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Params) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Params.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Params) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Params.Merge(m, src)
}
func (m *Params) XXX_Size() int {
	return m.Size()
}
func (m *Params) XXX_DiscardUnknown() {
	xxx_messageInfo_Params.DiscardUnknown(m)
}

var xxx_messageInfo_Params proto.InternalMessageInfo

func (m *Params) GetFeeCollectorAddress() string {
	if m != nil {
		return m.FeeCollectorAddress
	}
	return ""
}

func init() {
	proto.RegisterType((*Params)(nil), "crescent.referral.v1beta1.Params")
}

func init() {
	proto.RegisterFile("crescent/referral/v1beta1/params.proto", fileDescriptor_96228ce6359f5682)
}

var fileDescriptor_96228ce6359f5682 = []byte{
	// 399 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x92, 0xb1, 0x8e, 0xd3, 0x40,
	0x10, 0x86, 0x6d, 0x08, 0x91, 0x58, 0x04, 0x02, 0xdf, 0x81, 0x0c, 0x12, 0x0e, 0xa2, 0x38, 0xd1,
	0xc4, 0xd6, 0x81, 0x44, 0x41, 0xc7, 0x25, 0x8a, 0x44, 0x83, 0x22, 0x43, 0x81, 0x68, 0x56, 0x9b,
	0xf5, 0x38, 0x59, 0xe2, 0xdd, 0xb5, 0x66, 0x37, 0x21, 0xbc, 0x05, 0x25, 0x25, 0x1d, 0xaf, 0x92,
	0x32, 0x25, 0xa2, 0x88, 0x50, 0xf2, 0x22, 0xc8, 0x6b, 0x3b, 0xa1, 0xbd, 0x54, 0xb6, 0x34, 0xbf,
	0xbf, 0x6f, 0x66, 0x3c, 0xe4, 0x82, 0x23, 0x18, 0x0e, 0xca, 0x26, 0x08, 0x39, 0x20, 0xb2, 0x22,
	0x59, 0x5e, 0x4e, 0xc0, 0xb2, 0xcb, 0xa4, 0x64, 0xc8, 0xa4, 0x89, 0x4b, 0xd4, 0x56, 0x07, 0x8f,
	0xdb, 0x5c, 0xdc, 0xe6, 0xe2, 0x26, 0xf7, 0xe4, 0x7c, 0xaa, 0xa7, 0xda, 0xa5, 0x92, 0xea, 0xad,
	0xfe, 0xe0, 0xf9, 0xaf, 0x0e, 0xe9, 0x8e, 0x1d, 0x21, 0xf8, 0x44, 0xee, 0x4b, 0xb6, 0xa2, 0x25,
	0x43, 0x50, 0x96, 0x66, 0x50, 0xda, 0x59, 0xe8, 0x3f, 0xf3, 0x5f, 0xdc, 0xbe, 0x8a, 0xd7, 0xdb,
	0x9e, 0xf7, 0x67, 0xdb, 0xbb, 0x98, 0x0a, 0x3b, 0x5b, 0x4c, 0x62, 0xae, 0x65, 0xc2, 0xb5, 0x91,
	0xda, 0x34, 0x8f, 0xbe, 0xc9, 0xe6, 0x89, 0xfd, 0x56, 0x82, 0x89, 0xdf, 0x29, 0x9b, 0xde, 0x93,
	0x6c, 0x35, 0x76, 0x98, 0x61, 0x45, 0x09, 0x52, 0x72, 0xd7, 0x7c, 0x65, 0x25, 0xcd, 0x01, 0x28,
	0x32, 0x0b, 0xe1, 0x8d, 0x6b, 0x63, 0x87, 0xc0, 0xd3, 0x3b, 0x15, 0x64, 0x04, 0x90, 0x32, 0x0b,
	0x01, 0x27, 0x8f, 0x8e, 0xcc, 0x66, 0xd6, 0x1a, 0x7e, 0xf3, 0x24, 0xf8, 0x59, 0x0b, 0x6f, 0x58,
	0x4e, 0x42, 0xc9, 0xf9, 0x41, 0xd2, 0xec, 0xc5, 0x29, 0x3a, 0x27, 0x29, 0x1e, 0x34, 0x8a, 0x7a,
	0x35, 0x4e, 0xa0, 0xc9, 0x53, 0x29, 0x14, 0x55, 0x0b, 0x49, 0xf9, 0x4c, 0x14, 0x19, 0xcd, 0x35,
	0x52, 0x2b, 0x00, 0xa9, 0x11, 0xc5, 0x12, 0x30, 0xbc, 0x75, 0xd2, 0x0f, 0x08, 0xa5, 0x50, 0xef,
	0x17, 0x72, 0x50, 0x21, 0x47, 0x1a, 0x3f, 0x0a, 0xc0, 0x0f, 0x8e, 0x17, 0xbc, 0x24, 0x0f, 0xab,
	0x61, 0xb8, 0x2e, 0x0a, 0xe0, 0x56, 0x23, 0x65, 0x59, 0x86, 0x60, 0x4c, 0xd8, 0xad, 0x44, 0xe9,
	0x59, 0x0e, 0x30, 0x68, 0x6b, 0x6f, 0xeb, 0xd2, 0x9b, 0xce, 0x8f, 0x9f, 0x3d, 0xef, 0x6a, 0xbc,
	0xde, 0x45, 0xfe, 0x66, 0x17, 0xf9, 0x7f, 0x77, 0x91, 0xff, 0x7d, 0x1f, 0x79, 0x9b, 0x7d, 0xe4,
	0xfd, 0xde, 0x47, 0xde, 0xe7, 0xd7, 0xff, 0x75, 0xf5, 0x05, 0x8c, 0x81, 0x7e, 0x29, 0xd4, 0x5c,
	0x32, 0xd5, 0xe7, 0x08, 0xc9, 0xe1, 0x72, 0x57, 0xc7, 0xdb, 0x75, 0x9d, 0x4e, 0xba, 0xee, 0x04,
	0x5f, 0xfd, 0x0b, 0x00, 0x00, 0xff, 0xff, 0xd1, 0x24, 0xcb, 0xa3, 0xdd, 0x02, 0x00, 0x00,
}

func (m *Params) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Params) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Params) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.FeeCollectorAddress) > 0 {
		i -= len(m.FeeCollectorAddress)
		copy(dAtA[i:], m.FeeCollectorAddress)
		i = encodeVarintParams(dAtA, i, uint64(len(m.FeeCollectorAddress)))
		i--
		dAtA[i] = 0x32
	}
	{
		size := m.MinNumChildForTierSilver.Size()
		i -= size
		if _, err := m.MinNumChildForTierSilver.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x2a
	{
		size := m.SwapFeeParentRate.Size()
		i -= size
		if _, err := m.SwapFeeParentRate.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	{
		size := m.SwapFeeReferralRate.Size()
		i -= size
		if _, err := m.SwapFeeReferralRate.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	{
		size := m.SwapFeeRate.Size()
		i -= size
		if _, err := m.SwapFeeRate.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	{
		size := m.MaxParentDepth.Size()
		i -= size
		if _, err := m.MaxParentDepth.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintParams(dAtA []byte, offset int, v uint64) int {
	offset -= sovParams(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.MaxParentDepth.Size()
	n += 1 + l + sovParams(uint64(l))
	l = m.SwapFeeRate.Size()
	n += 1 + l + sovParams(uint64(l))
	l = m.SwapFeeReferralRate.Size()
	n += 1 + l + sovParams(uint64(l))
	l = m.SwapFeeParentRate.Size()
	n += 1 + l + sovParams(uint64(l))
	l = m.MinNumChildForTierSilver.Size()
	n += 1 + l + sovParams(uint64(l))
	l = len(m.FeeCollectorAddress)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
	}
	return n
}

func sovParams(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozParams(x uint64) (n int) {
	return sovParams(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowParams
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
			return fmt.Errorf("proto: Params: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Params: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxParentDepth", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.MaxParentDepth.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SwapFeeRate", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.SwapFeeRate.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SwapFeeReferralRate", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.SwapFeeReferralRate.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SwapFeeParentRate", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.SwapFeeParentRate.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinNumChildForTierSilver", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.MinNumChildForTierSilver.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FeeCollectorAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.FeeCollectorAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipParams(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthParams
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
func skipParams(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
				return 0, ErrInvalidLengthParams
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupParams
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthParams
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthParams        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowParams          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupParams = fmt.Errorf("proto: unexpected end of group")
)
