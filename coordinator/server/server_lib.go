package server_lib

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	fiber "github.com/gofiber/fiber/v2"
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
	clients := make([]*grpc.ClientConn, 0)
	indexedFields := make(map[string]struct{})
	return &Coordinator{
		clients,
		indexedFields,
	}, nil
}

func concatindexkey(index string, key string) string {
	fmt.Sprintf("%s:%s", index, key)
}

func (coordinator *Coordinator) GetHandler(ctx *fiber.Ctx) error {
	index := ctx.Query("index", "id")
	key := ctx.Query("key")
	query := concatindexkey(index, key)
	return nil
}

func (coordinator *Coordinator) PutHandler(ctx *fiber.Ctx) error {
	index := ctx.Query("index", "id")
	key := ctx.Query("key")
	query := concatindexkey(index, key)
	return nil
}

func (coordinator *Coordinator) DeleteHandler(ctx *fiber.Ctx) error {
	index := ctx.Query("index", "id")
	key := ctx.Query("key")
	query := concatindexkey(index, key)
	return nil
}

func (coordinator *Coordinator) PostHandler(ctx *fiber.Ctx) error {
	newIndex := ctx.Query("new-index")
	coordinator.IndexedFields[newIndex] = struct{}{}
	return nil
}
