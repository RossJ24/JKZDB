package server_lib

import (
	pb "JKZDB/db/proto"
	"context"
	"hash/fnv"
	"log"

	"google.golang.org/grpc"
)

func KeyHashFunc(key string) int {
	hash := fnv.New32()
	hash.Write([]byte(key))
	return int(hash.Sum32())
}

// Returns the client connection of the shard of the key
func (coordinator *Coordinator) ShardForKey(key string) *grpc.ClientConn {
	return coordinator.ShardMapping[KeyHashFunc(key)%len(coordinator.ShardMapping)]
}

// Sends Prepare RPC to given shard
// Returns true if shard successfully acquired necessary resources, false otherwise
func (coordinator *Coordinator) SendPrepareRPC(conn *grpc.ClientConn, key string, changes map[string]string, idempotencyKey int64, unique bool) bool {
	client := pb.NewJKZDBClient(conn)
	req := &pb.SetEntryPrepareRequest{
		IdempotencyKey: idempotencyKey,
		Key:            key,
		Updates:        changes,
		Unique:         unique,
	}
	_, err := client.SetEntryPrepare(context.TODO(), req)
	if err != nil {
		log.Printf("Prepare key(%s) failed: %v", idempotencyKey, err)
		return false
	}
	return true
}

func (coordinator *Coordinator) SendCommitRPC(conn *grpc.ClientConn, key string, changes map[string]string, idempotencyKey int64, unique bool) bool {
	committed := false
	client := pb.NewJKZDBClient(conn)
	req := &pb.SetEntryCommitRequest{
		IdempotencyKey: idempotencyKey,
		Key:            key,
		Updates:        changes,
		Unique:         false,
	}
	for !committed {
		_, err := client.SetEntryCommit(context.TODO(), req)
		if err != nil {
			log.Printf("Commit IdempotencyKey(%s) failed: %v", idempotencyKey, err)
		} else {
			committed = true
		}
	}
	return true
}

func (coordinator *Coordinator) SendAbortRPC(conn *grpc.ClientConn, idempotencyKey int64) bool {
	aborted := false
	client := pb.NewJKZDBClient(conn)
	req := &pb.SetEntryAbortRequest{
		IdempotencyKey: idempotencyKey,
		Unique:         false,
	}
	for !aborted {
		_, err := client.SetEntryAbort(context.TODO(), req)
		if err != nil {
			log.Printf("Abort IdempotencyKey(%s) failed: %v", idempotencyKey, err)
		} else {
			aborted = true
		}
	}
	return true
}

func (coordinator *Coordinator) SendPrepareBatchRPC(conn *grpc.ClientConn, keys []string, changes []map[string]string, idempotencyKey int64, unique []bool) bool {
	client := pb.NewJKZDBClient(conn)
	updates := make([]*pb.UpdateMap, 0)
	for i, m := range changes {
		updates = append(updates, &pb.UpdateMap{
			Updates: m,
			Unique:  unique[i],
		})
	}
	req := &pb.SetEntryPrepareBatchRequest{
		IdempotencyKey: idempotencyKey,
		Keys:           keys,
		Updates:        updates,
	}
	_, err := client.SetEntryPrepareBatch(context.TODO(), req)
	if err != nil {
		log.Printf("Prepare IdempotencyKey(%s) failed: %v", idempotencyKey, err)
		return false
	}
	return true
}

func (coordinator *Coordinator) SendCommitBatchRPC(conn *grpc.ClientConn, keys []string, changes []map[string]string, idempotencyKey int64, unique []bool) bool {
	committed := false
	client := pb.NewJKZDBClient(conn)
	updates := make([]*pb.UpdateMap, 0)
	for i, m := range changes {
		updates = append(updates, &pb.UpdateMap{
			Updates: m,
			Unique:  unique[i],
		})
	}
	req := &pb.SetEntryCommitBatchRequest{
		IdempotencyKey: idempotencyKey,
		Keys:           keys,
		Updates:        updates,
	}
	for !committed {
		_, err := client.SetEntryCommitBatch(context.TODO(), req)
		if err != nil {
			log.Printf("Commit IdempotencyKey(%s) failed: %v", idempotencyKey, err)
		} else {
			committed = true
		}
	}
	return true
}

func (coordinator *Coordinator) SendAbortBatchRPC(conn *grpc.ClientConn, idempotencyKey int64) bool {
	aborted := false
	client := pb.NewJKZDBClient(conn)
	req := &pb.SetEntryAbortBatchRequest{
		IdempotencyKey: idempotencyKey,
	}
	for !aborted {
		_, err := client.SetEntryAbortBatch(context.TODO(), req)
		if err != nil {
			log.Printf("Abort IdempotencyKey(%s) failed: %v", idempotencyKey, err)
		} else {
			aborted = true
		}
	}
	return true
}
