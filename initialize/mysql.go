package initialize

import (
	"os"
	"path/filepath"

	"github.com/GoAdminGroup/go-admin/modules/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB            *gorm.DB
	DefaultConfig config.Database
)

func InitMySQL() {
	dir, _ := os.Getwd()

	var configPath string
	if os.Getenv("GoTest") != "" {
		configPath = "test/config.yml"
	} else {
		configPath = "config.yml"
	}

	cf := config.ReadFromYaml(filepath.Join(dir, configPath))
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
