package overlock

import (
	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"

	modulev1 "overlock/api/overlock/crossplane"
)

// AutoCLIOptions implements the autocli.HasAutoCLIConfig interface.
func (am AppModule) AutoCLIOptions() *autocliv1.ModuleOptions {
	return &autocliv1.ModuleOptions{
		Query: &autocliv1.ServiceCommandDescriptor{
			Service: modulev1.Query_ServiceDesc.ServiceName,
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "Params",
					Use:       "params",
					Short:     "Shows the parameters of the module",
				},
				{
					RpcMethod:      "ShowConfiguration",
					Use:            "show-configuration [id]",
					Short:          "Query show-configuration",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "id"}},
				},

				{
					RpcMethod:      "ListConfiguration",
					Use:            "list-configuration",
					Short:          "Query list-configuration",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{},
				},

				{
					RpcMethod:      "ListComposition",
					Use:            "list-composition",
					Short:          "Query list-composition",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{},
				},

				{
					RpcMethod:      "ShowComposition",
					Use:            "show-composition [id]",
					Short:          "Query show-composition",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "id"}},
				},

				{
					RpcMethod:      "ShowComposition",
					Use:            "show-composition [id]",
					Short:          "Query show-composition",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "id"}},
				},

				{
					RpcMethod:      "ShowComposition",
					Use:            "show-composition [id]",
					Short:          "Query show-composition",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "id"}},
				},

				{
					RpcMethod:      "ShowComposition",
					Use:            "show-composition [id]",
					Short:          "Query show-composition",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "id"}},
				},

				{
					RpcMethod:      "ShowComposition",
					Use:            "show-composition [id]",
					Short:          "Query show-composition",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "id"}},
				},

				{
					RpcMethod:      "ShowXrd",
					Use:            "show-xrd [id]",
					Short:          "Query show-xrd",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "id"}},
				},

				{
					RpcMethod:      "ShowXrd",
					Use:            "show-xrd [id]",
					Short:          "Query show-xrd",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "id"}},
				},

				{
					RpcMethod:      "ShowXrd",
					Use:            "show-xrd [id]",
					Short:          "Query show-xrd",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "id"}},
				},

				{
					RpcMethod:      "ShowXrd",
					Use:            "show-xrd [id]",
					Short:          "Query show-xrd",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "id"}},
				},

				{
					RpcMethod:      "ShowXrd",
					Use:            "show-xrd [id]",
					Short:          "Query show-xrd",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "id"}},
				},

				{
					RpcMethod:      "ListXrd",
					Use:            "list-xrd",
					Short:          "Query list-xrd",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{},
				},

				// this line is used by ignite scaffolding # autocli/query
			},
		},
		Tx: &autocliv1.ServiceCommandDescriptor{
			Service:              modulev1.Msg_ServiceDesc.ServiceName,
			EnhanceCustomCommand: true, // only required if you want to use the custom command
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "UpdateParams",
					Skip:      true, // skipped because authority gated
				},
				{
					RpcMethod:      "CreateComposition",
					Use:            "create-composition [metadata] [spec]",
					Short:          "Send a create-composition tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "metadata"}, {ProtoField: "spec"}},
				},
				{
					RpcMethod:      "UpdateComposition",
					Use:            "update-composition [id] [metadata] [spec]",
					Short:          "Send a update-composition tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "id"}, {ProtoField: "metadata"}, {ProtoField: "spec"}},
				},
				{
					RpcMethod:      "DeleteComposition",
					Use:            "delete-composition [id]",
					Short:          "Send a delete-composition tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "id"}},
				},
				{
					RpcMethod:      "CreateXrd",
					Use:            "create-xrd [metadata] [spec]",
					Short:          "Send a create-xrd tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "metadata"}, {ProtoField: "spec"}},
				},
				{
					RpcMethod:      "UpdateXrd",
					Use:            "update-xrd [id] [metadata] [spec]",
					Short:          "Send a update-xrd tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "id"}, {ProtoField: "metadata"}, {ProtoField: "spec"}},
				},
				{
					RpcMethod:      "DeleteXrd",
					Use:            "delete-xrd [id]",
					Short:          "Send a delete-xrd tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "id"}},
				},
				{
					RpcMethod:      "CreateConfiguration",
					Use:            "create-configuration [metadata] [spec]",
					Short:          "Send a create-configuration tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "metadata"}, {ProtoField: "spec"}},
				},
				{
					RpcMethod:      "UpdateConfiguration",
					Use:            "update-configuration [id] [metadata] [spec]",
					Short:          "Send a update-configuration tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "id"}, {ProtoField: "metadata"}, {ProtoField: "spec"}},
				},
				{
					RpcMethod:      "DeleteConfiguration",
					Use:            "delete-configuration [id]",
					Short:          "Send a delete-configuration tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "id"}},
				},
				// this line is used by ignite scaffolding # autocli/tx
			},
		},
	}
}
