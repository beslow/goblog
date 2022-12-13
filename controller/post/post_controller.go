package post

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"github.com/beslow/goblog/controller"
	"github.com/beslow/goblog/db"
	"github.com/beslow/goblog/models"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/vcraescu/go-paginator/v2"
	"github.com/vcraescu/go-paginator/v2/adapter"
	"github.com/vcraescu/go-paginator/v2/view"
	"golang.org/x/exp/maps"
)

func PostIndex(router *gin.Engine) {
	router.GET("/blog", func(c *gin.Context) {
		var posts []models.Post
		db.DB.Find(&posts)

		pageStr, _ := c.GetQuery("page")
		page, _ := strconv.Atoi(pageStr)
		if page == 0 {
			page = 1
		}

		p := paginator.New(adapter.NewGORMAdapter(db.DB.Table("posts")), 10)
		p.SetPage(page)

		if err := p.Results(&posts); err != nil {
			log.Panic(err)
		}

		view := view.New(p)
		pages, _ := view.Pages()
		next, _ := view.Next()
		prev, _ := view.Prev()
		current, _ := view.Current()

		data := gin.H{
			"posts":   posts,
			"pages":   pages,
			"current": current,
			"next":    next,
			"prev":    prev,
		}

		maps.Copy(data, *controller.LayoutData())

		maps.Copy(data, *controller.BlogSideBarData())

		c.HTML(http.StatusOK, "views/blog.html", data)
	})
}

func PostShow(router *gin.Engine) {
	router.GET("/blog/:hashid", func(c *gin.Context) {
		id, _ := models.HashID.DecodeWithError(c.Param("hashid"))

		var post models.Post
		if len(id) > 0 {
			db.DB.Debug().Where("id = ?", id[0]).Find(&post)
		}

		if post.ID == 0 {
			c.HTML(http.StatusOK, "views/404.html", controller.LayoutData())
		} else {
			if err := addVisitCount(post, c); err != nil {
				log.Panic(err)
			}

			comments, err := post.GetAllComments()
			if err != nil {
				log.Panic(err)
			}

			data := gin.H{
				"hashid":        post.HashID(),
				"created_at":    post.CreatedAt,
				"title":         post.Title,
				"body":          template.HTML(post.Body),
				"visit_count":   post.VisitCount,
				"comment_count": post.CommentCount,
				"comments":      comments,
			}

			maps.Copy(data, *controller.LayoutData())

			maps.Copy(data, *controller.BlogSideBarData())

			c.HTML(http.StatusOK, "views/blog-single.html", data)
		}
	})
}

func PostComment(router *gin.Engine) {
	router.POST("/blog/:hashid/comments", func(c *gin.Context) {
		hashid := c.Param("hashid")
		id, _ := models.HashID.DecodeWithError(hashid)

		var post models.Post
		if len(id) > 0 {
			db.DB.Where("id = ?", id[0]).Find(&post)
		}

		if post.ID == 0 {
			c.HTML(http.StatusOK, "views/404.html", controller.LayoutData())
		} else {
			models.CreateComment(post, c.PostForm("name"), c.PostForm("email"), c.PostForm("body"))

			c.Redirect(http.StatusFound, fmt.Sprintf("/blog/%v", hashid))
		}
	})
}

func addVisitCount(post models.Post, c *gin.Context) (err error) {
	ipAddr := c.ClientIP()
	if ipAddr != "" {
		conn := db.RedisPool.Get()
		defer conn.Close()

		key := fmt.Sprintf("post#%d#ip#%s", post.ID, ipAddr)
		var reply interface{}
		reply, err = conn.Do("get", key)
		if err != nil {
			return
		}

		if reply == nil {
			// first visit
			post.IncreaseVisitCount()
			if _, err = conn.Do("setex", key, 24*3600, 1); err != nil {
				return
			}
		}
	}

	return
}
