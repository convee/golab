package main

import (
	"gin-demo/db"
	"gin-demo/route"
)

func main() {
	defer db.SqlDB.Close()
	router := route.InitRouter()
	router.Run(":8000")
}