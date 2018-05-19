package blog

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"gin-blog/conf"
)

type MainController struct {
	baseController
}

func (mc *MainController) Index(c *gin.Context)  {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"siteName": conf.AppConfig.SiteName,
	})
}