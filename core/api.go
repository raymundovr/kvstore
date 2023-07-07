package core

import (
	"errors"
	"sync"
)

/*
Handle with care
*/

type KVStore struct {
	sync.RWMutex
	m       map[string]string
	storage KVStorage
}

var ErrorNoSuchKey = errors.New("no such key")

func NewKVStore(storage KVStorage) *KVStore {
	store := KVStore{
		m:       make(map[string]string),
		storage: storage,
	}

	return &store
}

func (store *KVStore) Put(k, v string) error {
	store.Lock()
	store.m[k] = v
	store.storage.WritePut(k, v)
	store.Unlock()

	return nil
}

func (store *KVStore) Get(k string) (string, error) {
	store.RLock()
	value, ok := store.m[k]
	store.RUnlock()

	if !ok {
		return "", ErrorNoSuchKey
	}

	return value, nil
}

func (store *KVStore) Delete(k string) error {
	store.Lock()
	delete(store.m, k)
	store.storage.WriteDelete(k)
	store.Unlock()

	return nil
}

func (store *KVStore) Restore() error {
	return nil
}
