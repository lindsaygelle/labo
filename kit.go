package labo

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/PuerkitoBio/goquery"
)

type Kit struct {
	BoxImage    *Image
	KitImage    *Image
	Retailers   []*Retailer
	SoftwareBox *Image
	Status      string
	StatusCode  int
	Toycons     []*Toycon
	URL         *url.URL
}

var (
	kitFn = [](func(s *goquery.Selection, k *Kit)){
		getKitBoxImage,
		getKitImage,
		getKitRetailers,
		getKitToyCons}
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

func getKitBoxImage(s *goquery.Selection, k *Kit) {
	const (
		CSS string = ".product-hero .hero-content .kit.column .packshot picture img"
	)
	var (
		ok bool
	)
	s = s.Find(CSS)
	ok = (s.Length() > 0)
	if !ok {
		return
	}
	k.BoxImage = newImage(s)
}

func getKitImage(s *goquery.Selection, k *Kit) {
	const (
		CSS string = ".kit-contents > picture:nth-child(1) > img:nth-child(1)"
	)
	var (
		ok bool
	)
	s = s.Find(CSS)
	ok = (s.Length() > 0)
	if !ok {
		return
	}
	k.KitImage = newImage(s)
}

func getKitToyCons(s *goquery.Selection, k *Kit) {
	const (
		CSS string = ".toy-con-area > div:nth-child(1) > div"
	)
	var (
		ok bool
	)
	s = s.Find(CSS)
	ok = (s.Length() > 0)
	if !ok {
		return
	}
	k.ToyCons = newToyCons(s)
}

func getKitRetailers(s *goquery.Selection, k *Kit) {
	const (
		CSS string = ".retailers > ul:nth-child(1) li"
	)
	var (
		ok bool
	)
	s = s.Find(CSS)
	ok = (s.Length() > 0)
	if !ok {
		return
	}
	k.Retailers = newRetailers(s)
}

func getKitSoftwareBox(s *goquery.Selection, k *Kit) {
	const (
		CSS string = ".content:nth-child(2) > div:nth-child(1) > picture:nth-child(1) > img:nth-child(1)"
	)
	var (
		ok bool
	)
	s = s.Find(CSS)
	ok = (s.Length() > 0)
	if !ok {
		return
	}
	k.SoftwareBox = newImage(s)
}

func newKit(s *goquery.Selection, k *Kit) *Kit {
	for _, fn := range kitFn {
		fn(s, k)
	}
	return k
}
