package labo

import (
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// Href is a snapshot of HTTP reference that contains information about a Nintendo Labo product.
type Href struct {
	Link   string   `json:"link"`
	Target string   `json:"target"`
	URL    *url.URL `json:"URL"`
}

var (
	hrefFn = [](func(*goquery.Selection, *Href)){
		getHrefLink,
		getHrefTarget,
		getHrefURL}
)

// getHrefLink searches the *goquery.Selection for the HTML href attribute required for a labo.Href struct.
func getHrefLink(s *goquery.Selection, h *Href) {
	var (
		link string
	)
	link, _ = s.Attr(attrHref)
	h.Link = link
}

// getHrefTarget searches the *goquery.Selection for the HTML target attribute required for a labo.Href struct.
func getHrefTarget(s *goquery.Selection, h *Href) {
	var (
		ok     bool
		t      string
		target = defaultAttrTarget
	)
	t, ok = s.Attr(attrTarget)
	if ok {
		target = strings.ToUpper(t)
	}
	h.Target = target
}

// getHrefURL searches the *goquery.Selection for *url.URL required for a labo.Href struct.
func getHrefURL(s *goquery.Selection, h *Href) {
	var (
		err  error
		href string
		ok   bool
		URL  *url.URL
	)
	href, ok = s.Attr(attrHref)
	if !ok {
		return
	}
	URL, err = url.Parse(href)
	ok = (err == nil)
	if !ok {
		return
	}
	h.URL = URL
}

// newHref is a constructor function that instantiates a new Href struct pointer.
func newHref(s *goquery.Selection) *Href {
	var (
		ok   bool
		href = &Href{}
	)
	ok = (s.Length() > 0)
	if !ok {
		return nil
	}
	ok = (strings.ToLower(s.Nodes[0].Data) == htmlAnchor)
	if !ok {
		return newHref(s.Find(htmlAnchor))
	}
	for _, fn := range hrefFn {
		fn(s, href)
	}
	return href
}
