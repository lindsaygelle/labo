package labo

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

const (
	customizationPartImageCSSSelector string = "picture img"
)

// CustomizationPart is a struct that represents a part available for adding visual customization for a Nintendo Labo Kit.
type CustomizationPart struct {
	Amount int
	Image  *Image
	Name   string
}

// NewCustomizationPart is a constructor function that instantiates and returns a new CustomizationPart struct pointer.
func NewCustomizationPart(s *goquery.Selection) (*CustomizationPart, error) {
	if ok := (s != nil); !ok {
		return nil, fmt.Errorf(errorGoQuerySelectionNil)
	}
	if ok := (s.Length() > 0); !ok {
		return nil, fmt.Errorf(errorGoQuerySelectionEmptyHTMLNodes, s)
	}
	var (
		amount int
		image  *Image
		name   string
	)
	imageSelection := s.Find(customizationPartImageCSSSelector)
	if ok := (imageSelection.Length() > 0); !ok {
		return nil, fmt.Errorf("")
	}
	contents := strings.TrimSpace(s.Text())
	if ok := (len(contents) > 0); !ok {
		return nil, fmt.Errorf(errorPartEmptyText, s)
	}
	substring := regexpPartFindPartName.FindString(contents)
	if ok := (len(substring) > 0); !ok {
		return nil, fmt.Errorf(errorPartEmptyPartName, s)
	}
	substring = strings.TrimSpace(substring)
	substring = strings.TrimSuffix(substring, "x")
	substring = regexpMatchParenthesis.ReplaceAllString(substring, "")
	name = regexpMatchSequenceWhitespace.ReplaceAllString(substring, "")
	name = strings.TrimSpace(name)
	name = strings.ToUpper(name)
	substring = regexpMatchNumeric.FindString(contents)
	if ok := (len(substring) > 0); ok {
		amount, _ = strconv.Atoi(substring)
	}
	customizationPart := CustomizationPart{
		Amount: amount,
		Image:  image,
		Name:   name}
	return &customizationPart, nil
}
