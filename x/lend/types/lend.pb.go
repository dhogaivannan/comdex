// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: comdex/lend/v1beta1/lend.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/types"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
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

type LendAsset struct {
	ID                 uint64                                  `protobuf:"varint,1,opt,name=lending_id,json=lendingId,proto3" json:"lending_id,omitempty" yaml:"lending_id"`
	AssetId            uint64                                  `protobuf:"varint,2,opt,name=asset_id,json=assetId,proto3" json:"asset_id,omitempty" yaml:"asset_id"`
	PoolId             uint64                                  `protobuf:"varint,3,opt,name=pool_id,json=poolId,proto3" json:"pool_id,omitempty" yaml:"pool_id"`
	Owner              string                                  `protobuf:"bytes,4,opt,name=owner,proto3" json:"owner,omitempty" yaml:"owner"`
	AmountIn           github_com_cosmos_cosmos_sdk_types.Coin `protobuf:"bytes,5,opt,name=amount_in,json=amountIn,proto3,casttype=github.com/cosmos/cosmos-sdk/types.Coin" json:"amount_in" yaml:"amount_in"`
	LendingTime        time.Time                               `protobuf:"bytes,6,opt,name=lending_time,json=lendingTime,proto3,stdtime" json:"lending_time" yaml:"lending_time"`
	Reward_Accumulated github_com_cosmos_cosmos_sdk_types.Coin `protobuf:"bytes,7,opt,name=reward_Accumulated,json=rewardAccumulated,proto3,casttype=github.com/cosmos/cosmos-sdk/types.Coin" json:"reward_Accumulated" yaml:"reward"`
}

func (m *LendAsset) Reset()         { *m = LendAsset{} }
func (m *LendAsset) String() string { return proto.CompactTextString(m) }
func (*LendAsset) ProtoMessage()    {}
func (*LendAsset) Descriptor() ([]byte, []int) {
	return fileDescriptor_b87bb4bef8334ddd, []int{0}
}
func (m *LendAsset) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *LendAsset) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_LendAsset.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *LendAsset) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LendAsset.Merge(m, src)
}
func (m *LendAsset) XXX_Size() int {
	return m.Size()
}
func (m *LendAsset) XXX_DiscardUnknown() {
	xxx_messageInfo_LendAsset.DiscardUnknown(m)
}

var xxx_messageInfo_LendAsset proto.InternalMessageInfo

func (m *LendAsset) GetID() uint64 {
	if m != nil {
		return m.ID
	}
	return 0
}

func (m *LendAsset) GetAssetId() uint64 {
	if m != nil {
		return m.AssetId
	}
	return 0
}

func (m *LendAsset) GetPoolId() uint64 {
	if m != nil {
		return m.PoolId
	}
	return 0
}

func (m *LendAsset) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *LendAsset) GetAmountIn() github_com_cosmos_cosmos_sdk_types.Coin {
	if m != nil {
		return m.AmountIn
	}
	return github_com_cosmos_cosmos_sdk_types.Coin{}
}

func (m *LendAsset) GetLendingTime() time.Time {
	if m != nil {
		return m.LendingTime
	}
	return time.Time{}
}

func (m *LendAsset) GetReward_Accumulated() github_com_cosmos_cosmos_sdk_types.Coin {
	if m != nil {
		return m.Reward_Accumulated
	}
	return github_com_cosmos_cosmos_sdk_types.Coin{}
}

type Pool struct {
	PoolId    uint64                 `protobuf:"varint,1,opt,name=pool_id,json=poolId,proto3" json:"pool_id,omitempty" yaml:"pool_id"`
	AccName   string                 `protobuf:"bytes,2,opt,name=acc_name,json=accName,proto3" json:"acc_name,omitempty" yaml:"acc_name"`
	AssetData []AssetDataPoolMapping `protobuf:"bytes,3,rep,name=asset_data,json=assetData,proto3" json:"asset_data" yaml:"asset_data"`
}

