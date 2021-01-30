package controllers

import (
	"github.com/alcjohn/go_blog_app/models"
	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/gorm"
)

type ArticlesController struct {
	db *gorm.DB
}

func NewArticleController(db *gorm.DB) *ArticlesController {
	return &ArticlesController{
		db: db,
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
	var articles []models.Article
	if err := c.db.Find(&articles).Error; err != nil {
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
