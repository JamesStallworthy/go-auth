package config

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

type Config struct {
	KeyLocation          string `yaml:"KeysLocation"`
	ClientConfigLocation string `yaml:"ClientConfigFile"`
	Url                  string `yaml:"URL"`
}

func LoadConfig(configFile string) *Config {
	config := &Config{}

	yamlFile, err := ioutil.ReadFile(configFile)
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, config)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	return config
}
