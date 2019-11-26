package middlewares

import (
	"github.com/iris-contrib/middleware/jwt"
	"github.com/kataras/iris/v12"
	"note/src/models"
)

func HasAuth(auth string) iris.Handler {
	return func(ctx iris.Context) {
		var user models.User
		userInfo := ctx.Values().Get("jwt").(*jwt.Token).Claims.(jwt.MapClaims)
		uid := int(userInfo["uid"].(float64))
		permissions := user.GetPermissions(uid)
		for _, permission := range permissions {
			if permission == auth {
				ctx.Next()
				return
			}
		}
		ctx.StopExecution()
		ctx.StatusCode(iris.StatusForbidden)
		ctx.Values().Set("msg", "无权访问")
	}
}
