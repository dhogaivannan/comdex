package types

import (
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
)

const (
	ProposalLookupTableParams = "LookupTableParams"
)

func init() {
	govtypes.RegisterProposalType(ProposalLookupTableParams)
	govtypes.RegisterProposalTypeCodec(&LookupTableParams{}, "comdex/LookupTableParams")
}

var (
	_ govtypes.Content = &LookupTableParams{}
)

func NewLookupTableParamsProposal(title, description string, lookupTableData []CollectorLookupTable) govtypes.Content {
	return &LookupTableParams{
		Title:           title,
		Description:     description,
		LookupTableData: lookupTableData,
	}
}

func (p *LookupTableParams) ProposalRoute() string { return RouterKey }

func (p *LookupTableParams) ProposalType() string { return ProposalLookupTableParams }

func (p *LookupTableParams) ValidateBasic() error {

	err := govtypes.ValidateAbstract(p)
	if err != nil {
		return err
	}
	return nil
}

func NewAuctionLookupTableProposal(title, description string, appIdToAuctionLookup []AppIdToAuctionLookupTable) govtypes.Content {
	return &AuctionControlByAppIdProposal{
		Title:           title,
		Description:     description,
		AppIdToAuctionLookup: appIdToAuctionLookup,
	}
}

func (p *AuctionControlByAppIdProposal) ProposalRoute() string { return RouterKey }

func (p *AuctionControlByAppIdProposal) ProposalType() string { return ProposalLookupTableParams }

func (p *AuctionControlByAppIdProposal) ValidateBasic() error {

	err := govtypes.ValidateAbstract(p)
	if err != nil {
		return err
	}
	return nil
}
