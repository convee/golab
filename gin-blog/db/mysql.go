package db

import (
	_ "github.com/Go-SQL-Driver/MySQL"
	"database/sql"
	"log"
	"gin-blog/conf"
)

var SqlDB *sql.DB

func Init()  {
	var err error
	config := conf.AppConfig.DB
	dsn := config.User + ":" + config.Password + "@tcp(" + config.Host + ":" + config.Port + ")/" + config.Name + "?charset=utf8&loc=Asia%2FShanghai"
	SqlDB, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err.Error())
	}
	err = SqlDB.Ping()
	if err != nil {
		log.Fatal(err.Error())
	}
}