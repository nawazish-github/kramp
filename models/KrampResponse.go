package models

type KrampResponses struct {
	Responses []KrampResponse `json:"KrampResponses"`
}

type KrampResponse struct {
	Title       string   `json:"title"`
	Protagonist []string `json:"protagonist"`
	IsAlbum     bool     `json:"isAlbum"`
}
