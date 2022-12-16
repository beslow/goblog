package router

import (
	"embed"
	"html/template"
	"net/http"

	"github.com/GoAdminGroup/go-admin/engine"
	go_admin_template "github.com/GoAdminGroup/go-admin/template"
	"github.com/GoAdminGroup/go-admin/template/chartjs"
	"github.com/beslow/goblog/controller/post"
	"github.com/beslow/goblog/controller/resume"
	"github.com/beslow/goblog/helpers"
	"github.com/beslow/goblog/initialize"
	"github.com/beslow/goblog/middleware"
	"github.com/beslow/goblog/pages"
	"github.com/beslow/goblog/tables"
	sentrygin "github.com/getsentry/sentry-go/gin"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	ginlogrus "github.com/toorop/gin-logrus"
)

var TemplateFs embed.FS

var (
	router *gin.Engine
	eng    *engine.Engine
)

func SetRouter() (*gin.Engine, *engine.Engine) {
	if router != nil && eng != nil {
		return router, eng
	}

	router = gin.Default()

	go_admin_template.AddComp(chartjs.NewChart())

	eng = engine.Default()

	if err := eng.AddConfigFromYAML("./config.yml").
		AddGenerators(tables.Generators).
		Use(router); err != nil {
		panic(err)
	}

	eng.HTML("GET", "/admin", pages.GetDashBoard)

	log := log.New()

	router.Use(ginlogrus.Logger(log), gin.Recovery())
	router.Use(middleware.CountVisit())
	if initialize.GetSentryDsn() != "" {
		router.Use(sentrygin.New(sentrygin.Options{}))
	}

	router.StaticFS("/public", http.FS(TemplateFs))

	var funcMaps = template.FuncMap{
		"toHashID":     helpers.ToHashID,
		"formatAsDate": helpers.FormatAsDate,
		"avatarURL":    helpers.AvatarURL,
	}

	templ := template.Must(template.New("").Funcs(funcMaps).ParseFS(TemplateFs, "views/*.html"))
	router.SetHTMLTemplate(templ)

	// routes here
	resume.Resume(router)
	post.PostIndex(router)
	post.PostShow(router)
	post.PostComment(router)

	return router, eng
}
