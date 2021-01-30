package kernel

import (
	"fmt"
	"os"

	"github.com/alcjohn/go_blog_app/config"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
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
	dbURL := fmt.Sprintf(

		"host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.User,
		dbConfig.Database,
		dbConfig.Password,
		dbConfig.SSL,
	)
	fmt.Println(dbURL)

	database, err := gorm.Open("postgres", dbURL)

	if err != nil {
		panic(err)
	}
	database.AutoMigrate(config.Models[:]...)
	return database
}
