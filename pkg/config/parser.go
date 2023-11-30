package config

import (
	"os"
	"sync"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Database string   `yaml:"database"`
	Indexers []string `yaml:"indexers"`
	Address  string   `yaml:"address"`
	Port     int      `yaml:"port"`
}

var (
	instance     *Config
	instanceOnce sync.Once
)

func Instance() *Config {
	if instance == nil {
		instanceOnce.Do(func() {
			instance = &Config{
				Database: "mayoi.db",
			}
		})
	}
	return instance
}

func (c *Config) Load(path string) error {
	fd, err := os.Open(path)
	if err != nil {
		return err
	}

	if err := yaml.NewDecoder(fd).Decode(c); err != nil {
		return err
	}

	return nil
}
