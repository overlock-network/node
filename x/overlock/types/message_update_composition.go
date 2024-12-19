package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgUpdateComposition{}

func NewMsgUpdateComposition(creator string, id uint64, apiVersion string, kind string, metadata *Metadata, spec *CompositionSpec) *MsgUpdateComposition {
	return &MsgUpdateComposition{
		Creator:    creator,
		Id:         id,
		ApiVersion: apiVersion,
		Kind:       kind,
		Metadata:   metadata,
		Spec:       spec,
	}
}

func (msg *MsgUpdateComposition) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
