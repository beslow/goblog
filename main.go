package main

import (
	"embed"
	"flag"
	"html/template"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	_ "github.com/GoAdminGroup/go-admin/adapter/gin"              // web framework adapter
	_ "github.com/GoAdminGroup/go-admin/modules/db/drivers/mysql" // sql driver
	_ "github.com/GoAdminGroup/themes/adminlte"                   // ui theme

	"github.com/GoAdminGroup/go-admin/engine"
	go_admin_template "github.com/GoAdminGroup/go-admin/template"
	"github.com/GoAdminGroup/go-admin/template/chartjs"
	"github.com/beslow/goblog/controller/post"
	"github.com/beslow/goblog/controller/resume"
	"github.com/beslow/goblog/db"
	"github.com/beslow/goblog/db/migrations"
	"github.com/beslow/goblog/db/seed"
	"github.com/beslow/goblog/helpers"
	"github.com/beslow/goblog/middleware"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	ginlogrus "github.com/toorop/gin-logrus"

	"github.com/beslow/goblog/pages"
	"github.com/beslow/goblog/tables"
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

var (
	// ./goblog -db:create
	// create database if not exist
	dbCreate = flag.Bool("db:create", false, "create database")

	// ./goblog -db:migrate
	// execute the db migrations in the db/migrations dir
	migrate = flag.Bool("db:migrate", false, "execute the db/migrations")

	// ./goblog -db:seed
	// seed data, it wil skip if exist any data in the table
	dbSeed = flag.Bool("db:seed", false, "generate seed data")
)

func main() {
	gin.SetMode(gin.ReleaseMode)

	flag.Parse()

	supportDBCreateCommand()
	db.Init()
	supportDBMigrate()
	supportDBSeed()

	router := gin.Default()

	router.Use(middleware.CountVisit())

	go_admin_template.AddComp(chartjs.NewChart())

	eng := engine.Default()

	if err := eng.AddConfigFromYAML("./config.yml").
		AddGenerators(tables.Generators).
		Use(router); err != nil {
		panic(err)
	}

	eng.HTML("GET", "/admin", pages.GetDashBoard)

	log := log.New()

	router.Use(ginlogrus.Logger(log), gin.Recovery())

	router.StaticFS("/public", http.FS(f))

	var funcMaps = template.FuncMap{
		"toHashID":     helpers.ToHashID,
		"formatAsDate": helpers.FormatAsDate,
		"avatarURL":    helpers.AvatarURL,
	}

	templ := template.Must(template.New("").Funcs(funcMaps).ParseFS(f, "views/*.html"))
	router.SetHTMLTemplate(templ)

	// routes here
	resume.Resume(router)
	post.PostIndex(router)
	post.PostShow(router)
	post.PostComment(router)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	go releaseResource(quit, eng)

	router.Run(":80")
}

func supportDBCreateCommand() {
	if *dbCreate {
		db.CreateDatabase()
		os.Exit(0)
	}
}

func supportDBMigrate() {
	if *migrate {
		migrations.Migrate()
		os.Exit(0)
	}
}

func supportDBSeed() {
	if *dbSeed {
		seed.Seed()
		os.Exit(0)
	}
}

func releaseResource(quit chan os.Signal, eng *engine.Engine) {
	for s := range quit {
		switch s {
		case syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT:
			log.Print("closing database connection")
			eng.MysqlConnection().Close()

			log.Print("close redis connection")
			db.RedisPool.Close()

			os.Exit(0)
		default:
		}
	}
}
