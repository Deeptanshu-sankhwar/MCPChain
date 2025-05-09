package keeper

import (
	"context"
	"fmt"

	"mcpchain/x/toolregistry/types"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) RegisterServer(goCtx context.Context, msg *types.MsgRegisterServer) (*types.MsgRegisterServerResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Optional: enforce unique endpoint URL
	allIds := k.Keeper.GetAllRegisteredServers(ctx)
	for _, s := range allIds {
		if s.EndpointUrl == msg.EndpointUrl {
			return nil, errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "endpoint already registered: %s", msg.EndpointUrl)
		}
	}

	server := types.RegisteredServer{
		OwnerAddress:          msg.OwnerAddress,
		Description:           msg.Description,
		EndpointUrl:           msg.EndpointUrl,
		McpSchemaHash:         msg.McpSchemaHash,
		StakedAmount:          msg.StakedAmount,
		RegistrationTimestamp: ctx.BlockTime().Unix(),
	}

	id, err := k.Keeper.AppendRegisteredServer(ctx, server)
	if err != nil {
		return nil, err
	}

	// Emit event
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeRegisterServer,
			sdk.NewAttribute("id", fmt.Sprint(id)),
			sdk.NewAttribute("owner", msg.OwnerAddress),
			sdk.NewAttribute("endpoint", msg.EndpointUrl),
			sdk.NewAttribute("mcp_schema_hash", msg.McpSchemaHash),
			sdk.NewAttribute("staked_amount", msg.StakedAmount),
		),
	)

	return &types.MsgRegisterServerResponse{Id: id}, nil
}
