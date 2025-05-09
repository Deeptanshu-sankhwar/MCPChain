package keeper

import (
	"context"
	"fmt"

	"mcpchain/x/toolregistry/types"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) RemoveServer(goCtx context.Context, msg *types.MsgRemoveServer) (*types.MsgRemoveServerResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	server, found := k.GetServer(ctx, msg.Id)
	if !found {
		return nil, errorsmod.Wrapf(sdkerrors.ErrNotFound, "server %d not found", msg.Id)
	}

	if server.OwnerAddress != msg.Creator {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "only the owner can remove")
	}

	if err := k.DeleteServer(ctx, msg.Id); err != nil {
		return nil, err
	}

	// TODO: implement unstaking logic

	ctx.EventManager().EmitEvent(
		sdk.NewEvent("remove_server",
			sdk.NewAttribute("id", fmt.Sprint(msg.Id)),
			sdk.NewAttribute("owner", msg.Creator),
		),
	)

	return &types.MsgRemoveServerResponse{}, nil
}
