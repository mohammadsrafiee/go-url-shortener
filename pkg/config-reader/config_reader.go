package configReader

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"sync"
)

var (
	instance *Config
	once     sync.Once
)

func Instance() *Config {
	return instance
}

func ConfigFactory(path string) {
	once.Do(func() {
		if instance == nil {
			instance = &Config{} // Initialize the instance here
			yamlFile, err := ioutil.ReadFile(path)
			if err != nil {
				panic(err)
			}
			if err := yaml.Unmarshal(yamlFile, instance); err != nil {
				panic(err)
			}
		}
	})
}
