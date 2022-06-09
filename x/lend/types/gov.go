package types

import (
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
)

var (
	ProposalAddLendPairs    = "ProposalAddLendPairs"
	ProposalUpdateLenddPair = "ProposalUpdateLenddPair"
	ProposalAddPool         = "ProposalAddPool"
)

func init() {

	govtypes.RegisterProposalType(ProposalAddLendPairs)
	govtypes.RegisterProposalTypeCodec(&LendPairsProposal{}, "comdex/AddLendPairsProposal")

	govtypes.RegisterProposalType(ProposalUpdateLenddPair)
	govtypes.RegisterProposalTypeCodec(&UpdatePairProposal{}, "comdex/UpdateLenddPairProposal")

	govtypes.RegisterProposalType(ProposalAddPool)
	govtypes.RegisterProposalTypeCodec(&AddPoolsProposal{}, "comdex/AddPoolsProposal")

}

var (
	_ govtypes.Content = &LendPairsProposal{}
	_ govtypes.Content = &UpdatePairProposal{}
)

func NewAddLendPairsProposal(title, description string, pairs []Extended_Pair) govtypes.Content {
	return &LendPairsProposal{
		Title:       title,
		Description: description,
		Pairs:       pairs,
	}
}

func (p *LendPairsProposal) ProposalRoute() string { return RouterKey }

func (p *LendPairsProposal) ProposalType() string { return ProposalAddLendPairs }

func (p *LendPairsProposal) ValidateBasic() error {

	err := govtypes.ValidateAbstract(p)
	if err != nil {
		return err
	}
	if len(p.Pairs) == 0 {
		return ErrorEmptyProposalAssets
	}

	for _, pair := range p.Pairs {
		if err := pair.Validate(); err != nil {
			return err
		}
	}

	return nil
}

func NewUpdateLendPairProposal(title, description string, pair Extended_Pair) govtypes.Content {
	return &UpdatePairProposal{
		Title:       title,
		Description: description,
		Pair:        pair,
	}
}

func (p *UpdatePairProposal) ProposalRoute() string { return RouterKey }

func (p *UpdatePairProposal) ProposalType() string { return ProposalUpdateLenddPair }

func (p *UpdatePairProposal) ValidateBasic() error {

	err := govtypes.ValidateAbstract(p)
	if err != nil {
		return err
	}

	pair := p.Pair
	if err := pair.Validate(); err != nil {
		return err
	}

	return nil
}

func NewAddPoolProposal(title, description string, pool Pool) govtypes.Content {
	return &AddPoolsProposal{
		Title:       title,
		Description: description,
		Pool:        pool,
	}
}

func (p *AddPoolsProposal) ProposalRoute() string {
	return RouterKey
}

func (p *AddPoolsProposal) ProposalType() string {
	return ProposalAddPool
}

func (p *AddPoolsProposal) ValidateBasic() error {
	err := govtypes.ValidateAbstract(p)
	if err != nil {
		return err
	}

	pool := p.Pool
	if err := pool.Validate(); err != nil {
		return err
	}

	return nil
}
