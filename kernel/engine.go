package kernel

import (
	"os"

	"github.com/gofiber/template/django"
)

func NewEngine() *django.Engine {
	engine := django.New("./views", ".html")
	engine.Reload(os.Getenv("APP_ENV") != "prod")
	return engine
}
