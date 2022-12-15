package config

import (
	"os"
	"path/filepath"
	"strings"

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

	if os.Getenv("GO_TEST") != "" && !strings.HasSuffix(dir, "test") {
		panic("wrong position")
	}

	configPath := "sentry.yml"

	if c := os.Getenv("CONFIG_YAML"); c != "" {
		configPath = c
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
