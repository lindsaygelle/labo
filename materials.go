package labo

import (
	"fmt"

	"github.com/PuerkitoBio/goquery"
)

const (
	materialsImageCSSSelector string = ".kit-contents picture img"
	materialsPartsCSSSelector string = ".contents-content ul li"
	materialsRootCSSSelector  string = ".contents .right-column"
)

// Materials is the collection of parts needed to all projects within the Nintendo Labo Kit.
type Materials struct {
	Image *Image
	Parts []*Part
}

// NewMaterials is a constructor function that instantiates and returns a new Materials pointer.
func NewMaterials(s *goquery.Selection) (*Materials, error) {
	if ok := (s != nil); !ok {
		return nil, fmt.Errorf(errorGoQuerySelectionNil)
	}
	if ok := (s.Length() > 0); !ok {
		return nil, fmt.Errorf(errorGoQuerySelectionEmptyHTMLNodes, s)
	}
	var (
		image *Image
		parts []*Part
	)
	partsSelection := s.Find(materialsPartsCSSSelector)
	if ok := (partsSelection.Length() > 0); !ok {
		return nil, fmt.Errorf(errorGoQuerySelectionEmptyHTMLNodes, partsSelection)
	}
	partsSelection.Each(func(i int, s *goquery.Selection) {
		part, err := NewPart(s)
		if err != nil {
			return
		}
		parts = append(parts, part)
	})
	imageSelection := s.Find(materialsImageCSSSelector)
	if ok := (imageSelection.Length() > 0); !ok {
		return nil, fmt.Errorf(errorGoQuerySelectionEmptyHTMLNodes, imageSelection)
	}
	image, err := NewImage(imageSelection)
	if err != nil {
		return nil, err
	}
	materials := Materials{
		Image: image,
		Parts: parts}
	return &materials, nil
}
