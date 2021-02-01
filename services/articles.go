package services

import (
	"github.com/alcjohn/go_blog_app/models"
	"gorm.io/gorm"
)

// ArticlesService is a service for Article model
type ArticlesService struct {
	db *gorm.DB
}

// NewArticlesService Service returns ArticlesService
func NewArticlesService(db *gorm.DB) *ArticlesService {
	return &ArticlesService{
		db: db,
	}
}

// GetAll return all Articles
func (s *ArticlesService) GetAll() ([]models.Article, error) {
	var articles []models.Article
	err := s.db.Find(&articles).Error
	return articles, err
}

// GetByID return an Article
func (s *ArticlesService) GetByID(id uint) (models.Article, error) {
	var article models.Article
	err := s.db.First(&article, id).Error
	return article, err
}

// Create an Article and return this
func (s *ArticlesService) Create(article models.Article) (models.Article, error) {
	err := s.db.Create(&article).Error
	return article, err
}

// Update an article
func (s *ArticlesService) Update(id uint, article models.Article) (models.Article, error) {
	err := s.db.Model(models.Article{}).Where("id = ?", id).Updates(article).Error
	return article, err
}

// Delete an Article
func (s *ArticlesService) Delete(id uint) error {
	return s.db.Where("id = ?", id).Delete(&models.Article{}).Error
}
