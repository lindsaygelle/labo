package labo

import (
	"fmt"

	"github.com/PuerkitoBio/goquery"
)

const (
	errorKitCustomizationNoBoxImage  string = ""
	errorKitCustomizationNoMaterials string = ""
	errorKitCustomizationNoPrice     string = ""
)

const (
	kitCustomizationBoxImageCSSSelector  string = "div.packshot picture img"
	kitCustomizationMaterialsCSSSelector string = "section.contents .row"
	kitCustomizationPriceCSSSelector     string = "div.price .price-card .price"
	kitCustomizationOverviewCSSSelector  string = "div.hero-content"
)

// KitCustomization is a struct that describes a subset of the Nintendo Labo Kits dedicated to customizing a Nintendo Labo Kit.
type KitCustomization struct {
	BoxImage  *Image
	Href      string
	Materials []*CustomizationPart
	Overview  *Overview
	Price     *Price
}

// NewKitCustomization is a constructor function that instantiates and returns a new KitCustomization struct pointer.
func NewKitCustomization(d *goquery.Document) (*KitCustomization, error) {
	var (
		ok bool
	)
	ok = (d != nil)
	if !ok {
		return nil, fmt.Errorf(errorGoQueryDocumentNil)
	}
	s := d.Find("html")
	ok = (s != nil)
	if !ok {
		return nil, fmt.Errorf(errorGoQuerySelectionNil)
	}
	ok = (s.Length() > 0)
	if !ok {
		return nil, fmt.Errorf(errorGoQuerySelectionEmptyHTMLNodes, s)
	}
	var (
		err error
	)
	var (
		boxImage  *Image
		materials []*CustomizationPart
		overview  *Overview
		price     *Price
	)
	imageSelection := s.Find(kitCustomizationBoxImageCSSSelector)
	if ok := (imageSelection.Length() > 0); !ok {
		return nil, fmt.Errorf(errorKitCustomizationNoBoxImage)
	}
	boxImage, err = NewImage(imageSelection)
	materialsSelection := s.Find(kitCustomizationMaterialsCSSSelector)
	ok = (materialsSelection.Length() > 0)
	if !ok {
		return nil, fmt.Errorf(errorKitCustomizationNoMaterials)
	}
	materialsSelection.Each(func(i int, s *goquery.Selection) {
		s.Find(".item").Each(func(j int, s *goquery.Selection) {
			customizationPart, err := NewCustomizationPart(s)
			if err != nil {
				return
			}
			materials = append(materials, customizationPart)
		})
	})
	overviewSelection := s.Find(kitCustomizationOverviewCSSSelector)
	ok = (overviewSelection.Length() > 0)
	if !ok {
		return nil, err
	}
	overview, err = NewOverview(overviewSelection)
	if err != nil {
		return nil, err
	}
	priceSelection := s.Find(kitCustomizationPriceCSSSelector)
	ok = (priceSelection.Length() > 0)
	if !ok {
		return nil, fmt.Errorf(errorKitCustomizationNoPrice)
	}
	price, err = NewPrice(priceSelection)
	if err != nil {
		return nil, err
	}
	kitCustomization := KitCustomization{
		BoxImage:  boxImage,
		Href:      d.Url.String(),
		Materials: materials,
		Overview:  overview,
		Price:     price}
	return &kitCustomization, nil
}
