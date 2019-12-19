package routers

import "github.com/kataras/iris/v12"

func InitRouter(app *iris.Application) {
	root := app.Party("/api")
	//root.PartyFunc("/user", UsersRoutes)
	root.PartyFunc("/graphql", GraphqlRoute)
}
