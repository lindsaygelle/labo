package labo

import (
	"github.com/PuerkitoBio/goquery"
)

type Kit struct {
	Href      string
	ID        string
	Materials *Materials
	Name      string
	Overview  Overview
	Projects  []Project
	Price     Price
	Retailers []Retailer
	Software  *Software
	ToyCons   []ToyCon
}

func NewKit(s *goquery.Selection) (*Kit, error) {
	var (
		materials, _ = NewMaterials(s.Find(materialsRootCSSSelector))
		software, _  = NewSoftware(s.Find(softwareRootCSSSelector))
	)
	kit := Kit{
		Materials: materials,
		Software:  software}
	return &kit, nil
}
