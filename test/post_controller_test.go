package test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	_ "github.com/GoAdminGroup/go-admin/adapter/gin"              // web framework adapter
	_ "github.com/GoAdminGroup/go-admin/modules/db/drivers/mysql" // sql driver
	_ "github.com/GoAdminGroup/themes/adminlte"                   // ui theme
	"github.com/beslow/goblog/initialize"
	"github.com/beslow/goblog/models"
	"github.com/beslow/goblog/router"
	"github.com/stretchr/testify/assert"
)

func TestPostIndex(t *testing.T) {
	prepareTestDatabase()

	router.TemplateFs = f
	app, _ := router.SetRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/blog", nil)
	app.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)

	var post models.Post
	initialize.DB.First(&post)
	assert.Contains(t, w.Body.String(), post.Title)

	var category models.Category
	initialize.DB.Find(&category)
	assert.Contains(t, w.Body.String(), category.Name)
}

func TestPostShow(t *testing.T) {
	prepareTestDatabase()

	router.TemplateFs = f
	app, _ := router.SetRouter()

	w := httptest.NewRecorder()

	var post models.Post
	initialize.DB.Where("id = ?", 1).Find(&post)

	req, _ := http.NewRequest("GET", fmt.Sprintf("/blog/%s", post.HashID()), nil)
	app.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Contains(t, w.Body.String(), post.Title)
	assert.Contains(t, w.Body.String(), fmt.Sprintf("%d commnets", post.CommentCount))
}

func TestPostShow404(t *testing.T) {
	prepareTestDatabase()

	router.TemplateFs = f
	app, _ := router.SetRouter()

	w := httptest.NewRecorder()

	req, _ := http.NewRequest("GET", fmt.Sprintf("/blog/%s", "wrong-id"), nil)
	app.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)

	assert.Contains(t, w.Body.String(), "404")
	assert.Contains(t, w.Body.String(), "Back to Home Page")
}

func TestPostComment(t *testing.T) {
	prepareTestDatabase()

	router.TemplateFs = f
	app, _ := router.SetRouter()

	w := httptest.NewRecorder()

	var post models.Post
	initialize.DB.Where("id = ?", 1).Find(&post)

	originCommentCount := post.CommentCount

	data := url.Values{"name": {"test-guest"}, "email": {"test@test.com"}, "body": {"good job"}}
	body := strings.NewReader(data.Encode())

	req, _ := http.NewRequest("POST", fmt.Sprintf("/blog/%s/comments", post.HashID()), body)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	app.ServeHTTP(w, req)

	assert.Equal(t, http.StatusFound, w.Code)

	currentCount, _ := post.GetCommentCount()
	assert.Equal(t, originCommentCount+1, currentCount)
}
