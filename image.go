package labo

import (
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

const (
	defaultImageAlt string = "NIL"
)

type Image struct {
	Alt  string   `json:"alt"`
	Link string   `json:"link"`
	URL  *url.URL `json:"URL"`
}

func newImage(s *goquery.Selection) *Image {
	var (
		alt  = defaultImageAlt
		link string
		ok   bool
		URL  *url.URL
	)
	ok = (s.Length() > 0)
	if !ok {
		return nil
	}
	ok = (strings.ToUpper(s.Nodes[0].Data) == "IMG")
	if !ok {
		return newImage(s.Find("img"))
	}
	link, _ = s.Attr("src")
	ok = (strings.HasPrefix(link, "data:image") == false)
	if !ok {
		link, _ = s.Attr("data-src")
	}
	ok = (len(link) > 0)
	if !ok {
		return nil
	}
	URL, _ = url.Parse(link)
	return &Image{
		Alt:  alt,
		Link: link,
		URL:  URL}
}

func newImages(s *goquery.Selection) []*Image {
	var (
		images []*Image
	)
	s.Each(func(i int, s *goquery.Selection) {
		image := newImage(s)
		if image == nil {
			return
		}
		images = append(images, image)
	})
	return images
}
