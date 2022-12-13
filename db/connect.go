package db

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/GoAdminGroup/go-admin/modules/config"
	redis_config "github.com/beslow/goblog/config"
	"github.com/gomodule/redigo/redis"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB            *gorm.DB
	DefaultConfig config.Database
	RedisPool     *redis.Pool
)

func Init() {
	InitMySQL()
	InitRedis()
}

func InitMySQL() {
	dir, _ := os.Getwd()
	dir = strings.Replace(dir, "db/migrations", "", -1)
	dir = strings.Replace(dir, "/db", "", -1)

	cf := config.ReadFromYaml(filepath.Join(dir, "config.yml"))
	DefaultConfig = cf.Databases["default"]

	var err error
	DB, err = gorm.Open(mysql.Open(DefaultConfig.GetDSN()), &gorm.Config{})
	if err != nil {
		panic("open mysql connection fail.")
	}

	sqlDB, _ := DB.DB()

	sqlDB.SetMaxIdleConns(DefaultConfig.MaxIdleConns)
	sqlDB.SetMaxOpenConns(DefaultConfig.MaxOpenConns)
	sqlDB.SetConnMaxLifetime(DefaultConfig.ConnMaxLifetime)
}

func InitRedis() {
	RedisPool = &redis.Pool{
		MaxIdle:     redis_config.Redis.MaxIdel,
		IdleTimeout: time.Duration(redis_config.Redis.IdleTimeout) * time.Second,
		Dial: func() (redis.Conn, error) {
			return redis.Dial(
				"tcp",
				redis_config.Redis.Host+":"+redis_config.Redis.Port,
				redis.DialUsername(redis_config.Redis.Username),
				redis.DialPassword(redis_config.Redis.Password),
			)
		},
	}
}

func CreateDatabase() {
	dir, _ := os.Getwd()

	cf := config.ReadFromYaml(filepath.Join(dir, "config.yml"))
	database := cf.Databases["default"]
	targetDatabaseName := database.Name

	var err error
	_, err = gorm.Open(mysql.Open(database.GetDSN()), &gorm.Config{})
	if err == nil {
		fmt.Printf("database %s exist.\n", targetDatabaseName)
		return
	}

	if strings.Contains(err.Error(), "Unknown database") {
		database.Name = "mysql"
		mysqlDb, err := gorm.Open(mysql.Open(database.GetDSN()), &gorm.Config{})
		if err != nil {
			fmt.Printf("open connection fail. err: %#v\n", err)
			os.Exit(1)
		}

		err = mysqlDb.Exec(`CREATE DATABASE IF NOT EXISTS ` + targetDatabaseName + `
			DEFAULT CHARACTER SET utf8mb4
			DEFAULT COLLATE utf8mb4_general_ci;`).Error
		if err != nil {
			fmt.Printf("create database %s fail, err: %#v\n", targetDatabaseName, err)
			os.Exit(1)
		}

		fmt.Printf("create database %s successfully.\n", targetDatabaseName)
	}
}
