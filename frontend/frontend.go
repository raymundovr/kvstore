package frontend

import (
	"log"

	"github.com/raymundovr/kvstore/core"
	"github.com/raymundovr/kvstore/frontend/grpc"
	"github.com/raymundovr/kvstore/frontend/rest"
)

func InitializeFrontend(serverType string, store *core.KVStore) {
	switch serverType {
	case "rest":
		rest.InitializeRest(store)
	case "grpc":
		grpc.InitializeGRPC(store)
	default:
		log.Fatal("not a valid frontend")
	}
}