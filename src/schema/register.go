package schema

import (
	"fmt"
	"github.com/graphql-go/graphql"
	"note/src/models"
)

var registerInputType = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "registerInputType",
	Fields: graphql.InputObjectConfigFieldMap{
		"uname": &graphql.InputObjectFieldConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"password": &graphql.InputObjectFieldConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"passwordRepeat": &graphql.InputObjectFieldConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"mobile": &graphql.InputObjectFieldConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"email": &graphql.InputObjectFieldConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"motto": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
	},
	Description: "用户 input 类型",
})

// 定义查询对象的字段，支持嵌套
var registerOutputType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "Hello",
	Description: "Hello Model",
	Fields: graphql.Fields{
		"uname": &graphql.Field{
			Type: graphql.String,
		},
		"mobile": &graphql.Field{
			Type: graphql.String,
		},
		"email": &graphql.Field{
			Type: graphql.String,
		},
	},
})

// 处理查询请求
var queryRegister = graphql.Field{
	Name:        "QueryRegister",
	Description: "Query Register",
	Type:        registerOutputType,
	Args: graphql.FieldConfigArgument{
		"user": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(registerInputType),
		},
	},
	// Resolve是一个处理请求的函数，具体处理逻辑可在此进行
	Resolve: func(p graphql.ResolveParams) (result interface{}, err error) {
		//ctx := p.Context.Value("ctx").(myContext.MyContext)
		inputUser := p.Args["user"].(map[string]interface{})
		user := models.User{
			LoginUser: models.LoginUser{
				Uname:    inputUser["uname"].(string),
				Password: inputUser["password"].(string),
			},
			Mobile: inputUser["mobile"].(string),
			Email:  inputUser["email"].(string),
			Motto:  inputUser["motto"].(string),
		}
		// 解析
		registerUser := models.RegisterUser{
			User:           user,
			PasswordRepeat: inputUser["passwordRepeat"].(string),
		}
		// 校验
		if err := registerUser.Verify(); err != nil {
			return nil, err
		}
		//ctx.JSON(user)
		// 密码加盐
		// 入库
		if err := user.CreateUser(); err != nil {
			return nil, err
		}
		// 调用Hello这个model里面的Query方法查询数据
		//return user, errors.New("自定义错误")
		fmt.Println(user.Uname)
		return map[string]string{
			"uname":  registerUser.Uname,
			"email":  registerUser.Email,
			"mobile": registerUser.Mobile,
		}, nil
	},
}