func (m *Pool) Reset()         { *m = Pool{} }
func (m *Pool) String() string { return proto.CompactTextString(m) }
func (*Pool) ProtoMessage()    {}
func (*Pool) Descriptor() ([]byte, []int) {
	return fileDescriptor_b87bb4bef8334ddd, []int{1}
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

func (m *Pool) GetPoolId() uint64 {
	if m != nil {
		return m.PoolId
	}
	return 0
}

func (m *Pool) GetAccName() string {
	if m != nil {
		return m.AccName
	}
	return ""
}

func (m *Pool) GetAssetData() []AssetDataPoolMapping {
	if m != nil {
		return m.AssetData
	}
	return nil
}

type AssetDataPoolMapping struct {
	AssetId    uint64                                 `protobuf:"varint,1,opt,name=asset_id,json=assetId,proto3" json:"asset_id,omitempty" yaml:"asset_id"`
	LendRate   github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,2,opt,name=lend_rate,json=lendRate,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"lend_rate" yaml:"lend_rate"`
	BorrowRate github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,3,opt,name=borrow_rate,json=borrowRate,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"borrow_rate" yaml:"borrow_rate"`
}

func (m *AssetDataPoolMapping) Reset()         { *m = AssetDataPoolMapping{} }
func (m *AssetDataPoolMapping) String() string { return proto.CompactTextString(m) }
func (*AssetDataPoolMapping) ProtoMessage()    {}
func (*AssetDataPoolMapping) Descriptor() ([]byte, []int) {
	return fileDescriptor_b87bb4bef8334ddd, []int{2}
}
func (m *AssetDataPoolMapping) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AssetDataPoolMapping) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AssetDataPoolMapping.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *AssetDataPoolMapping) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AssetDataPoolMapping.Merge(m, src)
}
func (m *AssetDataPoolMapping) XXX_Size() int {
	return m.Size()
}
func (m *AssetDataPoolMapping) XXX_DiscardUnknown() {
	xxx_messageInfo_AssetDataPoolMapping.DiscardUnknown(m)
}

var xxx_messageInfo_AssetDataPoolMapping proto.InternalMessageInfo

func (m *AssetDataPoolMapping) GetAssetId() uint64 {
	if m != nil {
		return m.AssetId
	}
	return 0
}

func init() {
	proto.RegisterType((*LendAsset)(nil), "comdex.lend.v1beta1.LendAsset")
	proto.RegisterType((*Pool)(nil), "comdex.lend.v1beta1.Pool")
	proto.RegisterType((*AssetDataPoolMapping)(nil), "comdex.lend.v1beta1.AssetDataPoolMapping")
}

func init() { proto.RegisterFile("comdex/lend/v1beta1/lend.proto", fileDescriptor_b87bb4bef8334ddd) }

