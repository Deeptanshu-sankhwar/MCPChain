package types

const (
	// ModuleName defines the module name
	ModuleName = "mcpchain"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_mcpchain"
)

var (
	ParamsKey = []byte("p_mcpchain")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
