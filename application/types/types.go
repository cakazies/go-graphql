package types

import "github.com/graphql-go/graphql"

var ProductTypes = graphql.NewObject(graphql.ObjectConfig{
	Name: "Products",
	Fields: graphql.Fields{
		"ID_PRO": &graphql.Field{
			Type: graphql.String,
		},
		"PRO_NAME": &graphql.Field{
			Type: graphql.String,
		},
		"QTE_IN_STOCK": &graphql.Field{
			Type: graphql.Int,
		},
		"PRO_IMAGE": &graphql.Field{
			Type: graphql.String,
		},
	},
})

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
