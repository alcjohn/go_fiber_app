package kernel

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func NewApp(engine fiber.Views) *fiber.App {
	app := fiber.New(fiber.Config{
		Views: engine,
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			c.Status(fiber.StatusInternalServerError)

			return c.SendString(err.Error())
		},
	})
	app.Use(recover.New())
	return app
}
