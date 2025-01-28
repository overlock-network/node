package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgUpdateRegistry{}

func NewMsgUpdateRegistry(creator string, name string, provider string, id uint64) *MsgUpdateRegistry {
	return &MsgUpdateRegistry{
		Creator:  creator,
		Name:     name,
		Provider: provider,
		Id:       id,
	}
}

func (msg *MsgUpdateRegistry) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
