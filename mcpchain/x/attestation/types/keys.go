package types

const (
	// ModuleName defines the module name
	ModuleName = "attestation"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_attestation"
)

var (
	ParamsKey = []byte("p_attestation")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
