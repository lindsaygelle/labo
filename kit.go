package labo

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/PuerkitoBio/goquery"
)

type Kit struct {
	Hero       *Image
	Retailers  []*Retailer
	Status     string
	StatusCode int
	URL        *url.URL
}

var (
	kitFn = [](func(s *goquery.Selection, k *Kit)){
		getKitRetailers}
)

func GetKit(l *Labo) *Kit {
	var (
		doc *goquery.Document
		err error
		ok  bool
		req *http.Request
		res *http.Response
		s   *goquery.Selection

		k = &Kit{
			Status:     http.StatusText(http.StatusBadRequest),
			StatusCode: http.StatusBadRequest}
	)
	ok = (l.Ref != "NIL")
	if !ok {
		return k
	}
	req, err = http.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", laboURI, l.Ref), nil)
	ok = (err == nil)
	if !ok {
		return k
	}
	res, err = client.Do(req)
	ok = (err == nil)
	if !ok {
		return k
	}
	k.Status = res.Status
	k.StatusCode = res.StatusCode
	k.URL = req.URL
	ok = (res.StatusCode == http.StatusOK)
	if !ok {
		return k
	}
	doc, err = goquery.NewDocumentFromResponse(res)
	ok = (err == nil)
	if !ok {
		return k
	}
	s = doc.Find(htmlBody)
	ok = (s.Length() > 0)
	if !ok {
		return k
	}
	return newKit(s, k)
}

func getKitRetailers(s *goquery.Selection, k *Kit) {
	const (
		CSS string = ".retailers ul li"
	)
	var (
		ok        bool
		r         *Retailer
		retailers []*Retailer
	)
	s = s.Find(CSS)
	ok = (s.Length() > 0)
	if !ok {
		return
	}
	s.Each(func(i int, s *goquery.Selection) {
		r = newRetailer(s)
		ok = (r != nil)
		if !ok {
			return
		}
		retailers = append(retailers, r)
	})
	k.Retailers = retailers
}

func newKit(s *goquery.Selection, k *Kit) *Kit {
	for _, fn := range kitFn {
		fn(s, k)
	}
	return k
}
