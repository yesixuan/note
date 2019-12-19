package routers

import (
	"context"
	"github.com/kataras/iris/v12"
	"note/src/schema"
	myContext2 "note/src/util/myContext"
	"note/src/vicGraphql"
)

func GraphqlRoute(usersRouter iris.Party) {
	usersRouter.Get("/", GraphqlHandler)
	usersRouter.Post("/", GraphqlHandler)
}

func GraphqlHandler(ctx iris.Context) {
	h := vicGraphql.New(&vicGraphql.Config{
		Schema:   &schema.Schema,
		Pretty:   true,
		GraphiQL: true,
		//Playground: true,
	})
	// 传入自己封装的的 ctx
	myContext := context.WithValue(context.Background(), "ctx", myContext2.MyContext{Context: ctx})

	// 使用自己封装的 graphql handler
	result := h.ServeHTTP(ctx, myContext)
	if len(result.Errors) > 0 {
		ctx.StopExecution()
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.Values().Set("msg", result.Errors)
	} else {
		ctx.Values().Set("data", result.Data)
	}
	ctx.Next()
}
