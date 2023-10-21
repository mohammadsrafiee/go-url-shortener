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
