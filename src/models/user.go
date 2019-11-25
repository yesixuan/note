package models

import (
	"github.com/jameskeane/bcrypt"
	"github.com/kataras/iris/v12"
	"note/src/util"
	"note/src/validators"
)

var salt, _ = bcrypt.Salt(10)

type User struct {
	BaseModel
	Uname    string `gorm:"type:varchar(50);not null;unique_index"json:"uname"`
	Password string `gorm:"type:varchar(200);not null"json:"password"`
	Mobile   string `gorm:"type:varchar(11);not null"json:"mobile"`
	Email    string `gorm:"type:varchar(200);not null"json:"email"`
	Motto    string `gorm:"type:varchar(50)"json:"motto"`
}

func (user *User) CreateUser(ctx iris.Context) {
	hash, _ := bcrypt.Hash(user.Password, salt)
	user.Password = hash
	if err := DB.Create(user).Error; err != nil {
		ctx.StopExecution()
		ctx.StatusCode(iris.StatusConflict)
		ctx.Values().Set("msg", err.Error())
		return
	}
	ctx.Values().Set("data", user)
	ctx.Values().Set("msg", "注册成功")
}

func (user *User) Login(ctx iris.Context, loginUser *validators.LoginUser) {
	var result User
	if result := DB.Where("uname = ?", loginUser.Uname).First(&result); result.Error != nil {
		ctx.StopExecution()
		ctx.StatusCode(iris.StatusNotFound)
		ctx.Values().Set("msg", "不存在该用户")
		return
	}
	if !bcrypt.Match(loginUser.Password, result.Password) {
		ctx.StopExecution()
		ctx.StatusCode(iris.StatusNotFound)
		ctx.Values().Set("msg", "用户名或密码错误")
		return
	}
	ctx.Values().Set("data", util.GetToken(int(result.ID)))
	ctx.Next()
}
