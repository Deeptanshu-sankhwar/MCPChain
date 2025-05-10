package keeper

import (
	"context"
	"fmt"

	"mcpchain/x/attestation/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) LogAttestation(goCtx context.Context, msg *types.MsgLogAttestation) (*types.MsgLogAttestationResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Validate fields
	if msg.AgentId == "" || msg.ToolName == "" || msg.RequestHash == "" || msg.ResponseHash == "" {
		return nil, fmt.Errorf("missing required attestation fields")
	}

	// Optional: Verify server exists (you can connect to toolregistry keeper if needed)
	// For now, we assume ToolId is valid.

	record := types.AttestationRecord{
		InteractionId: fmt.Sprintf("ia_%d", ctx.BlockHeight()),
		AgentId:       msg.AgentId,
		ToolId:        msg.ToolId,
		ToolName:      msg.ToolName,
		RequestHash:   msg.RequestHash,
		ResponseHash:  msg.ResponseHash,
		Timestamp:     ctx.BlockTime().Unix(),
	}

	id, err := k.Keeper.AppendAttestation(ctx, record)
	if err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeLogAttestation,
			sdk.NewAttribute("id", fmt.Sprint(id)),
			sdk.NewAttribute("agent", msg.AgentId),
			sdk.NewAttribute("tool_id", fmt.Sprint(msg.ToolId)),
		),
	)

	return &types.MsgLogAttestationResponse{Id: id}, nil
}
