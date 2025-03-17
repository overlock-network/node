package keeper

import (
	"encoding/binary"

	"cosmossdk.io/store/prefix"
	"github.com/cosmos/cosmos-sdk/runtime"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/web-seven/overlock-api/go/node/overlock/crossplane/v1beta1"
)

func (k Keeper) AppendComposition(ctx sdk.Context, composition v1beta1.Composition) uint64 {
	count := k.GetCompositionCount(ctx)
	composition.Id = count
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, v1beta1.KeyPrefix(v1beta1.CompositionKey))
	appendedValue := k.cdc.MustMarshal(&composition)
	store.Set(GetCompositionIDBytes(composition.Id), appendedValue)
	k.SetCompositionCount(ctx, count+1)
	return count
}

func (k Keeper) GetCompositionCount(ctx sdk.Context) uint64 {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, []byte{})
	byteKey := v1beta1.KeyPrefix(v1beta1.CompositionCountKey)
	bz := store.Get(byteKey)
	if bz == nil {
		return 0
	}
	return binary.BigEndian.Uint64(bz)
}

func GetCompositionIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

func (k Keeper) SetCompositionCount(ctx sdk.Context, count uint64) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, []byte{})
	byteKey := v1beta1.KeyPrefix(v1beta1.CompositionCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

func (k Keeper) GetComposition(ctx sdk.Context, id uint64) (val v1beta1.Composition, found bool) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, v1beta1.KeyPrefix(v1beta1.CompositionKey))
	b := store.Get(GetCompositionIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

func (k Keeper) SetComposition(ctx sdk.Context, composition v1beta1.Composition) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, v1beta1.KeyPrefix(v1beta1.CompositionKey))
	b := k.cdc.MustMarshal(&composition)
	store.Set(GetCompositionIDBytes(composition.Id), b)
}

func (k Keeper) RemoveComposition(ctx sdk.Context, id uint64) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, v1beta1.KeyPrefix(v1beta1.CompositionKey))
	store.Delete(GetCompositionIDBytes(id))
}
