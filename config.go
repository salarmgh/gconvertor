package main

import (
	"github.com/alecthomas/gometalinter/_linters/src/gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

type ConfigMap struct {
	Host    string
	Port    string
	Path	string
}

func setEnv(key string, defaultValue string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return defaultValue
	}
	return value
}

func ConfigLoad() (ConfigMap, error) {
	configPath, _ := os.Getwd()
	configPath += "/config.yaml"
	c := ConfigMap{}
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		c.Host = setEnv("LISTEN_ADDR", "127.0.0.1")
		c.Port = setEnv("LISTEN_PORT", "8080")
		c.Path = setEnv("EXPORT_PATH", "/data")
	} else {
		yamlFile, err := ioutil.ReadFile(configPath)
		if err != nil {
			return c, err
		}
		err = yaml.Unmarshal(yamlFile, &c)
		if err != nil {
			return c, err
		}
	}
	return c, nil
}
