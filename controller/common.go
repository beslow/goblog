package controller

import (
	"github.com/beslow/goblog/initialize"
	"github.com/beslow/goblog/models"
	"github.com/gin-gonic/gin"
	"github.com/ungerik/go-gravatar"
)

func LayoutData() *gin.H {
	return &gin.H{
		"realName":   models.GetConst("real_name"),
		"profession": models.GetConst("profession"),
		"avatar":     gravatar.Url(models.GetConst("email")),
	}
}

func BlogSideBarData() *gin.H {
	var categories []models.Category
	initialize.DB.Find(&categories)
	var categoriesWithNum = make([]*models.CategoryWithNum, 0, 10)
	for _, c := range categories {
		var count int64
		initialize.DB.Table("posts").Where("category_id = ?", c.ID).Count(&count)
		categoriesWithNum = append(categoriesWithNum, &models.CategoryWithNum{c, count})
	}

	var lastPosts []models.Post
	initialize.DB.Order("created_at desc").Limit(3).Find(&lastPosts)

	return &gin.H{
		"categories": categoriesWithNum,
		"last_posts": lastPosts,
	}
}
