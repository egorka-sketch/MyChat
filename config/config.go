package config

import (
	"github.com/stretchr/testify/assert/yaml"
	"io/ioutil"
	"log"
)

type Config struct {
	Database struct {
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		Name     string `yaml:"name"`
	} `yaml:"database"`
	JWT struct {
		Key string `yaml:"key"`
	} `yaml:"jwt"`
}

func LoadConfig(path string) *Config {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal("load config file error:", err)
	}
	var config Config
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		log.Fatal("unmarshal config file error:", err)
	}
	return &config
}
