package grpc

import (
	"fmt"
	"log"
	"net"

	"github.com/raymundovr/kvstore/core"
	"google.golang.org/grpc"
)

func InitializeGRPC(store *core.KVStore) {
	server := grpc.NewServer()
	RegisterKeyValueServer(server, &KVServer{ store: store })

	listener, err := net.Listen("tcp", ":8081")

	if err != nil {
		log.Fatalf("failed to start TCP listener: %v", err)
	}

	fmt.Println("Starting gRPC server")
	if err := server.Serve(listener); err != nil {
		log.Fatalf("failed to start gRPC server: %v", err)
	}
}
