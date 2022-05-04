package server_lib

import (
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
	return nil, nil
}

func (coordinator *Coordinator) GetHandler(ctx *fiber.Ctx) error {
	key := ctx.Query("key")

}

func (coordinator *Coordinator) PutHandler(ctx *fiber.Ctx) error {
	key := ctx.Query("key")
}

func (coordinator *Coordinator) DeleteHandler(ctx *fiber.Ctx) error {
	key := ctx.Query("key")
}
