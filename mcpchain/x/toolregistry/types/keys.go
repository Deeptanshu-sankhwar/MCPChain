package types

const (
	// ModuleName defines the module name
	ModuleName = "toolregistry"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_toolregistry"
)

var (
	ParamsKey = []byte("p_toolregistry")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
