package server_lib

import (
	"fmt"

	fiber "github.com/gofiber/fiber/v2"
)

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

func (Coordinator *Coordinator) TransactionHandler(ctx *fiber.Ctx) error {

}
