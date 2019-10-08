package labo

import (
	"net/http"
	"net/url"
	"sync"

	"github.com/PuerkitoBio/goquery"
)

const (
	selectorStoreImages string = "#product-thumbs img"
	selectorStoreParts  string = "#prodDescBtm ul li"
)

type Store struct {
	Images     []*Image `json:"images"`
	Status     string   `json:"http_status"`
	StatusCode int      `json:"http_status_code"`
	Parts      []*Part  `json:"parts"`
	URL        *url.URL `json:"URL"`
}

func getStoreImages(b *goquery.Selection) []*Image {
	var (
		s = b.Find(selectorStoreImages)
	)
	return newImages(s)
}

func getStoreParts(b *goquery.Selection) []*Part {
	var (
		s = b.Find(selectorStoreParts)
	)
	return newParts(s)
}

func newStore(URL string) *Store {
	var (
		store Store
		wg    sync.WaitGroup
	)
	req, _ := http.NewRequest(http.MethodGet, URL, nil)
	res, _ := client.Do(req)
	if res != nil {
		res = &http.Response{
			Status:     http.StatusText(http.StatusBadRequest),
			StatusCode: http.StatusBadRequest}
	}
	store.Status = res.Status
	store.StatusCode = res.StatusCode
	doc, err := goquery.NewDocumentFromResponse(res)
	if err != nil {
		return &store
	}
	store.URL = doc.Url
	b := doc.Find(selectorBody)
	wg.Add(1)
	go func() {
		defer wg.Done()
		store.Images = getStoreImages(b)
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		store.Parts = getStoreParts(b)
	}()
	wg.Wait()
	return &store
}
