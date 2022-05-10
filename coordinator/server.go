package main

import (
	server "JKZDB/coordinator/server"
	"log"

	"flag"

	fiber "github.com/gofiber/fiber/v2"
)

var shardConfigPath = flag.String("shard-config", "../shard-config.json", "JSON file representing shard configuration.")

func main() {
	app := fiber.New()
	coordinator, err := server.NewCoordinator(*shardConfigPath)
	if err != nil {
		log.Fatal("Could not start server.\n")
	}
	transaction := app.Group("/transaction", func(ctx *fiber.Ctx) error {
		return ctx.Next()
	})
	api := app.Group("/api", func(ctx *fiber.Ctx) error {
		return ctx.Next()
	})
	api.Get("/user", coordinator.GetUserHandler)
	api.Put("/email", coordinator.EmailUpdateHandler)
	api.Delete("", coordinator.DeleteUserHandler)
	api.Post("/user", coordinator.CreateUserHandler)
	api.Put("/deposit", coordinator.DepositHandler)
	api.Put("/withdraw", coordinator.WithdrawalHandler)
	transaction.Post("", coordinator.TransactionHandler)

	app.Listen(":5000")
}
