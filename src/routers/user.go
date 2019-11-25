package routers

import (
	"github.com/iris-contrib/middleware/jwt"
	"github.com/kataras/iris/v12"
	"note/src/middlewares"
	"note/src/models"
	"note/src/validators"
)

func UsersRoutes(usersRouter iris.Party) {
	usersRouter.Post("/register", register)
	usersRouter.Post("/login", login)
	usersRouter.Post("/test", middlewares.GetJwtHandler(), func(ctx iris.Context) {
		var user models.User
		userInfo := ctx.Values().Get("jwt").(*jwt.Token).Claims.(jwt.MapClaims)
		uid := int(userInfo["uid"].(float64))
		permissions := user.GetPermissions(uid)
		//ctx.JSON(iris.Map{"hehe": "haha"})
		ctx.Values().Set("data", permissions)
		ctx.Next()
	})
}

func register(ctx iris.Context) {
	// 解析
	var user validators.User
	var userModel models.User
	if err := ctx.ReadJSON(&user); err != nil {
		ctx.StopExecution()
		ctx.StatusCode(iris.StatusUnauthorized)
		ctx.Values().Set("msg", "json 解析错误")
		return
	}
	// 校验
	if user.Verify(ctx) != nil {
		return
	}
	//ctx.JSON(user)
	// 密码加盐
	// 入库
	_ = ctx.ReadJSON(&userModel)
	userModel.CreateUser(ctx)
	ctx.Next()
}

func login(ctx iris.Context) {
	var loginUser validators.LoginUser
	var user models.User
	if err := ctx.ReadJSON(&loginUser); err != nil {
		ctx.StopExecution()
		ctx.StatusCode(iris.StatusUnauthorized)
		ctx.Values().Set("msg", "json 解析错误")
		return
	}
	if loginUser.Verify(ctx) != nil {
		return
	}
	user.Login(ctx, &loginUser)
}
