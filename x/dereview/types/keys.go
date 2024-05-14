package types

import "cosmossdk.io/collections"

const (
	// ModuleName defines the module name
	ModuleName = "dereview"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_dereview"
)

var (
	ParamsKey = collections.NewPrefix("p_dereview")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

var (
	PostKey      = collections.NewPrefix("post/value/")
	PostCountKey = collections.NewPrefix("post/count/")
)
