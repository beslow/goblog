package config

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

type MysqlConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
}

var MySQL MysqlConfig

func init() {
	dir, _ := os.Getwd()
	dir = strings.Replace(dir, "db/migrations", "", -1)
	dir = strings.Replace(dir, "/db", "", -1)

	data, err := ioutil.ReadFile(filepath.Join(dir, "database.yml"))
	if err != nil {
		log.Errorf("Read database.yml fail, err: %#v", err)
	}

	err = yaml.Unmarshal(data, &MySQL)
	if err != nil {
		log.Errorf("unmarshal database.yml fail, err: %#v", err)
	}
}
