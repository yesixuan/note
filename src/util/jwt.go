package util

import (
	"github.com/iris-contrib/middleware/jwt"
	"time"
)

var mySecret = []byte("My Secret")

func GetToken(uid string) string {
	token := jwt.NewTokenWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"uid": uid,
		"exp": time.Now().Add(30 * 24 * time.Hour * time.Duration(1)).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, _ := token.SignedString(mySecret)
	return tokenString
}
