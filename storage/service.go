package storage

import (
	"fmt"
	kv "github.com/raymundovr/kvstore/core"
)

// One pointer to share among the package
var ServiceStorage *KVLogStorage

func NewKVStorage(kind string) (kv.KVStorage, error) {
	switch kind {
	case "logs":
		return NewKVLogStorage("transactions.log")
	default:
		return nil, fmt.Errorf("not a valid storage")
	}
}

