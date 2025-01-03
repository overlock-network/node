package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgUpdateProvider{}

func NewMsgUpdateProvider(creator string, id uint64, name string, url string, environment uint64) *MsgUpdateProvider {
	return &MsgUpdateProvider{
		Creator:     creator,
		Id:          id,
		Name:        name,
		Url:         url,
		Environment: environment,
	}
}

func (msg *MsgUpdateProvider) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
