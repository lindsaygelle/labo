package labo

import (
	"fmt"

	"github.com/PuerkitoBio/goquery"
)

type Materials struct {
	Image *Image
	Parts []*Part
}

func NewMaterials(s *goquery.Selection) (*Materials, error) {
	if ok := (s.Length() > 0); !ok {
		return nil, fmt.Errorf("goquery.Selection is empty")
	}
	partsSelection := s.Find(".contents-content ul li")
	if ok := (partsSelection.Length() > 0); !ok {
		return nil, fmt.Errorf("goquery.Selection is empty")
	}
	parts := []*Part{}
	partsSelection.Each(func(i int, s *goquery.Selection) {
		part, err := NewPart(s)
		if err != nil {
			return
		}
		parts = append(parts, part)
	})
	imageSelection := s.Find(".kit-contents picture img")
	if ok := (imageSelection.Length() > 0); !ok {
		return nil, fmt.Errorf("goquery.Selection is empty")
	}
	image, err := NewImage(s)
	if err != nil {
		return nil, err
	}
	materials := Materials{
		Image: image,
		Parts: parts}
	return &materials, nil
}
