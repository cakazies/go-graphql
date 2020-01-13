package mutation

import (
	"github.com/cakazies/go-graphql/application/types"
	"github.com/graphql-go/graphql"
)

var UsersMutation = graphql.NewObject(graphql.ObjectConfig{
	Name: "UsersMutation",
	Fields: graphql.Fields{
		"CreateUser": &graphql.Field{
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
				"deleted_at": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
			},
			Resolve: UsersRevolver,
		},
	},
})
