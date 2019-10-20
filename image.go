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

var (
	imageFn = [](func(*goquery.Selection, *Image)){
		getImageAlt,
		getImageLink,
		getImageURL,
		getImageVariants}
)

// getImageAlt searches the *goquery.Selection for the alt attribute required for a labo.Image
func getImageAlt(s *goquery.Selection, i *Image) {
	var (
		alt = defaultAttrAlt
		ok  bool
	)
	_, ok = s.Attr(attrSrc)
	if ok {
		alt, _ = s.Attr(attrSrc)
	}
	i.Alt = alt
}

// getImageLink searches the *goquery.Selection for the link resource required for a labo.Image
func getImageLink(s *goquery.Selection, i *Image) {
	var (
		link string
		ok   bool
	)
	link, _ = s.Attr(attrSrc)
	ok = (link == imageBase64)
	if ok {
		link, _ = s.Attr(attrDataSrc)
	}
	ok = (len(link) > 0)
	if !ok {
		return
	}
	i.Link = link
}

// getImageURL searches the *goquery.Selection for the *url.URL required for a labo.Image
func getImageURL(s *goquery.Selection, i *Image) {
	var (
		err     error
		link, _ = s.Attr(attrSrc)
		ok      bool
		URL     *url.URL
	)
	URL, err = url.Parse(link)
	ok = (err == nil)
	if !ok {
		return
	}
	i.URL = URL
}

func getImageVariants(s *goquery.Selection, i *Image) {}

// newImage is a constructor function that instantiates a new labo.Image struct pointer.
func newImage(s *goquery.Selection) *Image {
	var (
		ok    bool
		image = &Image{}
	)
	ok = (s.Length() > 0)
	if !ok {
		return nil
	}
	ok = (strings.ToLower(s.Nodes[0].Data) == htmlImage)
	if !ok {
		return newImage(s.Find(htmlImage))
	}
	for _, fn := range imageFn {
		fn(s, image)
	}
	return image
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
