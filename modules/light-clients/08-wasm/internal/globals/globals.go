package globals

import (
	"errors"

	wasmvm "github.com/CosmWasm/wasmvm"

	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
)

var (
	WasmVM *wasmvm.VM
	// storeKeyMap stores the storeKey for the 08-wasm module. Using a single global storetypes.StoreKey fails in the context
	// of tests with multiple test chains utilized. As such, we utilize a workaround involving a mapping from the chains codec
	// to the storeKey which can be used to store a key per test chain.
	// This is required as a global so that the KV store can be retrieved in the ClientState Initialize function which doesn't
	// have access to the keeper. The storeKey is used to check the code hash of the contract and determine if the light client
	// is allowed to be instantiated.
	storeKeyMap = make(map[codec.BinaryCodec]storetypes.StoreKey)
)

// SetWasmStoreKey sets the store key for the 08-wasm module keyed by the chain's codec.
func SetWasmStoreKey(key codec.BinaryCodec, storeKey storetypes.StoreKey) {
	storeKeyMap[key] = storeKey
}

// GetWasmStoreKey returns the store key for the 08-wasm module keyed by the chain's codec.
func GetWasmStoreKey(key codec.BinaryCodec) storetypes.StoreKey {
	if storeKey, ok := storeKeyMap[key]; ok {
		return storeKey
	}
	panic(errors.New("store key not set"))
}
