package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreateXrd{}

func NewMsgCreateXrd(creator string, metadata *Metadata, spec *XrdSpec) *MsgCreateXrd {
	return &MsgCreateXrd{
		Creator:  creator,
		Metadata: metadata,
		Spec:     spec,
	}
}

func (msg *MsgCreateXrd) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
