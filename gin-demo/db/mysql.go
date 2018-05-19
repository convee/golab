package db

import (
	_ "github.com/Go-SQL-Driver/MySQL"
	"database/sql"
	"log"
)

var SqlDB * sql.DB

func init()  {
	var err error
	SqlDB, err = sql.Open("mysql", "root:wang1019@tcp(127.0.0.1:3306)/user?parseTime=true")
	if err != nil {
		log.Fatal(err.Error())
	}
	err = SqlDB.Ping()
	if err != nil {
		log.Fatal(err.Error())
	}
}