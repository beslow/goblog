package resume

import (
	"net/http"

	"github.com/beslow/goblog/controller"
	"github.com/beslow/goblog/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/exp/maps"
)

func Resume(router *gin.Engine) {
	resumeFunc := func(c *gin.Context) {
		data := gin.H{
			"history_works":      models.GetAllHistoryWorks(),
			"history_educations": models.GetAllHistoryEducations(),
		}
		maps.Copy(data, *controller.LayoutData())

		c.HTML(http.StatusOK, "views/resume.html", data)
	}

	router.GET("/resume", resumeFunc)
	router.GET("/", resumeFunc)
}
