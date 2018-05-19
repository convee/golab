package route

import (
	"github.com/gin-gonic/gin"
	"gin-demo/apis"
)

func InitRouter() *gin.Engine  {
	r := gin.Default()
	r.GET("/", apis.Index)
	r.POST("/user/add", apis.AddUser)
	return r
}