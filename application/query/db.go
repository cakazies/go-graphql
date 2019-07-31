package query

import (
	"log"

	"github.com/graphql-go/graphql"
	"github.com/import/Golang-mysql-graphql/schema/types"
	conf "github.com/local/go-graphql/application/config"
)

func ProductResolve(params graphql.ResolveParams) (interface{}, error) {
	var a types.Product
	var b []types.Product

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
