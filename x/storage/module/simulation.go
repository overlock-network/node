package storage

import (
	"math/rand"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"overlock/testutil/sample"
	storagesimulation "overlock/x/storage/simulation"

	"github.com/overlock-network/api/go/node/overlock/storage/v1beta1"
)

// avoid unused import issue
var (
	_ = storagesimulation.FindAccount
	_ = rand.Rand{}
	_ = sample.AccAddress
	_ = sdk.AccAddress{}
	_ = simulation.MsgEntryKind
)

const (
	opWeightMsgCreateRegistry = "op_weight_msg_create_registry"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateRegistry int = 100

	opWeightMsgUpdateRegistry = "op_weight_msg_update_registry"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateRegistry int = 100

	opWeightMsgDeleteRegistry = "op_weight_msg_delete_registry"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteRegistry int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module.
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	storageGenesis := v1beta1.GenesisState{
		Params: v1beta1.DefaultParams(),
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[v1beta1.ModuleName] = simState.Cdc.MustMarshalJSON(&storageGenesis)
}

// RegisterStoreDecoder registers a decoder.
func (am AppModule) RegisterStoreDecoder(_ simtypes.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgCreateRegistry int
	simState.AppParams.GetOrGenerate(opWeightMsgCreateRegistry, &weightMsgCreateRegistry, nil,
		func(_ *rand.Rand) {
			weightMsgCreateRegistry = defaultWeightMsgCreateRegistry
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateRegistry,
		storagesimulation.SimulateMsgCreateRegistry(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateRegistry int
	simState.AppParams.GetOrGenerate(opWeightMsgUpdateRegistry, &weightMsgUpdateRegistry, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateRegistry = defaultWeightMsgUpdateRegistry
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateRegistry,
		storagesimulation.SimulateMsgUpdateRegistry(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeleteRegistry int
	simState.AppParams.GetOrGenerate(opWeightMsgDeleteRegistry, &weightMsgDeleteRegistry, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteRegistry = defaultWeightMsgDeleteRegistry
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteRegistry,
		storagesimulation.SimulateMsgDeleteRegistry(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}

// ProposalMsgs returns msgs used for governance proposals for simulations.
func (am AppModule) ProposalMsgs(simState module.SimulationState) []simtypes.WeightedProposalMsg {
	return []simtypes.WeightedProposalMsg{
		simulation.NewWeightedProposalMsg(
			opWeightMsgCreateRegistry,
			defaultWeightMsgCreateRegistry,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				storagesimulation.SimulateMsgCreateRegistry(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgUpdateRegistry,
			defaultWeightMsgUpdateRegistry,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				storagesimulation.SimulateMsgUpdateRegistry(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgDeleteRegistry,
			defaultWeightMsgDeleteRegistry,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				storagesimulation.SimulateMsgDeleteRegistry(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		// this line is used by starport scaffolding # simapp/module/OpMsg
	}
}
