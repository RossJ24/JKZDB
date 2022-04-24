package server_lib

import (
	fiber "github.com/gofiber/fiber/v2"
)

type Coordinator struct {
	// Should map to grpc clients/connections (Lowkey might need to be a tree)
	ShardMap      map[string]int
	IndexedFields map[string]int
}

func NewCoordinator(shardConfigPath string) (*Coordinator, error) {
	return nil, nil
}

func (coordinator *Coordinator) GetHandler(ctx *fiber.Ctx) error {
	key := ctx.Query("key")

}

func (coordinator *Coordinator) PostHandler(ctx *fiber.Ctx) error {
	key := ctx.Query("key")
}

func (coordinator *Coordinator) DeleteHandler(ctx *fiber.Ctx) error {
	key := ctx.Query("key")
}
