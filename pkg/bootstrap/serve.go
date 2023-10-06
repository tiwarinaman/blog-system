package bootstrap

import (
	"blog-system/pkg/config"
	"blog-system/pkg/database"
	"blog-system/pkg/html"
	"blog-system/pkg/routing"
	"blog-system/pkg/static"
)

func Serve() {
	config.Set()

	database.Connect()

	routing.Init()

	static.LoadStatic(routing.GetRouter())

	html.LoadHtml(routing.GetRouter())

	routing.RegisterRoutes()

	routing.Serve()
}
