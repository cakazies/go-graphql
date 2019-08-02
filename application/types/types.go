package types

import "github.com/graphql-go/graphql"

var UsersTypes = graphql.NewObject(graphql.ObjectConfig{
	Name: "Users",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.String,
		},
		"name": &graphql.Field{
			Type: graphql.String,
		},
		"age": &graphql.Field{
			Type: graphql.Int,
		},
		"profession": &graphql.Field{
			Type: graphql.String,
		},
		"friendly": &graphql.Field{
			Type: graphql.String,
		},
	},
})
