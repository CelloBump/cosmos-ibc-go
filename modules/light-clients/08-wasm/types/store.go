package types

import (
	"bytes"
	"errors"
	"io"
	"reflect"
	"strings"

	wasmvmtypes "github.com/CosmWasm/wasmvm/types"

	errorsmod "cosmossdk.io/errors"
	"cosmossdk.io/store/cachekv"
	"cosmossdk.io/store/listenkv"
	storeprefix "cosmossdk.io/store/prefix"
	"cosmossdk.io/store/tracekv"
	storetypes "cosmossdk.io/store/types"

	"github.com/cosmos/ibc-go/v8/modules/core/exported"
)

// updateProposalWrappedStore combines two KVStores into one while transparently routing the calls based on key prefix
type updateProposalWrappedStore struct {
	subjectStore    storetypes.KVStore
	substituteStore storetypes.KVStore

	subjectPrefix    []byte
	substitutePrefix []byte
}

func newUpdateProposalWrappedStore(subjectStore, substituteStore storetypes.KVStore, subjectPrefix, substitutePrefix []byte) updateProposalWrappedStore {
	return updateProposalWrappedStore{
		subjectStore:     subjectStore,
		substituteStore:  substituteStore,
		subjectPrefix:    subjectPrefix,
		substitutePrefix: substitutePrefix,
	}
}

func (ws updateProposalWrappedStore) Get(key []byte) []byte {
	return ws.getStore(key).Get(ws.trimPrefix(key))
}

func (ws updateProposalWrappedStore) Has(key []byte) bool {
	return ws.getStore(key).Has(ws.trimPrefix(key))
}

func (ws updateProposalWrappedStore) Set(key, value []byte) {
	ws.getStore(key).Set(ws.trimPrefix(key), value)
}

func (ws updateProposalWrappedStore) Delete(key []byte) {
	ws.getStore(key).Delete(ws.trimPrefix(key))
}

func (ws updateProposalWrappedStore) GetStoreType() storetypes.StoreType {
	return ws.subjectStore.GetStoreType()
}

func (ws updateProposalWrappedStore) Iterator(start, end []byte) storetypes.Iterator {
	return ws.getStore(start).Iterator(ws.trimPrefix(start), ws.trimPrefix(end))
}

func (ws updateProposalWrappedStore) ReverseIterator(start, end []byte) storetypes.Iterator {
	return ws.getStore(start).ReverseIterator(ws.trimPrefix(start), ws.trimPrefix(end))
}

func (ws updateProposalWrappedStore) CacheWrap() storetypes.CacheWrap {
	return cachekv.NewStore(ws)
}

func (ws updateProposalWrappedStore) CacheWrapWithTrace(w io.Writer, tc storetypes.TraceContext) storetypes.CacheWrap {
	return cachekv.NewStore(tracekv.NewStore(ws, w, tc))
}

func (ws updateProposalWrappedStore) CacheWrapWithListeners(storeKey storetypes.StoreKey, listeners *storetypes.MemoryListener) storetypes.CacheWrap {
	return cachekv.NewStore(listenkv.NewStore(ws, storeKey, listeners))
}

func (ws updateProposalWrappedStore) trimPrefix(key []byte) []byte {
	if bytes.HasPrefix(key, ws.subjectPrefix) {
		key = bytes.TrimPrefix(key, ws.subjectPrefix)
	} else {
		key = bytes.TrimPrefix(key, ws.substitutePrefix)
	}

	return key
}

func (ws updateProposalWrappedStore) getStore(key []byte) storetypes.KVStore {
	if bytes.HasPrefix(key, ws.subjectPrefix) {
		return ws.subjectStore
	}

	return ws.substituteStore
}

var _ wasmvmtypes.KVStore = &storeAdapter{}

// storeAdapter adapter to bridge SDK store impl to wasmvm
type storeAdapter struct {
	parent storetypes.KVStore
}

// newStoreAdapter constructor
func newStoreAdapter(s storetypes.KVStore) *storeAdapter {
	if s == nil {
		panic(errors.New("store must not be nil"))
	}
	return &storeAdapter{parent: s}
}

func (s storeAdapter) Get(key []byte) []byte {
	return s.parent.Get(key)
}

func (s storeAdapter) Set(key, value []byte) {
	s.parent.Set(key, value)
}

func (s storeAdapter) Delete(key []byte) {
	s.parent.Delete(key)
}

func (s storeAdapter) Iterator(start, end []byte) wasmvmtypes.Iterator {
	return s.parent.Iterator(start, end)
}

func (s storeAdapter) ReverseIterator(start, end []byte) wasmvmtypes.Iterator {
	return s.parent.ReverseIterator(start, end)
}

// getClientID extracts and validates the clientID from the clientStore's prefix.
//
// Due to the 02-client module not passing the clientID to the 08-wasm module,
// this function was devised to infer it from the store's prefix.
// The expected format of the clientStore prefix is "<placeholder>/{clientID}/".
// If the clientStore is of type updateProposalWrappedStore, the subjectStore's prefix is utilized instead.
func getClientID(clientStore storetypes.KVStore) (string, error) {
	upws, isUpdateProposalWrappedStore := clientStore.(updateProposalWrappedStore)
	if isUpdateProposalWrappedStore {
		// if the clientStore is a updateProposalWrappedStore, we retrieve the subjectStore
		// because the contract call will be made on the client with the ID of the subjectStore
		clientStore = upws.subjectStore
	}

	store, ok := clientStore.(storeprefix.Store)
	if !ok {
		return "", errorsmod.Wrapf(ErrRetrieveClientID, "clientStore is not a prefix store")
	}

	// using reflect to retrieve the private prefix field
	r := reflect.ValueOf(&store).Elem()

	f := r.FieldByName("prefix")
	if !f.IsValid() {
		return "", errorsmod.Wrapf(ErrRetrieveClientID, "prefix field not found")
	}

	prefix := string(f.Bytes())

	split := strings.Split(prefix, "/")
	if len(split) < 3 {
		return "", errorsmod.Wrapf(ErrRetrieveClientID, "prefix is not of the expected form")
	}

	// the clientID is the second to last element of the prefix
	// the prefix is expected to be of the form "<placeholder>/{clientID}/"
	clientID := split[len(split)-2]
	isClientID := strings.HasPrefix(clientID, exported.Wasm)
	if !isClientID {
		return "", errorsmod.Wrapf(ErrRetrieveClientID, "prefix does not contain a %s clientID", exported.Wasm)
	}

	return clientID, nil
}
