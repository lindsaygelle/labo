package labo

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// Retailer is a snapshot of a first or third party retailer that stocks and sells Nintendo Labo products.
type Retailer struct {
	Href  *Href  `json:"href"`
	Image *Image `json:"image"`
	Name  string `json:"name"`
}

var (
	retailerFn = [](func(*goquery.Selection, *Retailer)){
		getRetailerHref,
		getRetailerImage,
		getRetailerName}
)

// getRetailerHref searches the *goquery.Selection for the *labo.Href required for a labo.Toycon struct.
func getRetailerHref(s *goquery.Selection, r *Retailer) {
	r.Href = newHref(s)
}

// getRetailerImage searches the *goquery.Selection for the *labo.Image struct required for a labo.Retailer struct.
func getRetailerImage(s *goquery.Selection, r *Retailer) {
	r.Image = newImage(s)
}

// getRetailerName searches the *goquery.Selection for the name of the retailer required for a labo.Retailer struct.
func getRetailerName(s *goquery.Selection, r *Retailer) {
	var (
		name      = defaultRetailerName
		ok        bool
		substring string
	)
	substring, ok = s.Attr(attrClass)
	if ok {
		name = strings.TrimSpace(substring)
	}
	r.Name = name
}

// newRetailer is a constructor function that instantiates and returns a new *labo.Retailer.
func newRetailer(s *goquery.Selection) *Retailer {
	var (
		retailer = &Retailer{}
	)
	for _, fn := range retailerFn {
		fn(s, retailer)
	}
	return retailer
}

// newRetailers is a constructor function that instantiates and returns a new slice of *labo.Retailer.
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
