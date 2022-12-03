package main

import (
	"embed"
	"flag"
	"html/template"
	"net/http"
	"os"

	"github.com/beslow/goblog/controller/home"
	"github.com/beslow/goblog/controller/resume"
	"github.com/beslow/goblog/db/migrations"
	"github.com/beslow/goblog/db/seed"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	ginlogrus "github.com/toorop/gin-logrus"
)

//go:embed views/* static/*
var f embed.FS

func init() {
	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.JSONFormatter{})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	log.SetLevel(log.WarnLevel)
}

var migrate = flag.Bool("db:migrate", false, "execute the db/migrations")
var dbSeed = flag.Bool("db:seed", false, "generate seed data")

func main() {
	flag.Parse()

	if *migrate {
		migrations.Migrate()
		return
	}

	if *dbSeed {
		seed.Seed()
		return
	}

	router := gin.Default()

	log := log.New()

	router.Use(ginlogrus.Logger(log), gin.Recovery())

	router.StaticFS("/public", http.FS(f))

	templ := template.Must(template.New("").ParseFS(f, "views/*.html"))
	router.SetHTMLTemplate(templ)

	home.Index(router)

	resume.Resume(router)

	router.GET("/portfolio", func(c *gin.Context) {
		c.HTML(http.StatusOK, "views/portfolio.html", gin.H{})
	})

	router.GET("/contacts", func(c *gin.Context) {
		c.HTML(http.StatusOK, "views/contacts.html", gin.H{})
	})

	router.GET("/blog", func(c *gin.Context) {
		c.HTML(http.StatusOK, "views/blog.html", gin.H{})
	})

	router.Run(":8080")
}
