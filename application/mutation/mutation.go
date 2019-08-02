package mutation

import (
	"github.com/graphql-go/graphql"
	"github.com/local/go-graphql/application/types"
)

var MutationProd = graphql.NewObject(graphql.ObjectConfig{
	Name: "MutationProd",
	Fields: graphql.Fields{
		"CreateProducts": &graphql.Field{
			Type: graphql.NewList(types.ProductTypes),
			Args: graphql.FieldConfigArgument{
				"ID_PRO": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"PRO_NAME": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"QTE_IN_STOCK": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.Int),
				},
				"PRO_IMAGE": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
			},
			Resolve: CreateProductMutation,
		},
	},
})

var UsersMutation = graphql.NewObject(graphql.ObjectConfig{
	Name: "UsersMutation",
	Fields: graphql.Fields{
		"CreateUser": &graphql.Field{
			Type: graphql.NewList(types.UsersTypes),
			Args: graphql.FieldConfigArgument{
				"name": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"age": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.Int),
				},
				"profession": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"friendly": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
			},
			Resolve: CreateUsers,
		},
	},
})
var UserUpdate = graphql.NewObject(graphql.ObjectConfig{
	Name: "UserUpdate",
	Fields: graphql.Fields{
		"UpdateUsers": &graphql.Field{
			Type: graphql.NewList(types.UsersTypes),
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.Int),
				},
				"name": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"age": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.Int),
				},
				"profession": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"friendly": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
			},
			Resolve: UpdateUsers,
		},
	},
})
