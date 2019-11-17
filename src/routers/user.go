package routers

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"note/src/controllers"
)

func UserMvc(app *mvc.Application) {
	//当然，你可以在MVC应用程序中使用普通的中间件。
	app.Router.Use(func(ctx iris.Context) {
		ctx.Application().Logger().Infof("Path: %s", ctx.Path())
		ctx.Next()
	})

	app.Handle(new(controllers.UserController))
}
