package server_lib

import (
	pb "JKZDB/db/proto"
	"JKZDB/models"
	"context"
	"fmt"
	"strconv"
	"sync/atomic"
	"time"

	fiber "github.com/gofiber/fiber/v2"
	"google.golang.org/grpc"
)

func CreateQuery(index string, key string) string {
	return fmt.Sprintf("%s:%s", index, key)
}

// Handles Get Requests to the
func (coordinator *Coordinator) ApiGetHandler(ctx *fiber.Ctx) error {
	index := ctx.Query("index", "id")
	key := ctx.Query("key")
	field := ctx.Query("field")
	query := CreateQuery(index, key)
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
		query = CreateQuery("id", res.GetEntry())
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
	oldEmailQuery := CreateQuery("email", oldEmail)
	newEmailQuery := CreateQuery("email", newEmail)
	idQuery := CreateQuery("id", id)
	// TODO: Write Logic for updating an email. This can be a 3, 2 or 1 shard update.
	return nil
}

func (coordinator *Coordinator) ApiDeleteHandler(ctx *fiber.Ctx) error {
	index := ctx.Query("index", "id")
	key := ctx.Query("key")
	var primaryQuery string
	var primaryShard *grpc.ClientConn
	var emailQuery string
	var emailShard *grpc.ClientConn
	if index != "email" {
		// Gioven Primary ID as Key to Delete, now we need the email too
		primaryQuery = CreateQuery(index, key)
		field := "email"
		primaryShard = coordinator.ShardForKey(primaryQuery)
		res, err := coordinator.SendGetRPC(primaryShard, primaryQuery, &field)
		if err != nil {
			return err
		}
		emailQuery = CreateQuery("email", res.Entry)
		emailShard = coordinator.ShardForKey(emailQuery)
	} else {
		emailQuery = CreateQuery(index, key)
		emailShard = coordinator.ShardForKey(emailQuery)
		res, err := coordinator.SendGetRPC(emailShard, emailQuery, nil)
		if err != nil {
			return err
		}
		primaryQuery = CreateQuery("id", res.Entry)
		primaryShard = coordinator.ShardForKey(primaryQuery)
	}
	idempotencyKey := time.Now().Unix()
	delMap := map[string]string{"del": ""}
	if primaryShard == emailShard {
		res := coordinator.SendPrepareBatchRPC(
			emailShard,
			[]string{
				primaryQuery,
				emailQuery,
			},
			[]map[string]string{
				delMap,
				delMap,
			},
			idempotencyKey,
			[]bool{false, false},
		)
		if res {
			coordinator.SendCommitBatchRPC(emailShard,
				[]string{
					primaryQuery,
					emailQuery,
				},
				[]map[string]string{
					delMap,
					delMap,
				},
				idempotencyKey,
				[]bool{false, false},
			)
		} else {
			coordinator.SendAbortBatchRPC(emailShard, idempotencyKey)
		}
	} else {
		fromRes := coordinator.SendPrepareRPC(
			primaryShard,
			primaryQuery,
			delMap,
			idempotencyKey,
			false,
		)
		toRes := coordinator.SendPrepareRPC(
			emailShard,
			emailQuery,
			delMap,
			idempotencyKey,
			false,
		)
		if fromRes && toRes {
			coordinator.SendCommitRPC(
				primaryShard,
				primaryQuery,
				delMap,
				idempotencyKey,
				false,
			)
			coordinator.SendCommitRPC(
				emailShard,
				emailQuery,
				delMap,
				idempotencyKey,
				false,
			)
		} else if fromRes && !toRes {
			coordinator.SendAbortBatchRPC(primaryShard, idempotencyKey)
		} else if toRes && !fromRes {
			coordinator.SendAbortBatchRPC(emailShard, idempotencyKey)
		}
	}
	return nil
}

func (coordinator *Coordinator) TransactionPostHandler(ctx *fiber.Ctx) error {
	// Id of the recipient of the transaction
	to := ctx.Query("to")
	// Id of the sender of the transaction
	from := ctx.Query("from")
	amount := ctx.Query("amount")
	fromQuery := CreateQuery("id", from)
	toQuery := CreateQuery("id", to)
	amt, err := strconv.ParseInt(amount, 10, 64)
	if err != nil {
		return err
	}
	fromShard := coordinator.ShardForKey(fromQuery)
	fromChanges := map[string]string{
		"transaction": fmt.Sprint(-amt),
	}
	toChanges := map[string]string{
		"transaction": fmt.Sprint(amt),
	}
	toShard := coordinator.ShardForKey(toQuery)
	idempotencyKey := time.Now().Unix()
	if fromShard == toShard {
		res := coordinator.SendPrepareBatchRPC(
			toShard,
			[]string{
				fromQuery,
				toQuery,
			},
			[]map[string]string{
				fromChanges,
				toChanges,
			},
			idempotencyKey,
			[]bool{false, false},
		)
		if res {
			coordinator.SendCommitBatchRPC(toShard,
				[]string{
					fromQuery,
					toQuery,
				},
				[]map[string]string{
					fromChanges,
					toChanges,
				},
				idempotencyKey,
				[]bool{false, false},
			)
		} else {
			coordinator.SendAbortBatchRPC(toShard, idempotencyKey)
		}
	} else {
		fromRes := coordinator.SendPrepareRPC(
			fromShard,
			fromQuery,
			fromChanges,
			idempotencyKey,
			false,
		)
		toRes := coordinator.SendPrepareRPC(
			toShard,
			toQuery,
			toChanges,
			idempotencyKey,
			false,
		)
		if fromRes && toRes {
			coordinator.SendCommitRPC(
				fromShard,
				fromQuery,
				fromChanges,
				idempotencyKey,
				false,
			)
			coordinator.SendCommitRPC(
				toShard,
				toQuery,
				toChanges,
				idempotencyKey,
				false,
			)
		} else if fromRes && !toRes {
			coordinator.SendAbortBatchRPC(fromShard, idempotencyKey)
		} else if toRes && !fromRes {
			coordinator.SendAbortBatchRPC(toShard, idempotencyKey)
		}
	}
	return nil
}

func (coordinator *Coordinator) CreateUserHandler(ctx *fiber.Ctx) error {
	user := &models.User{}
	if err := ctx.BodyParser(user); err != nil {
		return err
	}
	userId := atomic.AddInt64(&coordinator.nextId, 1)
	userMap := user.ToMap()
	emailKey := CreateQuery("email", userMap["email"])
	primaryKey := CreateQuery("id", fmt.Sprint(userId))
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
