package keeper

import (
	"context"
	"fmt"

	"mcpchain/x/toolregistry/types"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) UpdateServerDetails(goCtx context.Context, msg *types.MsgUpdateServerDetails) (*types.MsgUpdateServerDetailsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	server, found := k.GetServer(ctx, msg.Id)
	if !found {
		return nil, errorsmod.Wrapf(sdkerrors.ErrNotFound, "server %d not found", msg.Id)
	}

	if server.OwnerAddress != msg.Creator {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "only the owner can update")
	}

	server.Description = msg.Description
	server.EndpointUrl = msg.EndpointUrl

	if err := k.SetServer(ctx, server); err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvent(
		sdk.NewEvent("update_server",
			sdk.NewAttribute("id", fmt.Sprint(msg.Id)),
			sdk.NewAttribute("owner", msg.Creator),
		),
	)

	return &types.MsgUpdateServerDetailsResponse{}, nil
}
