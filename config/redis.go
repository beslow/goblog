package config

import (
	"os"
	"path/filepath"
	"strings"

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

var Redis RedisConfig

func init() {
	dir, _ := os.Getwd()

	if os.Getenv("GO_TEST") != "" && !strings.HasSuffix(dir, "test") {
		panic("wrong position")
	}

	configPath := "redis.yml"

	if c := os.Getenv("CONFIG_YAML"); c != "" {
		configPath = c
	}

	data, err := os.ReadFile(filepath.Join(dir, configPath))
	if err != nil {
		log.Errorf("Read redis.yml fail, err: %#v", err)
	}

	err = yaml.Unmarshal(data, &Redis)
	if err != nil {
		log.Errorf("unmarshal redis.yml fail, err: %#v", err)
	}
}
