package cache

import (
	"sync"

	"github.com/nawazish-github/kramp/models"
)

type Cache struct {
	Albums models.Album
	Books  models.Book
}

var instance *Cache
var once sync.Once
var mu sync.Mutex

func (c *Cache) UpdateCache(response *models.Response) {
	mu.Lock()
	defer mu.Unlock()
	conf := models.GetConfigInstance()
	maxBooks := conf.MaxBooks
	maxAlbums := conf.MaxAlbums

	tmpCache := GetCacheInstance()

	respBooks := response.Books.Items[0:maxBooks]
	respAlbums := response.Albums.Results[0:maxAlbums]

	tmpCache.Books.Items = respBooks
	tmpCache.Albums.Results = respAlbums
}

func GetCacheInstance() *Cache {
	once.Do(func() {
		instance = &Cache{}
	})
	return instance
}
