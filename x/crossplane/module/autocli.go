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

				{
					RpcMethod:      "ShowEnvironment",
					Use:            "show-environment [id]",
					Short:          "Query show-environment",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "id"}},
				},

				{
					RpcMethod:      "ListEnvironment",
					Use:            "list-environment",
					Short:          "Query list-environment",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{},
				},
				{
					RpcMethod:      "ShowProvider",
					Use:            "show-provider [id]",
					Short:          "Query show-provider",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "id"}},
				},

				{
					RpcMethod:      "ListProvider",
					Use:            "list-provider",
					Short:          "Query list-provider",
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
				{
					RpcMethod:      "CreateEnvironment",
					Use:            "create-environment [metadata] [spec]",
					Short:          "Send a create-environment tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "metadata"}, {ProtoField: "spec"}},
				},
				{
					RpcMethod:      "UpdateEnvironment",
					Use:            "update-environment [id] [metadata] [spec]",
					Short:          "Send a update-environment tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "id"}, {ProtoField: "metadata"}, {ProtoField: "spec"}},
				},
				{
					RpcMethod:      "DeleteEnvironment",
					Use:            "delete-environment [id]",
					Short:          "Send a delete-environment tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "id"}},
				},
				{
					RpcMethod:      "CreateProvider",
					Use:            "create-provider [metadata] [spec]",
					Short:          "Send a create-provider tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "metadata"}, {ProtoField: "spec"}},
				},
				{
					RpcMethod:      "UpdateProvider",
					Use:            "update-provider [id] [metadata] [spec]",
					Short:          "Send a update-provider tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "id"}, {ProtoField: "metadata"}, {ProtoField: "spec"}},
				},
				{
					RpcMethod:      "DeleteProvider",
					Use:            "delete-provider [id]",
					Short:          "Send a delete-provider tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "id"}},
				},
				{
					RpcMethod:      "CreateFunction",
					Use:            "create-function [name]",
					Short:          "Send a create-function tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "name"}},
				},
				{
					RpcMethod:      "UpdateFunction",
					Use:            "update-function [id]",
					Short:          "Send a update-function tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "id"}},
				},
				{
					RpcMethod:      "DeleteFunction",
					Use:            "delete-function [id]",
					Short:          "Send a delete-function tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "id"}},
				},
				// this line is used by ignite scaffolding # autocli/tx
			},
		},
	}
}
