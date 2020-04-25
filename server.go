package main

import (
	"api-book/src"
	"github.com/gofiber/fiber"
)

func main() {
	app := fiber.New()
	src.SetupRoutes(app)
	app.Listen("3000")
}
