package attestation

import (
	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"

	modulev1 "mcpchain/api/mcpchain/attestation"
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
					RpcMethod:      "AttestationById",
					Use:            "attestation-by-id [id]",
					Short:          "Query a single attestation by ID",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "id"}},
				},
				{
					RpcMethod:      "AttestationsByAgent",
					Use:            "attestations-by-agent [agent-id]",
					Short:          "List all attestations by an agent",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "agent_id"}},
				},
				{
					RpcMethod:      "AttestationsByTool",
					Use:            "attestations-by-tool [tool-id]",
					Short:          "List all attestations for a tool",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "tool_id"}},
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
					RpcMethod: "LogAttestation",
					Use:       "log-attestation [agent-id] [tool-id] [tool-name] [request-hash] [response-hash]",
					Short:     "Send a log-attestation tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "agent_id"},
						{ProtoField: "tool_id"},
						{ProtoField: "tool_name"},
						{ProtoField: "request_hash"},
						{ProtoField: "response_hash"},
					},
				},
				// this line is used by ignite scaffolding # autocli/tx
			},
		},
	}
}
