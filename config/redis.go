package config

import (
	"os"
	"path/filepath"

	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

type RedisConfig struct {
	Host        string
	Port        string
	Username    string
	Password    string
	MaxIdel     int
	IdleTimeout int
}

var (
	Redis RedisConfig
)

func init() {
	dir, _ := os.Getwd()

	data, err := os.ReadFile(filepath.Join(dir, "redis.yml"))
	if err != nil {
		log.Errorf("Read redis.yml fail, err: %#v", err)
	}

	err = yaml.Unmarshal(data, &Redis)
	if err != nil {
		log.Errorf("unmarshal redis.yml fail, err: %#v", err)
	}
}
