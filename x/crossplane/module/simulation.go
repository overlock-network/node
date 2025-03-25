package overlock

import (
	"math/rand"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"overlock/testutil/sample"
	crossplanesimulation "overlock/x/crossplane/simulation"

	"github.com/web-seven/overlock-api/go/node/overlock/crossplane/v1beta1"
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
	opWeightMsgCreateComposition = "op_weight_msg_create_composition"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateComposition int = 100

	opWeightMsgUpdateComposition = "op_weight_msg_update_composition"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateComposition int = 100

	opWeightMsgDeleteComposition = "op_weight_msg_delete_composition"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteComposition int = 100

	opWeightMsgCreateXrd = "op_weight_msg_create_xrd"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateXrd int = 100

	opWeightMsgUpdateXrd = "op_weight_msg_update_xrd"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateXrd int = 100

	opWeightMsgDeleteXrd = "op_weight_msg_delete_xrd"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteXrd int = 100

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

	opWeightMsgCreateFunction = "op_weight_msg_create_function"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateFunction int = 100

	opWeightMsgUpdateFunction = "op_weight_msg_update_function"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateFunction int = 100

	opWeightMsgDeleteFunction = "op_weight_msg_delete_function"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteFunction int = 100

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

	var weightMsgCreateComposition int
	simState.AppParams.GetOrGenerate(opWeightMsgCreateComposition, &weightMsgCreateComposition, nil,
		func(_ *rand.Rand) {
			weightMsgCreateComposition = defaultWeightMsgCreateComposition
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateComposition,
		crossplanesimulation.SimulateMsgCreateComposition(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateComposition int
	simState.AppParams.GetOrGenerate(opWeightMsgUpdateComposition, &weightMsgUpdateComposition, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateComposition = defaultWeightMsgUpdateComposition
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateComposition,
		crossplanesimulation.SimulateMsgUpdateComposition(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeleteComposition int
	simState.AppParams.GetOrGenerate(opWeightMsgDeleteComposition, &weightMsgDeleteComposition, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteComposition = defaultWeightMsgDeleteComposition
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteComposition,
		crossplanesimulation.SimulateMsgDeleteComposition(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgCreateXrd int
	simState.AppParams.GetOrGenerate(opWeightMsgCreateXrd, &weightMsgCreateXrd, nil,
		func(_ *rand.Rand) {
			weightMsgCreateXrd = defaultWeightMsgCreateXrd
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateXrd,
		crossplanesimulation.SimulateMsgCreateXrd(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateXrd int
	simState.AppParams.GetOrGenerate(opWeightMsgUpdateXrd, &weightMsgUpdateXrd, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateXrd = defaultWeightMsgUpdateXrd
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateXrd,
		crossplanesimulation.SimulateMsgUpdateXrd(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeleteXrd int
	simState.AppParams.GetOrGenerate(opWeightMsgDeleteXrd, &weightMsgDeleteXrd, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteXrd = defaultWeightMsgDeleteXrd
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteXrd,
		crossplanesimulation.SimulateMsgDeleteXrd(am.accountKeeper, am.bankKeeper, am.keeper),
	))

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

	var weightMsgCreateFunction int
	simState.AppParams.GetOrGenerate(opWeightMsgCreateFunction, &weightMsgCreateFunction, nil,
		func(_ *rand.Rand) {
			weightMsgCreateFunction = defaultWeightMsgCreateFunction
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateFunction,
		crossplanesimulation.SimulateMsgCreateFunction(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateFunction int
	simState.AppParams.GetOrGenerate(opWeightMsgUpdateFunction, &weightMsgUpdateFunction, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateFunction = defaultWeightMsgUpdateFunction
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateFunction,
		crossplanesimulation.SimulateMsgUpdateFunction(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeleteFunction int
	simState.AppParams.GetOrGenerate(opWeightMsgDeleteFunction, &weightMsgDeleteFunction, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteFunction = defaultWeightMsgDeleteFunction
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteFunction,
		crossplanesimulation.SimulateMsgDeleteFunction(am.accountKeeper, am.bankKeeper, am.keeper),
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
			opWeightMsgCreateComposition,
			defaultWeightMsgCreateComposition,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				crossplanesimulation.SimulateMsgCreateComposition(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgUpdateComposition,
			defaultWeightMsgUpdateComposition,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				crossplanesimulation.SimulateMsgUpdateComposition(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgDeleteComposition,
			defaultWeightMsgDeleteComposition,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				crossplanesimulation.SimulateMsgDeleteComposition(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgCreateXrd,
			defaultWeightMsgCreateXrd,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				crossplanesimulation.SimulateMsgCreateXrd(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgUpdateXrd,
			defaultWeightMsgUpdateXrd,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				crossplanesimulation.SimulateMsgUpdateXrd(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgDeleteXrd,
			defaultWeightMsgDeleteXrd,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				crossplanesimulation.SimulateMsgDeleteXrd(am.accountKeeper, am.bankKeeper, am.keeper)
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
		simulation.NewWeightedProposalMsg(
			opWeightMsgCreateFunction,
			defaultWeightMsgCreateFunction,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				crossplanesimulation.SimulateMsgCreateFunction(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgUpdateFunction,
			defaultWeightMsgUpdateFunction,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				crossplanesimulation.SimulateMsgUpdateFunction(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgDeleteFunction,
			defaultWeightMsgDeleteFunction,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				crossplanesimulation.SimulateMsgDeleteFunction(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		// this line is used by starport scaffolding # simapp/module/OpMsg
	}
}
