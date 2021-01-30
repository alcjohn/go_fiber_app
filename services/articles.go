package services

import (
	"github.com/alcjohn/go_blog_app/models"
	"gorm.io/gorm"
)

type ArticlesService struct {
	db *gorm.DB
}

func NewArticlesService(db *gorm.DB) *ArticlesService {
	return &ArticlesService{
		db: db,
	}
}

func (s *ArticlesService) GetAll() ([]models.Article, error) {
	var articles []models.Article
	s.db.Find(&articles)
	return articles, nil
}
