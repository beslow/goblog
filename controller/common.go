package controller

import (
	"github.com/beslow/goblog/db"
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
	db.DB.Find(&categories)
	var categoriesWithNum = make([]*models.CategoryWithNum, 0, 10)
	for _, c := range categories {
		var count int64
		db.DB.Table("posts").Where("category_id = ?", c.ID).Count(&count)
		categoriesWithNum = append(categoriesWithNum, &models.CategoryWithNum{c, count})
	}

	var lastPosts []models.Post
	db.DB.Order("created_at desc").Limit(3).Find(&lastPosts)

	return &gin.H{
		"categories": categoriesWithNum,
		"last_posts": lastPosts,
	}
}
