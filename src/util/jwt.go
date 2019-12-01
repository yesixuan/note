package util

import (
	"github.com/iris-contrib/middleware/jwt"
	"github.com/kataras/iris/v12"
	"time"
)

var mySecret = []byte("My Secret")

func GetToken(uid uint) string {
	token := jwt.NewTokenWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"uid": uid,
		"exp": time.Now().Add(30 * 24 * time.Hour * time.Duration(1)).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, _ := token.SignedString(mySecret)
	return tokenString
}

func GetUserId(ctx iris.Context) int {
	userInfo := ctx.Values().Get("jwt").(*jwt.Token).Claims.(jwt.MapClaims)
	return int(userInfo["uid"].(float64))
}
