package kernel

import (
	"github.com/alcjohn/go_blog_app/config"
	"github.com/alcjohn/go_blog_app/types"
	"github.com/goava/di"
)

var ServiceProvider = func() di.Option {
	var array []di.Option
	for _, s := range config.Services {
		array = append(array, di.Provide(s, di.As(new(types.Service))))
	}
	return di.Options(array[:]...)
}()
