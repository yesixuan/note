package routers

import (
	"fmt"
	"github.com/jmespath/go-jmespath"
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
)

func UsersRoutes(usersRouter iris.Party) {
	usersRouter.Get("/{name}", getAllUsersHandler)
}

func getAllUsersHandler(ctx iris.Context) {
	_, _ = ctx.JSON(map[string]string{"message": ctx.Params().Get("name")})

	//_ = sendJSON(ctx, map[string]string{"message": ctx.Params().Get("name")})
	//if err != nil {
	//	fail(ctx, iris.StatusInternalServerError, "unable to send a list of all users: %v", err)
	//	return
	//}
}

type httpError struct {
	Code   int    `json:"code"`
	Reason string `json:"reason"`
}

func (h httpError) Error() string {
	return fmt.Sprintf("Status Code: %d\nReason: %s", h.Code, h.Reason)
}

func fail(ctx context.Context, statusCode int, format string, a ...interface{}) {
	err := httpError{
		Code:   statusCode,
		Reason: fmt.Sprintf(format, a...),
	}
	//记录所有> = 500内部错误。
	if statusCode >= 500 {
		ctx.Application().Logger().Error(err)
	}
	ctx.StatusCode(statusCode)
	ctx.JSON(err)
	//没有下一个处理程序将运行。
	ctx.StopExecution()
}

func sendJSON(ctx iris.Context, resp interface{}) (err error) {
	indent := ctx.URLParamDefault("indent", "  ")
	// i.e [?Name == 'John Doe'].Age # to output the [age] of a user which his name is "John Doe".
	if query := ctx.URLParam("query"); query != "" && query != "[]" {
		resp, err = jmespath.Search(query, resp)
		if err != nil {
			return
		}
	}
	_, err = ctx.JSON(resp, context.JSON{Indent: indent, UnescapeHTML: true})
	return err
}
