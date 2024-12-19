package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgDeleteComposition{}

func NewMsgDeleteComposition(creator string, id uint64) *MsgDeleteComposition {
	return &MsgDeleteComposition{
		Creator: creator,
		Id:      id,
	}
}

func (msg *MsgDeleteComposition) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
