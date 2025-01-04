package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgUpdateFunction{}

func NewMsgUpdateFunction(creator string, id uint64, metadata Metadata, spec FunctionSpec) *MsgUpdateFunction {
	return &MsgUpdateFunction{
		Creator:  creator,
		Id:       id,
		Metadata: &metadata,
		Spec:     &spec,
	}
}

func (msg *MsgUpdateFunction) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
