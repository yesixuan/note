package routers

import (
	"context"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
	_ "github.com/graphql-go/handler"
	"github.com/kataras/iris/v12"
	"note/src/middlewares"
	"note/src/util"
)

func GraphqlRoute(graphqlRouter iris.Party) {
	graphqlRouter.Get("/", middlewares.GetJwtHandler(), graphqlHandler())
	graphqlRouter.Post("/", middlewares.GetJwtHandler(), graphqlHandler())
}

func graphqlHandler() iris.Handler {
	h := handler.New(&handler.Config{
		Schema:     &schema,
		Pretty:     true,
		GraphiQL:   true,
		Playground: true,
	})

	// 只需要通过 iris 简单封装即可
	return func(c iris.Context) {
		ctx := context.WithValue(context.Background(), "userId", util.GetUserId(c))
		h.ContextHandler(ctx, c.ResponseWriter(), c.Request())
	}
}

type user struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Vic struct {
	Emotion string `json:"emotion"`
	Happy   bool   `json:"happy"`
}

var data = map[string]user{
	"1": {"1", "Dan"},
	"2": {"2", "Lee"},
	"3": {"3", "Nick"},
}

var vicType = graphql.NewObject(graphql.ObjectConfig{
	Name: "vicType",
	Fields: graphql.Fields{
		"emotion": &graphql.Field{
			Type: graphql.String,
		},
		"happy": &graphql.Field{
			Type: graphql.Boolean,
		},
	},
})

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
			"vic": &graphql.Field{
				Type: vicType,
				Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
					return Vic{
						Emotion: "angry",
						Happy:   false,
					}, nil
				},
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
