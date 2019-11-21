# 路由

## 分组路由

方案一：  
```go
// 这里可以有第二个参数，传入中间件
users := app.Party("/users")
users.Get("/haha", func(ctx iris.Context) {
    ctx.JSON(map[string]string{ "name": "vic" })
})
```

方案二：  
```go
app.PartyFunc("/users", func(users iris.Party) {
    users.Use(myAuthMiddlewareHandler) // 添加中间件

    users.Get("/haha", func(ctx iris.Context) {
        ctx.JSON(map[string]string{ "name": "vic" })
    })
})
```


## 路径参数

```go
app.Get("/u/{username:string}", func(ctx iris.Context) {
	ctx.Writef(ctx.Params().Get("username"))
})
```

### 参数类型

**iris 会对参数进行类型转换，如果能转换，则匹配成功；如果不能转换，不会报错，只是路由不会匹配**

更多限定：  
```
"/haha/{name:int min(1)}"
"/haha/{name:int max(8)}"

"/haha/{name:string regexp(^1[0-9]{9}$)}"
```


## 中间件

**中间件本质上也是 handle 函数**

### 中间件的类型

- 前置中间件。必须要有 `ctx.Next()`  
- 后置中间件。不必 `ctx.Next()`  

### 中间件的使用

- 与 handle 函数一样，作为参数传入 `app.Get("/", before, handler, after)`  
- 使用 Party 路由的 `Use` `Done` 方法传入（在需要对某个模块的所有接口统一处理时很有用）  
- 统一使用 `app.UseGlobal(before)` `app.DoneGlobal(after)` 只能是 **app**   

### 中间件共享数据（仅限简单基础数据）

```go
// 存
ctx.Values().Set("data", "some data")
// 取
ctx.Values().GetString("info")
```