package db

import (
	"JKZDB/db/db"
	pb "JKZDB/db/proto"
	"JKZDB/models"
	"context"
	"encoding/json"
	"strconv"
	"sync"
	"sync/atomic"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type JKZDBServer struct {
	pb.UnimplementedJKZDBServer
	jkzdb         *db.JKZDB
	mx            sync.RWMutex
	currentUpdate int64
	emailDel      atomic.Value
}

func MakeJKZDBServer(port int) (*JKZDBServer, error) {
	jkzdb, err := db.CreateJKZDB(port)
	if err != nil {
		return nil, err
	}
	server := &JKZDBServer{
		jkzdb:         jkzdb,
		currentUpdate: 0,
	}
	server.emailDel.Store(nil)
	return server, nil
}

func (server *JKZDBServer) SetEntryPrepare(ctx context.Context, req *pb.SetEntryPrepareRequest) (*pb.SetEntryPrepareResponse, error) {
	// TODO: Might need to check again after the Write lock has been acquired?
	server.mx.RLock()
	val, err := server.jkzdb.GetValue(req.GetKey())
	if err != nil {
		return nil, err
	}
	if len(val) != 0 && req.Unique {
		return nil, status.Errorf(
			codes.AlreadyExists,
			"Key already exists",
		)
	}
	// add another option for del
	if _, exists := req.Updates["key"]; exists {
		// In this case it is just a key like email:mail@mail.com => 1 being added
	} else if _, exists := req.Updates["transaction"]; exists {
		// In this case it is just the balance being updated
		user := &models.User{}
		json.Unmarshal([]byte(val), &user)
		delta, err := strconv.ParseInt(req.Updates["transaction"], 10, 64)
		if err != nil {
			return nil, err
		}
		if user.Balance-delta < 0 {
			return nil, err
		}
	} else if _, exists := req.Updates["email"]; exists && len(req.Updates) == 1 {
		// In this case the user email is being updated
		user := &models.User{}
		json.Unmarshal([]byte(val), &user)
		if user.Email == req.Updates["email"] {
			return nil, status.Error(
				codes.AlreadyExists,
				"No change to email necessary",
			)
		}
	} else if _, exists := req.Updates["del"]; exists {
		// In this case a key is being deleted, nothing to check here. The check for existence is already done.
	}
	server.mx.RUnlock()
	server.mx.Lock()
	atomic.StoreInt64(&server.currentUpdate, req.IdempotencyKey)
	return &pb.SetEntryPrepareResponse{}, nil
}

func (server *JKZDBServer) SetEntryCommit(ctx context.Context, req *pb.SetEntryCommitRequest) (*pb.SetEntryCommitResponse, error) {
	prepareIdempotencyKey := atomic.LoadInt64(&server.currentUpdate)
	if prepareIdempotencyKey != req.IdempotencyKey {
		return nil, status.Errorf(
			codes.FailedPrecondition,
			"Required Idempotency Key: %d does not match the one provided: %d",
			prepareIdempotencyKey,
			req.IdempotencyKey,
		)
	}
	prepareEmailDel := server.emailDel.Load().(*string)
	if prepareEmailDel != nil {
		server.jkzdb.DeleteKey(*prepareEmailDel)
	}

	if _, exists := req.Updates["key"]; exists {
		server.jkzdb.UpdateEntry(req.GetKey(), req.Updates["key"])
	} else if _, exists := req.Updates["transaction"]; exists {
		// In this case it is just the balance being updated
		user := &models.User{}
		val, err := server.jkzdb.GetValue(req.GetKey())
		if err != nil {
			return nil, err
		}
		json.Unmarshal([]byte(val), &user)
		delta, err := strconv.ParseInt(req.Updates["transaction"], 10, 64)
		if err != nil {
			return nil, err
		}
		user.Balance -= delta
		newVal, err := json.Marshal(user)
		if err != nil {
			return nil, err
		}
		server.jkzdb.UpdateEntry(req.GetKey(), string(newVal))
	} else if _, exists := req.Updates["email"]; exists && len(req.Updates) == 1 {
		// In this case the user email is being updated
		user := &models.User{}
		val, err := server.jkzdb.GetValue(req.GetKey())
		if err != nil {
			return nil, err
		}
		json.Unmarshal([]byte(val), &user)
		user.Email = req.Updates["email"]
		newVal, err := json.Marshal(user)
		if err != nil {
			return nil, err
		}
		server.jkzdb.UpdateEntry(req.GetKey(), string(newVal))
	} else if _, exists := req.Updates["del"]; exists {
		// In this case a key is being deleted, nothing to check here
		server.jkzdb.DeleteKey(req.Updates["del"])
	}
	server.mx.Unlock()
	return &pb.SetEntryCommitResponse{}, nil
}

// Ensures that the Idempotency key is equal, and releases resources if they are
func (server *JKZDBServer) SetEntryAbort(ctx context.Context, req *pb.SetEntryAbortRequest) (*pb.SetEntryAbortResponse, error) {
	prepareIdempotencyKey := atomic.LoadInt64(&server.currentUpdate)
	if prepareIdempotencyKey != req.IdempotencyKey {
		return nil, status.Errorf(
			codes.FailedPrecondition,
			"Required Idempotency Key: %d does not match the one provided: %d",
			prepareIdempotencyKey,
			req.IdempotencyKey,
		)
	}
	server.emailDel.Store(nil)
	atomic.StoreInt64(&server.currentUpdate, 0)
	server.mx.Unlock()
	return &pb.SetEntryAbortResponse{}, nil
}

func (server *JKZDBServer) GetEntry(ctx context.Context, req *pb.GetEntryRequest) (*pb.GetEntryResponse, error) {
	server.mx.RLock()
	defer server.mx.RUnlock()
	value, err := server.jkzdb.GetValue(req.Query)
	if err != nil {
		return nil, err
	}
	entry := make(map[string]string)
	entry[req.Query] = value
	resp := &pb.GetEntryResponse{
		Entry: entry,
	}
	return resp, nil
}
