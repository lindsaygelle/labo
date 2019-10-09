package labo

import (
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type Href struct {
	Link   string   `json:"link"`
	Target string   `json:"target"`
	URL    *url.URL `json:"URL"`
}

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
