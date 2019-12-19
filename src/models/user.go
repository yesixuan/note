package models

import (
	"errors"
	"github.com/go-playground/validator/v10"
	"github.com/jameskeane/bcrypt"
	"note/src/util"
	"regexp"
)

var salt, _ = bcrypt.Salt(10)
var validate *validator.Validate

func init() {
	validate = validator.New()
}

type LoginUser struct {
	Uname    string `gorm:"type:varchar(50);not null;unique_index" json:"uname" validate:"required,min=1,max=50"`
	Password string `gorm:"type:varchar(200);not null" json:"password" validate:"required,min=1,max=50"`
}

type User struct {
	BaseModel
	LoginUser
	Mobile string  `gorm:"type:varchar(11);not null"json:"mobile" validate:"required"`
	Email  string  `gorm:"type:varchar(200);not null"json:"email" validate:"required,email"`
	Motto  string  `gorm:"type:varchar(50)"json:"motto"`
	Role   []*Role `gorm:"many2many:user_role;"`
}

type RegisterUser struct {
	User
	PasswordRepeat string `json:"passwordRepeat" validate:"required,eqfield=Password"`
}

// login user
func (loginUser LoginUser) Verify() (err error) {
	errs := validate.Struct(loginUser)
	if errs != nil && len(errs.(validator.ValidationErrors)) != 0 {
		return errs
	}
	return nil
}

// 数据库 user
func (user *User) CreateUser() error {
	hash, _ := bcrypt.Hash(user.Password, salt)
	user.Password = hash
	if err := DB.Create(user).Error; err != nil {
		return err
	}
	// 给用户一个默认角色
	DB.Model(user).Association("role").Append(Role{
		Name: "normal",
	})
	return nil
}

func (loginUser *LoginUser) Login() (string, error) {
	var result User
	if result := DB.Where("uname = ?", loginUser.Uname).First(&result); result.Error != nil {
		return "", result.Error
	}
	if !bcrypt.Match(loginUser.Password, result.Password) {
		return "", errors.New("密码不匹配")
	}
	return util.GetToken(result.ID), nil
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

func (user *User) GetUserById(uid int) User {
	var res User
	DB.First(&res, 10)
	return res
}

// register user
func (user RegisterUser) Verify() (err error) {
	validate.RegisterStructValidation(UserStructLevelValidation, RegisterUser{})
	errs := validate.Struct(user)
	if errs != nil && len(errs.(validator.ValidationErrors)) != 0 {
		println(errs.Error())
		return errs
	}
	return nil
}

func UserStructLevelValidation(sl validator.StructLevel) {
	user := sl.Current().Interface().(RegisterUser)
	reg := regexp.MustCompile("^1[0-9]{10}$")
	if !reg.MatchString(user.Mobile) {
		sl.ReportError(user.Mobile, "Mobile", "user.Mobile", "请输入正确的手机号", "")
	}
}
