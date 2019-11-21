package validators

import (
	"github.com/go-playground/validator/v10"
	"github.com/kataras/iris/v12"
	"regexp"
)

var validate *validator.Validate

type User struct {
	Uname          string `json:"uname" validate:"required,min=1,max=50"`
	Password       string `json:"password" validate:"required,min=1,max=50"`
	PasswordRepeat string `json:"passwordRepeat" validate:"required,eqfield=Password"`
	Mobile         string `json:"mobile" validate:"required"`
	Email          string `json:"email" validate:"required,email"`
	Motto          string `json:"motto"`
}

func init() {
	validate = validator.New()
}

func (user User) Verify(ctx iris.Context) {
	validate.RegisterStructValidation(UserStructLevelValidation, User{})
	errs := validate.Struct(user)
	if errs != nil && len(errs.(validator.ValidationErrors)) != 0 {
		ctx.StopExecution()
		ctx.StatusCode(iris.StatusUnauthorized)
		ctx.Values().Set("msg", errs.Error())
		println(errs.Error())
	}
}

func UserStructLevelValidation(sl validator.StructLevel) {
	user := sl.Current().Interface().(User)
	reg := regexp.MustCompile("^1[0-9]{10}$")
	if reg.MatchString(user.Mobile) {
		sl.ReportError(user.Mobile, "Mobile", "user.Mobile", "请输入正确的手机号", "")
	}
}
