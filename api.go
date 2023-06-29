package main

import (
	"errors"
	"sync"
)

/*
Handle with care
*/
var store = struct {
	sync.RWMutex
	m map[string]string
} {m: make(map[string]string) }

var ErrorNoSuchKey = errors.New("no such key")

func Put(k, v string) error {
	store.Lock()
	store.m[k] = v
	store.Unlock()

	return nil
}

func Get(k string) (string, error) {
	store.RLock()
	value, ok := store.m[k]
	store.RUnlock()

	if !ok {
		return "", ErrorNoSuchKey
	}

	return value, nil
}

func Delete(k string) error {
	store.Lock()
	delete(store.m, k)
	store.Unlock()

	return nil
}