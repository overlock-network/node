package overlock

import (
	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"

	apiv1 "github.com/overlock-network/api/go/api/overlock/crossplane/v1beta1"
	nodev1 "github.com/overlock-network/api/go/node/overlock/crossplane/v1beta1"
)

// AutoCLIOptions implements the autocli.HasAutoCLIConfig interface.
func (am AppModule) AutoCLIOptions() *autocliv1.ModuleOptions {
	return &autocliv1.ModuleOptions{
		Query: &autocliv1.ServiceCommandDescriptor{
			Service: apiv1.Query_ServiceDesc.ServiceName,
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
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
			Service:              nodev1.Msg_serviceDesc.ServiceName,
			EnhanceCustomCommand: true, // only required if you want to use the custom command
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod:      "CreateEnvironment",
					Use:            "create-environment [metadata] [spec]",
					Short:          "Send a create-environment tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "metadata"}},
				},
				{
					RpcMethod:      "UpdateEnvironment",
					Use:            "update-environment [id] [metadata] [spec]",
					Short:          "Send a update-environment tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "id"}, {ProtoField: "metadata"}},
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
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "metadata"}},
				},
				{
					RpcMethod:      "UpdateProvider",
					Use:            "update-provider [id] [metadata] [spec]",
					Short:          "Send a update-provider tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "id"}, {ProtoField: "metadata"}},
				},
				{
					RpcMethod:      "DeleteProvider",
					Use:            "delete-provider [id]",
					Short:          "Send a delete-provider tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "id"}},
				},
				// this line is used by ignite scaffolding # autocli/tx
			},
		},
	}
}
