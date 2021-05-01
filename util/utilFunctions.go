package util

import (
	"encoding/json"
	"io/ioutil"
	"fmt"
	"net/http"
)

func GoogleBooksAPI(isbn string) (body []byte, err error) {
	bookInfoUrl := fmt.Sprintf("https://www.googleapis.com/books/v1/volumes?q=isbn:%s&key=AIzaSyD6hE1YZy8dT0-x1025sGBuFW9A1gtVMhI", isbn)
	resp, err := http.Get(bookInfoUrl)
	if err != nil {
		return
	}

	body, err = ioutil.ReadAll(resp.Body)
	return
}

func ParseJSON(body []byte) VolumeInfo {
	bookDetails := BookDetailsJSON{}
	err := json.Unmarshal(body, &bookDetails)
	if err != nil {
		panic(err)
	}
	return bookDetails.Items[0].BookInfo
}

type BookDetailsJSON struct {
	Items []Item `json:"items"`
}

type Item struct {
	BookInfo VolumeInfo `json:"volumeInfo"`
}

type VolumeInfo struct {
	Title string `json:"title"`
	Authors []string `json:"authors"`
}