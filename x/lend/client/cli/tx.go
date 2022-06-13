package cli

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/gov/client/cli"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	"strconv"
	"time"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	// "github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/comdex-official/comdex/x/lend/types"
)

var (
	DefaultRelativePacketTimeoutTimestamp = uint64((time.Duration(10) * time.Minute).Nanoseconds())
)

const (
	flagPacketTimeoutTimestamp = "packet-timeout-timestamp"
	listSeparator              = ","
)

// GetTxCmd returns the transaction commands for this module
func GetTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("%s transactions subcommands", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(
		txLend(),
		txWithdraw(), //withdraw collateral partially or fully
		txDeposit(),
		txCloseLend(),
		txBorrowAsset(),
		txRepayAsset(), //including functionality of both repaying and closing position
		txFundModuleAccounts(),
	)

	return cmd
}

func txLend() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "lend [Asset_Id] [Amount] [Pool_Id]",
		Short: "lend a whitelisted asset",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			pair, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			asset, err := sdk.ParseCoinNormalized(args[1])
			if err != nil {
				return err
			}

			pool, err := strconv.ParseUint(args[2], 10, 64)
			if err != nil {
				return err
			}

			msg := types.NewMsgLend(ctx.GetFromAddress(), pair, asset, pool)

			return tx.GenerateOrBroadcastTxCLI(ctx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	return cmd

}

func txWithdraw() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "withdraw [lendId] [Amount]",
		Short: "withdraw lent asset",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			lendId, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			asset, err := sdk.ParseCoinNormalized(args[1])
			if err != nil {
				return err
			}

			msg := types.NewMsgWithdraw(ctx.GetFromAddress(), lendId, asset)

			return tx.GenerateOrBroadcastTxCLI(ctx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	return cmd

}

func txDeposit() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "deposit [lendId] [Amount]",
		Short: "deposit into a lent position",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			lendId, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			asset, err := sdk.ParseCoinNormalized(args[1])
			if err != nil {
				return err
			}

			msg := types.NewMsgDeposit(ctx.GetFromAddress(), lendId, asset)

			return tx.GenerateOrBroadcastTxCLI(ctx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	return cmd

}

func txCloseLend() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "close-lend [lendId]",
		Short: "close a lent position",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			lendId, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			msg := types.NewMsgCloseLend(ctx.GetFromAddress(), lendId)

			return tx.GenerateOrBroadcastTxCLI(ctx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	return cmd

}

func txBorrowAsset() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "borrow [lend-id] [pair-id] [Amount]",
		Short: "borrow a whitelisted asset",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			lendId, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			pairId, err := strconv.ParseUint(args[1], 10, 64)
			if err != nil {
				return err
			}

			asset, err := sdk.ParseCoinNormalized(args[2])
			if err != nil {
				return err
			}

			msg := types.NewMsgBorrow(ctx.GetFromAddress(), lendId, pairId, asset)

			return tx.GenerateOrBroadcastTxCLI(ctx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	return cmd

}

func txRepayAsset() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "repay [amount]",
		Short: "repay borrowed asset",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			asset, err := sdk.ParseCoinNormalized(args[0])
			if err != nil {
				return err
			}

			msg := types.NewMsgRepay(ctx.GetFromAddress(), asset)

			return tx.GenerateOrBroadcastTxCLI(ctx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	return cmd

}

func txFundModuleAccounts() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "fund-module [module-name] [amount]",
		Short: "Deposit amount to the respective module account",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			moduleName := args[0]

			amount, err := sdk.ParseCoinNormalized(args[1])
			if err != nil {
				return err
			}

			msg := types.NewMsgFundModuleAccounts(moduleName, ctx.GetFromAddress(), amount)
			err = msg.ValidateBasic()
			if err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(ctx, cmd.Flags(), msg)
		},
	}
	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

