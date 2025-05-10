package keeper

import (
	"context"

	"mcpchain/x/attestation/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) AttestationById(goCtx context.Context, req *types.QueryAttestationByIdRequest) (*types.QueryAttestationByIdResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	val, found := k.GetAttestation(ctx, req.Id)
	if !found {
		return nil, types.ErrAttestationNotFound
	}
	return &types.QueryAttestationByIdResponse{Attestation: &val}, nil
}

func (k Keeper) AttestationsByAgent(goCtx context.Context, req *types.QueryAttestationsByAgentRequest) (*types.QueryAttestationsByAgentResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	all := k.GetAllAttestations(ctx)
	out := make([]*types.AttestationRecord, 0)
	for _, att := range all {
		if att.AgentId == req.AgentId {
			attCopy := att
			out = append(out, &attCopy)
		}
	}
	return &types.QueryAttestationsByAgentResponse{Attestations: out}, nil
}

func (k Keeper) AttestationsByTool(goCtx context.Context, req *types.QueryAttestationsByToolRequest) (*types.QueryAttestationsByToolResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	all := k.GetAllAttestations(ctx)
	out := make([]*types.AttestationRecord, 0)
	for _, att := range all {
		if att.ToolId == req.ToolId {
			attCopy := att
			out = append(out, &attCopy)
		}
	}
	return &types.QueryAttestationsByToolResponse{Attestations: out}, nil
}
