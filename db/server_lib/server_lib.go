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

func MakeJKZDBServer() (*JKZDBServer, error) {
	jkzdb, err := db.CreateJKZDB()
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
func (server *JKZDBServer) GetEntryById(ctx context.Context, in *pb.GetEntryByIdRequest) (*pb.GetEntryByIdResponse, error) {
	return nil, nil
}
func (server *JKZDBServer) GetEntryByField(ctx context.Context, in *pb.GetEntryByFieldRequest) (*pb.GetEntryByFieldResponse, error) {
	return nil, nil
}
