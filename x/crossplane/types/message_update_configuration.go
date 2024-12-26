package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgUpdateConfiguration{}

func NewMsgUpdateConfiguration(creator string, id uint64, metadata *Metadata, spec *ConfigurationSpec) *MsgUpdateConfiguration {
	return &MsgUpdateConfiguration{
		Creator:  creator,
		Id:       id,
		Metadata: metadata,
		Spec:     spec,
	}
}

func (msg *MsgUpdateConfiguration) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
