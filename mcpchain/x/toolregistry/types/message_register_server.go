package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgRegisterServer{}

func NewMsgRegisterServer(creator string, ownerAddress string, description string, endpointUrl string, mcpSchemaHash string, stakedAmount string) *MsgRegisterServer {
	return &MsgRegisterServer{
		Creator:       creator,
		OwnerAddress:  ownerAddress,
		Description:   description,
		EndpointUrl:   endpointUrl,
		McpSchemaHash: mcpSchemaHash,
		StakedAmount:  stakedAmount,
	}
}

func (msg *MsgRegisterServer) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(msg.OwnerAddress); err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid owner address: %v", err)
	}
	if msg.EndpointUrl == "" {
		return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "endpoint URL cannot be empty")
	}
	if msg.StakedAmount == "" {
		return errorsmod.Wrap(sdkerrors.ErrInvalidCoins, "stake must be non-empty")
	}
	return nil
}
