package server_lib

import (
	"encoding/json"
	"log"
	"os"

	"google.golang.org/grpc"
)

type Coordinator struct {
	// ShardMap (Lowkey should be a tree for range-sharding)
	// Doing hash sharding for now
	ShardMapping  []*grpc.ClientConn
	IndexedFields map[string]struct{}
}

func NewCoordinator(shardConfigPath string) (*Coordinator, error) {
	configFile, err := os.ReadFile(shardConfigPath)
	if err != nil {
		log.Fatalf("Could not parse shard configuration file.")
	}
	var config map[string][]map[string]string
	json.Unmarshal(configFile, &config)
	// TODO: actually fill in clients
	clients := make([]*grpc.ClientConn, 0)
	indexedFields := make(map[string]struct{})
	return &Coordinator{
		clients,
		indexedFields,
	}, nil
}
