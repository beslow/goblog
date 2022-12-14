package models

import (
	"time"

	"github.com/beslow/goblog/initialize"
)

type Comment struct {
	ID        int
	PostID    int
	Name      string
	Email     string
	Body      string
	CreatedAt time.Time

	Post Post
}

func CreateComment(post Post, name, email, body string) error {
	// delete the post's comment count value from redis
	// the value will regenerate at next visit
	post.DeleteCacheCommentCount()

	comment := Comment{
		Post:  post,
		Name:  name,
		Email: email,
		Body:  body,
	}

	return initialize.DB.Create(&comment).Error
}
