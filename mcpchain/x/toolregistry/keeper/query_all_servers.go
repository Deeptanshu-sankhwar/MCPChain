package keeper

import (
	"context"

	"mcpchain/x/toolregistry/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) AllServers(goCtx context.Context, req *types.QueryAllServersRequest) (*types.QueryAllServersResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	all := k.GetAllRegisteredServers(ctx)
	servers := make([]*types.RegisteredServer, len(all))
	for i := range all {
		servers[i] = &all[i]
	}
	return &types.QueryAllServersResponse{Servers: servers}, nil
}
