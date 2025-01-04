package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreateFunction{}

func NewMsgCreateFunction(creator string, metadata Metadata, spec FunctionSpec) *MsgCreateFunction {
	return &MsgCreateFunction{
		Creator:  creator,
		Metadata: &metadata,
		Spec:     &spec,
	}
}

func (msg *MsgCreateFunction) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
