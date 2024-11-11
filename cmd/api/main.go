package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/wignn/Native/config"
	"github.com/wignn/Native/db"
	"github.com/wignn/Native/handlers"
	"github.com/wignn/Native/middlewares"
	"github.com/wignn/Native/repositories"
	"github.com/wignn/Native/services"
)

func main() {
	envConfig := config.NewEnvConfig()
	db := db.Init(envConfig, db.DBMigrator)

	app := fiber.New(fiber.Config{
		AppName:      "Native-Go-API",
		ServerHeader: "Fiber",
	})


	eventRepository := repositories.NewEventRepository(db)
	ticketRepository := repositories.NewTicketRepository(db)
	authRepository := repositories.NewAuthRepository(db)


	authService := services.NewAuthService(authRepository)

	
	server := app.Group("/api")
	handlers.NewAuthHandler(server.Group("/auth"), authService)

	privateRoutes := server.Use(middlewares.AuthProtected(db))

	handlers.NewEventHandler(privateRoutes.Group("/event"), eventRepository)
	handlers.NewTicketHandler(privateRoutes.Group("/ticket"), ticketRepository)

	app.Listen(fmt.Sprintf(":" + envConfig.ServerPort))
}
