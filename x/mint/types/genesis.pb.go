// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: crescent/mint/v1beta1/genesis.proto

package types

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	io "io"
	math "math"
	math_bits "math/bits"
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// GenesisState defines the mint module's genesis state.
type GenesisState struct {
	// last_block_time defines the last block time, which is used to calculate inflation.
	LastBlockTime *time.Time `protobuf:"bytes,1,opt,name=last_block_time,json=lastBlockTime,proto3,stdtime" json:"last_block_time,omitempty" yaml:"last_block_time"`
	// params defines all the parameters of the module.
	Params Params `protobuf:"bytes,2,opt,name=params,proto3" json:"params"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_a1d3423edceca2d4, []int{0}
}
func (m *GenesisState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisState.Merge(m, src)
}
func (m *GenesisState) XXX_Size() int {
	return m.Size()
}
func (m *GenesisState) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisState.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisState proto.InternalMessageInfo

func (m *GenesisState) GetLastBlockTime() *time.Time {
	if m != nil {
		return m.LastBlockTime
	}
	return nil
}

func (m *GenesisState) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "crescent.mint.v1beta1.GenesisState")
}

func init() {
	proto.RegisterFile("crescent/mint/v1beta1/genesis.proto", fileDescriptor_a1d3423edceca2d4)
}

var fileDescriptor_a1d3423edceca2d4 = []byte{
	// 297 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0x31, 0x4e, 0xc3, 0x30,
	0x14, 0x86, 0x63, 0x84, 0x3a, 0x04, 0x10, 0x52, 0x05, 0xa8, 0xaa, 0x84, 0x53, 0x85, 0x85, 0xa5,
	0xb6, 0x5a, 0x36, 0xd8, 0xb2, 0xb0, 0x30, 0xa0, 0xc2, 0xc4, 0x52, 0xd9, 0xd1, 0xc3, 0x84, 0xc6,
	0x71, 0x14, 0xbb, 0x88, 0xde, 0xa2, 0x37, 0xe1, 0x1a, 0x1d, 0x3b, 0x32, 0x15, 0x94, 0xdc, 0x80,
	0x13, 0x20, 0xdb, 0x0d, 0x03, 0xea, 0x66, 0x3f, 0x7d, 0xff, 0xa7, 0xf7, 0xfe, 0xf0, 0x22, 0xad,
	0x40, 0xa7, 0x50, 0x18, 0x2a, 0xb3, 0xc2, 0xd0, 0xb7, 0x11, 0x07, 0xc3, 0x46, 0x54, 0x40, 0x01,
	0x3a, 0xd3, 0xa4, 0xac, 0x94, 0x51, 0xdd, 0xd3, 0x16, 0x22, 0x16, 0x22, 0x5b, 0xa8, 0x7f, 0x22,
	0x94, 0x50, 0x8e, 0xa0, 0xf6, 0xe5, 0xe1, 0xfe, 0x60, 0xb7, 0xd1, 0x25, 0x3d, 0x11, 0x09, 0xa5,
	0x44, 0x0e, 0xd4, 0xfd, 0xf8, 0xfc, 0x99, 0x9a, 0x4c, 0x82, 0x36, 0x4c, 0x96, 0x1e, 0x88, 0x3f,
	0x50, 0x78, 0x78, 0xeb, 0x37, 0x78, 0x30, 0xcc, 0x40, 0x97, 0x87, 0xc7, 0x39, 0xd3, 0x66, 0xca,
	0x73, 0x95, 0xce, 0xa6, 0x16, 0xef, 0xa1, 0x01, 0xba, 0x3c, 0x18, 0xf7, 0x89, 0x77, 0x91, 0xd6,
	0x45, 0x1e, 0x5b, 0x57, 0x82, 0x7f, 0x36, 0xd1, 0xd9, 0x82, 0xc9, 0xfc, 0x3a, 0xfe, 0x17, 0x8e,
	0x97, 0x5f, 0x11, 0x9a, 0x1c, 0xd9, 0x69, 0x62, 0x87, 0x36, 0xd3, 0xbd, 0x09, 0x3b, 0x25, 0xab,
	0x98, 0xd4, 0xbd, 0x3d, 0xa7, 0x3e, 0x27, 0x3b, 0xaf, 0x26, 0xf7, 0x0e, 0x4a, 0xf6, 0x57, 0x9b,
	0x28, 0x98, 0x6c, 0x23, 0xc9, 0xdd, 0xaa, 0xc6, 0x68, 0x5d, 0x63, 0xf4, 0x5d, 0x63, 0xb4, 0x6c,
	0x70, 0xb0, 0x6e, 0x70, 0xf0, 0xd9, 0xe0, 0xe0, 0x69, 0x2c, 0x32, 0xf3, 0x32, 0xe7, 0x24, 0x55,
	0x92, 0xbe, 0x82, 0xd6, 0x30, 0x2c, 0xb3, 0x62, 0x26, 0x59, 0x31, 0x4c, 0x2b, 0xa0, 0x7f, 0x5d,
	0xbd, 0xfb, 0xb6, 0xcc, 0xa2, 0x04, 0xcd, 0x3b, 0xee, 0x9a, 0xab, 0xdf, 0x00, 0x00, 0x00, 0xff,
	0xff, 0xf6, 0x42, 0x85, 0xfb, 0x9d, 0x01, 0x00, 0x00,
}

func (m *GenesisState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if m.LastBlockTime != nil {
		n2, err2 := github_com_gogo_protobuf_types.StdTimeMarshalTo(*m.LastBlockTime, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(*m.LastBlockTime):])
		if err2 != nil {
			return 0, err2
		}
		i -= n2
		i = encodeVarintGenesis(dAtA, i, uint64(n2))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintGenesis(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenesis(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.LastBlockTime != nil {
		l = github_com_gogo_protobuf_types.SizeOfStdTime(*m.LastBlockTime)
		n += 1 + l + sovGenesis(uint64(l))
	}
	l = m.Params.Size()
	n += 1 + l + sovGenesis(uint64(l))
	return n
}

func sovGenesis(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenesis(x uint64) (n int) {
	return sovGenesis(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GenesisState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: GenesisState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastBlockTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.LastBlockTime == nil {
				m.LastBlockTime = new(time.Time)
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(m.LastBlockTime, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func skipGenesis(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
				return 0, ErrInvalidLengthGenesis
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGenesis
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGenesis
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGenesis        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenesis          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGenesis = fmt.Errorf("proto: unexpected end of group")
)
