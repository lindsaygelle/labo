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
)

// KitCustomization is a struct that describes a subset of the Nintendo Labo Kits dedicated to customizing a Nintendo Labo Kit.
type KitCustomization struct {
	BoxImage  *Image
	Href      string
	Materials []*CustomizationPart
	Price     *Price
}

// NewKitCustomization is a constructor function that instantiates and returns a new KitCustomization struct pointer.
func NewKitCustomization(s *goquery.Selection) (*KitCustomization, error) {
	if ok := (s != nil); !ok {
		return nil, fmt.Errorf(errorGoQuerySelectionNil)
	}
	if ok := (s.Length() > 0); !ok {
		return nil, fmt.Errorf(errorGoQuerySelectionEmptyHTMLNodes, s)
	}
	var (
		boxImage  *Image
		href      string
		materials []*CustomizationPart
		price     *Price
	)
	imageSelection := s.Find(kitCustomizationBoxImageCSSSelector)
	if ok := (imageSelection.Length() > 0); !ok {
		return nil, fmt.Errorf(errorKitCustomizationNoBoxImage)
	}
	materialsSelection := s.Find(kitCustomizationMaterialsCSSSelector)
	if ok := (materialsSelection.Length() > 0); !ok {
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
	priceSelection := s.Find(kitCustomizationPriceCSSSelector)
	if ok := (priceSelection.Length() > 0); !ok {
		return nil, fmt.Errorf(errorKitCustomizationNoPrice)
	}
	kitCustomization := KitCustomization{
		BoxImage:  boxImage,
		Href:      href,
		Materials: materials,
		Price:     price}
	return &kitCustomization, nil
}
