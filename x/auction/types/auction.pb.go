// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: comdex/auction/v1beta1/auction.proto

package types

import (
	fmt "fmt"
	types1 "github.com/comdex-official/comdex/x/asset/types"
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

type CollateralAuction struct {
	Id                  uint64                                        `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty" yaml:"id"`
	LockedVaultId       uint64                                        `protobuf:"varint,2,opt,name=locked_vault_id,json=lockedVaultId,proto3" json:"locked_vault_id,omitempty" yaml:"locked_vault_id"`
	AuctionedCollateral github_com_cosmos_cosmos_sdk_types.Coin       `protobuf:"bytes,3,opt,name=auctioned_collateral,json=auctionedCollateral,proto3,casttype=github.com/cosmos/cosmos-sdk/types.Coin" json:"auctioned_collateral" yaml:"auctioned_collateral"`
	DiscountQuantity    github_com_cosmos_cosmos_sdk_types.Coin       `protobuf:"bytes,4,opt,name=discount_quantity,json=discountQuantity,proto3,casttype=github.com/cosmos/cosmos-sdk/types.Coin" json:"discount_quantity" yaml:"discount_quantity"`
	ActiveBiddingId     uint64                                        `protobuf:"varint,5,opt,name=active_bidding_id,json=activeBiddingId,proto3" json:"active_bidding_id,omitempty" yaml:"active_bidding_id"`
	Bidder              github_com_cosmos_cosmos_sdk_types.AccAddress `protobuf:"bytes,6,opt,name=bidder,proto3,casttype=github.com/cosmos/cosmos-sdk/types.AccAddress" json:"bidder,omitempty" yaml:"owner"`
	Bid                 github_com_cosmos_cosmos_sdk_types.Coin       `protobuf:"bytes,7,opt,name=bid,proto3,casttype=github.com/cosmos/cosmos-sdk/types.Coin" json:"bid" yaml:"bid"`
	MinBid              github_com_cosmos_cosmos_sdk_types.Coin       `protobuf:"bytes,8,opt,name=min_bid,json=minBid,proto3,casttype=github.com/cosmos/cosmos-sdk/types.Coin" json:"min_bid" yaml:"min_bid"`
	MaxBid              github_com_cosmos_cosmos_sdk_types.Coin       `protobuf:"bytes,9,opt,name=max_bid,json=maxBid,proto3,casttype=github.com/cosmos/cosmos-sdk/types.Coin" json:"max_bid" yaml:"max_bid"`
	EndTime             time.Time                                     `protobuf:"bytes,10,opt,name=end_time,json=endTime,proto3,stdtime" json:"end_time" yaml:"end_time"`
	Pair                types1.Pair                                   `protobuf:"bytes,11,opt,name=pair,proto3" json:"pair" yaml:"pair"`
	BiddingIds          []uint64                                      `protobuf:"varint,12,rep,packed,name=bidding_ids,json=biddingIds,proto3" json:"bidding_ids,omitempty" yaml:"bidding_ids"`
}

func (m *CollateralAuction) Reset()         { *m = CollateralAuction{} }
func (m *CollateralAuction) String() string { return proto.CompactTextString(m) }
func (*CollateralAuction) ProtoMessage()    {}
func (*CollateralAuction) Descriptor() ([]byte, []int) {
	return fileDescriptor_4bb9aead25d5fe6c, []int{0}
}
func (m *CollateralAuction) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CollateralAuction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CollateralAuction.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CollateralAuction) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CollateralAuction.Merge(m, src)
}
func (m *CollateralAuction) XXX_Size() int {
	return m.Size()
}
func (m *CollateralAuction) XXX_DiscardUnknown() {
	xxx_messageInfo_CollateralAuction.DiscardUnknown(m)
}

var xxx_messageInfo_CollateralAuction proto.InternalMessageInfo

