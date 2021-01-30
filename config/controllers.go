package config

import (
	"github.com/alcjohn/go_blog_app/controllers"
	"github.com/alcjohn/go_blog_app/types"
)

var Controllers = []types.Controller{
	controllers.NewArticleController,
}
