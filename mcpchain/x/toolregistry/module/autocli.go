package toolregistry

import (
	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"

	modulev1 "mcpchain/api/mcpchain/toolregistry"
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
					RpcMethod:      "ServerById",
					Use:            "server-by-id [id]",
					Short:          "Query server-by-id",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "id"}},
				},

				{
					RpcMethod:      "ServersByOwner",
					Use:            "servers-by-owner [owner-address]",
					Short:          "Query servers-by-owner",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "ownerAddress"}},
				},

				{
					RpcMethod:      "AllServers",
					Use:            "all-servers",
					Short:          "Query all-servers",
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
					RpcMethod:      "RegisterServer",
					Use:            "register-server [owner-address] [description] [endpoint-url] [mcp-schema-hash] [staked-amount]",
					Short:          "Send a register-server tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "ownerAddress"}, {ProtoField: "description"}, {ProtoField: "endpointUrl"}, {ProtoField: "mcpSchemaHash"}, {ProtoField: "stakedAmount"}},
				},
				{
					RpcMethod:      "UpdateServerDetails",
					Use:            "update-server-details [id] [description] [endpoint-url]",
					Short:          "Send a update-server-details tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "id"}, {ProtoField: "description"}, {ProtoField: "endpointUrl"}},
				},
				{
					RpcMethod:      "RemoveServer",
					Use:            "remove-server [id]",
					Short:          "Send a remove-server tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "id"}},
				},
				// this line is used by ignite scaffolding # autocli/tx
			},
		},
	}
}
