package types

const (
	// ModuleName defines the module name
	ModuleName = "blocks"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_blocks"
)

var (
	ParamsKey = []byte("p_blocks")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
