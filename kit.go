package labo

import (
	"fmt"

	"github.com/PuerkitoBio/goquery"
)

const (
	kitRetailersCSSSelector string = ".retailers ul li"
)

// Kit is a struct that is the entirety of a Nintendo Labo Toy Con kit.
type Kit struct {
	Href      string
	ID        string
	Materials *Materials
	Name      string
	Overview  *Overview
	Projects  []*Project
	Price     Price
	Retailers []*Retailer
	Software  *Software
	ToyCons   []*ToyCon
}

func NewKit(s *goquery.Selection) (*Kit, error) {
	if ok := (s != nil); !ok {
		return nil, fmt.Errorf(errorGoQuerySelectionNil)
	}
	if ok := (s.Length() > 0); !ok {
		return nil, fmt.Errorf(errorGoQuerySelectionEmptyHTMLNodes, s)
	}
	var (
		retailers []*Retailer
	)
	var (
		materials, _ = NewMaterials(s.Find(materialsRootCSSSelector))
		software, _  = NewSoftware(s.Find(softwareRootCSSSelector))
	)
	retailersSelection := s.Find(kitRetailersCSSSelector)
	if ok := (retailersSelection.Length() > 0); !ok {

	}
	retailersSelection.Each(func(i int, s *goquery.Selection) {
		retailer, err := NewRetailer(s)
		if err != nil {
			return
		}
		retailers = append(retailers, retailer)
	})
	kit := Kit{
		Materials: materials,
		Retailers: retailers,
		Software:  software}
	return &kit, nil
}
