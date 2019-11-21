package routers

import (
	"github.com/kataras/iris/v12"
)

func UsersRoutes(usersRouter iris.Party) {
	usersRouter.Get("/{name}", getAllUsersHandler)
}

func getAllUsersHandler(ctx iris.Context) {
	_, _ = ctx.JSON(map[string]string{"message": ctx.Params().Get("name")})
}
