package overlock

import (
	"math/rand"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"overlock/testutil/sample"
	crossplanesimulation "overlock/x/crossplane/simulation"

	"github.com/overlock-network/api/go/node/overlock/crossplane/v1beta1"
)

// avoid unused import issue
var (
	_ = crossplanesimulation.FindAccount
	_ = rand.Rand{}
	_ = sample.AccAddress
	_ = sdk.AccAddress{}
	_ = simulation.MsgEntryKind
)

const (
	opWeightMsgCreateConfiguration = "op_weight_msg_create_configuration"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateConfiguration int = 100

	opWeightMsgUpdateConfiguration = "op_weight_msg_update_configuration"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateConfiguration int = 100

	opWeightMsgDeleteConfiguration = "op_weight_msg_delete_configuration"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteConfiguration int = 100

	opWeightMsgCreateEnvironment = "op_weight_msg_create_environment"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateEnvironment int = 100

	opWeightMsgUpdateEnvironment = "op_weight_msg_update_environment"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateEnvironment int = 100

	opWeightMsgDeleteEnvironment = "op_weight_msg_delete_environment"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteEnvironment int = 100

	opWeightMsgCreateProvider = "op_weight_msg_create_provider"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateProvider int = 100

	opWeightMsgUpdateProvider = "op_weight_msg_update_provider"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateProvider int = 100

	opWeightMsgDeleteProvider = "op_weight_msg_delete_provider"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteProvider int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module.
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	overlockGenesis := v1beta1.GenesisState{
		Params: v1beta1.DefaultParams(),
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[v1beta1.ModuleName] = simState.Cdc.MustMarshalJSON(&overlockGenesis)
}

// RegisterStoreDecoder registers a decoder.
func (am AppModule) RegisterStoreDecoder(_ simtypes.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgCreateConfiguration int
	simState.AppParams.GetOrGenerate(opWeightMsgCreateConfiguration, &weightMsgCreateConfiguration, nil,
		func(_ *rand.Rand) {
			weightMsgCreateConfiguration = defaultWeightMsgCreateConfiguration
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateConfiguration,
		crossplanesimulation.SimulateMsgCreateConfiguration(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateConfiguration int
	simState.AppParams.GetOrGenerate(opWeightMsgUpdateConfiguration, &weightMsgUpdateConfiguration, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateConfiguration = defaultWeightMsgUpdateConfiguration
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateConfiguration,
		crossplanesimulation.SimulateMsgUpdateConfiguration(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeleteConfiguration int
	simState.AppParams.GetOrGenerate(opWeightMsgDeleteConfiguration, &weightMsgDeleteConfiguration, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteConfiguration = defaultWeightMsgDeleteConfiguration
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteConfiguration,
		crossplanesimulation.SimulateMsgDeleteConfiguration(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgCreateEnvironment int
	simState.AppParams.GetOrGenerate(opWeightMsgCreateEnvironment, &weightMsgCreateEnvironment, nil,
		func(_ *rand.Rand) {
			weightMsgCreateEnvironment = defaultWeightMsgCreateEnvironment
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateEnvironment,
		crossplanesimulation.SimulateMsgCreateEnvironment(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateEnvironment int
	simState.AppParams.GetOrGenerate(opWeightMsgUpdateEnvironment, &weightMsgUpdateEnvironment, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateEnvironment = defaultWeightMsgUpdateEnvironment
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateEnvironment,
		crossplanesimulation.SimulateMsgUpdateEnvironment(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeleteEnvironment int
	simState.AppParams.GetOrGenerate(opWeightMsgDeleteEnvironment, &weightMsgDeleteEnvironment, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteEnvironment = defaultWeightMsgDeleteEnvironment
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteEnvironment,
		crossplanesimulation.SimulateMsgDeleteEnvironment(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgCreateProvider int
	simState.AppParams.GetOrGenerate(opWeightMsgCreateProvider, &weightMsgCreateProvider, nil,
		func(_ *rand.Rand) {
			weightMsgCreateProvider = defaultWeightMsgCreateProvider
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateProvider,
		crossplanesimulation.SimulateMsgCreateProvider(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateProvider int
	simState.AppParams.GetOrGenerate(opWeightMsgUpdateProvider, &weightMsgUpdateProvider, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateProvider = defaultWeightMsgUpdateProvider
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateProvider,
		crossplanesimulation.SimulateMsgUpdateProvider(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeleteProvider int
	simState.AppParams.GetOrGenerate(opWeightMsgDeleteProvider, &weightMsgDeleteProvider, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteProvider = defaultWeightMsgDeleteProvider
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteProvider,
		crossplanesimulation.SimulateMsgDeleteProvider(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}

// ProposalMsgs returns msgs used for governance proposals for simulations.
func (am AppModule) ProposalMsgs(simState module.SimulationState) []simtypes.WeightedProposalMsg {
	return []simtypes.WeightedProposalMsg{
		simulation.NewWeightedProposalMsg(
			opWeightMsgCreateConfiguration,
			defaultWeightMsgCreateConfiguration,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				crossplanesimulation.SimulateMsgCreateConfiguration(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgUpdateConfiguration,
			defaultWeightMsgUpdateConfiguration,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				crossplanesimulation.SimulateMsgUpdateConfiguration(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgDeleteConfiguration,
			defaultWeightMsgDeleteConfiguration,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				crossplanesimulation.SimulateMsgDeleteConfiguration(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgCreateConfiguration,
			defaultWeightMsgCreateConfiguration,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				crossplanesimulation.SimulateMsgCreateConfiguration(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgUpdateConfiguration,
			defaultWeightMsgUpdateConfiguration,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				crossplanesimulation.SimulateMsgUpdateConfiguration(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgDeleteConfiguration,
			defaultWeightMsgDeleteConfiguration,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				crossplanesimulation.SimulateMsgDeleteConfiguration(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgCreateEnvironment,
			defaultWeightMsgCreateEnvironment,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				crossplanesimulation.SimulateMsgCreateEnvironment(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgUpdateEnvironment,
			defaultWeightMsgUpdateEnvironment,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				crossplanesimulation.SimulateMsgUpdateEnvironment(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgDeleteEnvironment,
			defaultWeightMsgDeleteEnvironment,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				crossplanesimulation.SimulateMsgDeleteEnvironment(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgCreateProvider,
			defaultWeightMsgCreateProvider,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				crossplanesimulation.SimulateMsgCreateProvider(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgUpdateProvider,
			defaultWeightMsgUpdateProvider,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				crossplanesimulation.SimulateMsgUpdateProvider(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgDeleteProvider,
			defaultWeightMsgDeleteProvider,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				crossplanesimulation.SimulateMsgDeleteProvider(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		// this line is used by starport scaffolding # simapp/module/OpMsg
	}
}
