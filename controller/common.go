package controller

import (
	"github.com/beslow/goblog/models/consts"
	"github.com/gin-gonic/gin"
	"github.com/ungerik/go-gravatar"
)

func LayoutData() *gin.H {
	return &gin.H{
		"realName":   consts.GetConst("real_name"),
		"profession": consts.GetConst("profession"),
		"avatar":     gravatar.Url(consts.GetConst("email")),
	}
}
