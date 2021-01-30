package main

import (
	"log"

	"github.com/alcjohn/go_blog_app/config"
	"github.com/alcjohn/go_blog_app/kernel"
	"github.com/goava/di"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}
	c, err := di.New(
		di.Provide(kernel.NewApp),
		di.Provide(kernel.NewEngine, di.As(new(fiber.Views))),
		di.Provide(config.NewRoutes),
		di.Provide(kernel.NewDatabase),
		kernel.ControllerProvider,
		kernel.ServiceProvider,
	)
	if err != nil {
		log.Fatal(err)
	}
	log.Fatal(c.Invoke(start))
}

func start(app *fiber.App, db *gorm.DB, _ *config.Routes) error {
	app.Get("/monitor", monitor.New())
	return app.Listen(":4000")
}
