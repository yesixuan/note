package main

import (
	"github.com/kataras/iris/v12"
	"note/src/routers"
)

func createApp() *iris.Application {
	app := iris.Default()
	// 初始化分组路由
	routers.InitRouter(app)
	return app
}

func main() {
	app := createApp()

	_ = app.Run(iris.Addr(":8080"), iris.WithConfiguration(iris.YAML("./config/iris.yml")))
}