type DebtAuction struct {
	AuctionId           uint64                                        `protobuf:"varint,1,opt,name=auction_id,json=auctionId,proto3" json:"auction_id,omitempty" yaml:"id"`
	AuctionedToken      github_com_cosmos_cosmos_sdk_types.Coin       `protobuf:"bytes,2,opt,name=auctioned_token,json=auctionedToken,proto3,casttype=github.com/cosmos/cosmos-sdk/types.Coin" json:"auctioned_token" yaml:"auctioned_token"`
	ExpectedUserToken   github_com_cosmos_cosmos_sdk_types.Coin       `protobuf:"bytes,3,opt,name=expected_user_token,json=expectedUserToken,proto3,casttype=github.com/cosmos/cosmos-sdk/types.Coin" json:"expected_user_token" yaml:"expected_token"`
	ExpectedMintedToken github_com_cosmos_cosmos_sdk_types.Coin       `protobuf:"bytes,4,opt,name=expected_minted_token,json=expectedMintedToken,proto3,casttype=github.com/cosmos/cosmos-sdk/types.Coin" json:"expected_minted_token" yaml:"expected_token"`
	EndTime             time.Time                                     `protobuf:"bytes,5,opt,name=end_time,json=endTime,proto3,stdtime" json:"end_time" yaml:"end_time"`
	ActiveBiddingId     uint64                                        `protobuf:"varint,6,opt,name=active_bidding_id,json=activeBiddingId,proto3" json:"active_bidding_id,omitempty" yaml:"active_bidding_id"`
	Bidder              github_com_cosmos_cosmos_sdk_types.AccAddress `protobuf:"bytes,7,opt,name=bidder,proto3,casttype=github.com/cosmos/cosmos-sdk/types.AccAddress" json:"bidder,omitempty" yaml:"owner"`
	CurrentBidAmount    github_com_cosmos_cosmos_sdk_types.Coin       `protobuf:"bytes,8,opt,name=current_bid_amount,json=currentBidAmount,proto3,casttype=github.com/cosmos/cosmos-sdk/types.Coin" json:"current_bid_amount" yaml:"min_bid"`
	AuctionStatus       uint64                                        `protobuf:"varint,9,opt,name=auction_status,json=auctionStatus,proto3" json:"auction_status,omitempty" yaml:"auction_status"`
}

func (m *DebtAuction) Reset()         { *m = DebtAuction{} }
func (m *DebtAuction) String() string { return proto.CompactTextString(m) }
func (*DebtAuction) ProtoMessage()    {}
func (*DebtAuction) Descriptor() ([]byte, []int) {
	return fileDescriptor_4bb9aead25d5fe6c, []int{1}
}
func (m *DebtAuction) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DebtAuction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DebtAuction.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DebtAuction) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DebtAuction.Merge(m, src)
}
func (m *DebtAuction) XXX_Size() int {
	return m.Size()
}
func (m *DebtAuction) XXX_DiscardUnknown() {
	xxx_messageInfo_DebtAuction.DiscardUnknown(m)
}

var xxx_messageInfo_DebtAuction proto.InternalMessageInfo

func init() {
	proto.RegisterType((*CollateralAuction)(nil), "comdex.auction.v1beta1.CollateralAuction")
	proto.RegisterType((*DebtAuction)(nil), "comdex.auction.v1beta1.DebtAuction")
}

func init() {
	proto.RegisterFile("comdex/auction/v1beta1/auction.proto", fileDescriptor_4bb9aead25d5fe6c)
}

