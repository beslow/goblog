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
