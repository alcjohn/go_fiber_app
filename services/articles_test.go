package services

import (
	"database/sql"
	"regexp"
	"testing"
	"time"

	"github.com/alcjohn/go_blog_app/models"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/alcjohn/go_blog_app/utils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
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
		getAllRequest = `SELECT * FROM "articles" WHERE "articles"."deleted_at" IS NULL`
		articles      = []models.Article{
			models.Article{
				Title:   "this is title 1",
				Content: "this is content 1",
			},
			models.Article{
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

func (s *Suite) Test_ArticlesService_GetByID() {
	var getByIDResquest = `SELECT * FROM "articles" WHERE "articles"."id" = $1 AND "articles"."deleted_at" IS NULL ORDER BY "articles"."id" LIMIT 1`
	s.Run("ID EXIST", func() {
		a := models.Article{
			Model: gorm.Model{
				ID:        1,
				DeletedAt: gorm.DeletedAt{},
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			Title:   "this is title 1",
			Content: "this is content 1",
		}
		rows := s.mock.NewRows(
			[]string{"id", "title", "content", "created_at", "updated_at", "deleted_at"},
		).AddRow(a.ID, a.Title, a.Content, a.CreatedAt, a.UpdatedAt, a.DeletedAt)
		s.mock.ExpectQuery(regexp.QuoteMeta(getByIDResquest)).WithArgs(a.ID).WillReturnRows(rows)

		res, err := s.articlesService.GetByID(a.ID)
		require.NoError(s.T(), err)
		assert.Equal(s.T(), a, res)
	})
	s.Run("ID NOT EXIST", func() {

		var idToTest = uint(42)

		s.mock.ExpectQuery(regexp.QuoteMeta(getByIDResquest)).WithArgs(idToTest).WillReturnError(gorm.ErrRecordNotFound)

		res, err := s.articlesService.GetByID(idToTest)
		require.Error(s.T(), err)
		assert.Equal(s.T(), uint(0), res.ID)
	})
}

func (s *Suite) Test_ArticlesService_Create() {
	var (
		createRequest = `INSERT INTO "articles" ("created_at","updated_at","deleted_at","title","content") VALUES ($1,$2,$3,$4,$5) RETURNING "id"`
		a             = models.Article{
			Title:   "This is Title",
			Content: "This is Content",
		}
	)
	rows := s.mock.NewRows(
		[]string{"id", "title", "content", "created_at", "updated_at", "deleted_at"},
	)
	s.mock.ExpectQuery(
		regexp.QuoteMeta(createRequest),
	).WithArgs(utils.AnyTime{}, utils.AnyTime{}, nil, a.Title, a.Content).WillReturnRows(rows)

	res, err := s.articlesService.Create(a)
	require.NoError(s.T(), err)
	assert.Equal(s.T(), a.Title, res.Title)
	assert.Equal(s.T(), a.Content, res.Content)
}

func (s *Suite) Test_ArticlesService_Update() {
	var query = `UPDATE "articles" SET "updated_at"=$1,"title"=$2,"content"=$3 WHERE id = $4`
	s.Run("Update with article exist", func() {
		var (
			a = models.Article{
				Title:   "This is new Title",
				Content: "This is new Content",
			}
		)
		s.mock.ExpectExec(regexp.QuoteMeta(query)).WithArgs(sqlmock.AnyArg(), a.Title, a.Content, 1).WillReturnResult(sqlmock.NewResult(0, 1))
		res, err := s.articlesService.Update(1, a)
		require.NoError(s.T(), err)
		assert.Equal(s.T(), a, res)
	})
	s.Run("Update with not existed article", func() {
		var (
			a = models.Article{
				Title:   "This is new Title",
				Content: "This is new Content",
			}
		)
		s.mock.ExpectExec(regexp.QuoteMeta(query)).WithArgs(sqlmock.AnyArg(), a.Title, a.Content, 1).WillReturnError(gorm.ErrRecordNotFound)
		_, err := s.articlesService.Update(1, a)
		require.Error(s.T(), err)
	})
}
