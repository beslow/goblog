package db

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/GoAdminGroup/go-admin/modules/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func CreateDatabase() {
	dir, _ := os.Getwd()

	var configPath string
	if os.Getenv("GoTest") != "" {
		configPath = "test/config.yml"
	} else {
		configPath = "config.yml"
	}

	cf := config.ReadFromYaml(filepath.Join(dir, configPath))
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
