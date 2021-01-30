package kernel

import (
	"fmt"
	"os"

	"github.com/alcjohn/go_blog_app/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDatabase() *gorm.DB {
	dbConfig := struct {
		Host     string
		Port     string
		User     string
		Database string
		Password string
		SSL      string
	}{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Database: os.Getenv("DB_NAME"),
		Password: os.Getenv("DB_PASSWORD"),
		SSL:      os.Getenv("DB_SSL"),
	}
	dsn := fmt.Sprintf(

		"host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.User,
		dbConfig.Database,
		dbConfig.Password,
		dbConfig.SSL,
	)

	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}
	database.AutoMigrate(config.Models[:]...)
	return database
}
