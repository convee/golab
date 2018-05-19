package main

import (
	"gin-blog/conf"
	"github.com/gin-gonic/gin"
	"gin-blog/app/blog"
)

func init()  {
	conf.InitConfig()
}

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("./view/*")
	//r.Static("/", "static")
	mc := new (blog.MainController)
	r.GET("/", mc.Index)
	r.Run(":8000")
}