var fileDescriptor_b87bb4bef8334ddd = []byte{
	// 650 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x54, 0xbf, 0x6f, 0xd3, 0x4e,
	0x14, 0x8f, 0x9b, 0x34, 0xa9, 0x2f, 0xfd, 0x7e, 0x69, 0xdd, 0x0e, 0xa1, 0x20, 0x5f, 0x74, 0x43,
	0x09, 0x43, 0x6d, 0xb5, 0x88, 0x81, 0x4e, 0xd4, 0xcd, 0x40, 0x24, 0x40, 0xc8, 0x62, 0x62, 0x20,
	0xba, 0x9c, 0xaf, 0xc6, 0xc2, 0xf6, 0x59, 0xf6, 0x85, 0x52, 0xf1, 0x1f, 0x30, 0x75, 0xe5, 0x3f,
	0xaa, 0xc4, 0xd2, 0x11, 0x31, 0x18, 0xe4, 0x6e, 0x8c, 0x19, 0x99, 0xd0, 0xfd, 0x70, 0x93, 0x4a,
	0x95, 0x20, 0x53, 0xfc, 0xee, 0xf3, 0x3e, 0xef, 0x73, 0xf7, 0xde, 0xe7, 0x05, 0xd8, 0x84, 0x25,
	0x01, 0xfd, 0xe8, 0xc6, 0x34, 0x0d, 0xdc, 0x0f, 0xfb, 0x13, 0xca, 0xf1, 0xbe, 0x0c, 0x9c, 0x2c,
	0x67, 0x9c, 0x59, 0x5b, 0x0a, 0x77, 0xe4, 0x91, 0xc6, 0x77, 0xb6, 0x43, 0x16, 0x32, 0x89, 0xbb,
	0xe2, 0x4b, 0xa5, 0xee, 0xc0, 0x90, 0xb1, 0x30, 0xa6, 0xae, 0x8c, 0x26, 0xd3, 0x13, 0x97, 0x47,
	0x09, 0x2d, 0x38, 0x4e, 0x32, 0x9d, 0x60, 0x13, 0x56, 0x24, 0xac, 0x70, 0x27, 0xb8, 0xa0, 0xd7,
	0x5a, 0x84, 0x45, 0xa9, 0xc2, 0xd1, 0xd7, 0x16, 0x30, 0x9f, 0xd3, 0x34, 0x38, 0x2a, 0x0a, 0xca,
	0xad, 0x43, 0x00, 0x84, 0x68, 0x94, 0x86, 0xe3, 0x28, 0xe8, 0x19, 0x7d, 0x63, 0xd0, 0xf2, 0xee,
	0x55, 0x25, 0x5c, 0x19, 0x0d, 0x67, 0x25, 0xdc, 0x3c, 0xc3, 0x49, 0x7c, 0x88, 0xe6, 0x19, 0xc8,
	0x37, 0x75, 0x30, 0x0a, 0xac, 0x27, 0x60, 0x0d, 0x8b, 0x22, 0x82, 0xb9, 0x22, 0x99, 0x76, 0x55,
	0xc2, 0x8e, 0x2c, 0x3c, 0x0a, 0x66, 0x25, 0xbc, 0xa3, 0xe8, 0x75, 0x12, 0xf2, 0x3b, 0x58, 0x61,
	0xd6, 0x63, 0xd0, 0xc9, 0x18, 0x8b, 0x05, 0xb3, 0x29, 0x99, 0xf7, 0xab, 0x12, 0xb6, 0x5f, 0x31,
	0x16, 0x4b, 0xe2, 0xff, 0x8a, 0xa8, 0x53, 0x90, 0xdf, 0xce, 0x24, 0x62, 0xed, 0x82, 0x55, 0x76,
	0x9a, 0xd2, 0xbc, 0xd7, 0xea, 0x1b, 0x03, 0xd3, 0xdb, 0x98, 0x95, 0x70, 0x5d, 0xa5, 0xca, 0x63,
	0xe4, 0x2b, 0xd8, 0xfa, 0x04, 0x4c, 0x9c, 0xb0, 0x69, 0xca, 0xc7, 0x51, 0xda, 0x5b, 0xed, 0x1b,
	0x83, 0xee, 0xc1, 0x5d, 0x47, 0xf5, 0xc5, 0x11, 0x7d, 0xa9, 0x7b, 0xec, 0x1c, 0xb3, 0x28, 0xf5,
	0x8e, 0x2f, 0x4a, 0xd8, 0x98, 0x95, 0x70, 0x43, 0x5f, 0xb7, 0x66, 0xa2, 0xdf, 0x25, 0x7c, 0x10,
	0x46, 0xfc, 0xdd, 0x74, 0xe2, 0x10, 0x96, 0xb8, 0xba, 0xb1, 0xea, 0x67, 0xaf, 0x08, 0xde, 0xbb,
	0xfc, 0x2c, 0xa3, 0x85, 0x2c, 0xe2, 0xaf, 0x29, 0xda, 0x28, 0xb5, 0xde, 0x82, 0xf5, 0xba, 0x61,
	0x62, 0x36, 0xbd, 0xb6, 0xd4, 0xdf, 0x71, 0xd4, 0xe0, 0x9c, 0x7a, 0x70, 0xce, 0xeb, 0x7a, 0x70,
	0x1e, 0xd4, 0x17, 0xd8, 0xba, 0xd9, 0x6e, 0xc1, 0x46, 0xe7, 0x3f, 0xa0, 0xe1, 0x77, 0xf5, 0x91,
	0xa0, 0x58, 0x9f, 0x0d, 0x60, 0xe5, 0xf4, 0x14, 0xe7, 0xc1, 0xf8, 0x88, 0x90, 0x69, 0x32, 0x8d,
	0x31, 0xa7, 0x41, 0xaf, 0xf3, 0xb7, 0x67, 0x3e, 0xd5, 0x2a, 0xff, 0x29, 0x15, 0x55, 0x62, 0xa9,
	0x37, 0x6e, 0x2a, 0xce, 0x82, 0x2a, 0xfa, 0x65, 0x80, 0x96, 0x18, 0xdb, 0xe2, 0x44, 0x8d, 0x25,
	0x26, 0x2a, 0x3c, 0x44, 0xc8, 0x38, 0xc5, 0x09, 0x95, 0x1e, 0x32, 0xb5, 0x87, 0x08, 0x79, 0x89,
	0x13, 0xba, 0xe0, 0x21, 0x9d, 0x24, 0x3c, 0xa4, 0x30, 0x2b, 0x03, 0x40, 0x39, 0x2b, 0xc0, 0x1c,
	0xf7, 0x9a, 0xfd, 0xe6, 0xa0, 0x7b, 0xf0, 0xd0, 0xb9, 0x65, 0x93, 0x1c, 0xe9, 0xc8, 0x21, 0xe6,
	0x58, 0x5c, 0xe7, 0x05, 0xce, 0xb2, 0x28, 0x0d, 0xbd, 0x5d, 0xd1, 0x8e, 0xaa, 0x84, 0xe6, 0x35,
	0x3a, 0x37, 0xfc, 0xbc, 0x2e, 0xf2, 0x4d, 0x5c, 0xe3, 0xe8, 0xcb, 0x0a, 0xd8, 0xbe, 0xad, 0xd6,
	0x8d, 0x4d, 0x30, 0x96, 0xdb, 0x84, 0x31, 0x90, 0x1b, 0x35, 0xce, 0x31, 0xaf, 0x3b, 0xe0, 0x89,
	0x9b, 0x7d, 0x2f, 0xe1, 0xee, 0x3f, 0xcc, 0x65, 0x48, 0xc9, 0xdc, 0xb9, 0xd7, 0x85, 0x90, 0xbf,
	0x26, 0xbe, 0x7d, 0xcc, 0xa9, 0x45, 0x41, 0x77, 0xc2, 0xf2, 0x9c, 0x9d, 0x2a, 0x89, 0xa6, 0x94,
	0x18, 0x2e, 0x2d, 0x61, 0x29, 0x89, 0x85, 0x52, 0xc8, 0x07, 0x2a, 0x12, 0x32, 0xde, 0xb3, 0x8b,
	0xca, 0x36, 0x2e, 0x2b, 0xdb, 0xf8, 0x59, 0xd9, 0xc6, 0xf9, 0x95, 0xdd, 0xb8, 0xbc, 0xb2, 0x1b,
	0xdf, 0xae, 0xec, 0xc6, 0x1b, 0xe7, 0x86, 0x86, 0x98, 0xce, 0x1e, 0x3b, 0x39, 0x89, 0x48, 0x84,
	0x63, 0x1d, 0xbb, 0xfa, 0x9f, 0x51, 0xea, 0x4d, 0xda, 0x72, 0x43, 0x1e, 0xfd, 0x09, 0x00, 0x00,
	0xff, 0xff, 0xb4, 0xa4, 0x9e, 0x77, 0x35, 0x05, 0x00, 0x00,
}

