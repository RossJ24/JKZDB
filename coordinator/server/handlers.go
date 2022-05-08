package server_lib

import (
	pb "JKZDB/db/proto"
	"JKZDB/models"
	"context"
	"errors"
	"fmt"
	"sync/atomic"

	fiber "github.com/gofiber/fiber/v2"
)

func concatindexkey(index string, key string) string {
	return fmt.Sprintf("%s:%s", index, key)
}

// Handles Get Requests to the
func (coordinator *Coordinator) ApiGetHandler(ctx *fiber.Ctx) error {
	index := ctx.Query("index", "id")
	if !coordinator.IsIndexedField(index) {
		return errors.New("")
	}
	key := ctx.Query("key")
	field := ctx.Query("field")
	query := concatindexkey(index, key)
	if index != "id" {
		conn := coordinator.ShardForKey(query)
		client := pb.NewJKZDBClient(conn)
		req := &pb.GetEntryRequest{
			Query: query,
			Field: nil,
		}
		res, err := client.GetEntry(context.Background(), req)
		if err != nil {
			return err
		}
		query = concatindexkey("id", res.GetEntry()["id"])
	}
	conn := coordinator.ShardForKey(query)
	client := pb.NewJKZDBClient(conn)
	req := &pb.GetEntryRequest{
		Query: query,
		Field: &field,
	}
	res, err := client.GetEntry(context.Background(), req)
	if err != nil {
		return err
	}
	return ctx.Status(200).JSON(res.GetEntry())
}

func (coordinator *Coordinator) ApiPutHandler(ctx *fiber.Ctx) error {
	index := ctx.Query("index", "id")
	key := ctx.Query("key")
	query := concatindexkey(index, key)

	return nil
}

func (coordinator *Coordinator) ApiDeleteHandler(ctx *fiber.Ctx) error {
	index := ctx.Query("index", "id")
	key := ctx.Query("key")
	query := concatindexkey(index, key)
	return nil
}

func (coordinator *Coordinator) IndexPostHandler(ctx *fiber.Ctx) error {
	newIndex := ctx.Query("new-index", "")
	if len(newIndex) == 0 {
		return errors.New("No new index name supplied")
	}
	coordinator.IndexedFields[newIndex] = struct{}{}
	return nil
}

func (coordinator *Coordinator) TransactionPostHandler(ctx *fiber.Ctx) error {
	to := ctx.Query("to")
	from := ctx.Query("from")
	return nil
}

func (coordinator *Coordinator) CreateUserHandler(ctx *fiber.Ctx) error {
	user := &models.User{}
	if err := ctx.BodyParser(user); err != nil {
		return err
	}

	// How are we gonna map indexed fields to the same node

	//

	// TODO: check indexed fields first to make sure that they are unique
	userId := atomic.AddInt64(&coordinator.nextId, 1)
	return nil
}
