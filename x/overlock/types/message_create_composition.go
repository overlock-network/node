package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreateComposition{}

func NewMsgCreateComposition(creator string, metadata *Metadata, spec *CompositionSpec) *MsgCreateComposition {
	return &MsgCreateComposition{
		Creator:  creator,
		Metadata: metadata,
		Spec:     spec,
	}
}

func (msg *MsgCreateComposition) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
