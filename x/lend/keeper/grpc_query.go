package keeper

import (
	"context"
	"github.com/comdex-official/comdex/x/lend/types"
)

var (
	_ types.QueryServer = (*queryServer)(nil)
)

type queryServer struct {
	Keeper
}

func NewQueryServiceServer(k Keeper) types.QueryServer {
	return &queryServer{
		Keeper: k,
	}
}

func (q queryServer) QueryLends(ctx context.Context, request *types.QueryLendsRequest) (*types.QueryLendsResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (q queryServer) QueryLend(ctx context.Context, request *types.QueryLendRequest) (*types.QueryLendResponse, error) {
	//TODO implement me
	panic("implement me")
}
