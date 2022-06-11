// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: comdex/liquidity/v1beta1/params.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
	_ "google.golang.org/protobuf/types/known/durationpb"
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

// Params defines the parameters for the liquidity module.
type Params struct {
}

func (m *Params) Reset()         { *m = Params{} }
func (m *Params) String() string { return proto.CompactTextString(m) }
func (*Params) ProtoMessage()    {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_babec35f52b1356c, []int{0}
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

// Params defines the parameters for the liquidity module.
type GenericParams struct {
	BatchSize                uint32                                   `protobuf:"varint,1,opt,name=batch_size,json=batchSize,proto3" json:"batch_size,omitempty"`
	TickPrecision            uint32                                   `protobuf:"varint,2,opt,name=tick_precision,json=tickPrecision,proto3" json:"tick_precision,omitempty"`
	FeeCollectorAddress      string                                   `protobuf:"bytes,3,opt,name=fee_collector_address,json=feeCollectorAddress,proto3" json:"fee_collector_address,omitempty"`
	DustCollectorAddress     string                                   `protobuf:"bytes,4,opt,name=dust_collector_address,json=dustCollectorAddress,proto3" json:"dust_collector_address,omitempty"`
	MinInitialPoolCoinSupply github_com_cosmos_cosmos_sdk_types.Int   `protobuf:"bytes,5,opt,name=min_initial_pool_coin_supply,json=minInitialPoolCoinSupply,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"min_initial_pool_coin_supply"`
	PairCreationFee          github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,6,rep,name=pair_creation_fee,json=pairCreationFee,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"pair_creation_fee"`
	PoolCreationFee          github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,7,rep,name=pool_creation_fee,json=poolCreationFee,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"pool_creation_fee"`
	MinInitialDepositAmount  github_com_cosmos_cosmos_sdk_types.Int   `protobuf:"bytes,8,opt,name=min_initial_deposit_amount,json=minInitialDepositAmount,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"min_initial_deposit_amount"`
	MaxPriceLimitRatio       github_com_cosmos_cosmos_sdk_types.Dec   `protobuf:"bytes,9,opt,name=max_price_limit_ratio,json=maxPriceLimitRatio,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"max_price_limit_ratio"`
	MaxOrderLifespan         time.Duration                            `protobuf:"bytes,10,opt,name=max_order_lifespan,json=maxOrderLifespan,proto3,stdduration" json:"max_order_lifespan"`
	SwapFeeRate              github_com_cosmos_cosmos_sdk_types.Dec   `protobuf:"bytes,11,opt,name=swap_fee_rate,json=swapFeeRate,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"swap_fee_rate"`
	WithdrawFeeRate          github_com_cosmos_cosmos_sdk_types.Dec   `protobuf:"bytes,12,opt,name=withdraw_fee_rate,json=withdrawFeeRate,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"withdraw_fee_rate"`
	DepositExtraGas          github_com_cosmos_cosmos_sdk_types.Gas   `protobuf:"varint,13,opt,name=deposit_extra_gas,json=depositExtraGas,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Gas" json:"deposit_extra_gas"`
	WithdrawExtraGas         github_com_cosmos_cosmos_sdk_types.Gas   `protobuf:"varint,14,opt,name=withdraw_extra_gas,json=withdrawExtraGas,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Gas" json:"withdraw_extra_gas"`
	OrderExtraGas            github_com_cosmos_cosmos_sdk_types.Gas   `protobuf:"varint,15,opt,name=order_extra_gas,json=orderExtraGas,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Gas" json:"order_extra_gas"`
	SwapFeeDistrDenom        string                                   `protobuf:"bytes,16,opt,name=swap_fee_distr_denom,json=swapFeeDistrDenom,proto3" json:"swap_fee_distr_denom,omitempty"`
	SwapFeeBurnRate          github_com_cosmos_cosmos_sdk_types.Dec   `protobuf:"bytes,17,opt,name=swap_fee_burn_rate,json=swapFeeBurnRate,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"swap_fee_burn_rate"`
	AppId                    uint64                                   `protobuf:"varint,18,opt,name=app_id,json=appId,proto3" json:"app_id,omitempty"`
}

func (m *GenericParams) Reset()         { *m = GenericParams{} }
func (m *GenericParams) String() string { return proto.CompactTextString(m) }
func (*GenericParams) ProtoMessage()    {}
func (*GenericParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_babec35f52b1356c, []int{1}
}
func (m *GenericParams) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenericParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenericParams.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenericParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenericParams.Merge(m, src)
}
func (m *GenericParams) XXX_Size() int {
	return m.Size()
}
func (m *GenericParams) XXX_DiscardUnknown() {
	xxx_messageInfo_GenericParams.DiscardUnknown(m)
}

