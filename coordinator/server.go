package main

import (
	server "JKZDB/Coordinator/server"
	"log"

	fiber "github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	coordinator, err := server.NewCoordinator("")
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
