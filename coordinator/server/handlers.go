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

// Creates a formatted query
func CreateQuery(index string, key string) string {
	return fmt.Sprintf("%s:%s", index, key)
}

// Handles GET API requests
func (coordinator *Coordinator) GetUserHandler(ctx *fiber.Ctx) error {
	index := ctx.Query("index", "id")
	key := ctx.Query("key")
	field := ctx.Query("field")
	query := CreateQuery(index, key)

	// find primary key id if email index is used instead
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

	// Retrieve the user account entry
	conn := coordinator.ShardForKey(query)
	client := pb.NewJKZDBClient(conn)
	var req *pb.GetEntryRequest
	if len(field) != 0 {
		req = &pb.GetEntryRequest{
			Query: query,
			Field: &field,
		}
	} else {
		req = &pb.GetEntryRequest{
			Query: query,
			Field: nil,
		}
	}
	res, err := client.GetEntry(context.Background(), req)
	if err != nil {
		return err
	}
	return ctx.Status(fiber.StatusOK).JSON(res.GetEntry())
}

// Handles PUT API requests for updating a user's email
func (coordinator *Coordinator) EmailUpdateHandler(ctx *fiber.Ctx) error {
	oldEmail := ctx.Query("old-email")
	newEmail := ctx.Query("new-email")
	oldEmailQuery := CreateQuery("email", oldEmail)
	newEmailQuery := CreateQuery("email", newEmail)

	// Retrieve the account's primary key id associated with the current/old email
	conn := coordinator.ShardForKey(oldEmailQuery)
	client := pb.NewJKZDBClient(conn)
	req := &pb.GetEntryRequest{
		Query: oldEmailQuery,
		Field: nil,
	}
	res, err := client.GetEntry(context.Background(), req)
	if err != nil {
		return err
	}
	id := res.GetEntry()
	idQuery := CreateQuery("id", id)

	// Get relevant shards and prepare all the necessary changes (delete the old email, update key mapping, update new email)
	primaryShard := coordinator.ShardForKey(idQuery)
	oldEmailShard := coordinator.ShardForKey(oldEmailQuery)
	newEmailShard := coordinator.ShardForKey(newEmailQuery)

	oldEmailMap := map[string]string{"del": ""}
	emailChanges := map[string]string{"email": newEmail}
	newEmailMap := map[string]string{"key": fmt.Sprint(id)}

	idempotencyKey := time.Now().Unix()

	// 1 shard update
	if primaryShard == oldEmailShard && oldEmailShard == newEmailShard {
		keys := []string{idQuery, oldEmailQuery, newEmailQuery}
		changes := []map[string]string{emailChanges, oldEmailMap, newEmailMap}
		unique := []bool{false, false, true}

		if coordinator.SendPrepareBatchRPC(primaryShard, keys, changes, idempotencyKey, unique) {
			coordinator.SendCommitBatchRPC(primaryShard, keys, changes, idempotencyKey, unique)
		} else {
			coordinator.SendAbortBatchRPC(primaryShard, idempotencyKey)
			return ctx.SendStatus(fiber.StatusPreconditionFailed)
		}
	} else if primaryShard == oldEmailShard {
		// 2 shard update: primary and oldEmail
		keys := []string{idQuery, oldEmailQuery}
		changes := []map[string]string{emailChanges, oldEmailMap}
		unique := []bool{false, false}

		batchRes := coordinator.SendPrepareBatchRPC(primaryShard, keys, changes, idempotencyKey, unique)
		newEmailRes := coordinator.SendPrepareRPC(newEmailShard, newEmailQuery, newEmailMap, idempotencyKey, true)

		if batchRes && newEmailRes {
			coordinator.SendCommitBatchRPC(primaryShard, keys, changes, idempotencyKey, unique)
			coordinator.SendCommitRPC(newEmailShard, newEmailQuery, newEmailMap, idempotencyKey, true)
		} else if batchRes {
			coordinator.SendAbortRPC(newEmailShard, idempotencyKey)
			return ctx.SendStatus(fiber.StatusPreconditionFailed)
		} else {
			coordinator.SendAbortBatchRPC(primaryShard, idempotencyKey)
			return ctx.SendStatus(fiber.StatusPreconditionFailed)
		}
	} else if primaryShard == newEmailShard {
		// 2 shard update: primary and newEmail
		keys := []string{idQuery, newEmailQuery}
		changes := []map[string]string{emailChanges, newEmailMap}
		unique := []bool{false, true}

		batchRes := coordinator.SendPrepareBatchRPC(primaryShard, keys, changes, idempotencyKey, unique)
		oldEmailRes := coordinator.SendPrepareRPC(oldEmailShard, oldEmailQuery, oldEmailMap, idempotencyKey, false)

		if batchRes && oldEmailRes {
			coordinator.SendCommitBatchRPC(primaryShard, keys, changes, idempotencyKey, unique)
			coordinator.SendCommitRPC(oldEmailShard, oldEmailQuery, oldEmailMap, idempotencyKey, false)
		} else if batchRes {
			coordinator.SendAbortRPC(oldEmailShard, idempotencyKey)
			return ctx.SendStatus(fiber.StatusPreconditionFailed)
		} else {
			coordinator.SendAbortBatchRPC(primaryShard, idempotencyKey)
			return ctx.SendStatus(fiber.StatusPreconditionFailed)
		}
	} else if oldEmailShard == newEmailShard {
		// 2 shard update: oldEmail and newEmail
		keys := []string{oldEmailQuery, newEmailQuery}
		changes := []map[string]string{oldEmailMap, newEmailMap}
		unique := []bool{false, true}

		batchRes := coordinator.SendPrepareBatchRPC(oldEmailShard, keys, changes, idempotencyKey, unique)
		primaryRes := coordinator.SendPrepareRPC(primaryShard, idQuery, emailChanges, idempotencyKey, false)

		if batchRes && primaryRes {
			coordinator.SendCommitBatchRPC(oldEmailShard, keys, changes, idempotencyKey, unique)
			coordinator.SendCommitRPC(primaryShard, idQuery, emailChanges, idempotencyKey, false)
		} else if batchRes {
			coordinator.SendAbortRPC(primaryShard, idempotencyKey)
			return ctx.SendStatus(fiber.StatusPreconditionFailed)
		} else {
			coordinator.SendAbortBatchRPC(oldEmailShard, idempotencyKey)
			return ctx.SendStatus(fiber.StatusPreconditionFailed)
		}
	} else {
		// 3 shard update
		primaryRes := coordinator.SendPrepareRPC(primaryShard, idQuery, emailChanges, idempotencyKey, false)
		oldEmailRes := coordinator.SendPrepareRPC(oldEmailShard, oldEmailQuery, oldEmailMap, idempotencyKey, false)
		newEmailRes := coordinator.SendPrepareRPC(newEmailShard, newEmailQuery, newEmailMap, idempotencyKey, true)

		if primaryRes && oldEmailRes && newEmailRes {
			coordinator.SendCommitRPC(primaryShard, idQuery, emailChanges, idempotencyKey, false)
			coordinator.SendCommitRPC(oldEmailShard, oldEmailQuery, oldEmailMap, idempotencyKey, false)
			coordinator.SendCommitRPC(newEmailShard, newEmailQuery, newEmailMap, idempotencyKey, true)
		} else if primaryRes && oldEmailRes {
			coordinator.SendAbortRPC(newEmailShard, idempotencyKey)
			return ctx.SendStatus(fiber.StatusPreconditionFailed)
		} else if primaryRes && newEmailRes {
			coordinator.SendAbortRPC(oldEmailShard, idempotencyKey)
			return ctx.SendStatus(fiber.StatusPreconditionFailed)
		} else if oldEmailRes && newEmailRes {
			coordinator.SendAbortRPC(primaryShard, idempotencyKey)
			return ctx.SendStatus(fiber.StatusPreconditionFailed)
		}
	}

	return ctx.SendStatus(fiber.StatusOK)
}

