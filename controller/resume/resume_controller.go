package resume

import (
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/beslow/goblog/controller"
	"github.com/beslow/goblog/models"
	"github.com/beslow/goblog/rocketmq/produce"
	"github.com/gin-gonic/gin"
	"golang.org/x/exp/maps"
)

func Resume(router *gin.Engine) {
	resumeFunc := func(c *gin.Context) {
		var msg strings.Builder
		msg.WriteString(`{"action": "visit", "time": "`)
		msg.WriteString(strconv.Itoa(int(time.Now().Unix())))
		msg.WriteString(`"}`)
		produce.Do("resume", msg.String())

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
