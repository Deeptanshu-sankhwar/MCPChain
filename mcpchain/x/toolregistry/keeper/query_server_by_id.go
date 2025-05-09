package keeper

import (
	"context"

	"mcpchain/x/toolregistry/types"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k Keeper) ServerById(goCtx context.Context, req *types.QueryServerByIdRequest) (*types.QueryServerByIdResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	server, found := k.GetServer(ctx, req.Id)
	if !found {
		return nil, errorsmod.Wrapf(sdkerrors.ErrNotFound, "server with id %d not found", req.Id)
	}

	return &types.QueryServerByIdResponse{Server: &server}, nil
}