// Handles DELETE API requests
func (coordinator *Coordinator) DeleteUserHandler(ctx *fiber.Ctx) error {
	var primaryQuery string
	var primaryShard *grpc.ClientConn
	var emailQuery string
	var emailShard *grpc.ClientConn

	index := ctx.Query("index", "id")
	key := ctx.Query("key")

	// Given the primary key id to delete, we also need the email to delete
	if index != "email" {
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
			return ctx.SendStatus(fiber.StatusPreconditionFailed)
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
			return ctx.SendStatus(fiber.StatusPreconditionFailed)
		} else if toRes && !fromRes {
			coordinator.SendAbortBatchRPC(emailShard, idempotencyKey)
			return ctx.SendStatus(fiber.StatusPreconditionFailed)
		}
	}

	return ctx.SendStatus(fiber.StatusOK)
}

// Handles POST API requests for transactions
func (coordinator *Coordinator) TransactionHandler(ctx *fiber.Ctx) error {
	index := ctx.Query("index", "id")
	amount := ctx.Query("amount")

	// id of the recipient of the transaction
	to := ctx.Query("to")
	// id of the sender of the transaction
	from := ctx.Query("from")

	var fromQuery string
	var toQuery string

	// find primary key ids if email index is used instead
	if index != "id" {
		fromQuery = CreateQuery(index, from)
		toQuery = CreateQuery(index, to)

		fromConn := coordinator.ShardForKey(fromQuery)
		toConn := coordinator.ShardForKey(toQuery)
		fromClient := pb.NewJKZDBClient(fromConn)
		toClient := pb.NewJKZDBClient(toConn)

		res, err := fromClient.GetEntry(context.Background(), &pb.GetEntryRequest{
			Query: fromQuery,
			Field: nil,
		})
		if err != nil {
			return err
		}
		fromQuery = CreateQuery("id", res.GetEntry())

		res, err = toClient.GetEntry(context.Background(), &pb.GetEntryRequest{
			Query: toQuery,
			Field: nil,
		})
		if err != nil {
			return err
		}
		toQuery = CreateQuery("id", res.GetEntry())
	} else {
		fromQuery = CreateQuery("id", from)
		toQuery = CreateQuery("id", to)
	}

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
			coordinator.SendAbortRPC(fromShard, idempotencyKey)
			return ctx.SendStatus(fiber.StatusPreconditionFailed)
		} else if toRes && !fromRes {
			coordinator.SendAbortRPC(toShard, idempotencyKey)
			return ctx.SendStatus(fiber.StatusPreconditionFailed)
		}
	}
	
	return ctx.SendStatus(fiber.StatusOK)
}

