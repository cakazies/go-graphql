package mutation

import (
	"github.com/graphql-go/graphql"
	conf "github.com/local/go-graphql/application/config"
)

func CreateProductMutation(param graphql.ResolveParams) (interface{}, error) {
	idpro := param.Args["ID_PRO"].(string)
	name := param.Args["PRO_NAME"].(string)
	qty := param.Args["QTE_IN_STOCK"].(string)
	image := param.Args["PRO_IMAGE"].(string)

	_, err := conf.DB.Query("INSERT INTO products values (?,?,?,?)", idpro, name, qty, image)
	if err != nil {
		panic(err.Error())
	}
	return nil, err
}
