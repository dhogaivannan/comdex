package rest

import (
	"net/http"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
	govrest "github.com/cosmos/cosmos-sdk/x/gov/client/rest"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"

	"github.com/comdex-official/comdex/x/rewards/types"
)

type AddNewMintingRewardsRequest struct {
	BaseReq                    rest.BaseReq `json:"base_req" yaml:"base_req"`
	Title                      string       `json:"title" yaml:"title"`
	Description                string       `json:"description" yaml:"description"`
	Deposit                    sdk.Coins    `json:"deposit" yaml:"deposit"`
	CollateralDenom            string       `json:"collateral_denom" yaml:"collateral_denom"`
	CAssetDenom                string       `json:"casset_denoms" yaml:"casset_denom"`
	TotalRewards               sdk.Coin     `json:"total_rewards" yaml:"total_rewards"`
	CAssetMaxcap               uint64       `json:"casset_maxcap" yaml:"casset_maxcap"`
	DurationDays               uint64       `json:"duration_days" yaml:"duration_days"`
	minimumLockupTimeInSeconds uint64       `json:"minimum_lockup_time_in_seconds" yaml:"minimum_lockup_time_in_seconds"`
}

type DisableMintingRewardsRequest struct {
	BaseReq         rest.BaseReq `json:"base_req" yaml:"base_req"`
	Title           string       `json:"title" yaml:"title"`
	Description     string       `json:"description" yaml:"description"`
	Deposit         sdk.Coins    `json:"deposit" yaml:"deposit"`
	MintingRewardId uint64       `json:"minting_reward_id" yaml:"minting_reward_id"`
}

func AddNewMintingRewardsProposalRESTHandler(clientCtx client.Context) govrest.ProposalRESTHandler {
	return govrest.ProposalRESTHandler{
		SubRoute: "add-new-mint-rewards",
		Handler:  AddNewMintingRewardsRESTHandler(clientCtx),
	}
}

func DisableMintingRewardsProposalRESTHandler(clientCtx client.Context) govrest.ProposalRESTHandler {
	return govrest.ProposalRESTHandler{
		SubRoute: "disable-mint-rewards",
		Handler:  DisableMintingRewardsRESTHandler(clientCtx),
	}
}

func AddNewMintingRewardsRESTHandler(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req AddNewMintingRewardsRequest

		if !rest.ReadRESTReq(w, r, clientCtx.LegacyAmino, &req) {
			return
		}

		req.BaseReq = req.BaseReq.Sanitize()
		if !req.BaseReq.ValidateBasic(w) {
			return
		}

		fromAddr, err := sdk.AccAddressFromBech32(req.BaseReq.From)
		if rest.CheckBadRequestError(w, err) {
			return
		}

		content := types.AddNewMintRewardsProposalContent(
			req.Title,
			req.Description,
			req.CollateralDenom,
			req.CAssetDenom,
			req.TotalRewards,
			req.CAssetMaxcap,
			req.DurationDays,
			req.minimumLockupTimeInSeconds,
		)
		msg, err := govtypes.NewMsgSubmitProposal(content, req.Deposit, fromAddr)
		if rest.CheckBadRequestError(w, err) {
			return
		}
		if rest.CheckBadRequestError(w, msg.ValidateBasic()) {
			return
		}

		tx.WriteGeneratedTxResponse(clientCtx, w, req.BaseReq, msg)
	}
}

func DisableMintingRewardsRESTHandler(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req DisableMintingRewardsRequest

		if !rest.ReadRESTReq(w, r, clientCtx.LegacyAmino, &req) {
			return
		}

		req.BaseReq = req.BaseReq.Sanitize()
		if !req.BaseReq.ValidateBasic(w) {
			return
		}

		fromAddr, err := sdk.AccAddressFromBech32(req.BaseReq.From)
		if rest.CheckBadRequestError(w, err) {
			return
		}

		content := types.DisableMintRewardsProposalContent(
			req.Title,
			req.Description,
			req.MintingRewardId,
		)
		msg, err := govtypes.NewMsgSubmitProposal(content, req.Deposit, fromAddr)
		if rest.CheckBadRequestError(w, err) {
			return
		}
		if rest.CheckBadRequestError(w, msg.ValidateBasic()) {
			return
		}

		tx.WriteGeneratedTxResponse(clientCtx, w, req.BaseReq, msg)
	}
}
