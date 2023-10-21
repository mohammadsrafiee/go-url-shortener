package configReader

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"sync"
	"time"
)

var (
	instance *Config
	once     sync.Once
)

type Config struct {
	Server   ServerConfig   `yaml:"server"`
	Database DatabaseConfig `yaml:"database"`
	Log      Log            `yaml:"log"`
	Cache    Cache          `yaml:"cache"`
}

type ServerConfig struct {
	Address string        `yaml:"address"`
	Timeout time.Duration `yaml:"timeout"`
}

type DatabaseConfig struct {
	IsActive         bool   `yaml:"is-active"`
	ConnectionString string `yaml:"connection_string"`
}

type Log struct {
	Level int    `yaml:"level"`
	Path  string `yaml:"path"`
}

type Cache struct {
	Type     string `yaml:"type"`
	Address  string `yaml:"address"`
	Password string `yaml:"password"`
	Username string `yaml:"username"`
	Db       string `yaml:"db"`
}

func GetInstance(path ...string) *Config {
	once.Do(func() {
		if len(path) > 0 {
			yamlFile, err := ioutil.ReadFile(path[0])
			if err != nil {
				panic(err)
			}
			if err := yaml.Unmarshal(yamlFile, &instance); err != nil {
				panic(err)
			}
		}
	})
	return instance
}

func Configuration() *Config {
	return instance
}
