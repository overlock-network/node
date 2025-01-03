package types

import (
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
	// this line is used by starport scaffolding # 1
)

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateConfiguration{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgUpdateConfiguration{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgDeleteConfiguration{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateComposition{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgUpdateComposition{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgDeleteComposition{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateXrd{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgUpdateXrd{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgDeleteXrd{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateConfiguration{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgUpdateConfiguration{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgDeleteConfiguration{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateEnvironment{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgUpdateEnvironment{},
	)
	// this line is used by starport scaffolding # 3

	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgUpdateParams{},
	)
	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}
