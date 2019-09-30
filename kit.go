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

// Kit is a struct that details all the unique the specifics of a Nintendo Labo Kit.
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

// NewKit is a constructor function that instantiates and returns a Nintendo Labo Kit struct pointer.
func NewKit(d *goquery.Document) (*Kit, error) {
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
	ok = (nameSelection.Length() > 0)
	if ok {
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
		Href:      d.Url.String(),
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
