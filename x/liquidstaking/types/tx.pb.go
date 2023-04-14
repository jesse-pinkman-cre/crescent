// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: crescent/liquidstaking/v1beta1/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
	_ "github.com/regen-network/cosmos-proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// MsgLiquidStake defines a SDK message for performing a liquid stake of coins
// from a delegator to whitelisted validators.
type MsgLiquidStake struct {
	DelegatorAddress string     `protobuf:"bytes,1,opt,name=delegator_address,json=delegatorAddress,proto3" json:"delegator_address,omitempty" yaml:"delegator_address"`
	Amount           types.Coin `protobuf:"bytes,2,opt,name=amount,proto3" json:"amount"`
}

func (m *MsgLiquidStake) Reset()         { *m = MsgLiquidStake{} }
func (m *MsgLiquidStake) String() string { return proto.CompactTextString(m) }
func (*MsgLiquidStake) ProtoMessage()    {}
func (*MsgLiquidStake) Descriptor() ([]byte, []int) {
	return fileDescriptor_9fe270968086aea1, []int{0}
}
func (m *MsgLiquidStake) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgLiquidStake) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgLiquidStake.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgLiquidStake) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgLiquidStake.Merge(m, src)
}
func (m *MsgLiquidStake) XXX_Size() int {
	return m.Size()
}
func (m *MsgLiquidStake) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgLiquidStake.DiscardUnknown(m)
}

var xxx_messageInfo_MsgLiquidStake proto.InternalMessageInfo

// MsgLiquidStakeResponse defines the Msg/LiquidStake response type.
type MsgLiquidStakeResponse struct {
}

func (m *MsgLiquidStakeResponse) Reset()         { *m = MsgLiquidStakeResponse{} }
func (m *MsgLiquidStakeResponse) String() string { return proto.CompactTextString(m) }
func (*MsgLiquidStakeResponse) ProtoMessage()    {}
func (*MsgLiquidStakeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_9fe270968086aea1, []int{1}
}
func (m *MsgLiquidStakeResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgLiquidStakeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgLiquidStakeResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgLiquidStakeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgLiquidStakeResponse.Merge(m, src)
}
func (m *MsgLiquidStakeResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgLiquidStakeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgLiquidStakeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgLiquidStakeResponse proto.InternalMessageInfo

// MsgLiquidUnstake defines a SDK message for performing an undelegation of liquid staking from a
// delegate.
type MsgLiquidUnstake struct {
	DelegatorAddress string     `protobuf:"bytes,1,opt,name=delegator_address,json=delegatorAddress,proto3" json:"delegator_address,omitempty" yaml:"delegator_address"`
	Amount           types.Coin `protobuf:"bytes,2,opt,name=amount,proto3" json:"amount"`
}

func (m *MsgLiquidUnstake) Reset()         { *m = MsgLiquidUnstake{} }
func (m *MsgLiquidUnstake) String() string { return proto.CompactTextString(m) }
func (*MsgLiquidUnstake) ProtoMessage()    {}
func (*MsgLiquidUnstake) Descriptor() ([]byte, []int) {
	return fileDescriptor_9fe270968086aea1, []int{2}
}
func (m *MsgLiquidUnstake) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgLiquidUnstake) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgLiquidUnstake.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgLiquidUnstake) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgLiquidUnstake.Merge(m, src)
}
func (m *MsgLiquidUnstake) XXX_Size() int {
	return m.Size()
}
func (m *MsgLiquidUnstake) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgLiquidUnstake.DiscardUnknown(m)
}

var xxx_messageInfo_MsgLiquidUnstake proto.InternalMessageInfo

// MsgLiquidUnstakeResponse defines the Msg/LiquidUnstake response type.
type MsgLiquidUnstakeResponse struct {
	CompletionTime time.Time `protobuf:"bytes,1,opt,name=completion_time,json=completionTime,proto3,stdtime" json:"completion_time"`
}

