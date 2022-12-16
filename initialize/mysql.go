package initialize

import (
	"os"
	"path/filepath"
	"strings"

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

	if os.Getenv("GO_TEST") != "" && !strings.HasSuffix(dir, "test") {
		panic("wrong position")
	}

	configPath := "config.yml"

	if c := os.Getenv("CONFIG_DIR"); c != "" {
		configPath = filepath.Join(c, configPath)
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
