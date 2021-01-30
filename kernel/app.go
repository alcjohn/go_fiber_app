package kernel

import (
	"github.com/gofiber/fiber/v2"
)

func NewApp(engine fiber.Views) *fiber.App {
	app := fiber.New(fiber.Config{
		Views: engine,
	})
	return app
}
