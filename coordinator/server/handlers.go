package server_lib

import (
	pb "JKZDB/db/proto"
	"JKZDB/models"
	"context"
	"fmt"
	"sync/atomic"
	"time"

	fiber "github.com/gofiber/fiber/v2"
)

func concatindexkey(index string, key string) string {
	return fmt.Sprintf("%s:%s", index, key)
}

// Handles Get Requests to the
func (coordinator *Coordinator) ApiGetHandler(ctx *fiber.Ctx) error {
	index := ctx.Query("index", "id")
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
	return ctx.Status(fiber.StatusOK).JSON(res.GetEntry())
}

func (coordinator *Coordinator) ApiPutHandler(ctx *fiber.Ctx) error {
	id := ctx.Query("id")
	oldEmail := ctx.Query("old-email")
	newEmail := ctx.Query("new-email")
	query := concatindexkey(index, key)
	// TODO: Implement Updates
	return nil
}

func (coordinator *Coordinator) ApiDeleteHandler(ctx *fiber.Ctx) error {
	index := ctx.Query("index", "id")
	key := ctx.Query("key")
	query := concatindexkey(index, key)
	// TODO: Implement Deletes
	return nil
}

func (coordinator *Coordinator) TransactionPostHandler(ctx *fiber.Ctx) error {
	// Id of the recipient of the transaction
	to := ctx.Query("to")
	// Id of the sender of the transaction
	from := ctx.Query("from")
	amount := ctx.Query("amount")
	// TODO: Implement Transactions
	return nil
}

func (coordinator *Coordinator) CreateUserHandler(ctx *fiber.Ctx) error {
	user := &models.User{}
	if err := ctx.BodyParser(user); err != nil {
		return err
	}
	userId := atomic.AddInt64(&coordinator.nextId, 1)
	userMap := user.ToMap()
	emailKey := concatindexkey("email", userMap["email"])
	primaryKey := concatindexkey("id", fmt.Sprint(userId))
	emailShard := coordinator.ShardForKey(emailKey)
	primaryShard := coordinator.ShardForKey(primaryKey)
	idempotencyKey := time.Now().Unix()
	emailChanges := map[string]string{"key": fmt.Sprint(userId)}
	if emailShard == primaryShard {
		keys := []string{emailKey, primaryKey}
		changes := []map[string]string{emailChanges, userMap}
		unique := []bool{true, true}
		if coordinator.SendPrepareBatchRPC(primaryShard, keys, changes, idempotencyKey, unique) {
			coordinator.SendCommitBatchRPC(primaryShard, keys, changes, idempotencyKey, unique)
		} else {
			coordinator.SendAbortBatchRPC(primaryShard, idempotencyKey)
		}
	} else {
		emailRes := coordinator.SendPrepareRPC(emailShard, emailKey, emailChanges, idempotencyKey, true)
		primaryRes := coordinator.SendPrepareRPC(primaryShard, primaryKey, userMap, idempotencyKey, true)
		if emailRes && primaryRes {
			coordinator.SendCommitRPC(emailShard, emailKey, emailChanges, idempotencyKey, true)
			coordinator.SendCommitRPC(primaryShard, primaryKey, userMap, idempotencyKey, true)
		} else if emailRes && !primaryRes {
			coordinator.SendAbortRPC(emailShard, idempotencyKey)
			return ctx.SendStatus(fiber.StatusPreconditionFailed)
		} else {
			coordinator.SendAbortRPC(primaryShard, idempotencyKey)
			return ctx.SendStatus(fiber.StatusPreconditionFailed)
		}
	}

	return ctx.SendStatus(fiber.StatusCreated)
}
