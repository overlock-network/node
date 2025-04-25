package keeper

import (
	"encoding/binary"

	"cosmossdk.io/store/prefix"
	"github.com/cosmos/cosmos-sdk/runtime"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/overlock-network/api/go/node/overlock/crossplane/v1beta1"
)

func (k Keeper) AppendEnvironment(ctx sdk.Context, environment v1beta1.Environment) uint64 {
	count := k.GetEnvironmentCount(ctx)
	environment.Id = count
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, v1beta1.KeyPrefix(v1beta1.EnvironmentKey))
	appendedValue := k.cdc.MustMarshal(&environment)
	store.Set(GetEnvironmentIDBytes(environment.Id), appendedValue)
	k.SetEnvironmentCount(ctx, count+1)
	return count
}

func (k Keeper) GetEnvironmentCount(ctx sdk.Context) uint64 {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, []byte{})
	byteKey := v1beta1.KeyPrefix(v1beta1.EnvironmentCountKey)
	bz := store.Get(byteKey)
	if bz == nil {
		return 0
	}
	return binary.BigEndian.Uint64(bz)
}

func GetEnvironmentIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

func (k Keeper) SetEnvironmentCount(ctx sdk.Context, count uint64) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, []byte{})
	byteKey := v1beta1.KeyPrefix(v1beta1.EnvironmentCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

func (k Keeper) GetEnvironment(ctx sdk.Context, id uint64) (val v1beta1.Environment, found bool) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, v1beta1.KeyPrefix(v1beta1.EnvironmentKey))
	b := store.Get(GetEnvironmentIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

func (k Keeper) SetEnvironment(ctx sdk.Context, environment v1beta1.Environment) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, v1beta1.KeyPrefix(v1beta1.EnvironmentKey))
	b := k.cdc.MustMarshal(&environment)
	store.Set(GetEnvironmentIDBytes(environment.Id), b)
}

func (k Keeper) RemoveEnvironment(ctx sdk.Context, id uint64) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, v1beta1.KeyPrefix(v1beta1.EnvironmentKey))
	store.Delete(GetEnvironmentIDBytes(id))
}
