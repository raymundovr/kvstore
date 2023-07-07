package main

import (
	"fmt"
	"os"

	"github.com/raymundovr/kvstore/grpc"
	// "github.com/raymundovr/kvstore/server"
	"github.com/raymundovr/kvstore/storage"
)

func main() {
	// Initialize events storage before server
	if err := storage.InitializeEventsStorage(); err != nil {
		fmt.Println("Cannot start service: %w", err)
		// At the moment we want to crash
		os.Exit(1)
	}

	// server.InitializeServer()
	grpc.RunServer()
}
