package server_lib

import (
	"encoding/json"
	"log"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var LOCALHOST string = "[::1]:"

type Coordinator struct {
	// ShardMap (Lowkey should be a tree for range-sharding)
	// Doing hash sharding for now
	ShardMapping []*grpc.ClientConn
	// Monotonically increasing, and ONLY atomically updated
	nextId int64
}

func NewCoordinator(shardConfigPath string) (*Coordinator, error) {
	configFile, err := os.ReadFile(shardConfigPath)
	if err != nil {
		log.Fatalf("Could not parse shard configuration file.")
	}
	var config map[string][]map[string]string
	json.Unmarshal(configFile, &config)
	connections := make([]*grpc.ClientConn, 0)
	for _, sh := range config["shards"] {
		conn, err := grpc.Dial(LOCALHOST+sh["port"], grpc.DialOption(grpc.WithTransportCredentials(insecure.NewCredentials())))
		if err != nil {
			return nil, err
		}
		connections = append(connections, conn)
	}
	return &Coordinator{
		connections,
		0,
	}, nil
}
