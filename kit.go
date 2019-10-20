package labo

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"golang.org/x/text/language"
)

// Kit is a Nintendo Labo Kit. Kit structs are built from the Nintendo Labo official website.
//
// Kits contain extended information not found on the Nintendo Labo store and may contain
// varying levels of detail depending on the type of page scraped and the category of
// the Nintendo Labo Kit. Non VR kits will not contain the corresponding VR
// information.
type Kit struct {
	Category    string       `json:"category"`
	CategoryID  string       `json:"category_ID"`
	BoxImage    *Image       `json:"box_image"`
	ID          int          `json:"ID"`
	IsVR        bool         `json:"is_VR"`
	KitImage    *Image       `json:"kit_image"`
	Language    language.Tag `json:"language"`
	Projects    []*Project   `json:"projects"`
	Retailers   []*Retailer  `json:"retailers"`
	SoftwareBox *Image       `json:"software_box"`
	Status      string       `json:"status"`
	StatusCode  int          `json:"status_code"`
	Toycons     []*Toycon    `json:"toycons"`
	URL         *url.URL     `json:"URL"`
}

var (
	kitFn = [](func(s *goquery.Selection, k *Kit)){
		getKitBoxImage,
		getKitImage,
		getKitProjects,
		getKitRetailers,
		getKitToycons}
)

// GetKit gets the extended Nintendo Labo Kit information from the official Nintendo Labo website.
//
// The argument Labo must contain a non NIL string Labo.Ref to perform the look-up on the Nintendo Labo website.
// Some products on the Nintendo Labo store do not contain a official Nintendo Labo website
// counterpart and will return a non-nil Kit pointer with the HTTP status code and status fields being
// set to a http.StatusBadRequest. Successful lookups should contain a Labo.StatusCode value of http.StatusOK.
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
	ok = (l != nil)
	if !ok {
		return k
	}
	k.Category = l.Category
	k.CategoryID = l.CategoryID
	k.ID = l.ID
	k.IsVR = strings.Contains(l.Name, "vr")
	k.Language = l.Language
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

// MarshalKit marshals a Kit struct into an ordered byte sequence. On error returns an empty byte slice.
func MarshalKit(k *Kit) (b []byte) {
	b, _ = json.Marshal(k)
	return b
}

// getKitBoxImage searches the argument goquery.Selection pointer
// for the Nintendo Labo Kit product box image and assign it to
// the argument labo.Kit pointer if found.
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

// getKitImage searches the argument goquery.Selection pointer
// for the Nintendo Labo Kit image and assign it to
// the argument labo.Kit pointer if found.
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

// getKitProjects searches the argument *goquery.Selection
// for the Nintendo Labo Kit project information hosted by the Nintendo Labo
// official website.
func getKitProjects(s *goquery.Selection, k *Kit) {
	const (
		CSS string = ".main-toycon:nth-child(1) > .toycon-tag"
	)
	var (
		ok bool
	)
	s = s.Find(CSS)
	ok = (s.Length() > 0)
	if !ok {
		return
	}
	k.Projects = newProjects(s)
}

// getKitToycons searches the argument *goquery.Selection
// for the Nintendo Labo Kit Toycon information hosted by the Nintendo Labo
// official website.
func getKitToycons(s *goquery.Selection, k *Kit) {
	const (
		CSS string = ".toy-con-area .toy-con"
	)
	var (
		ok bool
	)
	s = s.Find(CSS)
	ok = (s.Length() > 0)
	if !ok {
		return
	}
	k.Toycons = newToycons(s)
}

// getKitRetailers searches the argument *goquery.Selection
// for the Nintendo Labo Kit retailers hosted by the Nintendo Labo
// official website.
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
