package controllers

import (
	"github.com/alcjohn/go_blog_app/services"
	"github.com/gofiber/fiber/v2"
)

type ArticlesController struct {
	articlesService *services.ArticlesService
}

func NewArticleController(articlesService *services.ArticlesService) *ArticlesController {
	return &ArticlesController{
		articlesService: articlesService,
	}
}

func (c *ArticlesController) RegisterRoutes(r fiber.Router) fiber.Router {
	articles := r.Group("/")
	{
		articles.Get("/", c.Index)
		articles.Get("/new", c.New)
	}
	return articles
}

func (c *ArticlesController) Index(ctx *fiber.Ctx) error {
	articles, err := c.articlesService.GetAll()
	if err != nil {
		return err
	}
	ctx.Render("articles/index", fiber.Map{
		"Title":    "Hello world !",
		"articles": articles,
	}, "layouts/base")
	return nil
}

func (c *ArticlesController) New(ctx *fiber.Ctx) error {
	ctx.Render("articles/new", fiber.Map{}, "layouts/base")
	return nil
}

func (c *ArticlesController) Create(ctx *fiber.Ctx) error {

	return nil
}
