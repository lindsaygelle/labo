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

type KitCustomization struct {
	BoxImage  *Image
	Href      string
	Materials []*KitCustomization
	Price     *Price
}

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
		materials []*KitCustomization
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
