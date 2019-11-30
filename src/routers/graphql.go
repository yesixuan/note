package routers

import (
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
	_ "github.com/graphql-go/handler"
	"github.com/kataras/iris/v12"
)

func GraphqlRoute(graphqlRouter iris.Party) {
	graphqlRouter.Get("/", graphqlHandler())
	graphqlRouter.Post("/", graphqlHandler())
}

func graphqlHandler() iris.Handler {
	h := handler.New(&handler.Config{
		Schema:     &schema,
		Pretty:     true,
		GraphiQL:   true,
		Playground: true,
	})

	// 只需要通过Gin简单封装即可
	return func(c iris.Context) {
		h.ServeHTTP(c.ResponseWriter(), c.Request())
	}
}

type user struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

var data = map[string]user{
	"1": {"1", "Dan"},
	"2": {"2", "Lee"},
	"3": {"3", "Nick"},
}

var userType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "User",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.String,
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)

var queryType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"user": &graphql.Field{
				Type: userType,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					idQuery, isOK := p.Args["id"].(string)
					if isOK {
						return data[idQuery], nil
					}
					return nil, nil
				},
			},
		},
	})

var schema, _ = graphql.NewSchema(
	graphql.SchemaConfig{
		Query:    queryType,
		Mutation: queryType,
	},
)
