package main

import (
	"github.com/kataras/iris"
)

func main() {
	app := iris.Default()
	app.Get("/ping", func(ctx iris.Context) {
		_, _ = ctx.JSON(iris.Map{
			"message": "pong",
		})
	})

	_ = app.Run(iris.Addr(":8080"))
}
