package controllers

import (
	"github.com/kataras/iris/sessions"
)

type UserController struct {
	Session *sessions.Session
}

// by 是关键字，可以用来获取路径参数
func (c *UserController) GetBy(username string) interface{} {
	return map[string]string{"message": username}
}

func (c *UserController) PostHello() interface{} {
	return map[string]string{"message": "Hello Vic!"}
}