func CmdAddWNewLendPairsProposal() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add-lend-pairs [asset_in] [asset_out] [is_inter_pool] [asset_out_pool_id] [liquidation_ratio] [min_cr] [Base_Borrow_Rate_Asset1] [Base_Borrow_Rate_Asset2] [Base_Lend_Rate_Asset1] [Base_Lend_Rate_Asset2]",
		Short: "Add lend asset pairs",
		Args:  cobra.ExactArgs(10),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			assetIn, err := ParseUint64SliceFromString(args[0], ",")
			if err != nil {
				return err
			}

			assetOut, err := ParseUint64SliceFromString(args[1], ",")
			if err != nil {
				return err
			}

			isInterPool, err := ParseUint64SliceFromString(args[2], ",")
			if err != nil {
				return err
			}

			assetOutPoolId, err := ParseUint64SliceFromString(args[3], ",")
			if err != nil {
				return err
			}

			liquidationRatio, err := ParseStringFromString(args[4], ",")
			if err != nil {
				return err
			}

			minCr, err := ParseStringFromString(args[5], ",")
			if err != nil {
				return err
			}

			baseborrowrateasset1, err := ParseStringFromString(args[6], ",")
			if err != nil {
				return err
			}
			baseborrowrateasset2, err := ParseStringFromString(args[7], ",")
			if err != nil {
				return err
			}
			baselendrateasset1, err := ParseStringFromString(args[8], ",")
			if err != nil {
				return err
			}
			baselendrateasset2, err := ParseStringFromString(args[9], ",")
			if err != nil {
				return err
			}

			var pairs []types.Extended_Pair
			for i := range assetIn {

				newbaseborrowrateasset1, _ := sdk.NewDecFromStr(baseborrowrateasset1[i])
				newbaseborrowrateasset2, _ := sdk.NewDecFromStr(baseborrowrateasset2[i])
				newbaselendrateasset1, _ := sdk.NewDecFromStr(baselendrateasset1[i])
				newbaselendrateasset2, _ := sdk.NewDecFromStr(baselendrateasset2[i])
				newLiquidationRatio, _ := sdk.NewDecFromStr(liquidationRatio[i])
				newMinCr, _ := sdk.NewDecFromStr(minCr[i])
				interPool := ParseBoolFromString(isInterPool[i])
				pairs = append(pairs, types.Extended_Pair{
					AssetIn:                assetIn[i],
					AssetOut:               assetOut[i],
					IsInterPool:            interPool,
					AssetOutPoolId:         assetOutPoolId[i],
					LiquidationRatio:       newLiquidationRatio,
					MinCr:                  newMinCr,
					BaseBorrowRateAssetIn:  newbaseborrowrateasset1,
					BaseBorrowRateAssetOut: newbaseborrowrateasset2,
					BaseLendRateAssetIn:    newbaselendrateasset1,
					BaseLendRateAssetOut:   newbaselendrateasset2,
				})
			}

			title, err := cmd.Flags().GetString(cli.FlagTitle)
			if err != nil {
				return err
			}

			description, err := cmd.Flags().GetString(cli.FlagDescription)
			if err != nil {
				return err
			}

			from := clientCtx.GetFromAddress()

			depositStr, err := cmd.Flags().GetString(cli.FlagDeposit)
			if err != nil {
				return err
			}
			deposit, err := sdk.ParseCoinsNormalized(depositStr)
			if err != nil {
				return err
			}

			content := types.NewAddLendPairsProposal(title, description, pairs)

			msg, err := govtypes.NewMsgSubmitProposal(content, deposit, from)
			if err != nil {
				return err
			}

			if err = msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	cmd.Flags().String(cli.FlagTitle, "", "title of proposal")
	cmd.Flags().String(cli.FlagDescription, "", "description of proposal")
	cmd.Flags().String(cli.FlagDeposit, "", "deposit of proposal")
	_ = cmd.MarkFlagRequired(cli.FlagTitle)
	_ = cmd.MarkFlagRequired(cli.FlagDescription)

	return cmd
}

func CmdUpdateLendPairProposal() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-lend-pair [len_pair_id]",
		Short: "Update a lend asset pair",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			id, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			baseborrowrateasset1, err := cmd.Flags().GetString(flagbaseborrowrateasset1)
			if err != nil {
				return err
			}
			newbaseborrowrateasset1, err := sdk.NewDecFromStr(baseborrowrateasset1)
			if err != nil {
				return err
			}

			baseborrowrateasset2, err := cmd.Flags().GetString(flagbaseborrowrateasset2)
			if err != nil {
				return err
			}
			newbaseborrowrateasset2, err := sdk.NewDecFromStr(baseborrowrateasset2)
			if err != nil {
				return err
			}

			baselendrateasset1, err := cmd.Flags().GetString(flagbaselendrateasset1)
			if err != nil {
				return err
			}
			newbaselendrateasset1, err := sdk.NewDecFromStr(baselendrateasset1)
			if err != nil {
				return err
			}

			baselendrateasset2, err := cmd.Flags().GetString(flagbaselendrateasset2)
			if err != nil {
				return err
			}
			newbaselendrateasset2, err := sdk.NewDecFromStr(baselendrateasset2)
			if err != nil {
				return err
			}

			pair := types.Extended_Pair{
				Id:                     id,
				BaseBorrowRateAssetIn:  newbaseborrowrateasset1,
				BaseBorrowRateAssetOut: newbaseborrowrateasset2,
				BaseLendRateAssetIn:    newbaselendrateasset1,
				BaseLendRateAssetOut:   newbaselendrateasset2,
			}

			title, err := cmd.Flags().GetString(cli.FlagTitle)
			if err != nil {
				return err
			}

			description, err := cmd.Flags().GetString(cli.FlagDescription)
			if err != nil {
				return err
			}

			from := clientCtx.GetFromAddress()

			depositStr, err := cmd.Flags().GetString(cli.FlagDeposit)
			if err != nil {
				return err
			}
			deposit, err := sdk.ParseCoinsNormalized(depositStr)
			if err != nil {
				return err
			}

			content := types.NewUpdateLendPairProposal(title, description, pair)

			msg, err := govtypes.NewMsgSubmitProposal(content, deposit, from)
			if err != nil {
				return err
			}

			if err = msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	cmd.Flags().String(cli.FlagTitle, "", "title of proposal")
	cmd.Flags().String(cli.FlagDescription, "", "description of proposal")
	cmd.Flags().String(cli.FlagDeposit, "", "deposit of proposal")
	cmd.Flags().String(flagbaseborrowrateasset1, "", "baseborrowrateasset1")
	cmd.Flags().String(flagbaseborrowrateasset2, "", "baseborrowrateasset2")
	cmd.Flags().String(flagbaselendrateasset1, "", "baselendrateasset1")
	cmd.Flags().String(flagbaselendrateasset2, "", "baselendrateasset2")

	_ = cmd.MarkFlagRequired(cli.FlagTitle)
	_ = cmd.MarkFlagRequired(cli.FlagDescription)

	return cmd
}

