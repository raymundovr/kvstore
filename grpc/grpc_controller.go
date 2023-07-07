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

func (s KVServer) Put(c context.Context, req *PutRequest) (*PutResponse, error) {
	err := core.Put(req.Key, req.Value)

	return &PutResponse{Key: req.Key, Value: req.Value}, err
}