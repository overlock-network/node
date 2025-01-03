package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgUpdateEnvironment{}

func NewMsgUpdateEnvironment(creator string, name string, provider string) *MsgUpdateEnvironment {
	return &MsgUpdateEnvironment{
		Creator:  creator,
		Name:     name,
		Provider: provider,
	}
}

func (msg *MsgUpdateEnvironment) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
