package query

import (
	"github.com/cakazies/go-graphql/utils"

	"github.com/graphql-go/graphql"
	conf "github.com/cakazies/go-graphql/application/config"
	"github.com/cakazies/go-graphql/application/types"
)

func UsersResolve(params graphql.ResolveParams) (interface{}, error) {
	var a types.Users
	var b []types.Users
	b = b[:0]
	result, err := conf.DB.Query("SELECT id,name,age,profession,friendly FROM users WHERE deleted_at IS NULL")
	utils.PanicError(err, "Failed query select users")

	for result.Next() {
		err = result.Scan(&a.Id, &a.Name, &a.Age, &a.Profession, &a.Friendly)
		utils.PanicError(err, "Failed Scan in Query select users")
		b = append(b, a)
	}
	return b, nil
}
