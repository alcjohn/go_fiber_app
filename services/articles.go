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
	err := s.db.Find(&articles).Error
	return articles, err
}

func (s *ArticlesService) GetByID(id uint) (models.Article, error) {
	var article models.Article
	err := s.db.First(&article, id).Error
	return article, err
}
