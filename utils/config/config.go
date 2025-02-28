package config

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Chunk struct {
		Size int64 `yaml:"size"`
	} `yaml:"Chunk"`
}

var config *Config

func LoadConfig() *Config {
	if config != nil {
		return config
	}
	data, err := ioutil.ReadFile("/Users/nagarajpoojari/Desktop/learn/nanoDFS/Master/config.yaml")
	if err != nil {
		log.Fatalf("error reading file: %v", err)
	}

	err = yaml.Unmarshal(data, &config)
	if err != nil {
		log.Fatalf("error unmarshalling yaml: %v", err)
	}
	log.Printf("Successfully loaded config: %v", config)
	return config
}
