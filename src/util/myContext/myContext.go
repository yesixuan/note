package myContext

import (
	"errors"
	"github.com/iris-contrib/middleware/jwt"
	"github.com/kataras/iris/v12"
	"note/src/models"
)

type MyContext struct {
	iris.Context
}

// 判断是否有权限
func (c MyContext) HasAuth(auth string) bool {
	var user models.User
	userInfo := c.Context.Values().Get("jwt").(*jwt.Token).Claims.(jwt.MapClaims)
	uid := int(userInfo["uid"].(float64))
	permissions := user.GetPermissions(uid)
	for _, permission := range permissions {
		if permission == auth {
			return true
		}
	}
	return false
}

type GraphqlBody struct {
	query         string
	variables     string
	operationName string
}

// 解析 variables
func (c MyContext) Variables() ([]byte, error) {
	var graphqlBody GraphqlBody
	if err := c.Context.ReadJSON(&graphqlBody); err != nil || graphqlBody.variables == "" {
		return []byte(""), errors.New("解析 variables 出错！")
	}
	return []byte(graphqlBody.variables), nil
}
