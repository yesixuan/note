package middlewares

import (
	"github.com/iris-contrib/middleware/cors"
	"github.com/kataras/iris/v12"
)

func InitMiddleware(app *iris.Application) {
	useCors(app)
	useHandleError(app)
	app.DoneGlobal(useGlobalAfter)
}

func useCors(app *iris.Application) {
	crs := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"}, // allows everything, use that to change the hosts.
		AllowCredentials: true,
		AllowedHeaders:   []string{"*"},
	})
	app.Use(crs)
	app.AllowMethods(iris.MethodOptions)
}

func useHandleError(app *iris.Application) {
	app.OnAnyErrorCode(func(ctx iris.Context) {
		_, _ = ctx.JSON(iris.Map{
			"code":    -1,
			"message": ctx.Values().GetString("msg"),
		})
	})
}

func useGlobalAfter(ctx iris.Context) {
	msg := "ok"
	if ctx.Values().GetString("msg") != "" {
		msg = ctx.Values().GetString("msg")
	}
	_, _ = ctx.JSON(iris.Map{
		"code":    0,
		"data":    ctx.Values().Get("data"),
		"message": msg,
	})
}
