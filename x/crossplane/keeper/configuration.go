package keeper

import (
	"encoding/binary"

	"cosmossdk.io/store/prefix"
	"github.com/cosmos/cosmos-sdk/runtime"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/web-seven/overlock-api/go/node/overlock/crossplane/v1beta1"
)

func (k Keeper) AppendConfiguration(ctx sdk.Context, configuration v1beta1.Configuration) uint64 {
	count := k.GetConfigurationCount(ctx)
	configuration.Id = count
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, v1beta1.KeyPrefix(v1beta1.ConfigurationKey))
	appendedValue := k.cdc.MustMarshal(&configuration)
	store.Set(GetConfigurationIDBytes(configuration.Id), appendedValue)
	k.SetConfigurationCount(ctx, count+1)
	return count
}

func (k Keeper) GetConfigurationCount(ctx sdk.Context) uint64 {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, []byte{})
	byteKey := v1beta1.KeyPrefix(v1beta1.ConfigurationCountKey)
	bz := store.Get(byteKey)
	if bz == nil {
		return 0
	}
	return binary.BigEndian.Uint64(bz)
}

func GetConfigurationIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

func (k Keeper) SetConfigurationCount(ctx sdk.Context, count uint64) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, []byte{})
	byteKey := v1beta1.KeyPrefix(v1beta1.ConfigurationCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

func (k Keeper) GetConfiguration(ctx sdk.Context, id uint64) (val v1beta1.Configuration, found bool) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, v1beta1.KeyPrefix(v1beta1.ConfigurationKey))
	b := store.Get(GetConfigurationIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

func (k Keeper) SetConfiguration(ctx sdk.Context, configuration v1beta1.Configuration) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, v1beta1.KeyPrefix(v1beta1.ConfigurationKey))
	b := k.cdc.MustMarshal(&configuration)
	store.Set(GetConfigurationIDBytes(configuration.Id), b)
}

func (k Keeper) RemoveConfiguration(ctx sdk.Context, id uint64) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, v1beta1.KeyPrefix(v1beta1.ConfigurationKey))
	store.Delete(GetConfigurationIDBytes(id))
}
