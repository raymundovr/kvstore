package main

import (
	"fmt"
	"os"

	"github.com/raymundovr/kvstore/core"
	"github.com/raymundovr/kvstore/frontend"
	"github.com/raymundovr/kvstore/storage"
)

func main() {
	// Initialize events storage before server
	kvStorage, err := storage.NewKVStorage("logs")
	if err != nil {
		fmt.Println("Cannot start service: %w", err)
		// At the moment we want to crash
		os.Exit(1)
	}

	kv := core.NewKVStore(kvStorage)

	kv.Restore()

	frontend.InitializeFrontend("grpc", kv)
}
