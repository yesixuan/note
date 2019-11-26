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
	Uname    string  `gorm:"type:varchar(50);not null;unique_index"json:"uname"`
	Password string  `gorm:"type:varchar(200);not null"json:"password"`
	Mobile   string  `gorm:"type:varchar(11);not null"json:"mobile"`
	Email    string  `gorm:"type:varchar(200);not null"json:"email"`
	Motto    string  `gorm:"type:varchar(50)"json:"motto"`
	Role     []*Role `gorm:"many2many:user_role;"`
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
	// 给用户一个默认角色
	DB.Model(user).Association("role").Append(Role{
		Name: "normal",
	})
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
	ctx.Values().Set("data", util.GetToken(result.ID))
	ctx.Next()
}

func (user *User) GetPermissions(uid int) []string {
	var permissions []Permission
	var result []string
	rows, _ := DB.Raw("SELECT DISTINCT permission.name FROM user LEFT JOIN user_role ON user.id = user_role.user_id LEFT JOIN role ON user_role.role_id = role.id LEFT JOIN role_permission on role.id = role_permission.role_id LEFT JOIN permission ON role_permission.permission_id = permission.id WHERE user.id = ?", uid).Rows()
	defer rows.Close()
	//DB.ScanRows(rows, &permissions)
	i := 0
	for rows.Next() {
		permissions = append(permissions, Permission{})
		_ = DB.ScanRows(rows, &permissions[i])
		result = append(result, permissions[i].Name)
		i++
	}

	return result
}
