package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"

	// this line is used by starport scaffolding # 1
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgLend{}, "comdex/lend/lend", nil)
	cdc.RegisterConcrete(&MsgWithdraw{}, "comdex/lend/withdraw", nil)
	cdc.RegisterConcrete(&MsgBorrow{}, "comdex/lend/borrow", nil)
	cdc.RegisterConcrete(&MsgRepay{}, "comdex/lend/repay", nil)
	cdc.RegisterConcrete(&MsgFundModuleAccounts{}, "comdex/lend/fund-module", nil)
	cdc.RegisterConcrete(&LendPairsProposal{}, "comdex/lend/add-lend-pairs", nil)
	cdc.RegisterConcrete(&UpdatePairProposal{}, "comdex/lend/update-lend-pairs", nil)
	cdc.RegisterConcrete(&AddPoolsProposal{}, "comdex/lend/add-lend-pools", nil)

}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations(
		(*govtypes.Content)(nil),
		&LendPairsProposal{},
		&UpdatePairProposal{},
		&AddPoolsProposal{},
	)
	registry.RegisterImplementations(
		(*sdk.Msg)(nil),
		&MsgLend{},
		&MsgWithdraw{},
		&MsgBorrow{},
		&MsgRepay{},
		&MsgFundModuleAccounts{},
	)

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
