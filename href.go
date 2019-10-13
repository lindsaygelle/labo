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

// getHrefLink inspects the argument *goquery.Selection and attempts to return the URL provided by the HTML anchor tag.
func getHrefLink(s *goquery.Selection) string {
	var (
		link string
	)
	link, _ = s.Attr(attrHref)
	return link
}

// getHrefTarget inspects the argument *goquery.Selection and attempts to return a normalized HTML target attribute.
func getHrefTarget(s *goquery.Selection) string {
	var (
		ok     bool
		t      string
		target = defaultAttrTarget
	)
	t, ok = s.Attr(attrTarget)
	if ok {
		target = strings.ToUpper(t)
	}
	return target
}

// getHrefURL inspects the argument *goquery.Selection and attempts to return a new url.URL struct pointer.
//
// getHrefURL assumes the argument *goquery.Selection is a valid HTML anchor element and contains
// a href attribute.
func getHrefURL(s *goquery.Selection) *url.URL {
	var (
		err  error
		href string
		ok   bool
		URL  *url.URL
	)
	href, ok = s.Attr(attrHref)
	if !ok {
		return nil
	}
	URL, err = url.Parse(href)
	ok = (err == nil)
	if !ok {
		return nil
	}
	return URL
}

// newHref is a constructor function that instantiates a new Href struct pointer.
//
// newHref requires the argument *goquery.Selection to be a valid
// HTML anchor element that contains a reference to an attribute that points
// to a URL. Should no URL be found no Href struct
// is returned. Assumes that the calling function will require the return
// nil struct pointer case.
func newHref(s *goquery.Selection) *Href {
	var (
		ok bool
	)
	ok = (s.Length() > 0)
	if !ok {
		return nil
	}
	ok = (strings.ToLower(s.Nodes[0].Data) == htmlAnchor)
	if !ok {
		return newHref(s.Find(htmlAnchor))
	}
	return &Href{
		Link:   getHrefLink(s),
		Target: getHrefTarget(s),
		URL:    getHrefURL(s)}
}
