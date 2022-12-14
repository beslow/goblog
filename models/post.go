package models

import (
	"fmt"
	"time"

	"github.com/gomodule/redigo/redis"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"github.com/beslow/goblog/initialize"
)

type Post struct {
	ID         int
	Title      string
	Summary    string
	Body       string
	VisitCount int
	CreatedAt  time.Time

	Comments     []Comment
	CommentCount int64 `gorm:"-:all"` // ignore CommentCount field when write, read and migrate with struct
}

// used for post's public url
func (post *Post) HashID() string {
	h, _ := HashID.Encode([]int{post.ID})

	return h
}

func (post *Post) GetAllComments() (comments []Comment, err error) {
	err = initialize.DB.Debug().Where("post_id = ?", post.ID).Find(&comments).Error
	return
}

// get post's comment count from redis or query database and cache it
func (post Post) GetCommentCount() (count int64, err error) {
	key := post.keyCommentCount()

	conn := initialize.RedisPool.Get()
	defer conn.Close()

	var val interface{}

	// first query comment count from redis
	if val, err = conn.Do("get", key); err != nil {
		return
	}

	if val == nil {
		// never cache in redis
		if err = initialize.DB.Table("comments").Where("post_id = ?", post.ID).
			Count(&count).Error; err != nil {
			return
		}

		// cache the comment count in redis
		if _, err = conn.Do("set", key, count); err != nil {
			return
		}
	} else {
		// fetch comment count from redis
		count, err = redis.Int64(val, err)
		if err != nil {
			return
		}
	}

	return
}

func (post *Post) keyCommentCount() string {
	return fmt.Sprintf("post#%v#comment_count", post.ID)
}

// delete post's comment count value from redis
func (post Post) DeleteCacheCommentCount() {
	key := post.keyCommentCount()
	var err error

	conn := initialize.RedisPool.Get()
	defer conn.Close()
	if _, err = conn.Do("del", key); err != nil {
		log.Panic(err)
	}
}

// add post's visit count
func (post Post) IncreaseVisitCount() {
	err := initialize.DB.Model(&post).
		Clauses(clause.Locking{Strength: "UPDATE"}). // Locking (FOR UPDATE)
		Update("visit_count", post.VisitCount+1).    // only update visit_count field
		Error
	if err != nil {
		fmt.Printf("err: %#v.\n\n", err)
	}
}

// after find post, set the post's CommentCount attribute
func (p *Post) AfterFind(tx *gorm.DB) (err error) {
	p.CommentCount, err = p.GetCommentCount()
	return err
}
