package core

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
	yamlFile, err := ioutil.ReadFile(".http-beat.yaml")
	if err != nil {
		log.Fatalln("read error: ", err)
	}

	var c config
	err = yaml.Unmarshal(yamlFile, &c)
	if err != nil {
		log.Fatalln("unmarshal error: ", err)
	}

	return c
}

func writeConfig(c config) {
	d, err := yaml.Marshal(&c)
	if err != nil {
		log.Fatalf("marshal error: ", err)
	}

	err = ioutil.WriteFile(".http-beat.yaml", d, 0644)
	if err != nil {
		log.Fatalln("write error:", err)
	}

	instance = readConfig()
}

func GetUrls() []string {
	return get().Urls
}

func GetBeatSeconds() int {
	return get().BeatSeconds
}

func AddUrls(urls []string) {
	writeConfig(config{
		Urls:        append(GetUrls(), urls...),
		BeatSeconds: GetBeatSeconds(),
	})
}

func SetBeatSeconds(seconds int) {
	writeConfig(config{
		Urls:        GetUrls(),
		BeatSeconds: seconds,
	})
}
