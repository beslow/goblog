package config

import (
	"os"
	"path/filepath"
	"strings"

	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

type RocketmqConfig struct {
	NamesrvAddr string `yaml:"namesrv_addr"`
}

var Rocketmq RocketmqConfig

func init() {
	dir, _ := os.Getwd()

	if os.Getenv("GO_TEST") != "" && !strings.HasSuffix(dir, "test") {
		panic("wrong position")
	}

	configPath := "rocketmq.yml"

	if c := os.Getenv("CONFIG_DIR"); c != "" {
		configPath = filepath.Join(c, configPath)
	}

	data, err := os.ReadFile(filepath.Join(dir, configPath))
	if err != nil {
		log.Errorf("Read rocketmq.yml fail, err: %#v", err)
	}

	err = yaml.Unmarshal(data, &Rocketmq)
	if err != nil {
		log.Errorf("unmarshal rocketmq.yml fail, err: %#v", err)
	}
}
