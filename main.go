package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/favicon"
	"github.com/gofiber/fiber/v2/middleware/logger"
	//"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"go.uber.org/zap"

	"github.com/mecitsemerci/go-todo-app/internal/rest"
)

// @title Todo API
// @version 1.0.0
// @description This is a sample todo restful api server.
// @host localhost:8080
// @BasePath /
func main() {
	app := fiber.New()

	app.Use(logger.New())
	app.Use(requestid.New())
	app.Use(favicon.New())
	app.Use(cors.New())
	//app.Use(recover.New())

	logger, err := zap.NewProduction()

	if err != nil {
		log.Fatal(err)
	}

	defer logger.Sync()

	err = rest.Register(app, logger)

	if err != nil {

		logger.Error("failed to register routes", zap.Error(err))
	}

	err = app.Listen(":8080")

	if err != nil {
		logger.Error("failed to start server", zap.Error(err))
	}
}
