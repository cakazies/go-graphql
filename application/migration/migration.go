package main

import (
	"fmt"
	"log"
	"strconv"
	"time"

	cf "github.com/cakazies/go-graphql/application/config"
	"github.com/cakazies/go-graphql/cmd"
	"github.com/cakazies/go-graphql/utils"
)

// go run application/migration/migration.go
func main() {
	cmd.InitViper()
	var limit int
	limit = 10
	cf.Connect()
	MigrationUsers(limit)
	cf.DB.Close()
}

func MigrationUsers(limit int) {
	tableName := "users"
	drop := fmt.Sprintf("DROP TABLE IF EXISTS %s;", tableName)
	_, err := cf.DB.Query(drop)
	utils.FailError(err, fmt.Sprintf("Error Query Drop table %s", tableName))
	queryCreate := fmt.Sprintf(`
					CREATE TABLE public.%s
					(
						id SERIAL NOT NULL,
						name character varying(200) COLLATE pg_catalog."default" NOT NULL,
						age integer NOT NULL,
						profession character varying(100) COLLATE pg_catalog."default" NOT NULL,
						friendly character varying(100) COLLATE pg_catalog."default" NOT NULL,
						created_at timestamp without time zone	,
						updated_at timestamp without time zone ,
						deleted_at timestamp without time zone ,
						CONSTRAINT %s_pk PRIMARY KEY (id)
					);`, tableName, tableName)
	stmt, err := cf.DB.Prepare(queryCreate)
	utils.FailError(err, fmt.Sprintf("Error Create Table %s ", tableName))
	_, err = stmt.Exec()
	utils.FailError(err, fmt.Sprintf("Error Create Table %s ", tableName))
	log.Println(fmt.Sprintf("Import Table %s Succesfull", tableName))

	for i := 1; i <= limit; i++ {
		name := "name-" + strconv.Itoa(i)
		age := strconv.Itoa(i)
		profession := "profession-" + strconv.Itoa(i)
		friendly := "friendly-" + strconv.Itoa(i)
		createdAt := time.Now().Format("2006-01-02 15:04:05")
		sql := fmt.Sprintf("INSERT INTO %s ( name, age, profession, friendly, created_at) VALUES ('%s', %v, '%s', '%s', '%s'); ",
			tableName, name, age, profession, friendly, createdAt)
		stmt, err := cf.DB.Query(sql)
		utils.FailError(err, fmt.Sprintf("Error Insert Data Table %s ", tableName))
		stmt.Close()
		time.Sleep(time.Second / 10)
	}
	log.Println(fmt.Sprintf("Insert Data Dummy table %s successfull", tableName))
}
