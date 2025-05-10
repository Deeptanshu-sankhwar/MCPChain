package toolregistry

import (
	"math/rand"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"mcpchain/testutil/sample"
	toolregistrysimulation "mcpchain/x/toolregistry/simulation"
	"mcpchain/x/toolregistry/types"
)

// avoid unused import issue
var (
	_ = toolregistrysimulation.FindAccount
	_ = rand.Rand{}
	_ = sample.AccAddress
	_ = sdk.AccAddress{}
	_ = simulation.MsgEntryKind
)

const (
	opWeightMsgRegisterServer = "op_weight_msg_register_server"
	// TODO: Determine the simulation weight value
	defaultWeightMsgRegisterServer int = 100

	opWeightMsgUpdateServerDetails = "op_weight_msg_update_server_details"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateServerDetails int = 100

	opWeightMsgRemoveServer = "op_weight_msg_remove_server"
	// TODO: Determine the simulation weight value
	defaultWeightMsgRemoveServer int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module.
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	toolregistryGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&toolregistryGenesis)
}

// RegisterStoreDecoder registers a decoder.
func (am AppModule) RegisterStoreDecoder(_ simtypes.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgRegisterServer int
	simState.AppParams.GetOrGenerate(opWeightMsgRegisterServer, &weightMsgRegisterServer, nil,
		func(_ *rand.Rand) {
			weightMsgRegisterServer = defaultWeightMsgRegisterServer
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgRegisterServer,
		toolregistrysimulation.SimulateMsgRegisterServer(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateServerDetails int
	simState.AppParams.GetOrGenerate(opWeightMsgUpdateServerDetails, &weightMsgUpdateServerDetails, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateServerDetails = defaultWeightMsgUpdateServerDetails
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateServerDetails,
		toolregistrysimulation.SimulateMsgUpdateServerDetails(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgRemoveServer int
	simState.AppParams.GetOrGenerate(opWeightMsgRemoveServer, &weightMsgRemoveServer, nil,
		func(_ *rand.Rand) {
			weightMsgRemoveServer = defaultWeightMsgRemoveServer
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgRemoveServer,
		toolregistrysimulation.SimulateMsgRemoveServer(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}

// ProposalMsgs returns msgs used for governance proposals for simulations.
func (am AppModule) ProposalMsgs(simState module.SimulationState) []simtypes.WeightedProposalMsg {
	return []simtypes.WeightedProposalMsg{
		simulation.NewWeightedProposalMsg(
			opWeightMsgRegisterServer,
			defaultWeightMsgRegisterServer,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				toolregistrysimulation.SimulateMsgRegisterServer(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgUpdateServerDetails,
			defaultWeightMsgUpdateServerDetails,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				toolregistrysimulation.SimulateMsgUpdateServerDetails(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgRemoveServer,
			defaultWeightMsgRemoveServer,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				toolregistrysimulation.SimulateMsgRemoveServer(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		// this line is used by starport scaffolding # simapp/module/OpMsg
	}
}
