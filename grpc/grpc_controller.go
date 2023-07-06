package grpc

import (
	"context"

	"github.com/raymundovr/kvstore/core"
)

type KVServer struct {
	UnimplementedKeyValueServer
}

func (s KVServer) Get(c context.Context, req *GetRequest) (*GetResponse, error) {
	value, err := core.Get(req.Key)

	return &GetResponse{ Key: req.Key, Value: value }, err
}