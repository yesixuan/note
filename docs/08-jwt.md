# jwt

## iris jwt 中间件

`"github.com/iris-contrib/middleware/jwt"`


## 生成 token

```go
var mySecret = []byte("My Secret")

// 传入 uid 存到 jwt 中
func GetToken(uid int) string {
	token := jwt.NewTokenWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"uid": uid,
		"exp": time.Now().Add(30*24*time.Hour * time.Duration(1)).Unix(), // 过期时间
	})
	tokenString, _ := token.SignedString(mySecret)
	return tokenString
}
```


## 获取 jwt handler

```go
var mySecret = []byte("My Secret")

func GetJwtHandler() iris.Handler {
	j := jwt.New(jwt.Config{
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			return mySecret, nil
		},
		// 这里很关键，header 必须是 { Authorization： bearer xxx } 的格式（当然，我们依旧可以设置从请求参数中取得）
		Extractor: jwt.FromAuthHeader,
		SigningMethod: jwt.SigningMethodHS256,
        // 自定义错误处理函数，方便后面进行统一的数据格式化
		ErrorHandler: jwtErrorHandler,
        // 开启过期验证（过期时间在 jwt.MapClaims 中设置）
		Expiration: true,
	})
    
	return j.Serve
}

// jwt token 校验失败的处理函数
func jwtErrorHandler(ctx iris.Context, e error) {
	ctx.StopExecution()
	ctx.StatusCode(iris.StatusUnauthorized)
	ctx.Values().Set("msg", e.Error())
}
```