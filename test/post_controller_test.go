package test

import (
	"embed"
	"net/http"
	"net/http/httptest"
	"testing"

	_ "github.com/GoAdminGroup/go-admin/adapter/gin"              // web framework adapter
	_ "github.com/GoAdminGroup/go-admin/modules/db/drivers/mysql" // sql driver
	_ "github.com/GoAdminGroup/themes/adminlte"                   // ui theme
	"github.com/beslow/goblog/initialize"
	"github.com/beslow/goblog/router"
	"github.com/stretchr/testify/assert"
)

//go:generate cp -r ../views .
//go:generate cp -r ../static .
//go:embed views/* static/**
var f embed.FS

func TestPostIndex(t *testing.T) {
	initialize.InitMySQL()
	router.TemplateFs = f
	app, _ := router.SetRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/blog", nil)
	app.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}
