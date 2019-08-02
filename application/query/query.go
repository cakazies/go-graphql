package query

import (
	"github.com/graphql-go/graphql"
	"github.com/local/go-graphql/application/types"
)

var QueryUsers = graphql.NewObject(graphql.ObjectConfig{
	Name: "Queryopolek",
	Fields: graphql.Fields{
		"Users": &graphql.Field{
			Type: graphql.NewList(types.UsersTypes),
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
			},
			Resolve: UsersResolve,
		},
	},
})
