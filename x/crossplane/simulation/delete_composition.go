package simulation

import (
	"math/rand"

	"overlock/x/crossplane/keeper"
	"overlock/x/crossplane/types"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
)

func SimulateMsgDeleteComposition(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgDeleteComposition{
			Creator: simAccount.Address.String(),
		}

		// TODO: Handling the DeleteComposition simulation

		return simtypes.NoOpMsg(types.ModuleName, sdk.MsgTypeURL(msg), "DeleteComposition simulation not implemented"), nil, nil
	}
}
