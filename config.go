package main

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"sync"
)

type config struct {
	BeatSeconds int      `yaml:"beat_seconds"`
	Urls        []string `yaml:"urls"`
}

var (
	instance config
)

var once sync.Once

func get() config {
	once.Do(func() {
		instance = readConfig()
	})
	return instance
}

func readConfig() config {
	yamlFile, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		log.Fatalln("Error reading YAML file: %s\n", err)
	}

	var c config
	err = yaml.Unmarshal(yamlFile, &c)
	if err != nil {
		log.Fatalln("Error parsing YAML file: %s\n", err)
	}

	return c
}

func GetUrls() []string {
	return get().Urls
}

func GetBeatSeconds() int {
	return get().BeatSeconds
}
