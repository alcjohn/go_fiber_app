package kernel

import (
	"fmt"

	"github.com/alcjohn/go_blog_app/config"
	"github.com/alcjohn/go_blog_app/types"
	"github.com/goava/di"
)

var ControllerProvider = func() di.Option {
	fmt.Println("Provider controllers")
	var array []di.Option
	for _, c := range config.Controllers {
		array = append(array, di.Provide(c, di.As(new(types.Controller))))
	}
	return di.Options(array[:]...)
}()
