package labo

import (
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// Image is a snapshot of an image resource belonging to a Nintendo Labo product.
type Image struct {
	Alt      string   `json:"alt"`
	Link     string   `json:"link"`
	URL      *url.URL `json:"URL"`
	Variants []*Image `json:"variants"`
}

// newImage is a constructor function that take an argument goquery.Selection pointer
// to build a new Image pointer. When building a new Image struct
// the function checks whether the argument goquery.Selection pointer
// contains a valid HTML src attribute or its derivatives. If no corresponding image
// reference address can be found, no pointer is returned.
func newImage(s *goquery.Selection) *Image {
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
	ok = (strings.ToLower(s.Nodes[0].Data) == htmlImage)
	if !ok {
		return newImage(s.Find(htmlImage))
	}
	link, _ = s.Attr(attrSrc)
	ok = (link == imageBase64)
	if ok {
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
