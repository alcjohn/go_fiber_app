package config

import (
	"github.com/alcjohn/go_blog_app/controllers"
	"github.com/gofiber/fiber/v2"
)

type Routes struct{}

func NewRoutes(
	app *fiber.App,
	articlesController *controllers.ArticlesController,
) *Routes {
	router := app.Group("/")
	router.Get("/err", func(ctx *fiber.Ctx) error {
		panic("Tugudu")
	})
	articlesController.RegisterRoutes(router)
	return &Routes{}
}