var fileDescriptor_4bb9aead25d5fe6c = []byte{
	// 846 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x56, 0x4f, 0x6f, 0xe3, 0x44,
	0x1c, 0x8d, 0xdb, 0x6c, 0xd2, 0x4e, 0xb6, 0xdb, 0xad, 0xbb, 0x5b, 0x79, 0xbb, 0x60, 0x47, 0x16,
	0x12, 0x39, 0x50, 0x5b, 0x85, 0x03, 0x12, 0x12, 0x12, 0x71, 0x39, 0x6c, 0x84, 0x40, 0xe0, 0x5d,
	0x90, 0xe0, 0x62, 0x8d, 0x3d, 0xd3, 0x30, 0xd4, 0xf6, 0x04, 0xcf, 0xb8, 0xa4, 0x5f, 0x01, 0x09,
	0x69, 0x25, 0x0e, 0x5c, 0xf8, 0x00, 0x7c, 0x0f, 0x2e, 0x3d, 0xee, 0x91, 0x53, 0x80, 0xf4, 0x1b,
	0xe4, 0x88, 0x38, 0xa0, 0xf9, 0xe7, 0xb0, 0x7f, 0x20, 0x44, 0x54, 0x7b, 0x6a, 0x67, 0xe6, 0xfd,
	0xde, 0x7b, 0xf3, 0x9b, 0x9f, 0x9f, 0x02, 0x5e, 0xcb, 0x68, 0x81, 0xf0, 0x34, 0x84, 0x75, 0xc6,
	0x09, 0x2d, 0xc3, 0xf3, 0xe3, 0x14, 0x73, 0x78, 0x6c, 0xd6, 0xc1, 0xa4, 0xa2, 0x9c, 0xda, 0x07,
	0x0a, 0x15, 0x98, 0x5d, 0x8d, 0x3a, 0xbc, 0x33, 0xa6, 0x63, 0x2a, 0x21, 0xa1, 0xf8, 0x4f, 0xa1,
	0x0f, 0xbd, 0x31, 0xa5, 0xe3, 0x1c, 0x87, 0x72, 0x95, 0xd6, 0xa7, 0x21, 0x27, 0x05, 0x66, 0x1c,
	0x16, 0x13, 0x0d, 0x70, 0x33, 0xca, 0x0a, 0xca, 0xc2, 0x14, 0x32, 0xdc, 0x28, 0x66, 0x94, 0x94,
	0x86, 0xc0, 0x98, 0x62, 0x0c, 0xf3, 0x06, 0x30, 0x81, 0xa4, 0x52, 0x00, 0xff, 0xcf, 0x2d, 0xb0,
	0x77, 0x42, 0xf3, 0x1c, 0x72, 0x5c, 0xc1, 0x7c, 0xa8, 0x5c, 0xd9, 0xaf, 0x82, 0x0d, 0x82, 0x1c,
	0xab, 0x6f, 0x0d, 0xda, 0xd1, 0xce, 0x62, 0xe6, 0x6d, 0x5f, 0xc0, 0x22, 0x7f, 0xc7, 0x27, 0xc8,
	0x8f, 0x37, 0x08, 0xb2, 0x23, 0xb0, 0x9b, 0xd3, 0xec, 0x0c, 0xa3, 0xe4, 0x1c, 0xd6, 0x39, 0x4f,
	0x08, 0x72, 0x36, 0x24, 0xf6, 0x70, 0x31, 0xf3, 0x0e, 0x14, 0xf6, 0x19, 0x80, 0x1f, 0xef, 0xa8,
	0x9d, 0xcf, 0xc4, 0xc6, 0x08, 0xd9, 0x3f, 0x5a, 0xe0, 0x8e, 0x6e, 0x02, 0x46, 0x49, 0xd6, 0x58,
	0x70, 0x36, 0xfb, 0xd6, 0xa0, 0xf7, 0xe6, 0xbd, 0x40, 0xdd, 0x2c, 0x10, 0x37, 0x33, 0x5d, 0x0a,
	0x4e, 0x28, 0x29, 0xa3, 0x8f, 0x2e, 0x67, 0x5e, 0x6b, 0x31, 0xf3, 0xee, 0x2b, 0xa1, 0x17, 0x91,
	0xf8, 0x7f, 0xcc, 0xbc, 0xd7, 0xc7, 0x84, 0x7f, 0x59, 0xa7, 0x41, 0x46, 0x8b, 0x50, 0x77, 0x49,
	0xfd, 0x39, 0x62, 0xe8, 0x2c, 0xe4, 0x17, 0x13, 0xcc, 0x24, 0x5f, 0xbc, 0xdf, 0x30, 0x2c, 0x1b,
	0x61, 0x7f, 0x6f, 0x81, 0x3d, 0x44, 0x58, 0x46, 0xeb, 0x92, 0x27, 0x5f, 0xd7, 0xb0, 0xe4, 0x84,
	0x5f, 0x38, 0xed, 0x55, 0xde, 0x3e, 0xd0, 0xde, 0x1c, 0xe5, 0xed, 0x39, 0x86, 0xb5, 0x8c, 0xdd,
	0x36, 0xe5, 0x9f, 0xe8, 0x6a, 0xfb, 0x01, 0xd8, 0x83, 0x19, 0x27, 0xe7, 0x38, 0x49, 0x09, 0x42,
	0xa4, 0x1c, 0x8b, 0xd6, 0xdf, 0x90, 0xad, 0x7f, 0x65, 0xa9, 0xfa, 0x1c, 0xc4, 0x8f, 0x77, 0xd5,
	0x5e, 0xa4, 0xb6, 0x46, 0xc8, 0xfe, 0x1c, 0x74, 0xc4, 0x39, 0xae, 0x9c, 0x4e, 0xdf, 0x1a, 0x6c,
	0x47, 0xc3, 0xc5, 0xcc, 0xbb, 0xa9, 0xca, 0xe9, 0x37, 0x25, 0xae, 0x84, 0xd1, 0xa3, 0xff, 0x60,
	0x74, 0x98, 0x65, 0x43, 0x84, 0x2a, 0xcc, 0x58, 0xac, 0x09, 0xed, 0xaf, 0xc0, 0x66, 0x4a, 0x90,
	0xd3, 0x5d, 0xd5, 0xab, 0x77, 0x75, 0xaf, 0x80, 0x92, 0x4d, 0x09, 0x5a, 0xab, 0x3b, 0x42, 0xc4,
	0xae, 0x41, 0xb7, 0x20, 0xa5, 0xb8, 0xaa, 0xb3, 0xb5, 0x4a, 0x6f, 0xa8, 0xf5, 0x6e, 0x29, 0x3d,
	0x5d, 0xb7, 0x96, 0x66, 0xa7, 0x20, 0x65, 0xa4, 0x65, 0xe1, 0x54, 0xca, 0x6e, 0xaf, 0x2b, 0xab,
	0xea, 0xd6, 0x94, 0x85, 0x53, 0x21, 0x1b, 0x83, 0x2d, 0x5c, 0xa2, 0x44, 0x84, 0x80, 0x03, 0xa4,
	0xee, 0x61, 0xa0, 0x12, 0x22, 0x30, 0x09, 0x11, 0x3c, 0x32, 0x09, 0x11, 0xdd, 0xd7, 0xc2, 0xbb,
	0x4a, 0xd8, 0x54, 0xfa, 0x8f, 0x7f, 0xf5, 0xac, 0xb8, 0x8b, 0x4b, 0x24, 0xa0, 0xf6, 0x09, 0x68,
	0x8b, 0x38, 0x70, 0x7a, 0x9a, 0xcf, 0xe4, 0x93, 0x08, 0x8c, 0xe6, 0x22, 0x1f, 0x43, 0x52, 0x45,
	0xfb, 0x9a, 0xaf, 0xa7, 0xf8, 0x44, 0x95, 0x1f, 0xcb, 0x62, 0xfb, 0x6d, 0xd0, 0x5b, 0x4e, 0x1b,
	0x73, 0x6e, 0xf6, 0x37, 0x07, 0xed, 0xe8, 0x60, 0x31, 0xf3, 0xec, 0xe6, 0x6d, 0xcd, 0xa1, 0x1f,
	0x83, 0xd4, 0x4c, 0x21, 0xf3, 0x7f, 0xee, 0x82, 0xde, 0xfb, 0x38, 0xe5, 0x26, 0x78, 0xde, 0x00,
	0x40, 0x7f, 0x8d, 0xc9, 0x3f, 0x05, 0xd0, 0xb6, 0x06, 0x8c, 0x90, 0xfd, 0x9d, 0x05, 0x76, 0x97,
	0x9f, 0x3f, 0xa7, 0x67, 0xb8, 0x94, 0x41, 0xf4, 0xaf, 0xef, 0x31, 0xd2, 0xd7, 0x38, 0x78, 0x36,
	0x3e, 0x64, 0xfd, 0x5a, 0xef, 0x72, 0xab, 0x29, 0x7e, 0x24, 0x6a, 0x45, 0x68, 0xec, 0xe3, 0xe9,
	0x04, 0x67, 0x1c, 0xa3, 0xa4, 0x66, 0xb8, 0xd2, 0x9e, 0x56, 0x46, 0xda, 0x03, 0xed, 0xe9, 0xae,
	0x7e, 0x2a, 0xc3, 0xb1, 0xbe, 0xa5, 0x3d, 0x53, 0xfb, 0x29, 0xc3, 0x95, 0x72, 0xf5, 0x83, 0x05,
	0xee, 0x36, 0x8c, 0x05, 0x29, 0x1b, 0xe2, 0xd5, 0x71, 0x76, 0x7d, 0xbe, 0x9a, 0xbe, 0x7c, 0x28,
	0x0d, 0x28, 0x67, 0x7f, 0x9f, 0xe7, 0x1b, 0xd7, 0x34, 0xcf, 0x2f, 0x8c, 0xc8, 0xce, 0xff, 0x8b,
	0xc8, 0xee, 0x75, 0x47, 0xe4, 0xb7, 0x16, 0xb0, 0xb3, 0xba, 0xaa, 0x70, 0xc9, 0x85, 0x87, 0x04,
	0x16, 0x22, 0xe6, 0x5f, 0x4a, 0x84, 0xdd, 0xd6, 0xba, 0x11, 0x41, 0x43, 0xa9, 0x6a, 0xbf, 0x07,
	0xcc, 0x1c, 0x27, 0x8c, 0x43, 0x5e, 0x33, 0x99, 0x69, 0xed, 0xe8, 0xde, 0xf2, 0xe1, 0x9f, 0x3e,
	0xf7, 0xe3, 0x1d, 0xbd, 0xf1, 0x50, 0xae, 0xa3, 0x87, 0x97, 0xbf, 0xbb, 0xad, 0x9f, 0xe6, 0x6e,
	0xeb, 0x72, 0xee, 0x5a, 0x4f, 0xe6, 0xae, 0xf5, 0xdb, 0xdc, 0xb5, 0x1e, 0x5f, 0xb9, 0xad, 0x27,
	0x57, 0x6e, 0xeb, 0x97, 0x2b, 0xb7, 0xf5, 0xc5, 0xf1, 0x53, 0xf6, 0x44, 0xc2, 0x1c, 0xd1, 0xd3,
	0x53, 0x92, 0x11, 0x98, 0xeb, 0x75, 0xb8, 0xfc, 0xe5, 0x24, 0xdd, 0xa6, 0x1d, 0x39, 0x02, 0x6f,
	0xfd, 0x15, 0x00, 0x00, 0xff, 0xff, 0xaf, 0xd3, 0x8d, 0xde, 0x58, 0x09, 0x00, 0x00,
}

