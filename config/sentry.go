package config

import (
	"os"
	"path/filepath"

	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

type SentryConfig struct {
	Enable bool
	Dsn    string
}

var Sentry SentryConfig

func init() {
	dir, _ := os.Getwd()

	var configPath string
	if os.Getenv("GoTest") != "" {
		configPath = "test/sentry.yml"
	} else {
		configPath = "sentry.yml"
	}

	data, err := os.ReadFile(filepath.Join(dir, configPath))
	if err != nil {
		log.Errorf("Read sentry.yml fail, err: %#v", err)
	}

	err = yaml.Unmarshal(data, &Sentry)
	if err != nil {
		log.Errorf("unmarshal sentry.yml fail, err: %#v", err)
	}
}
