package labo

import (
	"fmt"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

const (
	kitBoxImageCSSSelector  string = "div.packshot picture img"
	kitNameCSSSelector      string = "h1.visually-hidden"
	kitPriceCSSSelector     string = "p.price"
	kitProjectsCSSSelector  string = ".toycon-container .toycon-tag"
	kitRetailersCSSSelector string = "div.retailers ul li"
	kitToyConCSSSelector    string = "section.toy-con-area .toy-con"
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
		projects  []*Project
		retailers []*Retailer
		toyCons   []*ToyCon
	)
	var (
		boxImage, _  = NewImage(s.Find(kitBoxImageCSSSelector))
		materials, _ = NewMaterials(s.Find(materialsRootCSSSelector))
		overview, _  = NewOverview(s.Find(overviewRootCSSSelector))
		software, _  = NewSoftware(s.Find(softwareRootCSSSelector))
		price, _     = NewPrice(s.Find(kitPriceCSSSelector))
	)
	projectsSelection := s.Find(kitProjectsCSSSelector)
	projectsSelection.Each(func(i int, s *goquery.Selection) {
		project, err := NewProject(s)
		if err != nil {
			return
		}
		projects = append(projects, project)
	})
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
	toyConSelection := s.Find(kitToyConCSSSelector)
	toyConSelection.Each(func(i int, s *goquery.Selection) {
		toyCon, err := NewToyCon(s)
		if err != nil {
			return
		}
		toyCons = append(toyCons, toyCon)
	})
	kit := Kit{
		BoxImage:  boxImage,
		Materials: materials,
		Name:      name,
		Overview:  overview,
		Price:     price,
		Projects:  projects,
		Retailers: retailers,
		Software:  software,
		ToyCons:   toyCons}
	return &kit, nil
}
