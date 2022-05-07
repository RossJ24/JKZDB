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
	api := app.Group("/api", func(ctx *fiber.Ctx) error {
		return ctx.Next()
	})
	api.Get("", coordinator.GetHandler)
	api.Put("", coordinator.PutHandler)
	api.Delete("", coordinator.DeleteHandler)
}
