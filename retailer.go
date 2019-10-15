package labo

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type Retailer struct {
	Href  *Href  `json:"href"`
	Image *Image `json:"image"`
	Name  string `json:"name"`
}

func getRetailerName(s *goquery.Selection) string {
	var (
		name      = defaultRetailerName
		ok        bool
		substring string
	)
	substring, ok = s.Attr(attrClass)
	if ok {
		name = strings.TrimSpace(substring)
	}
	return name
}

func newRetailer(s *goquery.Selection) *Retailer {

	return &Retailer{
		Href:  newHref(s),
		Image: newImage(s),
		Name:  getRetailerName(s)}
}

func newRetailers(s *goquery.Selection) []*Retailer {
	var (
		retailer  *Retailer
		retailers []*Retailer
		ok        bool
	)
	s.Each(func(i int, s *goquery.Selection) {
		retailer = newRetailer(s)
		ok = (retailer != nil)
		if !ok {
			return
		}
		retailers = append(retailers, retailer)
	})
	return retailers
}
