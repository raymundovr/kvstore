package grpc

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
)

func RunServer() {
	server := grpc.NewServer()
	RegisterKeyValueServer(server, &KVServer{})

	listener, err := net.Listen("tcp", ":8081")

	if err != nil {
		log.Fatalf("failed to start TCP listener: %v", err)
	}

	fmt.Println("Starting gRPC server")
	if err := server.Serve(listener); err != nil {
		log.Fatalf("failed to start gRPC server: %v", err)
	}
}
