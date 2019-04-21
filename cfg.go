package main

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

type RedirectionInfo struct {
	path string
	url  string
}

type Config map[string]string

func LoadConfigStore(filePath string) *Config {
	cfg := make(Config)

	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Print(err)
		return &cfg
	}

	if err := yaml.Unmarshal(data, &cfg); err != nil {
		log.Fatalf("error: %v", err)
	}

	return &cfg
}

func (c *Config) Save(filePath string) {
	if c == nil || len(*c) == 0 {
		return
	}

	data, err := yaml.Marshal(c)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	if err := ioutil.WriteFile(filePath, data, 0644); err != nil {
		log.Fatal(err)
	}
}

func (c *Config) Add(info *RedirectionInfo) {
	(*c)[info.path] = info.url
}

func (c *Config) Remove(path string) {
	if _, ok := (*c)[path]; ok {
		delete(*c, path)
	}
}
