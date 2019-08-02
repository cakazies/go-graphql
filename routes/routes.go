package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/local/go-graphql/utils"

	"github.com/graphql-go/graphql"
	"github.com/local/go-graphql/application/mutation"
	"github.com/local/go-graphql/application/query"
)

func Route() {

	http.HandleFunc("/product", getProduct)
	http.HandleFunc("/users", getUsers)

	http.ListenAndServe(":8088", nil)
}

func getProduct(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query().Get("Query")
	var sch, err = graphql.NewSchema(graphql.SchemaConfig{
		Query:    query.RootQuery,
		Mutation: mutation.MutationProd,
	})
	utils.PanicError(err, "Failed New Schema")

	result := graphql.Do(graphql.Params{
		Schema:        sch,
		RequestString: params,
	})
	if len(result.Errors) > 0 {
		fmt.Printf("Wrong Result, unexpected errors : %v", result.Errors)
	}

	json.NewEncoder(w).Encode(result)
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query().Get("Query")
	var sch, err = graphql.NewSchema(graphql.SchemaConfig{
		Query:    query.QueryUsers,
		Mutation: mutation.UsersMutation,
	})
	utils.PanicError(err, "Failed New Schema")

	result := graphql.Do(graphql.Params{
		Schema:        sch,
		RequestString: params,
	})
	if len(result.Errors) > 0 {
		fmt.Printf("Wrong Result, unexpected errors : %v", result.Errors)
	}
	json.NewEncoder(w).Encode(result)
}
