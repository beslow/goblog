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
	"github.com/beslow/goblog/logger"
	"github.com/beslow/goblog/middleware"
	"github.com/beslow/goblog/pages"
	"github.com/beslow/goblog/tables"
	sentrygin "github.com/getsentry/sentry-go/gin"
	"github.com/gin-gonic/gin"
	ginlogrus "github.com/toorop/gin-logrus"
)

var TemplateFs embed.FS

var (
	r   *gin.Engine
	eng *engine.Engine
)

func SetRouter() (*gin.Engine, *engine.Engine) {
	if r != nil && eng != nil {
		return r, eng
	}

	r = gin.Default()

	go_admin_template.AddComp(chartjs.NewChart())

	eng = engine.Default()

	if err := eng.AddConfigFromYAML("./config.yml").
		AddGenerators(tables.Generators).
		Use(r); err != nil {
		panic(err)
	}

	eng.HTML("GET", "/admin", pages.GetDashBoard)

	r.Use(ginlogrus.Logger(logger.Log), gin.CustomRecovery(middleware.RecoverHandle))

	r.Use(middleware.CountVisit())

	if initialize.GetSentryDsn() != "" {
		r.Use(sentrygin.New(sentrygin.Options{Repanic: true})) // repanic for custom recovery
	}

	r.Use(middleware.ErrorHandler(logger.Log))

	r.StaticFS("/public", http.FS(TemplateFs))

	var funcMaps = template.FuncMap{
		"toHashID":     helpers.ToHashID,
		"formatAsDate": helpers.FormatAsDate,
		"avatarURL":    helpers.AvatarURL,
	}

	templ := template.Must(template.New("").Funcs(funcMaps).ParseFS(TemplateFs, "views/*.html"))
	r.SetHTMLTemplate(templ)

	// routes here
	resume.Resume(r)
	post.PostIndex(r)
	post.PostShow(r)
	post.PostComment(r)

	return r, eng
}