// Handles POST API requests to create users
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
			return ctx.SendStatus(fiber.StatusPreconditionFailed)
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
		} else if !emailRes && primaryRes {
			coordinator.SendAbortRPC(primaryShard, idempotencyKey)
			return ctx.SendStatus(fiber.StatusPreconditionFailed)
		}
	}
	err := ctx.JSON(
		&fiber.Map{
			"id": userId,
		},
	)
	if err != nil {
		return nil
	}

	return ctx.SendStatus(fiber.StatusCreated)
}

// Handles PUT API requests for withdraws
func (coordinator *Coordinator) WithdrawalHandler(ctx *fiber.Ctx) error {
	index := ctx.Query("index", "id")
	key := ctx.Query("key")
	amount := ctx.Query("amount")
	amt, err := strconv.ParseInt(amount, 10, 64)
	if err != nil {
		return nil
	}

	query := CreateQuery(index, key)
	shard := coordinator.ShardForKey(query)

	if index != "id" {
		res, err := coordinator.SendGetRPC(shard, query, nil)
		if err != nil {
			return err
		}
		query = CreateQuery("id", res.Entry)
		shard = coordinator.ShardForKey(query)
	}

	idempotencyKey := time.Now().Unix()
	if coordinator.SendPrepareRPC(
		shard,
		query,
		map[string]string{
			"withdrawal": fmt.Sprint(-amt),
		},
		idempotencyKey,
		false,
	) {
		coordinator.SendCommitRPC(shard,
			query,
			map[string]string{
				"withdrawal": fmt.Sprint(-amt),
			},
			idempotencyKey,
			false)
	}

	return ctx.SendStatus(fiber.StatusOK)
}

// Handles PUT API requests for deposits
func (coordinator *Coordinator) DepositHandler(ctx *fiber.Ctx) error {
	index := ctx.Query("index", "id")
	key := ctx.Query("key")
	amount := ctx.Query("amount")
	amt, err := strconv.ParseInt(amount, 10, 64)
	if err != nil {
		return nil
	}

	query := CreateQuery(index, key)
	shard := coordinator.ShardForKey(query)

	if index != "id" {
		res, err := coordinator.SendGetRPC(shard, query, nil)
		if err != nil {
			return err
		}
		query = CreateQuery("id", res.Entry)
		shard = coordinator.ShardForKey(query)
	}
	
	idempotencyKey := time.Now().Unix()
	if coordinator.SendPrepareRPC(
		shard,
		query,
		map[string]string{
			"deposit": fmt.Sprint(amt),
		},
		idempotencyKey,
		false,
	) {
		coordinator.SendCommitRPC(shard,
			query,
			map[string]string{
				"deposit": fmt.Sprint(amt),
			},
			idempotencyKey,
			false)
	}

	return ctx.SendStatus(fiber.StatusOK)
}
