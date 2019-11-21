package main

import (
	"github.com/iris-contrib/middleware/cors"
	"github.com/kataras/iris/v12"
	"note/src/routers"
)

func createApp() *iris.Application {
	app := iris.Default()
	// 初始化分组路由
	routers.InitRouter(app)
	crs := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"}, // allows everything, use that to change the hosts.
		AllowCredentials: true,
		AllowedHeaders:   []string{"*"},
	})
	app.Use(crs)
	app.AllowMethods(iris.MethodOptions)
	return app
}

func main() {
	app := createApp()

	_ = app.Run(iris.Addr(":8080"), iris.WithConfiguration(iris.YAML("./src/config/iris.yml")))
}
