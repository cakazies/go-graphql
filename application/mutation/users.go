package mutation

import (
	"fmt"
	"time"

	"github.com/graphql-go/graphql"
	conf "github.com/cakazies/go-graphql/application/config"
	"github.com/cakazies/go-graphql/utils"
)

func UsersRevolver(params graphql.ResolveParams) (interface{}, error) {
	id := params.Args["id"]
	name := params.Args["name"]
	age := params.Args["age"]
	profession := params.Args["profession"]
	friendly := params.Args["friendly"]
	createdAt := time.Now().Format("2006-01-02 15:04:05")
	deletedAt := params.Args["deleted_at"]
	sql := fmt.Sprintf("INSERT INTO users (name,age,profession,friendly, created_at) values ('%s',%v,'%s',%v, '%s')", name, age, profession, friendly, createdAt)

	if id != nil && deletedAt == nil {
		updatedAt := time.Now().Format("2006-01-02 15:04:05")
		sql = fmt.Sprintf("UPDATE users SET name='%s', age=%v, profession='%s', friendly='%s', updated_at = '%s' WHERE id = %v ;", name, age, profession, friendly, updatedAt, id)
	}

	if id != nil && deletedAt != nil {
		sql = fmt.Sprintf("UPDATE users SET deleted_at = '%s' WHERE id = %v ;", deletedAt, id)
	}

	_, err := conf.DB.Query(sql)
	utils.PanicError(err, "Failed query insert users")
	return nil, err
}
