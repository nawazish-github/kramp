package models

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"sync"
)

type Config struct {
	MaxBooks   int `json:"maxBooks"`
	MaxAlbums  int `json:"maxAlbums"`
	MaxTimeout int `json:"maxTimeout"`
}

var instance *Config
var once sync.Once

func GetConfigInstance() *Config {
	once.Do(func() {
		data, err := ioutil.ReadFile("../config/config.json")

		if err != nil {
			log.Println("Unable to read configuration file; using default values")
			instance = &Config{
				MaxAlbums:  5,
				MaxBooks:   5,
				MaxTimeout: 1,
			}
		}
		var conf Config
		json.Unmarshal(data, &conf)
		instance = &conf
	})
	return instance
}
