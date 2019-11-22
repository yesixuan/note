package models

import (
	"github.com/jameskeane/bcrypt"
	"github.com/jinzhu/gorm"
	"github.com/kataras/iris/v12"
)

type User struct {
	gorm.Model
	Uname    string `gorm:"type:varchar(50);not null;unique_index"`
	Password string `gorm:"type:varchar(200);not null"`
	Mobile   string `gorm:"type:varchar(11);not null"`
	Email    string `gorm:"type:varchar(200);not null"`
	Motto    string `gorm:"type:varchar(50)"`
	Status   string `gorm:"type:enum('normal', 'abnormal');default:'normal'"`
}

func (user *User) CreateUser(ctx iris.Context) {
	salt, _ := bcrypt.Salt(10)
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
