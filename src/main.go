package main

import (
	"github.com/kataras/iris/v12"
	"note/src/middlewares"
	"note/src/models"
	"note/src/routers"
	"note/src/util"
)

func main() {
	app := createApp()
	_ = app.Run(iris.Addr(util.Configs.AppPort), iris.WithConfiguration(iris.YAML("./src/config/iris.yml")))
}

func createApp() *iris.Application {
	app := iris.New()
	// 初始化分组路由
	routers.InitRouter(app)
	models.InitTables()
	iris.RegisterOnInterrupt(func() {
		_ = models.DB.Close()
	})
	middlewares.InitMiddleware(app)
	return app
}
