package main

import (
	"fmt"

	_ "github.com/brain-flowing-company/pprp-backend/docs"
	"github.com/brain-flowing-company/pprp-backend/internal/greeting"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/swagger"
)

// @title        Bangkok Property Matchmaking Platform
// @description  Bangkok Property Matchmaking Platform API docs
// @version      1.0
// @host         localhost:3000
// @BasePath     /
func main() {
	app := fiber.New()

	app.Use(logger.New(logger.Config{
		TimeFormat: "02-01-2006 15:04:05",
		TimeZone:   "Asia/Bangkok",
	}))

	app.Get("/docs/*", swagger.HandlerDefault)

	hwService := greeting.NewService()
	hwHandler := greeting.NewHandler(hwService)

	app.Get("/greeting", hwHandler.Greeting)

	err := app.Listen(":3000")
	if err != nil {
		panic(fmt.Sprintf("Server cannot start with error: %v", err.Error()))
	}
}
