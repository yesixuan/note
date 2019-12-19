package schema

import (
	"github.com/graphql-go/graphql"
	"note/src/models"
)

var loginInputType = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "loginInputType",
	Fields: graphql.InputObjectConfigFieldMap{
		"uname": &graphql.InputObjectFieldConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"password": &graphql.InputObjectFieldConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
	},
	Description: "登录用户 input 类型",
})

// 定义查询对象的字段，支持嵌套
var loginOutputType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "loginOutputType",
	Description: "Login Model",
	Fields: graphql.Fields{
		"token": &graphql.Field{
			Type: graphql.String,
		},
	},
})

// 处理查询请求
var queryLogin = graphql.Field{
	Name:        "QueryLogin",
	Description: "Query Login",
	Type:        loginOutputType,
	Args: graphql.FieldConfigArgument{
		"user": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(loginInputType),
		},
	},
	// Resolve是一个处理请求的函数，具体处理逻辑可在此进行
	Resolve: func(p graphql.ResolveParams) (result interface{}, err error) {
		//ctx := p.Context.Value("ctx").(myContext.MyContext)
		inputUser := p.Args["user"].(map[string]interface{})
		loginUser := models.LoginUser{
			Uname:    inputUser["uname"].(string),
			Password: inputUser["password"].(string),
		}
		if err := loginUser.Verify(); err != nil {
			return nil, err
		}
		token, err := loginUser.Login()
		if err != nil {
			return nil, err
		}
		return map[string]string{
			"token": token,
		}, nil
	},
}
