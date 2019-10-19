package json

import (
	"encoding/json"

	"github.com/nawazish-github/kramp/models"
)

func UnmarshalBooks(data []byte) models.Book {
	book := models.Book{}
	json.Unmarshal(data, &book)
	return book
}

func UnmarshalAlbums(data []byte) models.Album {
	//log.Println("Unmarshalling Albums...")
	album := models.Album{}
	json.Unmarshal(data, &album)
	//log.Println("Albums unmarshelled, ", album)
	return album
}
