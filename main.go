package main

import (
	"embed"
	"flag"
	"os"
	"os/signal"
	"syscall"

	_ "github.com/GoAdminGroup/go-admin/adapter/gin"              // web framework adapter
	_ "github.com/GoAdminGroup/go-admin/modules/db/drivers/mysql" // sql driver
	_ "github.com/GoAdminGroup/themes/adminlte"                   // ui theme
	"github.com/gin-gonic/gin"

	"github.com/GoAdminGroup/go-admin/engine"
	"github.com/beslow/goblog/db"
	"github.com/beslow/goblog/db/migrations"
	"github.com/beslow/goblog/db/seed"
	"github.com/beslow/goblog/initialize"
	"github.com/beslow/goblog/router"
	log "github.com/sirupsen/logrus"
)

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

//go:embed views/* static/*
var f embed.FS

func main() {
	gin.SetMode(gin.ReleaseMode)

	flag.Parse()

	supportDBCreateCommand()
	initialize.InitMySQL()
	supportDBMigrate()
	supportDBSeed()

	router.TemplateFs = f

	app, eng := router.SetRouter()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	go releaseResource(quit, eng)

	app.Run(":80")
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
			initialize.RedisPool.Close()

			os.Exit(0)
		default:
		}
	}
}
