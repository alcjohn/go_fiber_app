package config

import (
	"github.com/alcjohn/go_blog_app/services"
	"github.com/alcjohn/go_blog_app/types"
)

var Services = []types.Service{
	services.NewArticlesService,
}
