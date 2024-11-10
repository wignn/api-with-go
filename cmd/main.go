package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/wignn/Native/handlers"
	"github.com/wignn/Native/repository"
)

func main() {
	app := fiber.New( fiber.Config{
		AppName: "Native",
		ServerHeader: "Fiber",
	})

	eventRepository := repository.NewEventRepository(nil)

	server := app.Group("api")

	handlers.NewEventHandler(server.Group("/event"), eventRepository)

	app.Listen(":3000")
}
