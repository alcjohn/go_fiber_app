package services

import (
	"database/sql"
	"regexp"
	"testing"
	"time"

	"github.com/alcjohn/go_blog_app/models"
	"github.com/go-playground/assert"

	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type Suite struct {
	suite.Suite
	DB              *gorm.DB
	mock            sqlmock.Sqlmock
	articlesService *ArticlesService
	article         *models.Article
}

func (s *Suite) SetupSuite() {
	var (
		db  *sql.DB
		err error
	)

	db, s.mock, err = sqlmock.New()
	require.NoError(s.T(), err)

	s.DB, err = gorm.Open(postgres.New(postgres.Config{
		Conn: db,
	}),
		&gorm.Config{})
	require.NoError(s.T(), err)

	s.articlesService = NewArticlesService(s.DB)
}

func (s *Suite) AfterTest(_, _ string) {
	require.NoError(s.T(), s.mock.ExpectationsWereMet())
}

func TestInit(t *testing.T) {
	suite.Run(t, new(Suite))
}

func (s *Suite) Test_ArticlesService_GetAll() {
	var (
		getAllRequest = `SELECT * FROM "articles"`
		articles      = []models.Article{
			models.Article{
				Model: gorm.Model{
					ID:        1,
					DeletedAt: gorm.DeletedAt{},
					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
				},
				Title:   "this is title 1",
				Content: "this is content 1",
			},
			models.Article{
				Model: gorm.Model{
					ID:        2,
					DeletedAt: gorm.DeletedAt{},
					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
				},
				Title:   "this is title 2",
				Content: "this is content 2",
			},
		}
	)
	rows := s.mock.NewRows([]string{"id", "title", "content", "created_at", "updated_at", "deleted_at"})

	for _, a := range articles {
		rows.AddRow(a.ID, a.Title, a.Content, a.CreatedAt, a.UpdatedAt, a.DeletedAt)
	}
	s.mock.ExpectQuery(regexp.QuoteMeta(getAllRequest)).WillReturnRows(rows)

	res, err := s.articlesService.GetAll()
	require.NoError(s.T(), err)
	assert.Equal(s.T(), articles, res)
}