func (m *CollateralAuction) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CollateralAuction) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CollateralAuction) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.BiddingIds) > 0 {
		dAtA2 := make([]byte, len(m.BiddingIds)*10)
		var j1 int
		for _, num := range m.BiddingIds {
			for num >= 1<<7 {
				dAtA2[j1] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j1++
			}
			dAtA2[j1] = uint8(num)
			j1++
		}
		i -= j1
		copy(dAtA[i:], dAtA2[:j1])
		i = encodeVarintAuction(dAtA, i, uint64(j1))
		i--
		dAtA[i] = 0x62
	}
	{
		size, err := m.Pair.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintAuction(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x5a
	n4, err4 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.EndTime, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.EndTime):])
	if err4 != nil {
		return 0, err4
	}
	i -= n4
	i = encodeVarintAuction(dAtA, i, uint64(n4))
	i--
	dAtA[i] = 0x52
	{
		size, err := m.MaxBid.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintAuction(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x4a
	{
		size, err := m.MinBid.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintAuction(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x42
	{
		size, err := m.Bid.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintAuction(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x3a
	if len(m.Bidder) > 0 {
		i -= len(m.Bidder)
		copy(dAtA[i:], m.Bidder)
		i = encodeVarintAuction(dAtA, i, uint64(len(m.Bidder)))
		i--
		dAtA[i] = 0x32
	}
	if m.ActiveBiddingId != 0 {
		i = encodeVarintAuction(dAtA, i, uint64(m.ActiveBiddingId))
		i--
		dAtA[i] = 0x28
	}
	{
		size, err := m.DiscountQuantity.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintAuction(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	{
		size, err := m.AuctionedCollateral.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintAuction(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	if m.LockedVaultId != 0 {
		i = encodeVarintAuction(dAtA, i, uint64(m.LockedVaultId))
		i--
		dAtA[i] = 0x10
	}
	if m.Id != 0 {
		i = encodeVarintAuction(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *DebtAuction) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DebtAuction) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DebtAuction) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.AuctionStatus != 0 {
		i = encodeVarintAuction(dAtA, i, uint64(m.AuctionStatus))
		i--
		dAtA[i] = 0x48
	}
	{
		size, err := m.CurrentBidAmount.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintAuction(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x42
	if len(m.Bidder) > 0 {
		i -= len(m.Bidder)
		copy(dAtA[i:], m.Bidder)
		i = encodeVarintAuction(dAtA, i, uint64(len(m.Bidder)))
		i--
		dAtA[i] = 0x3a
	}
	if m.ActiveBiddingId != 0 {
		i = encodeVarintAuction(dAtA, i, uint64(m.ActiveBiddingId))
		i--
		dAtA[i] = 0x30
	}
	n11, err11 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.EndTime, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.EndTime):])
	if err11 != nil {
		return 0, err11
	}
	i -= n11
	i = encodeVarintAuction(dAtA, i, uint64(n11))
	i--
	dAtA[i] = 0x2a
	{
		size, err := m.ExpectedMintedToken.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintAuction(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	{
		size, err := m.ExpectedUserToken.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintAuction(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	{
		size, err := m.AuctionedToken.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintAuction(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if m.AuctionId != 0 {
		i = encodeVarintAuction(dAtA, i, uint64(m.AuctionId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintAuction(dAtA []byte, offset int, v uint64) int {
	offset -= sovAuction(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *CollateralAuction) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovAuction(uint64(m.Id))
	}
	if m.LockedVaultId != 0 {
		n += 1 + sovAuction(uint64(m.LockedVaultId))
	}
	l = m.AuctionedCollateral.Size()
	n += 1 + l + sovAuction(uint64(l))
	l = m.DiscountQuantity.Size()
	n += 1 + l + sovAuction(uint64(l))
	if m.ActiveBiddingId != 0 {
		n += 1 + sovAuction(uint64(m.ActiveBiddingId))
	}
	l = len(m.Bidder)
	if l > 0 {
		n += 1 + l + sovAuction(uint64(l))
	}
	l = m.Bid.Size()
	n += 1 + l + sovAuction(uint64(l))
	l = m.MinBid.Size()
	n += 1 + l + sovAuction(uint64(l))
	l = m.MaxBid.Size()
	n += 1 + l + sovAuction(uint64(l))
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.EndTime)
	n += 1 + l + sovAuction(uint64(l))
	l = m.Pair.Size()
	n += 1 + l + sovAuction(uint64(l))
	if len(m.BiddingIds) > 0 {
		l = 0
		for _, e := range m.BiddingIds {
			l += sovAuction(uint64(e))
		}
		n += 1 + sovAuction(uint64(l)) + l
	}
	return n
}

func (m *DebtAuction) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.AuctionId != 0 {
		n += 1 + sovAuction(uint64(m.AuctionId))
	}
	l = m.AuctionedToken.Size()
	n += 1 + l + sovAuction(uint64(l))
	l = m.ExpectedUserToken.Size()
	n += 1 + l + sovAuction(uint64(l))
	l = m.ExpectedMintedToken.Size()
	n += 1 + l + sovAuction(uint64(l))
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.EndTime)
	n += 1 + l + sovAuction(uint64(l))
	if m.ActiveBiddingId != 0 {
		n += 1 + sovAuction(uint64(m.ActiveBiddingId))
	}
	l = len(m.Bidder)
	if l > 0 {
		n += 1 + l + sovAuction(uint64(l))
	}
	l = m.CurrentBidAmount.Size()
	n += 1 + l + sovAuction(uint64(l))
	if m.AuctionStatus != 0 {
		n += 1 + sovAuction(uint64(m.AuctionStatus))
	}
	return n
}

func sovAuction(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozAuction(x uint64) (n int) {
	return sovAuction(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *CollateralAuction) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAuction
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
			return fmt.Errorf("proto: CollateralAuction: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CollateralAuction: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuction
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
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LockedVaultId", wireType)
			}
			m.LockedVaultId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuction
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.LockedVaultId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AuctionedCollateral", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuction
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
				return ErrInvalidLengthAuction
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthAuction
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.AuctionedCollateral.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DiscountQuantity", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuction
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
				return ErrInvalidLengthAuction
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthAuction
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.DiscountQuantity.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ActiveBiddingId", wireType)
			}
			m.ActiveBiddingId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuction
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ActiveBiddingId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Bidder", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuction
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
				return ErrInvalidLengthAuction
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAuction
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Bidder = github_com_cosmos_cosmos_sdk_types.AccAddress(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Bid", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuction
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
				return ErrInvalidLengthAuction
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthAuction
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Bid.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinBid", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuction
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
				return ErrInvalidLengthAuction
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthAuction
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.MinBid.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxBid", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuction
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
				return ErrInvalidLengthAuction
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthAuction
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.MaxBid.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EndTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuction
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
				return ErrInvalidLengthAuction
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthAuction
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.EndTime, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 11:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pair", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuction
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
				return ErrInvalidLengthAuction
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthAuction
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Pair.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 12:
			if wireType == 0 {
				var v uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowAuction
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.BiddingIds = append(m.BiddingIds, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowAuction
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthAuction
				}
				postIndex := iNdEx + packedLen
				if postIndex < 0 {
					return ErrInvalidLengthAuction
				}
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				var elementCount int
				var count int
				for _, integer := range dAtA[iNdEx:postIndex] {
					if integer < 128 {
						count++
					}
				}
				elementCount = count
				if elementCount != 0 && len(m.BiddingIds) == 0 {
					m.BiddingIds = make([]uint64, 0, elementCount)
				}
				for iNdEx < postIndex {
					var v uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowAuction
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.BiddingIds = append(m.BiddingIds, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field BiddingIds", wireType)
			}
		default:
			iNdEx = preIndex
			skippy, err := skipAuction(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthAuction
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
func (m *DebtAuction) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAuction
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
			return fmt.Errorf("proto: DebtAuction: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DebtAuction: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AuctionId", wireType)
			}
			m.AuctionId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuction
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.AuctionId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AuctionedToken", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuction
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
				return ErrInvalidLengthAuction
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthAuction
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.AuctionedToken.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ExpectedUserToken", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuction
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
				return ErrInvalidLengthAuction
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthAuction
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ExpectedUserToken.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ExpectedMintedToken", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuction
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
				return ErrInvalidLengthAuction
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthAuction
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ExpectedMintedToken.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EndTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuction
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
				return ErrInvalidLengthAuction
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthAuction
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.EndTime, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ActiveBiddingId", wireType)
			}
			m.ActiveBiddingId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuction
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ActiveBiddingId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Bidder", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuction
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
				return ErrInvalidLengthAuction
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAuction
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Bidder = github_com_cosmos_cosmos_sdk_types.AccAddress(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CurrentBidAmount", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuction
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
				return ErrInvalidLengthAuction
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthAuction
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.CurrentBidAmount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 9:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AuctionStatus", wireType)
			}
			m.AuctionStatus = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuction
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.AuctionStatus |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipAuction(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthAuction
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
func skipAuction(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowAuction
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
					return 0, ErrIntOverflowAuction
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
					return 0, ErrIntOverflowAuction
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
				return 0, ErrInvalidLengthAuction
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupAuction
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthAuction
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthAuction        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowAuction          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupAuction = fmt.Errorf("proto: unexpected end of group")
)
