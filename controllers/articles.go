package controllers

import (
	"github.com/alcjohn/go_blog_app/services"
	"github.com/gofiber/fiber/v2"
)

// ArticlesController is controller for model Article
type ArticlesController struct {
	articlesService *services.ArticlesService
}

// NewArticleController return an instance of ArticlesController
func NewArticleController(articlesService *services.ArticlesService) *ArticlesController {
	return &ArticlesController{
		articlesService: articlesService,
	}
}

// RegisterRoutes register routes for articles controller and return fiber.Router
func (c *ArticlesController) RegisterRoutes(r fiber.Router) fiber.Router {
	articles := r.Group("/")
	{
		articles.Get("/", c.Index)
		articles.Get("/new", c.New)
		articles.Get("/show/:articles_id")
	}
	return articles
}

// Index render page "articles/index" with all articles
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

// New render page "articles/new"
func (c *ArticlesController) New(ctx *fiber.Ctx) error {
	ctx.Render("articles/new", fiber.Map{}, "layouts/base")
	return nil
}

func (c *ArticlesController) Create(ctx *fiber.Ctx) error {

	return nil
}