func (m *LendAsset) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *LendAsset) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *LendAsset) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Reward_Accumulated.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintLend(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x3a
	n2, err2 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.LendingTime, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.LendingTime):])
	if err2 != nil {
		return 0, err2
	}
	i -= n2
	i = encodeVarintLend(dAtA, i, uint64(n2))
	i--
	dAtA[i] = 0x32
	{
		size, err := m.AmountIn.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintLend(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x2a
	if len(m.Owner) > 0 {
		i -= len(m.Owner)
		copy(dAtA[i:], m.Owner)
		i = encodeVarintLend(dAtA, i, uint64(len(m.Owner)))
		i--
		dAtA[i] = 0x22
	}
	if m.PoolId != 0 {
		i = encodeVarintLend(dAtA, i, uint64(m.PoolId))
		i--
		dAtA[i] = 0x18
	}
	if m.AssetId != 0 {
		i = encodeVarintLend(dAtA, i, uint64(m.AssetId))
		i--
		dAtA[i] = 0x10
	}
	if m.ID != 0 {
		i = encodeVarintLend(dAtA, i, uint64(m.ID))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
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
	if len(m.AssetData) > 0 {
		for iNdEx := len(m.AssetData) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.AssetData[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintLend(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.AccName) > 0 {
		i -= len(m.AccName)
		copy(dAtA[i:], m.AccName)
		i = encodeVarintLend(dAtA, i, uint64(len(m.AccName)))
		i--
		dAtA[i] = 0x12
	}
	if m.PoolId != 0 {
		i = encodeVarintLend(dAtA, i, uint64(m.PoolId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *AssetDataPoolMapping) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AssetDataPoolMapping) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AssetDataPoolMapping) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.BorrowRate.Size()
		i -= size
		if _, err := m.BorrowRate.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintLend(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	{
		size := m.LendRate.Size()
		i -= size
		if _, err := m.LendRate.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintLend(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if m.AssetId != 0 {
		i = encodeVarintLend(dAtA, i, uint64(m.AssetId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintLend(dAtA []byte, offset int, v uint64) int {
	offset -= sovLend(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *LendAsset) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ID != 0 {
		n += 1 + sovLend(uint64(m.ID))
	}
	if m.AssetId != 0 {
		n += 1 + sovLend(uint64(m.AssetId))
	}
	if m.PoolId != 0 {
		n += 1 + sovLend(uint64(m.PoolId))
	}
	l = len(m.Owner)
	if l > 0 {
		n += 1 + l + sovLend(uint64(l))
	}
	l = m.AmountIn.Size()
	n += 1 + l + sovLend(uint64(l))
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.LendingTime)
	n += 1 + l + sovLend(uint64(l))
	l = m.Reward_Accumulated.Size()
	n += 1 + l + sovLend(uint64(l))
	return n
}

func (m *Pool) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.PoolId != 0 {
		n += 1 + sovLend(uint64(m.PoolId))
	}
	l = len(m.AccName)
	if l > 0 {
		n += 1 + l + sovLend(uint64(l))
	}
	if len(m.AssetData) > 0 {
		for _, e := range m.AssetData {
			l = e.Size()
			n += 1 + l + sovLend(uint64(l))
		}
	}
	return n
}

func (m *AssetDataPoolMapping) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.AssetId != 0 {
		n += 1 + sovLend(uint64(m.AssetId))
	}
	l = m.LendRate.Size()
	n += 1 + l + sovLend(uint64(l))
	l = m.BorrowRate.Size()
	n += 1 + l + sovLend(uint64(l))
	return n
}

func sovLend(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozLend(x uint64) (n int) {
	return sovLend(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *LendAsset) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLend
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
			return fmt.Errorf("proto: LendAsset: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: LendAsset: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
			}
			m.ID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLend
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ID |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AssetId", wireType)
			}
			m.AssetId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLend
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.AssetId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PoolId", wireType)
			}
			m.PoolId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLend
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PoolId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Owner", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLend
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
				return ErrInvalidLengthLend
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLend
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Owner = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AmountIn", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLend
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
				return ErrInvalidLengthLend
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthLend
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.AmountIn.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LendingTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLend
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
				return ErrInvalidLengthLend
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthLend
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.LendingTime, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Reward_Accumulated", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLend
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
				return ErrInvalidLengthLend
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthLend
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Reward_Accumulated.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipLend(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthLend
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
func (m *Pool) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLend
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
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PoolId", wireType)
			}
			m.PoolId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLend
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PoolId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AccName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLend
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
				return ErrInvalidLengthLend
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLend
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AccName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AssetData", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLend
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
				return ErrInvalidLengthLend
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthLend
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AssetData = append(m.AssetData, AssetDataPoolMapping{})
			if err := m.AssetData[len(m.AssetData)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipLend(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthLend
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
func (m *AssetDataPoolMapping) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLend
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
			return fmt.Errorf("proto: AssetDataPoolMapping: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AssetDataPoolMapping: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AssetId", wireType)
			}
			m.AssetId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLend
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.AssetId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LendRate", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLend
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
				return ErrInvalidLengthLend
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLend
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.LendRate.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BorrowRate", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLend
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
				return ErrInvalidLengthLend
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLend
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.BorrowRate.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipLend(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthLend
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
func skipLend(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowLend
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
					return 0, ErrIntOverflowLend
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
					return 0, ErrIntOverflowLend
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
				return 0, ErrInvalidLengthLend
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupLend
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthLend
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthLend        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowLend          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupLend = fmt.Errorf("proto: unexpected end of group")
)
