package labo

import (
	"fmt"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

const (
	kitBoxImageCSSSelector  string = "div.packshot picture img"
	kitNameCSSSelector      string = "h1.visually-hidden"
	kitRetailersCSSSelector string = "div.retailers ul li"
	kitPriceCSSSelector     string = "p.price"
)

// Kit is a struct that is the entirety of a Nintendo Labo Toy Con kit.
type Kit struct {
	BoxImage  *Image
	Href      string
	Materials *Materials
	Name      string
	Overview  *Overview
	Projects  []*Project
	Price     *Price
	Retailers []*Retailer
	Software  *Software
	ToyCons   []*ToyCon
}

// NewKit is a constructor function that instantiates a Nintendo Labo Kit struct pointer.
func NewKit(s *goquery.Selection) (*Kit, error) {
	if ok := (s != nil); !ok {
		return nil, fmt.Errorf(errorGoQuerySelectionNil)
	}
	if ok := (s.Length() > 0); !ok {
		return nil, fmt.Errorf(errorGoQuerySelectionEmptyHTMLNodes, s)
	}
	var (
		name      string
		retailers []*Retailer
	)
	var (
		boxImage, _  = NewImage(s.Find(kitBoxImageCSSSelector))
		materials, _ = NewMaterials(s.Find(materialsRootCSSSelector))
		overview, _  = NewOverview(s.Find(overviewRootCSSSelector))
		software, _  = NewSoftware(s.Find(softwareRootCSSSelector))
		price, e     = NewPrice(s.Find(kitPriceCSSSelector))
	)
	fmt.Println(e)
	nameSelection := s.Find(kitNameCSSSelector)
	if ok := (nameSelection.Length() > 0); ok {
		name = strings.TrimSpace(nameSelection.Text())
		name = strings.ToUpper(name)
	}
	retailersSelection := s.Find(kitRetailersCSSSelector)
	retailersSelection.Each(func(i int, s *goquery.Selection) {
		retailer, err := NewRetailer(s)
		if err != nil {
			return
		}
		retailers = append(retailers, retailer)
	})
	kit := Kit{
		BoxImage:  boxImage,
		Materials: materials,
		Name:      name,
		Overview:  overview,
		Price:     price,
		Retailers: retailers,
		Software:  software}
	return &kit, nil
}
