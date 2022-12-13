package home

import (
	"net/http"

	"github.com/beslow/goblog/controller"
	"github.com/beslow/goblog/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/exp/maps"
)

func Index(router *gin.Engine) {
	homeFunc := func(c *gin.Context) {
		data := gin.H{
			"age":   models.GetConst("age"),
			"city":  models.GetConst("city"),
			"hobby": models.GetConst("hobby"),
		}
		maps.Copy(data, *controller.LayoutData())
		c.HTML(http.StatusOK, "views/index.html", data)
	}

	router.GET("/", homeFunc)
	router.GET("/index", homeFunc)
}
