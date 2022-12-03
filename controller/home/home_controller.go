package home

import (
	"net/http"

	"github.com/beslow/goblog/controller"
	"github.com/beslow/goblog/models/consts"
	"github.com/gin-gonic/gin"
	"golang.org/x/exp/maps"
)

func Index(router *gin.Engine) {
	homeFunc := func(c *gin.Context) {
		data := gin.H{
			"age":   consts.GetConst("age"),
			"city":  consts.GetConst("city"),
			"hobby": consts.GetConst("hobby"),
		}
		maps.Copy(data, *controller.LayoutData())
		c.HTML(http.StatusOK, "views/index.html", data)
	}

	router.GET("/", homeFunc)
	router.GET("/index", homeFunc)
}
