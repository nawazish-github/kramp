package json

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/nawazish-github/kramp/models"
)

func UnmarshalBooks(data []byte) models.Book {
	book := models.Book{}
	json.Unmarshal(data, &book)
	return book
}

func UnmarshalAlbums(data []byte) models.Album {
	album := models.Album{}
	json.Unmarshal(data, &album)
	return album
}

func UnmarshalConfig() models.Config {
	data, err := ioutil.ReadFile("../config/config.json")

	if err != nil {
		log.Println("Unable to read configuration file; using default values")
		return models.Config{
			MaxAlbums:  5,
			MaxBooks:   5,
			MaxTimeout: 1,
		}
	}
	var conf models.Config
	json.Unmarshal(data, &conf)
	return conf
}
