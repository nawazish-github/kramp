package executor

import (
	"context"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/nawazish-github/kramp/constants"
	"github.com/nawazish-github/kramp/json"
	"github.com/nawazish-github/kramp/models"
)

type SimpleExecutor struct{}

func (simpleExec *SimpleExecutor) Fetch(searchString string) models.Response {
	log.Println("Inside Simple Executor")

	backGroundCtx := context.Background()
	timeOutCtx, cancel := context.WithTimeout(backGroundCtx, 1*time.Minute)
	defer cancel()
	albumChan := make(chan models.Album)
	bookChan := make(chan models.Book)

	go fetchAlbumDetails(timeOutCtx, searchString, albumChan)
	go fetchBookDetails(timeOutCtx, searchString, bookChan)

	var albums models.Album = <-albumChan
	var books models.Book = <-bookChan
	var response = models.Response{
		Albums: albums,
		Books:  books,
	}
	return response
}

func fetchAlbumDetails(timeOutCtx context.Context, searchString string, albumChan chan models.Album) {
	defer func() {
		close(albumChan)
	}()
	req, err := http.NewRequest(http.MethodGet, constants.ITUNES_API+"term="+searchString+"&&entity=album", nil)
	if err != nil {
		log.Println("albun req formation failed", err)
		return
	}
	req = req.WithContext(timeOutCtx)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println("album request execution failed", err)
		return
	}
	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Failed to call ", constants.ITUNES_API+"term="+searchString+"&&entity=album")
		log.Println(err)
		return
	}
	albums := json.UnmarshalAlbums(respBytes)
	albumChan <- albums
}

func fetchBookDetails(timeOutCtx context.Context, searchString string, bookChan chan models.Book) {
	defer func() {
		close(bookChan)
	}()
	req, err := http.NewRequest(http.MethodGet, constants.GOOGLE_BOOKS_API+"q="+searchString, nil)
	if err != nil {
		log.Println("book req formation failed", err)
		return
	}

	req = req.WithContext(timeOutCtx)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println("book request execution failed", err)
		return
	}
	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Failed to call ", constants.GOOGLE_BOOKS_API+"q="+searchString)
		log.Println(err)
		return
	}
	books := json.UnmarshalBooks(respBytes)
	bookChan <- books
}
