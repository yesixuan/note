# 错误处理

## 统一处理异常的 HTTP 状态码

```go
/**
iris.StatusUnauthorized 401
iris.StatusNotFound 404
iris.StatusInternalServerError 500
*/
app.OnErrorCode(iris.StatusUnauthorized, handler)
```


## 所有错误统一处理

```go
app.OnAnyErrorCode(handler)
```


## 手动抛出错误（如鉴权不通过时，抛出401）

```go
func handler(ctx iris.Context) {
    // 有了这一行，ctx.Next() 将不会被执行
    ctx.StopExecution()
    ctx.StatusCode(iris.StatusUnauthorized)
    ctx.Values().Set("msg", "传递给错误拦截中间件")
}

// 统一异常处理
app.OnErrorCode(iris.StatusUnauthorized, func(ctx iris.Context) {
    ctx.JSON(iris.Map{
        "code": 0,
        "message": ctx.Values().GetString("msg"),
        "data": nil,
    })
})
```