package middleware

import (
	"github.com/beslow/goblog/controller"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func ErrorHandler(logger *logrus.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		// for _, err := range c.Errors {
		// 	switch err.Err {
		// 	case ErrNotFound:
		// 		c.JSON(-1, gin.H{"error": ErrNotFound.Error()})
		// 	}
		// 	etc...
		// }

		if len(c.Errors) > 0 {
			render500page(c)
		}
	}
}

func RecoverHandle(c *gin.Context, err any) {
	if err != nil {
		render500page(c)
	}
}

func render500page(c *gin.Context) {
	c.HTML(500, "views/500.html", *controller.LayoutData())
}
