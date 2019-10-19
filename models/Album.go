package models

type Album struct {
	ResultCount int64 `json: "resultCount"`
	Results     []AlbumInfo
}

type AlbumInfo struct {
	CollectionType string `json: "collectionType"`
	ArtistName     string `json: "artistName"`
	CollectionName string `json: "collectionName"`
}
