package server_lib

import (
	"hash/fnv"

	"google.golang.org/grpc"
)

func KeyHashFunc(key string) int {
	hash := fnv.New32()
	hash.Write([]byte(key))
	return int(hash.Sum32())
}

// Returns whether or not the field is indexed
// If it is not indexed, we don't need to have a second lookup for the field
func (coordinator *Coordinator) IsIndexedField(field string) bool {
	_, exists := coordinator.IndexedFields[field]
	return exists
}

// Returns the client connection of the shard of the key
func (coordinator *Coordinator) ShardForKey(key string) *grpc.ClientConn {
	return coordinator.ShardMapping[KeyHashFunc(key)]
}

// Sends Prepare RPC to given shard
// Returns true if shard successfully acquired necessary resources, false otherwise
func (coordinator *Coordinator) SendPrepareRPC(conn *grpc.ClientConn, changes map[string]string, idempotencyKey string) bool {
	//TODO: Create client from the connection, and send RPC, wait for response, and return bool
}

//
func (coordinator *Coordinator) SendCommitRPC(conn *grpc.ClientConn, changes map[string]string, idempotencyKey string) bool {
	//TODO: Create client from the connection, and send the RPC to Commit, returns true if successful, retry if not
}

func (coordinator *Coordinator) SendAbortRPC(conn *grpc.ClientConn, idempotencyKey string) bool {
	// TODO: Create client from the connection, and send the RPC to Abort the RPC, returns true if successful, retry if not
}
