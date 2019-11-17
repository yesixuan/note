package main

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"note/src/routers"
)

func createApp() *iris.Application {
	app := iris.New()

	mvc.Configure(app.Party("/api/v1/user"), routers.UserMvc)
	//mvc.New(app).Handle(new(controllers.UserController))

	return app
}

func main() {
	app := createApp()

	_ = app.Run(iris.Addr(":8080"))
}
