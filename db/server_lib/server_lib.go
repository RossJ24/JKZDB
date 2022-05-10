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
	server.mx.Lock()
	val, err := server.jkzdb.GetValue(req.GetKey())
	if err != nil {
		return nil, err
	}
	// If the value is supposed to be unique, but it already exists on this shard, then it's a bad request.
	if len(val) != 0 && req.Unique {
		server.mx.Unlock()
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
		err := json.Unmarshal([]byte(val), &user)
		if err != nil {
			return nil, err
		}
		delta, err := strconv.ParseInt(req.Updates["transaction"], 10, 64)
		if err != nil {

			server.mx.Unlock()
			return nil, err
		}
		if user.Balance-delta < 0 {

			server.mx.Unlock()
			return nil, err
		}
	} else if _, exists := req.Updates["email"]; exists && len(req.Updates) == 1 {
		// In this case the user email is being updated
		user := &models.User{}
		err := json.Unmarshal([]byte(val), &user)
		if err != nil {

			server.mx.Unlock()
			return nil, err
		}
		if user.Email == req.Updates["email"] {

			server.mx.Unlock()
			return nil, status.Error(
				codes.AlreadyExists,
				"No change to email necessary",
			)
		}
	} else if _, exists := req.Updates["del"]; exists {
		// In this case a key is being deleted
		if len(val) == 0 {
			server.mx.Unlock()
			return nil, status.Errorf(
				codes.NotFound,
				"Key (%s) not found",
				req.Key,
			)
		}
	} else if req.Unique && len(req.Updates) == 5 {
		// In this case a user is being created, nothing else to check here. We already know the new key doesn't exist
	}
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
		err = json.Unmarshal([]byte(val), &user)
		if err != nil {
			return nil, err
		}
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
		err = json.Unmarshal([]byte(val), &user)
		if err != nil {
			return nil, err
		}
		user.Email = req.Updates["email"]
		newVal, err := json.Marshal(user)
		if err != nil {
			return nil, err
		}
		server.jkzdb.UpdateEntry(req.GetKey(), string(newVal))
	} else if _, exists := req.Updates["del"]; exists {
		// In this case a key is being deleted, nothing to check here
		server.jkzdb.DeleteKey(req.Key)
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
	if len(value) == 0 {
		return nil, status.Errorf(
			codes.NotFound,
			"Key(%s) was not found",
			req.Query,
		)
	}
	user := &models.User{}
	err = json.Unmarshal([]byte(value), user)
	if err != nil {
		return nil, err
	}
	if req.Field != nil {
		value = user.ToMap()[*req.Field]
	}
	if err != nil {
		return nil, err
	}
	resp := &pb.GetEntryResponse{
		Entry: value,
	}
	return resp, nil
}

