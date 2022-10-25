// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: osmosis/concentrated-liquidity/concentratedPool.proto

// this is a legacy package that requires additional migration logic
// in order to use the correct packge. Decision made to use legacy package path
// until clear steps for migration logic and the unknowns for state breaking are
// investigated for changing proto package.

package concentrated_liquidity

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
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

type Pool struct {
	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty" yaml:"address"`
	Id      uint64 `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	// Amount of total liquidity
	Liquidity        github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,3,opt,name=liquidity,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"liquidity" yaml:"liquidity"`
	Token0           string                                 `protobuf:"bytes,4,opt,name=token0,proto3" json:"token0,omitempty"`
	Token1           string                                 `protobuf:"bytes,5,opt,name=token1,proto3" json:"token1,omitempty"`
	CurrentSqrtPrice github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,6,opt,name=current_sqrt_price,json=currentSqrtPrice,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"current_sqrt_price" yaml:"spot_price"`
	CurrentTick      github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,7,opt,name=current_tick,json=currentTick,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"current_tick" yaml:"current_tick"`
	GlobalFees       *Fee                                   `protobuf:"bytes,8,opt,name=globalFees,proto3" json:"globalFees,omitempty"`
}

func (m *Pool) Reset()      { *m = Pool{} }
func (*Pool) ProtoMessage() {}
func (*Pool) Descriptor() ([]byte, []int) {
	return fileDescriptor_b144264ce94bcf63, []int{0}
}
func (m *Pool) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Pool) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Pool.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Pool) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Pool.Merge(m, src)
}
func (m *Pool) XXX_Size() int {
	return m.Size()
}
func (m *Pool) XXX_DiscardUnknown() {
	xxx_messageInfo_Pool.DiscardUnknown(m)
}

var xxx_messageInfo_Pool proto.InternalMessageInfo

type TickInfo struct {
	LiquidityGross       github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,1,opt,name=liquidity_gross,json=liquidityGross,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"liquidity_gross" yaml:"liquidity_gross"`
	LiquidityNet         github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,2,opt,name=liquidity_net,json=liquidityNet,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"liquidity_net" yaml:"liquidity_net"`
	FeesAccruedAboveTick *Fee                                   `protobuf:"bytes,3,opt,name=feesAccruedAboveTick,proto3" json:"feesAccruedAboveTick,omitempty"`
}

func (m *TickInfo) Reset()         { *m = TickInfo{} }
func (m *TickInfo) String() string { return proto.CompactTextString(m) }
func (*TickInfo) ProtoMessage()    {}
func (*TickInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_b144264ce94bcf63, []int{1}
}
func (m *TickInfo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TickInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TickInfo.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TickInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TickInfo.Merge(m, src)
}
func (m *TickInfo) XXX_Size() int {
	return m.Size()
}
func (m *TickInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_TickInfo.DiscardUnknown(m)
}

var xxx_messageInfo_TickInfo proto.InternalMessageInfo

func (m *TickInfo) GetFeesAccruedAboveTick() *Fee {
	if m != nil {
		return m.FeesAccruedAboveTick
	}
	return nil
}

