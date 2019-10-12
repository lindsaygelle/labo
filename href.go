package labo

import (
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// Href is a hyperlink reference to a Nintendo Labo product resource provided by Nintendo.
type Href struct {
	Link   string   `json:"link"`
	Target string   `json:"target"`
	URL    *url.URL `json:"URL"`
}

// newHref is a constructor function that instantiates a new Href struct pointer.
//
// newHref requires the argument goquery.Selection pointer to be a valid
// HTML anchor element that contains a reference to an attribute that points
// to a URL. Should no URL be found no Href struct
// is returned.
func newHref(s *goquery.Selection) *Href {
	var (
		err    error
		link   string
		ok     bool
		target = defaultAttrTarget
		URL    *url.URL
	)
	ok = (s.Length() > 0)
	if !ok {
		return nil
	}
	ok = (strings.ToLower(s.Nodes[0].Data) == htmlAnchor)
	if !ok {
		return newHref(s.Find(htmlAnchor))
	}
	link, ok = s.Attr(attrHref)
	if !ok {
		return nil
	}
	_, ok = s.Attr(attrTarget)
	if ok {
		target, _ = s.Attr(attrTarget)
		target = strings.ToUpper(target)
	}
	URL, err = url.Parse(link)
	ok = (err == nil)
	if !ok {
		return nil
	}
	return &Href{
		Link:   link,
		Target: target,
		URL:    URL}
}
