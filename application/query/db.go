package query

import (
	"log"

	"github.com/graphql-go/graphql"
	conf "github.com/local/go-graphql/application/config"
	"github.com/local/go-graphql/application/types"
)

func ProductResolve(params graphql.ResolveParams) (interface{}, error) {
	var a types.Products
	var b []types.Products

	b = b[:0]
	result, err := conf.DB.Query("select ID_PRO,PRO_NAME,QTE_IN_STOCK,PRO_IMAGE from products")
	if err != nil {
		log.Println("errornya gan hehe:", err)
		panic(err.Error())
	}
	for result.Next() {
		err = result.Scan(&a.IdPro, &a.ProName, &a.QteStock, &a.ProImg)
		if err != nil {
			panic(err.Error())
		}
		b = append(b, a)
	}
	return b, nil
}

func UsersResolve(params graphql.ResolveParams) (interface{}, error) {
	var a types.Users
	var b []types.Users
	b = b[:0]
	result, err := conf.DB.Query("select id,name,age,profession,friendly from users")
	if err != nil {
		panic(err.Error())
	}
	for result.Next() {
		err = result.Scan(&a.Id, &a.Name, &a.Age, &a.Profession, &a.Friendly)
		if err != nil {
			panic(err.Error())
		}
		b = append(b, a)
	}
	return b, nil
}
