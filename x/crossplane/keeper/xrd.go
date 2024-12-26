package keeper

import (
	"encoding/binary"

	"cosmossdk.io/store/prefix"
	"github.com/cosmos/cosmos-sdk/runtime"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"overlock/x/crossplane/types"
)

func (k Keeper) AppendCompositeResourceDefinition(ctx sdk.Context, xrd types.CompositeResourceDefinition) uint64 {
	count := k.GetCompositeResourceDefinitionCount(ctx)
	xrd.Id = count
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.CompositeResourceDefinitionKey))
	appendedValue := k.cdc.MustMarshal(&xrd)
	store.Set(GetCompositeResourceDefinitionIDBytes(xrd.Id), appendedValue)
	k.SetCompositeResourceDefinitionCount(ctx, count+1)
	return count
}

func (k Keeper) GetCompositeResourceDefinitionCount(ctx sdk.Context) uint64 {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, []byte{})
	byteKey := types.KeyPrefix(types.CompositeResourceDefinitionCountKey)
	bz := store.Get(byteKey)
	if bz == nil {
		return 0
	}
	return binary.BigEndian.Uint64(bz)
}

func GetCompositeResourceDefinitionIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

func (k Keeper) SetCompositeResourceDefinitionCount(ctx sdk.Context, count uint64) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, []byte{})
	byteKey := types.KeyPrefix(types.CompositeResourceDefinitionCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

func (k Keeper) GetCompositeResourceDefinition(ctx sdk.Context, id uint64) (val types.CompositeResourceDefinition, found bool) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.CompositeResourceDefinitionKey))
	b := store.Get(GetCompositeResourceDefinitionIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

func (k Keeper) SetCompositeResourceDefinition(ctx sdk.Context, xrd types.CompositeResourceDefinition) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.CompositeResourceDefinitionKey))
	b := k.cdc.MustMarshal(&xrd)
	store.Set(GetCompositeResourceDefinitionIDBytes(xrd.Id), b)
}

func (k Keeper) RemoveCompositeResourceDefinition(ctx sdk.Context, id uint64) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.CompositeResourceDefinitionKey))
	store.Delete(GetCompositeResourceDefinitionIDBytes(id))
}
