package keeper

import (
	"encoding/binary"

	"github.com/overlock-network/api/go/node/overlock/storage/v1beta1"

	"cosmossdk.io/store/prefix"
	"github.com/cosmos/cosmos-sdk/runtime"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) AppendRegistry(ctx sdk.Context, composition v1beta1.Registry) uint64 {
	count := k.GetRegistryCount(ctx)
	composition.Id = count
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, v1beta1.KeyPrefix(v1beta1.RegistryKey))
	appendedValue := k.cdc.MustMarshal(&composition)
	store.Set(GetRegistryIDBytes(composition.Id), appendedValue)
	k.SetRegistryCount(ctx, count+1)
	return count
}

func (k Keeper) GetRegistryCount(ctx sdk.Context) uint64 {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, []byte{})
	byteKey := v1beta1.KeyPrefix(v1beta1.RegistryCountKey)
	bz := store.Get(byteKey)
	if bz == nil {
		return 0
	}
	return binary.BigEndian.Uint64(bz)
}

func GetRegistryIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

func (k Keeper) SetRegistryCount(ctx sdk.Context, count uint64) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, []byte{})
	byteKey := v1beta1.KeyPrefix(v1beta1.RegistryCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

func (k Keeper) GetRegistry(ctx sdk.Context, id uint64) (val v1beta1.Registry, found bool) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, v1beta1.KeyPrefix(v1beta1.RegistryKey))
	b := store.Get(GetRegistryIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

func (k Keeper) SetRegistry(ctx sdk.Context, composition v1beta1.Registry) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, v1beta1.KeyPrefix(v1beta1.RegistryKey))
	b := k.cdc.MustMarshal(&composition)
	store.Set(GetRegistryIDBytes(composition.Id), b)
}

func (k Keeper) RemoveRegistry(ctx sdk.Context, id uint64) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, v1beta1.KeyPrefix(v1beta1.RegistryKey))
	store.Delete(GetRegistryIDBytes(id))
}
