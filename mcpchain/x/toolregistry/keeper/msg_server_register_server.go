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
	allServers := k.Keeper.GetAllRegisteredServers(ctx)
	for _, s := range allServers {
		if s.EndpointUrl == msg.EndpointUrl {
			return nil, errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "endpoint already registered: %s", msg.EndpointUrl)
		}
	}

	// Validate sender address
	sender, err := sdk.AccAddressFromBech32(msg.OwnerAddress)
	if err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidAddress, "invalid owner address")
	}

	// Parse staked amount
	amount, err := sdk.ParseCoinNormalized(msg.StakedAmount)
	if err != nil {
		return nil, errorsmod.Wrapf(sdkerrors.ErrInvalidCoins, "invalid staked amount: %v", err)
	}

	// Transfer tokens from owner to module account
	if err := k.bankKeeper.SendCoinsFromAccountToModule(goCtx, sender, types.ModuleName, sdk.NewCoins(amount)); err != nil {
		return nil, err
	}

	// Build RegisteredServer object
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

	// Store stake record
	stake := types.StakeRecord{
		ServerId:      id,
		Amount:        msg.StakedAmount,
		StakerAddress: msg.OwnerAddress,
	}
	if err := k.Keeper.SetStakeRecord(ctx, stake); err != nil {
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
