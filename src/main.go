package src

import (
	"github.com/gofiber/fiber"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx){
		c.Send("hello world")
	})
}