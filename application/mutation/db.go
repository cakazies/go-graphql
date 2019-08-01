package mutation

import (
	"fmt"
	"log"

	"github.com/graphql-go/graphql"
	conf "github.com/local/go-graphql/application/config"
)

func CreateProductMutation(param graphql.ResolveParams) (interface{}, error) {
	idpro := param.Args["ID_PRO"].(string)
	name := param.Args["PRO_NAME"].(string)
	qty := param.Args["QTE_IN_STOCK"].(int)
	image := param.Args["PRO_IMAGE"].(string)
	sql := fmt.Sprintf("INSERT INTO products values ('%s','%s',%v,'%s')", idpro, name, qty, image)
	_, err := conf.DB.Query(sql)
	if err != nil {
		panic(err.Error())
	}
	return nil, err
}

func CreateUsers(params graphql.ResolveParams) (interface{}, error) {
	name := params.Args["name"]
	age := params.Args["age"]
	profession := params.Args["profession"]
	friendly := params.Args["friendly"]
	// sql := fmt.Sprintf("UPDATE users set name='%s', age=%v, profession='%s', friendly='%s' WHERE id = 1 ;", name, age, profession, friendly)
	sql := fmt.Sprintf("INSERT INTO users (name,age,profession,friendly) values ('%s',%v,'%s',%v)", name, age, profession, friendly)
	_, err := conf.DB.Query(sql)
	if err != nil {
		panic(err.Error())
	}
	return nil, err
}

func UpdateUsers(params graphql.ResolveParams) (interface{}, error) {
	// id := params.Args["id"]
	name := params.Args["name"]
	age := params.Args["age"]
	profession := params.Args["profession"]
	friendly := params.Args["friendly"]
	sql := fmt.Sprintf("UPDATE users set name='%s', age=%v, profession='%s', friendly='%s' WHERE id = %s ;", name, age, profession, friendly, "1")
	log.Println(sql)
	_, err := conf.DB.Query(sql)
	if err != nil {
		panic(err.Error())
	}
	return nil, err
}
