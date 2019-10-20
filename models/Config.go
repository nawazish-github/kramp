package models

type Config struct {
	MaxBooks   int64 `json:"maxBooks"`
	MaxAlbums  int64 `json:"maxAlbums"`
	MaxTimeout int64 `json:"maxTimeout"`
}
