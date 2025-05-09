package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgUpdateServerDetails{}

func NewMsgUpdateServerDetails(creator string, id uint64, description string, endpointUrl string) *MsgUpdateServerDetails {
	return &MsgUpdateServerDetails{
		Creator:     creator,
		Id:          id,
		Description: description,
		EndpointUrl: endpointUrl,
	}
}

func (msg *MsgUpdateServerDetails) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(msg.Creator); err != nil {
		return errorsmod.Wrap(sdkerrors.ErrInvalidAddress, "invalid creator address")
	}
	if msg.EndpointUrl == "" {
		return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "endpoint URL cannot be empty")
	}
	return nil
}
