package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/graphql-go/graphql"
	"github.com/local/go-graphql/application/mutation"
	"github.com/local/go-graphql/application/query"
)

func Route() {
	var Schame, err = graphql.NewSchema(graphql.SchemaConfig{
		Query:    query.RootQuery,
		Mutation: mutation.Mutation,
	})
	if err != nil {
		panic(err.Error())
	}
	http.HandleFunc("/graphql", func(w http.ResponseWriter, r *http.Request) {
		result := executeQuery(r.URL.Query().Get("Query"), Schame)
		json.NewEncoder(w).Encode(result)
	})
	http.ListenAndServe(":8088", nil)

	// r := mux.NewRouter()
	// routers := r.PathPrefix("/api").Subrouter()

	// // cek for middleware
	// // routers.Use(middleware.JwtAuthentication)

	// // modul rooms
	// // routers.Handle("/getrooms", HandlerFunc(api.GetRooms)).Methods(http.MethodGet)
	// // routers.Handle("/insertroom", HandlerFunc(api.InsertRooms)).Methods(http.MethodPost)
	// // routers.Handle("/getroom/{rm_id}", HandlerFunc(api.GetRoom)).Methods(http.MethodGet)
	// // routers.Handle("/updateroom/{rm_id}", HandlerFunc(api.UpdateRooms)).Methods(http.MethodPost)
	// // routers.Handle("/deleteroom/{rm_id}", HandlerFunc(api.DeleteRoom)).Methods(http.MethodGet)
	// // routers.Handle("/deleteallroom", HandlerFunc(api.DeleteAllRoom)).Methods(http.MethodGet)

	// // // module borrow
	// // routers.Handle("/borrow", HandlerFunc(api.GetBorrows)).Methods(http.MethodGet)

	// // // modul users
	// // routers.Handle("/user/register", HandlerFunc(api.Register)).Methods(http.MethodPost)
	// // routers.Handle("/user/login", HandlerFunc(api.Login)).Methods(http.MethodPost)
	// // routers.Handle("/user", HandlerFunc(api.Me)).Methods(http.MethodGet)

	// host := fmt.Sprintf(viper.GetString("app.host"))

	// srv := &http.Server{
	// 	Handler:      routers,
	// 	Addr:         host,
	// 	WriteTimeout: 15 * time.Second,
	// 	ReadTimeout:  15 * time.Second,
	// }
	// log.Fatal(srv.ListenAndServe())
}

func executeQuery(query string, schema graphql.Schema) *graphql.Result {
	result := graphql.Do(graphql.Params{
		Schema:        schema,
		RequestString: query,
	})
	if len(result.Errors) > 0 {
		fmt.Printf("Wrong Result, unexpected errors : %v", result.Errors)
	}
	return result
}
