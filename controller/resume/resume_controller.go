package resume

import (
	"net/http"

	"github.com/beslow/goblog/controller"
	"github.com/gin-gonic/gin"
	"golang.org/x/exp/maps"
)

func Resume(router *gin.Engine) {
	router.GET("/resume", func(c *gin.Context) {
		data := gin.H{}
		maps.Copy(data, *controller.LayoutData())

		c.HTML(http.StatusOK, "views/resume.html", data)
	})
}
