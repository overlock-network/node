package keeper

import (
	"encoding/binary"

	"cosmossdk.io/store/prefix"
	"github.com/cosmos/cosmos-sdk/runtime"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/web-seven/overlock-api/go/node/overlock/crossplane/v1beta1"
)

func (k Keeper) AppendFunction(ctx sdk.Context, function v1beta1.Function) uint64 {
	count := k.GetFunctionCount(ctx)
	function.Id = count
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, v1beta1.KeyPrefix(v1beta1.FunctionKey))
	appendedValue := k.cdc.MustMarshal(&function)
	store.Set(GetFunctionIDBytes(function.Id), appendedValue)
	k.SetFunctionCount(ctx, count+1)
	return count
}

func (k Keeper) GetFunctionCount(ctx sdk.Context) uint64 {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, []byte{})
	byteKey := v1beta1.KeyPrefix(v1beta1.FunctionCountKey)
	bz := store.Get(byteKey)
	if bz == nil {
		return 0
	}
	return binary.BigEndian.Uint64(bz)
}

func GetFunctionIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

func (k Keeper) SetFunctionCount(ctx sdk.Context, count uint64) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, []byte{})
	byteKey := v1beta1.KeyPrefix(v1beta1.FunctionCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

func (k Keeper) GetFunction(ctx sdk.Context, id uint64) (val v1beta1.Function, found bool) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, v1beta1.KeyPrefix(v1beta1.FunctionKey))
	b := store.Get(GetFunctionIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

func (k Keeper) SetFunction(ctx sdk.Context, function v1beta1.Function) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, v1beta1.KeyPrefix(v1beta1.FunctionKey))
	b := k.cdc.MustMarshal(&function)
	store.Set(GetFunctionIDBytes(function.Id), b)
}

func (k Keeper) RemoveFunction(ctx sdk.Context, id uint64) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, v1beta1.KeyPrefix(v1beta1.FunctionKey))
	store.Delete(GetFunctionIDBytes(id))
}