type Position struct {
	Liquidity      github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,1,opt,name=liquidity,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"liquidity" yaml:"liquidity"`
	FeesAtCreation *Fee                                   `protobuf:"bytes,2,opt,name=feesAtCreation,proto3" json:"feesAtCreation,omitempty"`
}

func (m *Position) Reset()         { *m = Position{} }
func (m *Position) String() string { return proto.CompactTextString(m) }
func (*Position) ProtoMessage()    {}
func (*Position) Descriptor() ([]byte, []int) {
	return fileDescriptor_b144264ce94bcf63, []int{2}
}
func (m *Position) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Position) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Position.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Position) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Position.Merge(m, src)
}
func (m *Position) XXX_Size() int {
	return m.Size()
}
func (m *Position) XXX_DiscardUnknown() {
	xxx_messageInfo_Position.DiscardUnknown(m)
}

var xxx_messageInfo_Position proto.InternalMessageInfo

func (m *Position) GetFeesAtCreation() *Fee {
	if m != nil {
		return m.FeesAtCreation
	}
	return nil
}

type Fee struct {
	Token0 github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,1,opt,name=token0,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"token0" yaml:"token0"`
	Token1 github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,2,opt,name=token1,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"token1" yaml:"token1"`
}

func (m *Fee) Reset()         { *m = Fee{} }
func (m *Fee) String() string { return proto.CompactTextString(m) }
func (*Fee) ProtoMessage()    {}
func (*Fee) Descriptor() ([]byte, []int) {
	return fileDescriptor_b144264ce94bcf63, []int{3}
}
func (m *Fee) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Fee) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Fee.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Fee) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Fee.Merge(m, src)
}
func (m *Fee) XXX_Size() int {
	return m.Size()
}
func (m *Fee) XXX_DiscardUnknown() {
	xxx_messageInfo_Fee.DiscardUnknown(m)
}

var xxx_messageInfo_Fee proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Pool)(nil), "osmosis.concentratedliquidity.v1beta1.Pool")
	proto.RegisterType((*TickInfo)(nil), "osmosis.concentratedliquidity.v1beta1.TickInfo")
	proto.RegisterType((*Position)(nil), "osmosis.concentratedliquidity.v1beta1.Position")
	proto.RegisterType((*Fee)(nil), "osmosis.concentratedliquidity.v1beta1.Fee")
}

func init() {
	proto.RegisterFile("osmosis/concentrated-liquidity/concentratedPool.proto", fileDescriptor_b144264ce94bcf63)
}

var fileDescriptor_b144264ce94bcf63 = []byte{
	// 608 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x94, 0x4d, 0x6f, 0xd3, 0x4c,
	0x10, 0xc7, 0xed, 0x34, 0x7d, 0xdb, 0xb6, 0x79, 0xda, 0x7d, 0xaa, 0xca, 0xf4, 0x60, 0x57, 0x2b,
	0x81, 0x22, 0x44, 0x1c, 0x5c, 0xc4, 0xa5, 0x42, 0x42, 0x4d, 0x21, 0x10, 0x0e, 0x28, 0x32, 0x88,
	0x4a, 0x1c, 0x08, 0x7e, 0x99, 0xa6, 0xab, 0x38, 0xde, 0xc4, 0xbb, 0x89, 0xc8, 0x37, 0xe0, 0x88,
	0x38, 0x71, 0xec, 0x87, 0x40, 0xe2, 0xca, 0xb1, 0xea, 0xa9, 0x47, 0xc4, 0x21, 0x42, 0xc9, 0x37,
	0xc8, 0x27, 0x40, 0x7e, 0x89, 0x9b, 0x46, 0x20, 0x61, 0x09, 0x4e, 0xc9, 0x8c, 0x77, 0x7e, 0xff,
	0xff, 0x8e, 0x67, 0x8c, 0xee, 0x33, 0xde, 0x66, 0x9c, 0xf2, 0xb2, 0xc3, 0x7c, 0x07, 0x7c, 0x11,
	0x58, 0x02, 0xdc, 0x92, 0x47, 0xbb, 0x3d, 0xea, 0x52, 0x31, 0xb8, 0x96, 0xae, 0x33, 0xe6, 0xe9,
	0x9d, 0x80, 0x09, 0x86, 0x6f, 0x26, 0x65, 0xfa, 0xec, 0xf3, 0xb4, 0x4a, 0xef, 0x1b, 0x36, 0x08,
	0xcb, 0xd8, 0xbd, 0xe1, 0x44, 0xe7, 0x1a, 0x51, 0x51, 0x39, 0x0e, 0x62, 0xc2, 0xee, 0x76, 0x93,
	0x35, 0x59, 0x9c, 0x0f, 0xff, 0xc5, 0x59, 0xf2, 0x31, 0x8f, 0xf2, 0xa1, 0x0c, 0xbe, 0x83, 0x96,
	0x2d, 0xd7, 0x0d, 0x80, 0x73, 0x45, 0xde, 0x93, 0x8b, 0xab, 0x15, 0x3c, 0x19, 0x6a, 0x85, 0x81,
	0xd5, 0xf6, 0x0e, 0x48, 0xf2, 0x80, 0x98, 0xd3, 0x23, 0xb8, 0x80, 0x72, 0xd4, 0x55, 0x72, 0x7b,
	0x72, 0x31, 0x6f, 0xe6, 0xa8, 0x8b, 0xdf, 0xa2, 0xd5, 0xd4, 0x8c, 0xb2, 0x10, 0xd5, 0x57, 0xce,
	0x87, 0x9a, 0xf4, 0x7d, 0xa8, 0xdd, 0x6a, 0x52, 0x71, 0xda, 0xb3, 0x75, 0x87, 0xb5, 0x13, 0x43,
	0xc9, 0x4f, 0x89, 0xbb, 0xad, 0xb2, 0x18, 0x74, 0x80, 0xeb, 0x35, 0x5f, 0x4c, 0x86, 0xda, 0x66,
	0xac, 0x96, 0x82, 0x88, 0x79, 0x05, 0xc5, 0x3b, 0x68, 0x49, 0xb0, 0x16, 0xf8, 0x77, 0x95, 0x7c,
	0x88, 0x37, 0x93, 0x28, 0xcd, 0x1b, 0xca, 0xe2, 0x4c, 0xde, 0xc0, 0x5d, 0x84, 0x9d, 0x5e, 0x10,
	0x80, 0x2f, 0x1a, 0xbc, 0x1b, 0x88, 0x46, 0x27, 0xa0, 0x0e, 0x28, 0x4b, 0x91, 0xb5, 0xa3, 0x0c,
	0xd6, 0x1e, 0x81, 0x33, 0x19, 0x6a, 0x5b, 0xb1, 0x35, 0xde, 0x61, 0x09, 0x89, 0x98, 0x9b, 0x09,
	0xfe, 0x45, 0x37, 0x10, 0xf5, 0x30, 0x85, 0x4f, 0xd1, 0xfa, 0x54, 0x52, 0x50, 0xa7, 0xa5, 0x2c,
	0x47, 0x62, 0x8f, 0x33, 0xf7, 0xe1, 0xff, 0x58, 0x6c, 0x96, 0x45, 0xcc, 0xb5, 0x24, 0x7c, 0x49,
	0x9d, 0x16, 0x7e, 0x86, 0x50, 0xd3, 0x63, 0xb6, 0xe5, 0x55, 0x01, 0xb8, 0xb2, 0xb2, 0x27, 0x17,
	0xd7, 0xf6, 0x6f, 0xeb, 0x7f, 0x34, 0x22, 0x7a, 0x15, 0xc0, 0x9c, 0xa9, 0x3e, 0xd8, 0x7a, 0x7f,
	0xa6, 0x49, 0x9f, 0xce, 0x34, 0xe9, 0xe2, 0x73, 0x69, 0x31, 0x1c, 0x85, 0x1a, 0xb9, 0xc8, 0xa1,
	0x95, 0x50, 0xa7, 0xe6, 0x9f, 0x30, 0xdc, 0x45, 0xff, 0xa5, 0x90, 0x46, 0x33, 0x60, 0xe9, 0x80,
	0x3c, 0xcd, 0x7c, 0xb1, 0x9d, 0xb9, 0x17, 0x1c, 0xe3, 0x88, 0x59, 0x48, 0x33, 0x4f, 0xc2, 0x04,
	0x6e, 0xa1, 0x8d, 0xab, 0x33, 0x3e, 0x88, 0x68, 0xd0, 0x56, 0x2b, 0xd5, 0xcc, 0x82, 0xdb, 0xf3,
	0x82, 0x3e, 0x08, 0x62, 0xae, 0xa7, 0xf1, 0x73, 0x10, 0xf8, 0x0d, 0xda, 0x3e, 0x01, 0xe0, 0x87,
	0x8e, 0x13, 0xf4, 0xc0, 0x3d, 0xb4, 0x59, 0x1f, 0xc2, 0xbb, 0x47, 0x53, 0x9c, 0xad, 0xab, 0xbf,
	0xe4, 0x90, 0xaf, 0x32, 0x5a, 0xa9, 0x33, 0x4e, 0x05, 0x65, 0xfe, 0xf5, 0x3d, 0x91, 0xff, 0xc5,
	0x9e, 0x98, 0xa8, 0x10, 0xd9, 0x10, 0x47, 0x01, 0x58, 0xa1, 0x66, 0xd4, 0xbc, 0x6c, 0x17, 0x99,
	0x23, 0x90, 0x2f, 0x32, 0x5a, 0xa8, 0x02, 0xe0, 0xe3, 0x74, 0x07, 0x63, 0xeb, 0x0f, 0x33, 0xef,
	0xd1, 0x46, 0x6c, 0x3d, 0xa6, 0x90, 0x74, 0x89, 0x8f, 0xd3, 0x25, 0xce, 0xfd, 0x05, 0xb0, 0x31,
	0x05, 0x1b, 0x95, 0x57, 0xe7, 0x23, 0x55, 0xbe, 0x1c, 0xa9, 0xf2, 0x8f, 0x91, 0x2a, 0x7f, 0x18,
	0xab, 0xd2, 0xe5, 0x58, 0x95, 0xbe, 0x8d, 0x55, 0xe9, 0xf5, 0x83, 0x19, 0x74, 0xd2, 0x99, 0x92,
	0x67, 0xd9, 0x7c, 0x1a, 0x94, 0xfb, 0xc6, 0x7e, 0xf9, 0xdd, 0x6f, 0xbe, 0xd2, 0xf6, 0x52, 0xf4,
	0xf5, 0xbc, 0xf7, 0x33, 0x00, 0x00, 0xff, 0xff, 0x39, 0x79, 0x4d, 0x60, 0xce, 0x05, 0x00, 0x00,
}

func (m *Pool) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Pool) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Pool) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.GlobalFees != nil {
		{
			size, err := m.GlobalFees.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintConcentratedPool(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x42
	}
	{
		size := m.CurrentTick.Size()
		i -= size
		if _, err := m.CurrentTick.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintConcentratedPool(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x3a
	{
		size := m.CurrentSqrtPrice.Size()
		i -= size
		if _, err := m.CurrentSqrtPrice.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintConcentratedPool(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x32
	if len(m.Token1) > 0 {
		i -= len(m.Token1)
		copy(dAtA[i:], m.Token1)
		i = encodeVarintConcentratedPool(dAtA, i, uint64(len(m.Token1)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Token0) > 0 {
		i -= len(m.Token0)
		copy(dAtA[i:], m.Token0)
		i = encodeVarintConcentratedPool(dAtA, i, uint64(len(m.Token0)))
		i--
		dAtA[i] = 0x22
	}
	{
		size := m.Liquidity.Size()
		i -= size
		if _, err := m.Liquidity.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintConcentratedPool(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	if m.Id != 0 {
		i = encodeVarintConcentratedPool(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintConcentratedPool(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *TickInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TickInfo) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TickInfo) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.FeesAccruedAboveTick != nil {
		{
			size, err := m.FeesAccruedAboveTick.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintConcentratedPool(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	{
		size := m.LiquidityNet.Size()
		i -= size
		if _, err := m.LiquidityNet.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintConcentratedPool(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	{
		size := m.LiquidityGross.Size()
		i -= size
		if _, err := m.LiquidityGross.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintConcentratedPool(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *Position) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Position) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Position) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.FeesAtCreation != nil {
		{
			size, err := m.FeesAtCreation.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintConcentratedPool(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	{
		size := m.Liquidity.Size()
		i -= size
		if _, err := m.Liquidity.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintConcentratedPool(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *Fee) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Fee) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Fee) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.Token1.Size()
		i -= size
		if _, err := m.Token1.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintConcentratedPool(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	{
		size := m.Token0.Size()
		i -= size
		if _, err := m.Token0.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintConcentratedPool(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintConcentratedPool(dAtA []byte, offset int, v uint64) int {
	offset -= sovConcentratedPool(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Pool) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovConcentratedPool(uint64(l))
	}
	if m.Id != 0 {
		n += 1 + sovConcentratedPool(uint64(m.Id))
	}
	l = m.Liquidity.Size()
	n += 1 + l + sovConcentratedPool(uint64(l))
	l = len(m.Token0)
	if l > 0 {
		n += 1 + l + sovConcentratedPool(uint64(l))
	}
	l = len(m.Token1)
	if l > 0 {
		n += 1 + l + sovConcentratedPool(uint64(l))
	}
	l = m.CurrentSqrtPrice.Size()
	n += 1 + l + sovConcentratedPool(uint64(l))
	l = m.CurrentTick.Size()
	n += 1 + l + sovConcentratedPool(uint64(l))
	if m.GlobalFees != nil {
		l = m.GlobalFees.Size()
		n += 1 + l + sovConcentratedPool(uint64(l))
	}
	return n
}

func (m *TickInfo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.LiquidityGross.Size()
	n += 1 + l + sovConcentratedPool(uint64(l))
	l = m.LiquidityNet.Size()
	n += 1 + l + sovConcentratedPool(uint64(l))
	if m.FeesAccruedAboveTick != nil {
		l = m.FeesAccruedAboveTick.Size()
		n += 1 + l + sovConcentratedPool(uint64(l))
	}
	return n
}

func (m *Position) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Liquidity.Size()
	n += 1 + l + sovConcentratedPool(uint64(l))
	if m.FeesAtCreation != nil {
		l = m.FeesAtCreation.Size()
		n += 1 + l + sovConcentratedPool(uint64(l))
	}
	return n
}

func (m *Fee) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Token0.Size()
	n += 1 + l + sovConcentratedPool(uint64(l))
	l = m.Token1.Size()
	n += 1 + l + sovConcentratedPool(uint64(l))
	return n
}

func sovConcentratedPool(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozConcentratedPool(x uint64) (n int) {
	return sovConcentratedPool(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Pool) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowConcentratedPool
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
			return fmt.Errorf("proto: Pool: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Pool: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConcentratedPool
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
				return ErrInvalidLengthConcentratedPool
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthConcentratedPool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConcentratedPool
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
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Liquidity", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConcentratedPool
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
				return ErrInvalidLengthConcentratedPool
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthConcentratedPool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Liquidity.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Token0", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConcentratedPool
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
				return ErrInvalidLengthConcentratedPool
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthConcentratedPool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Token0 = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Token1", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConcentratedPool
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
				return ErrInvalidLengthConcentratedPool
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthConcentratedPool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Token1 = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CurrentSqrtPrice", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConcentratedPool
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
				return ErrInvalidLengthConcentratedPool
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthConcentratedPool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.CurrentSqrtPrice.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CurrentTick", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConcentratedPool
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
				return ErrInvalidLengthConcentratedPool
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthConcentratedPool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.CurrentTick.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GlobalFees", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConcentratedPool
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
				return ErrInvalidLengthConcentratedPool
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthConcentratedPool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.GlobalFees == nil {
				m.GlobalFees = &Fee{}
			}
			if err := m.GlobalFees.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipConcentratedPool(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthConcentratedPool
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
func (m *TickInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowConcentratedPool
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
			return fmt.Errorf("proto: TickInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TickInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LiquidityGross", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConcentratedPool
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
				return ErrInvalidLengthConcentratedPool
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthConcentratedPool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.LiquidityGross.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LiquidityNet", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConcentratedPool
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
				return ErrInvalidLengthConcentratedPool
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthConcentratedPool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.LiquidityNet.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FeesAccruedAboveTick", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConcentratedPool
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
				return ErrInvalidLengthConcentratedPool
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthConcentratedPool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.FeesAccruedAboveTick == nil {
				m.FeesAccruedAboveTick = &Fee{}
			}
			if err := m.FeesAccruedAboveTick.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipConcentratedPool(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthConcentratedPool
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
func (m *Position) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowConcentratedPool
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
			return fmt.Errorf("proto: Position: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Position: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Liquidity", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConcentratedPool
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
				return ErrInvalidLengthConcentratedPool
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthConcentratedPool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Liquidity.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FeesAtCreation", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConcentratedPool
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
				return ErrInvalidLengthConcentratedPool
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthConcentratedPool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.FeesAtCreation == nil {
				m.FeesAtCreation = &Fee{}
			}
			if err := m.FeesAtCreation.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipConcentratedPool(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthConcentratedPool
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
func (m *Fee) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowConcentratedPool
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
			return fmt.Errorf("proto: Fee: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Fee: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Token0", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConcentratedPool
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
				return ErrInvalidLengthConcentratedPool
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthConcentratedPool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Token0.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Token1", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConcentratedPool
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
				return ErrInvalidLengthConcentratedPool
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthConcentratedPool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Token1.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipConcentratedPool(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthConcentratedPool
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
func skipConcentratedPool(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowConcentratedPool
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
					return 0, ErrIntOverflowConcentratedPool
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
					return 0, ErrIntOverflowConcentratedPool
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
				return 0, ErrInvalidLengthConcentratedPool
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupConcentratedPool
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthConcentratedPool
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthConcentratedPool        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowConcentratedPool          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupConcentratedPool = fmt.Errorf("proto: unexpected end of group")
)
