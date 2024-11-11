package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/wignn/api-with-go/config"
	"github.com/wignn/api-with-go/db"
	"github.com/wignn/api-with-go/handlers"
	"github.com/wignn/api-with-go/middlewares"
	"github.com/wignn/api-with-go/repositories"
	"github.com/wignn/api-with-go/services"
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
