package query

import (
	"github.com/graphql-go/graphql"
	"github.com/local/go-graphql/application/types"
)

var RootQuery = graphql.NewObject(graphql.ObjectConfig{
	Name: "Query",
	Fields: graphql.Fields{
		"Product": &graphql.Field{
			Type: graphql.NewList(types.ProductTypes),
			Args: graphql.FieldConfigArgument{
				"ID_PRO": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
			},
			Resolve: ProductResolve,
		},
	},
})
