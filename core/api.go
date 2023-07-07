package core

import (
	"errors"
	"fmt"
	"log"
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
	store.Unlock()

	store.storage.WritePut(k, v)
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
	store.Unlock()
	
	store.storage.WriteDelete(k)
	return nil
}

func (store *KVStore) DeleteThrough(k string) error {
	store.Lock()
	delete(store.m, k)
	store.Unlock()

	return nil
}

func (store *KVStore) PutThrough(k, v string) error {
	store.Lock()
	store.m[k] = v
	store.Unlock()

	return nil
}

func (store *KVStore) Restore() error {
	fmt.Println("[Storage] Initializing Events Storage");
	// We'll reuse this
	var err error

	events, errors := store.storage.LoadEvents()
	// We'll reuse the same variables set
	event, isChannelOpen := KVStorageEvent{}, true

	fmt.Println("[Storage] Loading events from storage", isChannelOpen, err);
	for isChannelOpen && err == nil {
		select {
		// the <-channel syntax allows isChannelOpen to get 'false' when channel is closed
		// the consequent `case` here are not like those in a `switch`
		case err, isChannelOpen = <-errors:
		case event, isChannelOpen = <-events:
			switch event.EventType {
			case DeleteEvent:
				err = store.DeleteThrough(event.Key)
			case PutEvent:
				err = store.PutThrough(event.Key, event.Value)
			}
		}
	}

	fmt.Println("[Storage] Running storage");
	store.storage.Run()

	go func() {
		for err := range store.storage.Err() {
			log.Print(err)
		}
	}()

	return err
}
