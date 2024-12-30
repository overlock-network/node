package types

const (
	// ModuleName defines the module name
	ModuleName = "storage"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_storage"
)

var (
	ParamsKey = []byte("p_storage")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
