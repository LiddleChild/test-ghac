package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := fiber.New()

	app.Use(logger.New(logger.Config{
		TimeFormat: "02-01-2006 15:04:05",
		TimeZone:   "Asia/Bangkok",
	}))

	app.Get("/greeting", func(c *fiber.Ctx) error {
		return c.SendString("Hello, Worldgo")
	})

	err := app.Listen(fmt.Sprintf(":3000"))
	if err != nil {
		panic(fmt.Sprintf("Server cannot start with error: %v", err.Error()))
	}
}