var xxx_messageInfo_GenericParams proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Params)(nil), "comdex.liquidity.v1beta1.Params")
	proto.RegisterType((*GenericParams)(nil), "comdex.liquidity.v1beta1.GenericParams")
}

func init() {
	proto.RegisterFile("comdex/liquidity/v1beta1/params.proto", fileDescriptor_babec35f52b1356c)
}

var fileDescriptor_babec35f52b1356c = []byte{
	// 740 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x95, 0xdf, 0x6e, 0xeb, 0x34,
	0x1c, 0xc7, 0x1b, 0xce, 0x39, 0x65, 0x73, 0xe9, 0xe9, 0x6a, 0x36, 0x08, 0x13, 0xa4, 0xd5, 0x91,
	0x86, 0x7a, 0xb3, 0x84, 0x6d, 0xbc, 0xc0, 0xba, 0xb2, 0x69, 0xd2, 0x24, 0xba, 0x4c, 0xe2, 0x62,
	0x20, 0x59, 0x6e, 0xf2, 0x6b, 0x67, 0x2d, 0x89, 0x8d, 0xed, 0xd0, 0x6e, 0x4f, 0xc1, 0x25, 0xcf,
	0x00, 0x2f, 0xb2, 0xcb, 0x5d, 0x22, 0x2e, 0x36, 0xd8, 0x5e, 0x04, 0xd9, 0x49, 0xda, 0x22, 0xb8,
	0x60, 0x95, 0xce, 0x55, 0x9b, 0xdf, 0x9f, 0xcf, 0xd7, 0xfe, 0x7d, 0xe3, 0x18, 0xed, 0x44, 0x3c,
	0x8d, 0x61, 0x16, 0x24, 0xec, 0xc7, 0x9c, 0xc5, 0x4c, 0xdf, 0x04, 0x3f, 0xed, 0x8d, 0x40, 0xd3,
	0xbd, 0x40, 0x50, 0x49, 0x53, 0xe5, 0x0b, 0xc9, 0x35, 0xc7, 0x6e, 0x51, 0xe6, 0xcf, 0xcb, 0xfc,
	0xb2, 0x6c, 0x7b, 0x73, 0xc2, 0x27, 0xdc, 0x16, 0x05, 0xe6, 0x5f, 0x51, 0xbf, 0xed, 0x45, 0x5c,
	0xa5, 0x5c, 0x05, 0x23, 0xaa, 0x60, 0x4e, 0x8c, 0x38, 0xcb, 0xaa, 0xfc, 0x84, 0xf3, 0x49, 0x02,
	0x81, 0x7d, 0x1a, 0xe5, 0xe3, 0x20, 0xce, 0x25, 0xd5, 0x8c, 0x97, 0xf9, 0x77, 0x6b, 0xa8, 0x3e,
	0xb4, 0xfa, 0xef, 0x7e, 0x6b, 0xa0, 0xe6, 0x09, 0x64, 0x20, 0x59, 0x54, 0x44, 0xf0, 0x17, 0x08,
	0x8d, 0xa8, 0x8e, 0xae, 0x88, 0x62, 0xb7, 0xe0, 0x3a, 0x5d, 0xa7, 0xd7, 0x0c, 0xd7, 0x6d, 0xe4,
	0x82, 0xdd, 0x02, 0xde, 0x41, 0x6f, 0x35, 0x8b, 0xae, 0x89, 0x90, 0x10, 0x31, 0xc5, 0x78, 0xe6,
	0x7e, 0x60, 0x4b, 0x9a, 0x26, 0x3a, 0xac, 0x82, 0x78, 0x1f, 0x6d, 0x8d, 0x01, 0x48, 0xc4, 0x93,
	0x04, 0x22, 0xcd, 0x25, 0xa1, 0x71, 0x2c, 0x41, 0x29, 0xf7, 0x55, 0xd7, 0xe9, 0xad, 0x87, 0x1f,
	0x8f, 0x01, 0x8e, 0xaa, 0xdc, 0x61, 0x91, 0xc2, 0x5f, 0xa3, 0x4f, 0xe2, 0x5c, 0xe9, 0xff, 0x68,
	0x7a, 0x6d, 0x9b, 0x36, 0x4d, 0xf6, 0x5f, 0x5d, 0x19, 0xfa, 0x3c, 0x65, 0x19, 0x61, 0x19, 0xd3,
	0x8c, 0x26, 0x44, 0x70, 0x9e, 0x10, 0x33, 0x0a, 0xa2, 0x72, 0x21, 0x92, 0x1b, 0xf7, 0x8d, 0xe9,
	0xed, 0xfb, 0x77, 0x0f, 0x9d, 0xda, 0x1f, 0x0f, 0x9d, 0x2f, 0x27, 0x4c, 0x5f, 0xe5, 0x23, 0x3f,
	0xe2, 0x69, 0x50, 0x0e, 0xb1, 0xf8, 0xd9, 0x55, 0xf1, 0x75, 0xa0, 0x6f, 0x04, 0x28, 0xff, 0x34,
	0xd3, 0xa1, 0x9b, 0xb2, 0xec, 0xb4, 0x40, 0x0e, 0x39, 0x4f, 0x8e, 0x38, 0xcb, 0x2e, 0x2c, 0x0f,
	0x4f, 0x51, 0x5b, 0x50, 0x26, 0x49, 0x24, 0xc1, 0x8e, 0x94, 0x8c, 0x01, 0xdc, 0x7a, 0xf7, 0x55,
	0xaf, 0xb1, 0xff, 0x99, 0x5f, 0xb0, 0x7c, 0xe3, 0x4b, 0x65, 0xa1, 0x6f, 0x7a, 0xfb, 0x5f, 0x19,
	0xfd, 0x5f, 0x1f, 0x3b, 0xbd, 0xff, 0xa1, 0x6f, 0x1a, 0x54, 0xd8, 0x32, 0x2a, 0x47, 0xa5, 0xc8,
	0x31, 0x80, 0x15, 0xb6, 0x9b, 0x5b, 0x16, 0xfe, 0xf0, 0x7d, 0x08, 0x9b, 0x0d, 0x2f, 0x09, 0x5f,
	0xa3, 0xed, 0xe5, 0x09, 0xc7, 0x20, 0xb8, 0x62, 0x9a, 0xd0, 0x94, 0xe7, 0x99, 0x76, 0xd7, 0x56,
	0x9a, 0xef, 0xa7, 0x8b, 0xf9, 0x0e, 0x0a, 0xde, 0xa1, 0xc5, 0x61, 0x8a, 0xb6, 0x52, 0x3a, 0x23,
	0x42, 0xb2, 0x08, 0x48, 0xc2, 0x52, 0xa6, 0x89, 0x7d, 0x75, 0xdd, 0xf5, 0x17, 0xeb, 0x0c, 0x20,
	0x0a, 0x71, 0x4a, 0x67, 0x43, 0xc3, 0x3a, 0x33, 0xa8, 0xd0, 0x90, 0xf0, 0x39, 0x32, 0x51, 0xc2,
	0x65, 0x0c, 0x92, 0x24, 0x6c, 0x0c, 0x4a, 0xd0, 0xcc, 0x45, 0x5d, 0xc7, 0x4e, 0xb2, 0x38, 0x3a,
	0x7e, 0x75, 0x74, 0xfc, 0x41, 0x79, 0x74, 0xfa, 0x6b, 0x46, 0xfa, 0x97, 0xc7, 0x8e, 0x13, 0x6e,
	0xa4, 0x74, 0xf6, 0xad, 0xe9, 0x3e, 0x2b, 0x9b, 0x71, 0x88, 0x9a, 0x6a, 0x4a, 0x85, 0xb1, 0xc4,
	0x2c, 0x17, 0xdc, 0xc6, 0x4a, 0xab, 0x6d, 0x18, 0xc8, 0x31, 0x40, 0x48, 0x35, 0xe0, 0x4b, 0xd4,
	0x9e, 0x32, 0x7d, 0x15, 0x4b, 0x3a, 0x5d, 0x70, 0x3f, 0x5a, 0x89, 0xdb, 0xaa, 0x40, 0x4b, 0xec,
	0xca, 0x46, 0x98, 0x69, 0x49, 0xc9, 0x84, 0x2a, 0xb7, 0xd9, 0x75, 0x7a, 0xaf, 0x5f, 0xc4, 0x3e,
	0xa1, 0x2a, 0x6c, 0x95, 0xa0, 0x6f, 0x0c, 0xe7, 0x84, 0x2a, 0xfc, 0x03, 0xc2, 0xf3, 0x75, 0x2f,
	0xe0, 0x6f, 0x57, 0x82, 0x6f, 0x54, 0xa4, 0x39, 0xfd, 0x3b, 0xd4, 0x2a, 0x8c, 0x5b, 0xa0, 0x5b,
	0x2b, 0xa1, 0x9b, 0x16, 0x33, 0xe7, 0x06, 0x68, 0x73, 0xee, 0x60, 0xcc, 0x94, 0x96, 0x24, 0x86,
	0x8c, 0xa7, 0xee, 0x86, 0xfd, 0xf4, 0xb4, 0x4b, 0x63, 0x06, 0x26, 0x33, 0x30, 0x09, 0xfc, 0x3d,
	0xc2, 0xf3, 0x86, 0x51, 0x2e, 0xb3, 0xc2, 0x9f, 0xf6, 0x6a, 0xfe, 0x94, 0xf8, 0x7e, 0x2e, 0x33,
	0xeb, 0xcf, 0x16, 0xaa, 0x53, 0x21, 0x08, 0x8b, 0x5d, 0x6c, 0x36, 0x17, 0xbe, 0xa1, 0x42, 0x9c,
	0xc6, 0xfd, 0xf3, 0xbb, 0xbf, 0xbc, 0xda, 0xdd, 0x93, 0xe7, 0xdc, 0x3f, 0x79, 0xce, 0x9f, 0x4f,
	0x9e, 0xf3, 0xf3, 0xb3, 0x57, 0xbb, 0x7f, 0xf6, 0x6a, 0xbf, 0x3f, 0x7b, 0xb5, 0xcb, 0x83, 0x7f,
	0xa8, 0x99, 0x0b, 0x65, 0x97, 0x8f, 0xc7, 0x2c, 0x62, 0x34, 0x29, 0x9f, 0x83, 0xe5, 0x9b, 0xc8,
	0xca, 0x8f, 0xea, 0xf6, 0x45, 0x3f, 0xf8, 0x3b, 0x00, 0x00, 0xff, 0xff, 0x56, 0x2a, 0xaa, 0x33,
	0xaa, 0x06, 0x00, 0x00,
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
	return len(dAtA) - i, nil
}

