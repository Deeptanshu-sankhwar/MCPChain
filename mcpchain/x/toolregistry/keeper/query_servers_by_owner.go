package keeper

import (
	"context"

	"mcpchain/x/toolregistry/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) ServersByOwner(goCtx context.Context, req *types.QueryServersByOwnerRequest) (*types.QueryServersByOwnerResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	all := k.GetAllRegisteredServers(ctx)
	var filtered []types.RegisteredServer
	for _, s := range all {
		if s.OwnerAddress == req.OwnerAddress {
			filtered = append(filtered, s)
		}
	}

	servers := make([]*types.RegisteredServer, len(filtered))
	for i := range filtered {
		servers[i] = &filtered[i]
	}
	return &types.QueryServersByOwnerResponse{Servers: servers}, nil
}
