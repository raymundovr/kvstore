package grpc

import (
	"context"

	"github.com/raymundovr/kvstore/core"
)

type KVServer struct {
	UnimplementedKeyValueServer
	store *core.KVStore
}

func (s KVServer) Get(c context.Context, req *GetRequest) (*GetResponse, error) {
	value, err := s.store.Get(req.Key)

	return &GetResponse{ Key: req.Key, Value: value }, err
}

func (s KVServer) Put(c context.Context, req *PutRequest) (*PutResponse, error) {
	err := s.store.Put(req.Key, req.Value)

	return &PutResponse{Key: req.Key, Value: req.Value}, err
}