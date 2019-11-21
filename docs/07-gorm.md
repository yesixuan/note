## 名字映射规则

表名和字段名都遵循：驼峰变下划线，其他变小写

## 自定义字段名

```go
type Animal struct {
  AnimalId    int64     `gorm:"column:beast_id"`
}
```

## 查询

### 根据 id 查 
`db.First(&user, 110)`

### 一些好的 api

- FirstOrInit 有记录则返回，没有就初始化一个  
- Attrs 查到某个记录，并更改这个记录的某些属性  
- Assign 无论是否找到，都给他分配属性  
- FirstOrCreate 类似 FirstOrInit  
- SubQuery 
