package test

import (
	"net/http"
	"net/http/httptest"
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
	initialize.DB.Find(&post)
	assert.Contains(t, w.Body.String(), post.Title)

	var category models.Category
	initialize.DB.Find(&category)
	assert.Contains(t, w.Body.String(), category.Name)
}
