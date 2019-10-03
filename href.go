package labo

import (
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

const (
	defaultTargetAttr string = "self"
)

type Href struct {
	Link   string   `json:"link"`
	Target string   `json:"target"`
	URL    *url.URL `json:"URL"`
}

func newHref(s *goquery.Selection) *Href {
	var (
		link   string
		ok     bool
		target = defaultTargetAttr
		URL    *url.URL
	)
	ok = (s.Length() > 0)
	if !ok {
		return nil
	}
	ok = (strings.ToUpper(s.Nodes[0].Data) == "A")
	if !ok {
		return newHref(s.Find("a"))
	}
	link, ok = s.Attr("href")
	if !ok {
		return nil
	}
	_, ok = s.Attr("target")
	if ok {
		target, _ = s.Attr("target")
		target = strings.ToUpper(target)
	}
	URL, _ = url.Parse(link)
	return &Href{
		Link:   link,
		Target: target,
		URL:    URL}
}
