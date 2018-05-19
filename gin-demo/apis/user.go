package apis

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"gin-demo/models"
	"log"
	"fmt"
)

func Index(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"error":0,
		"msg":"Ok",
		"data":"",
	})
}

func AddUser(c *gin.Context)  {
	username := c.Request.FormValue("username")
	u := models.User{Username:username}
	result, err := u.AddUser()
	if err != nil {
		log.Fatalln(err)
	}
	msg := fmt.Sprintf("insert user successful %d", result)
	c.JSON(http.StatusOK, gin.H{
		"error":0,
		"msg":msg,
		"data":"",
	})
}