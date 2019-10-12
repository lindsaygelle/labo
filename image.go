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

// newImage is a constructor function that instantiates a new Image struct pointer.
//
// newImage requires the argument goquery.Selection pointer to be a valid
// HTML image element that contains a reference to an attribute that points
// to a image resource. Should no image resource be found no Image struct
// is returned.
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

// newImages is a constructor function that instantiates and returns a slice of Image struct pointers.
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
