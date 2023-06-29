package main

import (
	"errors"
	"sync"
)

/*
Handle with care
*/
var Store = struct {
	sync.RWMutex
	m map[string]string
} {m: make(map[string]string) }

var ErrorNoSuchKey = errors.New("no such key")

func Put(k, v string) error {
	Store.Lock()
	Store.m[k] = v
	Store.Unlock()

	return nil
}

func Get(k string) (string, error) {
	Store.RLock()
	value, ok := Store.m[k]
	Store.RUnlock()

	if !ok {
		return "", ErrorNoSuchKey
	}

	return value, nil
}

func Delete(k string) error {
	Store.Lock()
	delete(Store.m, k)
	Store.Unlock()

	return nil
}