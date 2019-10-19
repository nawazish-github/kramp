package models

type Book struct {
	Kind       string     `json:"kind"`
	TotalItems int64      `json:"totalItems"`
	Items      []BookInfo `json:"items"`
}

type BookInfo struct {
	Kind       string      `json:"kind"`
	VolumeInfo BookDetails `json:"volumeInfo"`
}

type BookDetails struct {
	Tile    string   `json:"title"`
	Authors []string `json:"authors"`
}
