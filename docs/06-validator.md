# validator

## 包

`github.com/go-playground/validator/v10`


## 基础用法

```go
// 这个变量建议放在模块中供其它函数共享
var validate *validator.Validate
validate = validator.New()

// 验证基础类型数据
myEmail := "vic.gmail.com"
errs := validate.Var(myEmail, "required,email")
// 由于只校验一个数据，所以 errs 直接就是一条错误信息
if errs != nil {
    fmt.Println(errs) // output: Key: "" Error:Field validation for "" failed on the "email" tag
    return
}

// 结构体校验
type User struct {
	LastName       string     `json:"lname"`
	Age            uint8      `json:"age" validate:"gte=0,lte=130"`
	Email          string     `json:"email" validate:"required,email"`
}
user := &User{
    LastName:       "Smith",
    Age:            130,
    Email:          "Badger.Smith@gmail.com",
}
errs := validate.Struct(user)
if errs != nil && len(errs.(validator.ValidationErrors)) != 0 {
    fmt.Println(errs.Error())
}
```


## 自定义校验函数

```go
type User struct {
	FirstName      string     `json:"fname"`
	LastName       string     `json:"lname"`
}

var validate *validator.Validate
validate = validator.New()
// 第二个参数表示对 User 这个结构体添加校验函数
validate.RegisterStructValidation(UserStructLevelValidation, User{})

user := &User{
    FirstName:      "",
    LastName:       "",
}

// 自定义校验函数
func UserStructLevelValidation(sl validator.StructLevel) {
	user := sl.Current().Interface().(User)
	if len(user.FirstName) == 0 && len(user.LastName) == 0 {
		sl.ReportError(user.FirstName, "FirstName", "fname", "fname 和 lname 至少一个不为空", "")
		sl.ReportError(user.LastName, "LastName", "lname", "fname 和 lname 至少一个不为空", "")
	}
}
```


## 自定义 tag 校验

> 针对单个字段
```go
type MyStruct struct {
	String string `validate:"is-awesome"`
}

var validate *validator.Validate
validate = validator.New()
validate.RegisterValidation("is-awesome", ValidateMyVal)
s := MyStruct{String: "awesome"}
if err = validate.Struct(s); err != nil {}


func ValidateMyVal(fl validator.FieldLevel) bool {
	return fl.Field().String() == "awesome"
}
```