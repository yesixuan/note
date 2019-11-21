## 拉取依赖

```bash
go mod vendor
```

## 热重启

> go build 启动的脚本为 go build note/src/main.go  
> rizla 的启动脚本为 rizla src/main.go  
> 这会导致项目中的相对路径出现问题  
```bash
go get -u github.com/kataras/rizla

rizla src/main.go
```