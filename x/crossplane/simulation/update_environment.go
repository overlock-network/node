package simulation

import (
	"math/rand"

	"overlock/x/crossplane/keeper"

	"github.com/web-seven/overlock-api/go/node/overlock/crossplane/v1beta1"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
)

func SimulateMsgUpdateEnvironment(
	ak v1beta1.AccountKeeper,
	bk v1beta1.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &v1beta1.MsgUpdateEnvironment{
			Creator: simAccount.Address.String(),
		}

		// TODO: Handling the UpdateEnvironment simulation

		return simtypes.NoOpMsg(v1beta1.ModuleName, sdk.MsgTypeURL(msg), "UpdateEnvironment simulation not implemented"), nil, nil
	}
}