func (server *JKZDBServer) SetEntryPrepareBatch(ctx context.Context, req *pb.SetEntryPrepareBatchRequest) (*pb.SetEntryPrepareBatchResponse, error) {
	// TODO: Implement, can lowkey just copy logic from above and loop, or abstract that^ logic and call a helper function and loop
	// Matthew: Ross, I basically copied the code above and re-used it in a loop, you should check it

	keys := req.GetKeys()
	updateMap := req.GetUpdates()
	server.mx.Lock()

	for i := 0; i < len(keys); i++ {
		val, err := server.jkzdb.GetValue(keys[i])
		if err != nil {
			return nil, err
		}

		unique := updateMap[i].GetUnique()
		updates := updateMap[i].GetUpdates()

		// If the value is supposed to be unique, but it already exists on this shard, then it's a bad request.
		if len(val) != 0 && unique {
			server.mx.Unlock()
			return nil, status.Errorf(
				codes.AlreadyExists,
				"Batched: Key already exists",
			)
		}
		// add another option for del
		if _, exists := updates["key"]; exists {
			// In this case it is just a key like email:mail@mail.com => 1 being added
		} else if _, exists := updates["transaction"]; exists {
			// In this case it is just the balance being updated
			user := &models.User{}
			err := json.Unmarshal([]byte(val), &user)
			if err != nil {
				return nil, err
			}
			delta, err := strconv.ParseInt(updates["transaction"], 10, 64)
			if err != nil {

				server.mx.Unlock()
				return nil, err
			}
			if user.Balance-delta < 0 {

				server.mx.Unlock()
				return nil, err
			}
		} else if _, exists := updates["email"]; exists && len(req.Updates) == 1 {
			// In this case the user email is being updated
			user := &models.User{}
			err := json.Unmarshal([]byte(val), &user)
			if err != nil {

				server.mx.Unlock()
				return nil, err
			}
			if user.Email == updates["email"] {

				server.mx.Unlock()
				return nil, status.Error(
					codes.AlreadyExists,
					"No change to email necessary",
				)
			}
		} else if _, exists := updates["del"]; exists {
			// In this case a key is being deleted
			if len(val) == 0 {
				server.mx.Unlock()
				return nil, status.Errorf(
					codes.NotFound,
					"Key (%s) not found",
					keys[i],
				)
			}
		} else if unique && len(req.Updates) == 5 {
			// In this case a user is being created, nothing else to check here. We already know the new key doesn't exist
		}

		atomic.StoreInt64(&server.currentUpdate, req.IdempotencyKey)
	}

	return &pb.SetEntryPrepareBatchResponse{}, nil
}

func (server *JKZDBServer) SetEntryCommitBatch(ctx context.Context, req *pb.SetEntryCommitBatchRequest) (*pb.SetEntryCommitBatchResponse, error) {
	// TODO: Implement, can lowkey just copy logic from above and loop, or abstract that^ logic and call a helper function and loop
	keys := req.GetKeys()
	updateMap := req.GetUpdates()

	for i := 0; i < len(keys); i++ {
		updates := updateMap[i].GetUpdates()

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
	
		if _, exists := updates["key"]; exists {
			server.jkzdb.UpdateEntry(keys[i], updates["key"])
		} else if _, exists := updates["transaction"]; exists {
			// In this case it is just the balance being updated
			user := &models.User{}
			val, err := server.jkzdb.GetValue(keys[i])
			if err != nil {
				return nil, err
			}
			err = json.Unmarshal([]byte(val), &user)
			if err != nil {
				return nil, err
			}
			delta, err := strconv.ParseInt(updates["transaction"], 10, 64)
			if err != nil {
				return nil, err
			}
			user.Balance -= delta
			newVal, err := json.Marshal(user)
			if err != nil {
				return nil, err
			}
			server.jkzdb.UpdateEntry(keys[i], string(newVal))
		} else if _, exists := updates["email"]; exists && len(req.Updates) == 1 {
			// In this case the user email is being updated
			user := &models.User{}
			val, err := server.jkzdb.GetValue(keys[i])
			if err != nil {
				return nil, err
			}
			err = json.Unmarshal([]byte(val), &user)
			if err != nil {
				return nil, err
			}
			user.Email = updates["email"]
			newVal, err := json.Marshal(user)
			if err != nil {
				return nil, err
			}
			server.jkzdb.UpdateEntry(keys[i], string(newVal))
		} else if _, exists := updates["del"]; exists {
			// In this case a key is being deleted, nothing to check here
			server.jkzdb.DeleteKey(keys[i])
		}
	}

	server.mx.Unlock()

	return &pb.SetEntryCommitBatchResponse{}, nil
}

func (server *JKZDBServer) SetEntryAbortBatch(ctx context.Context, req *pb.SetEntryAbortBatchRequest) (*pb.SetEntryAbortBatchResponse, error) {
	// TODO: Implement, can lowkey just copy logic from above and loop, or abstract that^ logic and call a helper function and loop

	// Matthew: Ross, I don't think we need a loop here? We just unlock the shard even if it's batched.
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

	return &pb.SetEntryAbortBatchResponse{}, nil
}
