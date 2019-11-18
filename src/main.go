package main

import (
	"github.com/kataras/iris"
	"note/src/routers"
)

func createApp() *iris.Application {
	app := iris.New()

	app.PartyFunc("/user", routers.UsersRoutes)

	return app
}

func main() {
	app := createApp()

	_ = app.Run(iris.Addr(":8080"), iris.WithConfiguration(iris.YAML("./config/iris.yml")))
}
