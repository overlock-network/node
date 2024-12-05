package types

const (
	// ModuleName defines the module name
	ModuleName = "overlock"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_overlock"
)

var (
	ParamsKey = []byte("p_overlock")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
