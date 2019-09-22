package labo

import (
	"fmt"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

const (
	defaultRetailerAttrTarget string = "NIL"
)

const (
	errorRetailerEmptyImgChild string = "argument (*%p) does not contain an img tag"
)

const (
	retailerAnchorCSSSelector string = "a[href]"
	retailerLogoCSSSelector   string = "a img"
)

type Retailer struct {
	Href   string
	Logo   *Image
	Name   string
	Target string
}

func NewRetailer(s *goquery.Selection) (*Retailer, error) {
	if ok := (s != nil); !ok {
		return nil, fmt.Errorf(errorGoQuerySelectionNil)
	}
	if ok := (s.Length() > 0); !ok {
		return nil, fmt.Errorf(errorGoQuerySlectionEmptyHTMLNodes, s)
	}
	var (
		href   string
		logo   *Image
		name   string
		target string
	)
	hrefSelection := s.Find(retailerAnchorCSSSelector)
	if ok := (hrefSelection.Length() > 0); !ok {
		return nil, fmt.Errorf(errorGoQuerySlectionEmptyHTMLNodes, hrefSelection)
	}
	href, ok := hrefSelection.Attr(attrHref)
	if !ok {
		return nil, fmt.Errorf(errorEmptyHrefAlt, s)
	}
	logoSelection := s.Find(retailerLogoCSSSelector)
	if ok := (logoSelection.Length() > 0); !ok {
		return nil, fmt.Errorf(errorRetailerEmptyImgChild, s)
	}
	logo, err := NewImage(logoSelection)
	if err != nil {
		return nil, err
	}
	name, ok = s.Attr(attrClass)
	if !ok {
		return nil, fmt.Errorf(errorEmptyAttrClass, s)
	}
	name = strings.ToUpper(name)
	target = hrefSelection.AttrOr(attrTarget, defaultRetailerAttrTarget)
	target = strings.ToUpper(target)
	retailer := Retailer{
		Href:   href,
		Logo:   logo,
		Name:   name,
		Target: target}
	return &retailer, nil
}
