package mutation

import (
	"github.com/graphql-go/graphql"
	"github.com/local/go-graphql/application/types"
)

var Mutation = graphql.NewObject(graphql.ObjectConfig{
	Name: "Mutation",
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