func (m *GenericParams) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenericParams) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenericParams) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.AppId != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.AppId))
		i--
		dAtA[i] = 0x1
		i--
		dAtA[i] = 0x90
	}
	{
		size := m.SwapFeeBurnRate.Size()
		i -= size
		if _, err := m.SwapFeeBurnRate.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1
	i--
	dAtA[i] = 0x8a
	if len(m.SwapFeeDistrDenom) > 0 {
		i -= len(m.SwapFeeDistrDenom)
		copy(dAtA[i:], m.SwapFeeDistrDenom)
		i = encodeVarintParams(dAtA, i, uint64(len(m.SwapFeeDistrDenom)))
		i--
		dAtA[i] = 0x1
		i--
		dAtA[i] = 0x82
	}
	if m.OrderExtraGas != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.OrderExtraGas))
		i--
		dAtA[i] = 0x78
	}
	if m.WithdrawExtraGas != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.WithdrawExtraGas))
		i--
		dAtA[i] = 0x70
	}
	if m.DepositExtraGas != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.DepositExtraGas))
		i--
		dAtA[i] = 0x68
	}
	{
		size := m.WithdrawFeeRate.Size()
		i -= size
		if _, err := m.WithdrawFeeRate.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x62
	{
		size := m.SwapFeeRate.Size()
		i -= size
		if _, err := m.SwapFeeRate.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x5a
	n1, err1 := github_com_gogo_protobuf_types.StdDurationMarshalTo(m.MaxOrderLifespan, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdDuration(m.MaxOrderLifespan):])
	if err1 != nil {
		return 0, err1
	}
	i -= n1
	i = encodeVarintParams(dAtA, i, uint64(n1))
	i--
	dAtA[i] = 0x52
	{
		size := m.MaxPriceLimitRatio.Size()
		i -= size
		if _, err := m.MaxPriceLimitRatio.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x4a
	{
		size := m.MinInitialDepositAmount.Size()
		i -= size
		if _, err := m.MinInitialDepositAmount.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x42
	if len(m.PoolCreationFee) > 0 {
		for iNdEx := len(m.PoolCreationFee) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.PoolCreationFee[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintParams(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x3a
		}
	}
	if len(m.PairCreationFee) > 0 {
		for iNdEx := len(m.PairCreationFee) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.PairCreationFee[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintParams(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x32
		}
	}
	{
		size := m.MinInitialPoolCoinSupply.Size()
		i -= size
		if _, err := m.MinInitialPoolCoinSupply.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x2a
	if len(m.DustCollectorAddress) > 0 {
		i -= len(m.DustCollectorAddress)
		copy(dAtA[i:], m.DustCollectorAddress)
		i = encodeVarintParams(dAtA, i, uint64(len(m.DustCollectorAddress)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.FeeCollectorAddress) > 0 {
		i -= len(m.FeeCollectorAddress)
		copy(dAtA[i:], m.FeeCollectorAddress)
		i = encodeVarintParams(dAtA, i, uint64(len(m.FeeCollectorAddress)))
		i--
		dAtA[i] = 0x1a
	}
	if m.TickPrecision != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.TickPrecision))
		i--
		dAtA[i] = 0x10
	}
	if m.BatchSize != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.BatchSize))
		i--
		dAtA[i] = 0x8
	}
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
	return n
}

func (m *GenericParams) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.BatchSize != 0 {
		n += 1 + sovParams(uint64(m.BatchSize))
	}
	if m.TickPrecision != 0 {
		n += 1 + sovParams(uint64(m.TickPrecision))
	}
	l = len(m.FeeCollectorAddress)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
	}
	l = len(m.DustCollectorAddress)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
	}
	l = m.MinInitialPoolCoinSupply.Size()
	n += 1 + l + sovParams(uint64(l))
	if len(m.PairCreationFee) > 0 {
		for _, e := range m.PairCreationFee {
			l = e.Size()
			n += 1 + l + sovParams(uint64(l))
		}
	}
	if len(m.PoolCreationFee) > 0 {
		for _, e := range m.PoolCreationFee {
			l = e.Size()
			n += 1 + l + sovParams(uint64(l))
		}
	}
	l = m.MinInitialDepositAmount.Size()
	n += 1 + l + sovParams(uint64(l))
	l = m.MaxPriceLimitRatio.Size()
	n += 1 + l + sovParams(uint64(l))
	l = github_com_gogo_protobuf_types.SizeOfStdDuration(m.MaxOrderLifespan)
	n += 1 + l + sovParams(uint64(l))
	l = m.SwapFeeRate.Size()
	n += 1 + l + sovParams(uint64(l))
	l = m.WithdrawFeeRate.Size()
	n += 1 + l + sovParams(uint64(l))
	if m.DepositExtraGas != 0 {
		n += 1 + sovParams(uint64(m.DepositExtraGas))
	}
	if m.WithdrawExtraGas != 0 {
		n += 1 + sovParams(uint64(m.WithdrawExtraGas))
	}
	if m.OrderExtraGas != 0 {
		n += 1 + sovParams(uint64(m.OrderExtraGas))
	}
	l = len(m.SwapFeeDistrDenom)
	if l > 0 {
		n += 2 + l + sovParams(uint64(l))
	}
	l = m.SwapFeeBurnRate.Size()
	n += 2 + l + sovParams(uint64(l))
	if m.AppId != 0 {
		n += 2 + sovParams(uint64(m.AppId))
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
func (m *GenericParams) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: GenericParams: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenericParams: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BatchSize", wireType)
			}
			m.BatchSize = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.BatchSize |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TickPrecision", wireType)
			}
			m.TickPrecision = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TickPrecision |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
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
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DustCollectorAddress", wireType)
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
			m.DustCollectorAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinInitialPoolCoinSupply", wireType)
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
			if err := m.MinInitialPoolCoinSupply.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PairCreationFee", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PairCreationFee = append(m.PairCreationFee, types.Coin{})
			if err := m.PairCreationFee[len(m.PairCreationFee)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PoolCreationFee", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PoolCreationFee = append(m.PoolCreationFee, types.Coin{})
			if err := m.PoolCreationFee[len(m.PoolCreationFee)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinInitialDepositAmount", wireType)
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
			if err := m.MinInitialDepositAmount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxPriceLimitRatio", wireType)
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
			if err := m.MaxPriceLimitRatio.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxOrderLifespan", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdDurationUnmarshal(&m.MaxOrderLifespan, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 11:
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
		case 12:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field WithdrawFeeRate", wireType)
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
			if err := m.WithdrawFeeRate.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 13:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field DepositExtraGas", wireType)
			}
			m.DepositExtraGas = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.DepositExtraGas |= github_com_cosmos_cosmos_sdk_types.Gas(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 14:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field WithdrawExtraGas", wireType)
			}
			m.WithdrawExtraGas = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.WithdrawExtraGas |= github_com_cosmos_cosmos_sdk_types.Gas(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 15:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field OrderExtraGas", wireType)
			}
			m.OrderExtraGas = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.OrderExtraGas |= github_com_cosmos_cosmos_sdk_types.Gas(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 16:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SwapFeeDistrDenom", wireType)
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
			m.SwapFeeDistrDenom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 17:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SwapFeeBurnRate", wireType)
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
			if err := m.SwapFeeBurnRate.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 18:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AppId", wireType)
			}
			m.AppId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.AppId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
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