func CmdAddPoolProposal() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add-lend-pool [module_name] [first_bridged_asset_id] [second_bridged_asset_id] [asset_id] [is_bridged_asset] ",
		Short: "Add lend pool ",
		Args:  cobra.ExactArgs(5),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			moduleName := args[0]

			firstBridgedAssetId, err := strconv.ParseUint(args[1], 10, 64)
			if err != nil {
				return err
			}

			secondBridgedAssetId, err := strconv.ParseUint(args[2], 10, 64)
			if err != nil {
				return err
			}

			assetId, err := ParseUint64SliceFromString(args[3], ",")
			if err != nil {
				return err
			}

			isBridgedAsset, err := ParseUint64SliceFromString(args[4], ",")
			if err != nil {
				return err
			}
			var pool types.Pool
			var assetData []types.AssetDataPoolMapping

			for i := range assetId {
				bridged := ParseBoolFromString(isBridgedAsset[i])
				assetData = append(assetData, types.AssetDataPoolMapping{
					AssetId:    assetId[i],
					IsBridged:  bridged,
					LendRate:   sdk.ZeroDec(),
					BorrowRate: sdk.ZeroDec(),
				})
			}
			pool = types.Pool{
				ModuleName:           moduleName,
				FirstBridgedAssetId:  firstBridgedAssetId,
				SecondBridgedAssetId: secondBridgedAssetId,
				AssetData:            assetData,
			}

			title, err := cmd.Flags().GetString(cli.FlagTitle)
			if err != nil {
				return err
			}

			description, err := cmd.Flags().GetString(cli.FlagDescription)
			if err != nil {
				return err
			}

			from := clientCtx.GetFromAddress()

			depositStr, err := cmd.Flags().GetString(cli.FlagDeposit)
			if err != nil {
				return err
			}
			deposit, err := sdk.ParseCoinsNormalized(depositStr)
			if err != nil {
				return err
			}

			content := types.NewAddPoolProposal(title, description, pool)

			msg, err := govtypes.NewMsgSubmitProposal(content, deposit, from)
			if err != nil {
				return err
			}

			if err = msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	cmd.Flags().String(cli.FlagTitle, "", "title of proposal")
	cmd.Flags().String(cli.FlagDescription, "", "description of proposal")
	cmd.Flags().String(cli.FlagDeposit, "", "deposit of proposal")
	_ = cmd.MarkFlagRequired(cli.FlagTitle)
	_ = cmd.MarkFlagRequired(cli.FlagDescription)

	return cmd
}

func CmdAddAssetToPairProposal() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add-asset-to-pair-mapping [asset_id] [pair_id] ",
		Short: "Add asset to pair mapping ",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			assetId, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			rawPairId, _ := ParseUint64SliceFromString(args[1], ",")
			if err != nil {
				return err
			}
			var pairId []uint64
			for i := range rawPairId {

				pairId = append(pairId, rawPairId[i])
			}
			assetToPairMapping := types.AssetToPairMapping{
				AssetId: assetId,
				PairId:  pairId,
			}

			title, err := cmd.Flags().GetString(cli.FlagTitle)
			if err != nil {
				return err
			}

			description, err := cmd.Flags().GetString(cli.FlagDescription)
			if err != nil {
				return err
			}

			from := clientCtx.GetFromAddress()

			depositStr, err := cmd.Flags().GetString(cli.FlagDeposit)
			if err != nil {
				return err
			}
			deposit, err := sdk.ParseCoinsNormalized(depositStr)
			if err != nil {
				return err
			}

			content := types.NewAddAssetToPairProposal(title, description, assetToPairMapping)

			msg, err := govtypes.NewMsgSubmitProposal(content, deposit, from)
			if err != nil {
				return err
			}

			if err = msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	cmd.Flags().String(cli.FlagTitle, "", "title of proposal")
	cmd.Flags().String(cli.FlagDescription, "", "description of proposal")
	cmd.Flags().String(cli.FlagDeposit, "", "deposit of proposal")
	_ = cmd.MarkFlagRequired(cli.FlagTitle)
	_ = cmd.MarkFlagRequired(cli.FlagDescription)

	return cmd
}
