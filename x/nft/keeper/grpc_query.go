package keeper

import (
	"context"
	"github.com/comdex-official/comdex/x/nft/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"strings"
)

var _ types.QueryServer = Keeper{}

func (k Keeper) Supply(c context.Context, request *types.QuerySupplyRequest) (*types.QuerySupplyResponse, error) {
	denom := strings.ToLower(strings.TrimSpace(request.DenomId))
	ctx := sdk.UnwrapSDKContext(c)

	var supply uint64
	switch {
	case len(request.Owner) == 0 && len(denom) > 0:
		supply = k.GetTotalSupply(ctx, denom)
	default:
		owner, err := sdk.AccAddressFromBech32(request.Owner)
		if err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "invalid owner address %s", request.Owner)
		}
		supply = k.GetTotalSupplyOfOwner(ctx, denom, owner)
	}
	return &types.QuerySupplyResponse{
		Amount: supply,
	}, nil
}

func (k Keeper) Collection(c context.Context, request *types.QueryCollectionRequest) (*types.QueryCollectionResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	collection, pagination, err := k.GetPaginateCollection(ctx, request, request.DenomId)
	if err != nil {
		return nil, err
	}
	return &types.QueryCollectionResponse{
		Collection: &collection,
		Pagination: pagination,
	}, nil
}

func (k Keeper) Denom(c context.Context, request *types.QueryDenomRequest) (*types.QueryDenomResponse, error) {
	denom := strings.ToLower(strings.TrimSpace(request.DenomId))
	ctx := sdk.UnwrapSDKContext(c)

	denomObject, err := k.GetDenom(ctx, denom)
	if err != nil {
		return nil, err
	}

	return &types.QueryDenomResponse{
		Denom: &denomObject,
	}, nil
}

func (k Keeper) Denoms(c context.Context, request *types.QueryDenomsRequest) (*types.QueryDenomsResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	var (
		denoms     []types.Denom
		pagination *query.PageResponse
		err        error
	)
	store := ctx.KVStore(k.storeKey)

	if request.Owner != "" {
		owner, err := sdk.AccAddressFromBech32(request.Owner)
		if err != nil {
			return nil, err
		}
		denomStore := prefix.NewStore(store, types.KeyDenomCreator(owner, ""))
		pagination, err = query.Paginate(denomStore, request.Pagination, func(key []byte, value []byte) error {
			denomId := types.MustUnMarshalDenomID(k.cdc, value)
			denom, _ := k.GetDenom(ctx, denomId)
			denoms = append(denoms, denom)
			return nil
		})
		if err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "paginate: %v", err)
		}

	} else {
		denomStore := prefix.NewStore(store, types.KeyDenomID(""))
		pagination, err = query.Paginate(denomStore, request.Pagination, func(key []byte, value []byte) error {
			var denom types.Denom
			k.cdc.MustUnmarshal(value, &denom)
			denoms = append(denoms, denom)
			return nil
		})
		if err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "paginate: %v", err)
		}
	}
	return &types.QueryDenomsResponse{
		Denoms:     denoms,
		Pagination: pagination,
	}, nil
}

func (k Keeper) NFT(c context.Context, request *types.QueryNFTRequest) (*types.QueryNFTResponse, error) {
	denom := strings.ToLower(strings.TrimSpace(request.DenomId))
	nftID := strings.ToLower(strings.TrimSpace(request.Id))
	ctx := sdk.UnwrapSDKContext(c)

	nft, err := k.GetNFT(ctx, denom, nftID)
	if err != nil {
		return nil, sdkerrors.Wrapf(types.ErrUnknownNFT, "invalid NFT %s from collection %s", request.Id, request.DenomId)
	}

	NFT, ok := nft.(types.NFT)
	if !ok {
		return nil, sdkerrors.Wrapf(types.ErrUnknownNFT, "invalid type NFT %s from collection %s", request.Id, request.DenomId)
	}

	return &types.QueryNFTResponse{
		NFT: &NFT,
	}, nil
}
func (k Keeper) OwnerNFTs(c context.Context, request *types.QueryOwnerNFTsRequest) (*types.QueryOwnerNFTsResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	address, err := sdk.AccAddressFromBech32(request.Owner)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid owner address %s", request.Owner)
	}

	owner := types.Owner{
		Address:       address.String(),
		IDCollections: types.IDCollections{},
	}
	var ownerCollections []types.OwnerNFTCollection
	idsMap := make(map[string][]string)
	store := ctx.KVStore(k.storeKey)
	nftStore := prefix.NewStore(store, types.KeyOwner(address, request.DenomId, ""))
	pagination, err := query.Paginate(nftStore, request.Pagination, func(key []byte, value []byte) error {
		denomId := request.DenomId
		nftId := string(key)
		if len(denomId) == 0 {
			denomId, nftId, _ = types.SplitKeyDenom(key)
		}
		if ids, ok := idsMap[denomId]; ok {
			idsMap[denomId] = append(ids, nftId)
		} else {
			idsMap[denomId] = []string{nftId}
			owner.IDCollections = append(
				owner.IDCollections,
				types.IDCollection{DenomId: denomId},
			)
		}
		return nil
	})
	for i := 0; i < len(owner.IDCollections); i++ {
		owner.IDCollections[i].NftIds = idsMap[owner.IDCollections[i].DenomId]
		denom, _ := k.GetDenom(ctx, owner.IDCollections[i].DenomId)
		var nfts []types.NFT
		for _, nftid := range owner.IDCollections[i].NftIds {
			nft, _ := k.GetNFT(ctx, denom.Id, nftid)
			nfts = append(nfts, nft.(types.NFT))
		}
		ownerCollection := types.OwnerNFTCollection{
			Denom: denom,
			Nfts:  nfts,
		}
		ownerCollections = append(ownerCollections, ownerCollection)
	}
	return &types.QueryOwnerNFTsResponse{
		Owner:       address.String(),
		Collections: ownerCollections,
		Pagination:  pagination,
	}, nil
}
