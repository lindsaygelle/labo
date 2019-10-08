package labo

import (
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type Image struct {
	Alt      string
	Link     string
	URL      *url.URL
	Variants []*Image
}

func newImage(s *goquery.Selection) *Image {
	const (
		dataPrefix string = "data:image"
		HTML       string = "img"
	)
	var (
		alt      = defaultAttrAlt
		err      error
		link     string
		ok       bool
		URL      *url.URL
		variants []*Image
	)
	ok = (s.Length() > 0)
	if !ok {
		return nil
	}
	ok = (strings.ToLower(s.Nodes[0].Data) == HTML)
	if !ok {
		return newImage(s.Find(HTML))
	}
	link, _ = s.Attr(attrSrc)
	ok = (strings.HasPrefix(link, dataPrefix) == false)
	if !ok {
		link, _ = s.Attr(attrDataSrc)
	}
	ok = (len(link) > 0)
	if !ok {
		return nil
	}
	URL, err = url.Parse(link)
	ok = (err == nil)
	if !ok {
		return nil
	}
	return &Image{
		Alt:      alt,
		Link:     link,
		URL:      URL,
		Variants: variants}
}

func newImages(s *goquery.Selection) []*Image {
	var (
		image  *Image
		images []*Image
		ok     bool
	)
	s.Each(func(i int, s *goquery.Selection) {
		image = newImage(s)
		ok = (image != nil)
		if !ok {
			return
		}
		images = append(images, image)
	})
	return images
}
