package middlewares

import (
	"github.com/iris-contrib/middleware/jwt"
	"github.com/kataras/iris/v12"
)

var mySecret = []byte("My Secret")

func GetJwtHandler() iris.Handler {
	j := jwt.New(jwt.Config{
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			return mySecret, nil
		},
		// 这里很关键，header 必须是 { Authorization： bearer xxx } 的格式（当然，我们依旧可以设置从请求参数中取得）
		Extractor:     jwt.FromAuthHeader,
		SigningMethod: jwt.SigningMethodHS256,
		ErrorHandler:  jwtErrorHandler,
		Expiration:    true,
	})

	return j.Serve
}

// jwt token 校验失败的处理函数
func jwtErrorHandler(ctx iris.Context, e error) {
	ctx.StopExecution()
	ctx.StatusCode(iris.StatusUnauthorized)
	ctx.Values().Set("msg", e.Error())
}
