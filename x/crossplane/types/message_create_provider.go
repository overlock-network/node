package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreateProvider{}

func NewMsgCreateProvider(creator string, name string, url string, environment uint64) *MsgCreateProvider {
	return &MsgCreateProvider{
		Creator:     creator,
		Name:        name,
		Url:         url,
		Environment: environment,
	}
}

func (msg *MsgCreateProvider) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
