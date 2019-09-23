package labo

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

const (
	errorToyConEmptyName string = "argument (*%p) does not contain text nodes"
)
const (
	toyConDescriptionCSSSelector string = ".left-column .toy-con-sub-header"
	toyConNameCSSSelector        string = ".left-column .toy-con-header"
)

var (
	regexpToyConMatchToyCon = regexp.MustCompile(`Toy-Con`)
)

type ToyCon struct {
	Description string
	Features    []*Feature
	Icon        *Image
	Image       *Image
	Name        string
}

func NewToyCon(s *goquery.Selection) (*ToyCon, error) {
	if ok := (s != nil); !ok {
		return nil, fmt.Errorf(errorGoQuerySelectionNil)
	}
	if ok := (s.Length() > 0); !ok {
		return nil, fmt.Errorf(errorGoQuerySelectionEmptyHTMLNodes, s)
	}
	var (
		description string
		name        string
	)
	nameSelection := s.Find(toyConNameCSSSelector)
	if ok := (nameSelection.Length() > 0); !ok {
		return nil, fmt.Errorf(errorGoQuerySelectionEmptyHTMLNodes, nameSelection)
	}
	name = nameSelection.Text()
	if ok := (len(name) > 0); !ok {
		return nil, fmt.Errorf(errorToyConEmptyName, nameSelection)
	}
	name = regexpToyConMatchToyCon.ReplaceAllString(name, "")
	name = strings.TrimSpace(name)
	name = strings.ToUpper(name)
	descriptionSelection := s.Find(toyConDescriptionCSSSelector)
	if ok := (descriptionSelection.Length() > 0); ok {
		description = strings.TrimSpace(descriptionSelection.Text())
		description = regexpMatchLineBreaks.ReplaceAllString(description, " ")
		description = regexpMatchSequenceWhitespace.ReplaceAllString(description, "")
	}
	toyCon := ToyCon{
		Description: description,
		Name:        name}
	return &toyCon, nil
}
