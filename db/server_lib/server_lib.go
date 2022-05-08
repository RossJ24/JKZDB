package db

import (
	"JKZDB/db/db"
	pb "JKZDB/db/proto"
	"context"
	"sync"
)

type JKZDBServer struct {
	pb.UnimplementedJKZDBServer
	jkzdb *db.JKZDB
	mx    sync.RWMutex
}

func MakeJKZDBServer(port int) (*JKZDBServer, error) {
	jkzdb, err := db.CreateJKZDB(port)
	if err != nil {
		return nil, err
	}
	return &JKZDBServer{
		jkzdb: jkzdb,
	}, nil
}

func (server *JKZDBServer) SetEntryPrepare(ctx context.Context, in *pb.SetEntryPrepareRequest) (*pb.SetEntryPrepareResponse, error) {
	return nil, nil
}
func (server *JKZDBServer) SetEntryCommit(ctx context.Context, in *pb.SetEntryCommitRequest) (*pb.SetEntryCommitResponse, error) {
	return nil, nil
}
func (server *JKZDBServer) GetEntry(ctx context.Context, in *pb.GetEntryRequest) (*pb.GetEntryResponse, error) {
	server.mx.RLock()
	defer server.mx.RUnlock()
	value, err := server.jkzdb.GetValue(in.Query)
	if err != nil {
		return nil, err
	}
	entry := make(map[string]string)
	entry[in.Query] = value
	resp := &pb.GetEntryResponse{
		Entry: entry,
	}
	return resp, nil
}

func (server *JKZDBServer) GetEntryByIndexedField(ctx context.Context, in *pb.GetEntryByIndexedFieldRequest) (*pb.GetEntryByIndexedFieldResponse, error) {
	return nil, nil
}
