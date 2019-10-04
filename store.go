package labo

import (
	"net/url"
	"sync"

	"github.com/PuerkitoBio/goquery"
)

type Store struct {
	Images []*Image `json:"images"`
	Parts  []*Part  `json:"parts"`
	URL    *url.URL `json:"URL"`
}

func newStore(d *goquery.Document) *Store {
	const (
		CSS string = "body"
	)
	var (
		b = d.Find(CSS)
	)
	var (
		images []*Image
		parts  []*Part

		wg sync.WaitGroup
	)
	wg.Add(1)
	go func() {
		defer wg.Done()
		const (
			CSS = "#product-thumbs img"
		)
		images = newImages(b.Find(CSS))
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		const (
			CSS = "#prodDescBtm ul li"
		)
		parts = newParts(b.Find(CSS))
	}()
	wg.Wait()
	return &Store{
		Images: images,
		Parts:  parts,
		URL:    d.Url}
}