func (m *MsgLiquidUnstakeResponse) Reset()         { *m = MsgLiquidUnstakeResponse{} }
func (m *MsgLiquidUnstakeResponse) String() string { return proto.CompactTextString(m) }
func (*MsgLiquidUnstakeResponse) ProtoMessage()    {}
func (*MsgLiquidUnstakeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_9fe270968086aea1, []int{3}
}
func (m *MsgLiquidUnstakeResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgLiquidUnstakeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgLiquidUnstakeResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgLiquidUnstakeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgLiquidUnstakeResponse.Merge(m, src)
}
func (m *MsgLiquidUnstakeResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgLiquidUnstakeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgLiquidUnstakeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgLiquidUnstakeResponse proto.InternalMessageInfo

func (m *MsgLiquidUnstakeResponse) GetCompletionTime() time.Time {
	if m != nil {
		return m.CompletionTime
	}
	return time.Time{}
}

func init() {
	proto.RegisterType((*MsgLiquidStake)(nil), "crescent.liquidstaking.v1beta1.MsgLiquidStake")
	proto.RegisterType((*MsgLiquidStakeResponse)(nil), "crescent.liquidstaking.v1beta1.MsgLiquidStakeResponse")
	proto.RegisterType((*MsgLiquidUnstake)(nil), "crescent.liquidstaking.v1beta1.MsgLiquidUnstake")
	proto.RegisterType((*MsgLiquidUnstakeResponse)(nil), "crescent.liquidstaking.v1beta1.MsgLiquidUnstakeResponse")
}

func init() {
	proto.RegisterFile("crescent/liquidstaking/v1beta1/tx.proto", fileDescriptor_9fe270968086aea1)
}

var fileDescriptor_9fe270968086aea1 = []byte{
	// 451 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x93, 0x3f, 0x6f, 0xd4, 0x30,
	0x18, 0xc6, 0x63, 0x40, 0x55, 0xf1, 0x89, 0x52, 0x22, 0x84, 0xd2, 0x08, 0x25, 0x55, 0x16, 0xba,
	0xd4, 0xa6, 0x87, 0x04, 0xa8, 0x4c, 0x1c, 0x13, 0x12, 0xb7, 0x1c, 0x74, 0x61, 0x39, 0x39, 0xb9,
	0x17, 0x63, 0x2e, 0xb6, 0x43, 0xec, 0xa0, 0x56, 0x7c, 0x01, 0xc6, 0x7e, 0x02, 0x74, 0x1f, 0xa7,
	0x63, 0x47, 0xa6, 0x82, 0xee, 0x16, 0x66, 0x36, 0x36, 0x94, 0xbf, 0x90, 0x22, 0x21, 0xba, 0x75,
	0xcb, 0xeb, 0xe7, 0x79, 0xf3, 0xfe, 0x9c, 0x27, 0x2f, 0xbe, 0x97, 0xe4, 0x60, 0x12, 0x50, 0x96,
	0xa6, 0xe2, 0x7d, 0x21, 0x66, 0xc6, 0xb2, 0xb9, 0x50, 0x9c, 0x7e, 0xd8, 0x8b, 0xc1, 0xb2, 0x3d,
	0x6a, 0x0f, 0x49, 0x96, 0x6b, 0xab, 0xdd, 0xa0, 0x35, 0x92, 0x9e, 0x91, 0x34, 0x46, 0xff, 0x36,
	0xd7, 0x5c, 0x57, 0x56, 0x5a, 0x3e, 0xd5, 0x5d, 0xfe, 0x56, 0xa2, 0x8d, 0xd4, 0x66, 0x5a, 0x0b,
	0x75, 0xd1, 0x48, 0x41, 0x5d, 0xd1, 0x98, 0x19, 0xe8, 0xc6, 0x25, 0x5a, 0xa8, 0x46, 0x0f, 0xb9,
	0xd6, 0x3c, 0x05, 0x5a, 0x55, 0x71, 0xf1, 0x86, 0x5a, 0x21, 0xc1, 0x58, 0x26, 0xb3, 0xda, 0x10,
	0x7d, 0x46, 0x78, 0x63, 0x6c, 0xf8, 0x8b, 0x0a, 0xe7, 0xa5, 0x65, 0x73, 0x70, 0x9f, 0xe3, 0x5b,
	0x33, 0x48, 0x81, 0x33, 0xab, 0xf3, 0x29, 0x9b, 0xcd, 0x72, 0x30, 0xc6, 0x43, 0xdb, 0x68, 0xe7,
	0xfa, 0xe8, 0xee, 0x8f, 0xb3, 0xd0, 0x3b, 0x62, 0x32, 0xdd, 0x8f, 0xfe, 0xb2, 0x44, 0x93, 0xcd,
	0xee, 0xec, 0x69, 0x7d, 0xe4, 0x3e, 0xc2, 0x6b, 0x4c, 0xea, 0x42, 0x59, 0xef, 0xca, 0x36, 0xda,
	0x19, 0x0c, 0xb7, 0x48, 0x43, 0x5f, 0xf2, 0xb6, 0xb7, 0x26, 0xcf, 0xb4, 0x50, 0xa3, 0x6b, 0x27,
	0x67, 0xa1, 0x33, 0x69, 0xec, 0xfb, 0xeb, 0x9f, 0x16, 0xa1, 0xf3, 0x7d, 0x11, 0x3a, 0x91, 0x87,
	0xef, 0xf4, 0xf9, 0x26, 0x60, 0x32, 0xad, 0x0c, 0x44, 0x0b, 0x84, 0x37, 0x3b, 0xe9, 0x40, 0x99,
	0x4b, 0x08, 0x2f, 0xb0, 0x77, 0x9e, 0xb0, 0xc5, 0x77, 0xc7, 0xf8, 0x66, 0xa2, 0x65, 0x96, 0x82,
	0x15, 0x5a, 0x4d, 0xcb, 0x5c, 0x2a, 0xce, 0xc1, 0xd0, 0x27, 0x75, 0x68, 0xa4, 0x0d, 0x8d, 0xbc,
	0x6a, 0x43, 0x1b, 0xad, 0x97, 0x83, 0x8e, 0xbf, 0x86, 0x68, 0xb2, 0xf1, 0xbb, 0xb9, 0x94, 0x87,
	0x3f, 0x11, 0xbe, 0x3a, 0x36, 0xdc, 0x2d, 0xf0, 0xe0, 0xcf, 0x30, 0x09, 0xf9, 0xf7, 0x2f, 0x47,
	0xfa, 0x1f, 0xd7, 0x7f, 0x78, 0x31, 0x7f, 0x77, 0x9b, 0x8f, 0xf8, 0x46, 0x3f, 0x88, 0xfb, 0xff,
	0xfd, 0xa2, 0xa6, 0xc3, 0x7f, 0x7c, 0xd1, 0x8e, 0x76, 0xf8, 0xe8, 0xe0, 0x64, 0x19, 0xa0, 0xd3,
	0x65, 0x80, 0xbe, 0x2d, 0x03, 0x74, 0xbc, 0x0a, 0x9c, 0xd3, 0x55, 0xe0, 0x7c, 0x59, 0x05, 0xce,
	0xeb, 0x27, 0x5c, 0xd8, 0xb7, 0x45, 0x4c, 0x12, 0x2d, 0xe9, 0x3b, 0x30, 0x06, 0x76, 0x33, 0xa1,
	0xe6, 0x92, 0xa9, 0xdd, 0x24, 0x07, 0xda, 0xad, 0xed, 0xe1, 0xb9, 0xc5, 0xb5, 0x47, 0x19, 0x98,
	0x78, 0xad, 0x0a, 0xe0, 0xc1, 0xaf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xde, 0xef, 0x0e, 0xa4, 0xdf,
	0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgClient interface {
	// LiquidStake defines a method for performing a delegation of coins
	// from a delegator to whitelisted validators.
	LiquidStake(ctx context.Context, in *MsgLiquidStake, opts ...grpc.CallOption) (*MsgLiquidStakeResponse, error)
	// LiquidUnstake defines a method for performing an undelegation of liquid staking from a
	// delegate.
	LiquidUnstake(ctx context.Context, in *MsgLiquidUnstake, opts ...grpc.CallOption) (*MsgLiquidUnstakeResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) LiquidStake(ctx context.Context, in *MsgLiquidStake, opts ...grpc.CallOption) (*MsgLiquidStakeResponse, error) {
	out := new(MsgLiquidStakeResponse)
	err := c.cc.Invoke(ctx, "/crescent.liquidstaking.v1beta1.Msg/LiquidStake", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) LiquidUnstake(ctx context.Context, in *MsgLiquidUnstake, opts ...grpc.CallOption) (*MsgLiquidUnstakeResponse, error) {
	out := new(MsgLiquidUnstakeResponse)
	err := c.cc.Invoke(ctx, "/crescent.liquidstaking.v1beta1.Msg/LiquidUnstake", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	// LiquidStake defines a method for performing a delegation of coins
	// from a delegator to whitelisted validators.
	LiquidStake(context.Context, *MsgLiquidStake) (*MsgLiquidStakeResponse, error)
	// LiquidUnstake defines a method for performing an undelegation of liquid staking from a
	// delegate.
	LiquidUnstake(context.Context, *MsgLiquidUnstake) (*MsgLiquidUnstakeResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) LiquidStake(ctx context.Context, req *MsgLiquidStake) (*MsgLiquidStakeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LiquidStake not implemented")
}
func (*UnimplementedMsgServer) LiquidUnstake(ctx context.Context, req *MsgLiquidUnstake) (*MsgLiquidUnstakeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LiquidUnstake not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_LiquidStake_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgLiquidStake)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).LiquidStake(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/crescent.liquidstaking.v1beta1.Msg/LiquidStake",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).LiquidStake(ctx, req.(*MsgLiquidStake))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_LiquidUnstake_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgLiquidUnstake)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).LiquidUnstake(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/crescent.liquidstaking.v1beta1.Msg/LiquidUnstake",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).LiquidUnstake(ctx, req.(*MsgLiquidUnstake))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "crescent.liquidstaking.v1beta1.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "LiquidStake",
			Handler:    _Msg_LiquidStake_Handler,
		},
		{
			MethodName: "LiquidUnstake",
			Handler:    _Msg_LiquidUnstake_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "crescent/liquidstaking/v1beta1/tx.proto",
}

func (m *MsgLiquidStake) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgLiquidStake) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgLiquidStake) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Amount.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTx(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.DelegatorAddress) > 0 {
		i -= len(m.DelegatorAddress)
		copy(dAtA[i:], m.DelegatorAddress)
		i = encodeVarintTx(dAtA, i, uint64(len(m.DelegatorAddress)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgLiquidStakeResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgLiquidStakeResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgLiquidStakeResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *MsgLiquidUnstake) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgLiquidUnstake) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgLiquidUnstake) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Amount.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTx(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.DelegatorAddress) > 0 {
		i -= len(m.DelegatorAddress)
		copy(dAtA[i:], m.DelegatorAddress)
		i = encodeVarintTx(dAtA, i, uint64(len(m.DelegatorAddress)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgLiquidUnstakeResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgLiquidUnstakeResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgLiquidUnstakeResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	n3, err3 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.CompletionTime, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.CompletionTime):])
	if err3 != nil {
		return 0, err3
	}
	i -= n3
	i = encodeVarintTx(dAtA, i, uint64(n3))
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintTx(dAtA []byte, offset int, v uint64) int {
	offset -= sovTx(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgLiquidStake) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.DelegatorAddress)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = m.Amount.Size()
	n += 1 + l + sovTx(uint64(l))
	return n
}

func (m *MsgLiquidStakeResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *MsgLiquidUnstake) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.DelegatorAddress)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = m.Amount.Size()
	n += 1 + l + sovTx(uint64(l))
	return n
}

func (m *MsgLiquidUnstakeResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.CompletionTime)
	n += 1 + l + sovTx(uint64(l))
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgLiquidStake) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgLiquidStake: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgLiquidStake: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DelegatorAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DelegatorAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Amount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgLiquidStakeResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgLiquidStakeResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgLiquidStakeResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgLiquidUnstake) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgLiquidUnstake: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgLiquidUnstake: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DelegatorAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DelegatorAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Amount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgLiquidUnstakeResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgLiquidUnstakeResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgLiquidUnstakeResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CompletionTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.CompletionTime, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func skipTx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
				return 0, ErrInvalidLengthTx
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTx
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTx
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTx        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTx          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTx = fmt.Errorf("proto: unexpected end of group")
)
