package labo

import "github.com/PuerkitoBio/goquery"

type Retailer struct {
	Href  *Href  `json:"href"`
	Image *Image `json:"image"`
	Name  string `json:"name"`
}

func getRetailerName(s *goquery.Selection) {

}

func newRetailer(s *goquery.Selection) *Retailer {

	return &Retailer{
		Href:  newHref(s),
		Image: newImage(s)}
}
