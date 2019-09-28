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
	toyConFeatureRootCSSSelector string = ".toy-con-slider"
	toyConFeatureS1CSSSelector   string = ".slider-pagination li"
	toyConFeatureS2CSSSelector   string = ".slider-content > div"
	toyConFeatureS3CSSSelector   string = ".caption-content > div"
	toyConIconCSSSelector        string = ".icon img"
	toyConImageCSSSelector       string = ".main-image picture img"
	toyConNameCSSSelector        string = ".left-column .toy-con-header"
	toyConVerboseCSSSelector     string = ".toy-con-info .copy"
)

var (
	regexpToyConMatchToyCon = regexp.MustCompile(`Toy-Con`)
)

// ToyCon is a struct that contains the details about a Nintendo Labo ToyCon.
// A ToyCon is one of the kit specific projects that can be built with a Nintendo Labo Kit.
// Each Nintendo Labo kit will come with a unique series of ToyCons.
type ToyCon struct {
	Description string
	Features    []*Feature
	Icon        *Image
	Image       *Image
	Name        string
	Verbose     string
}

// NewToyCon is a constructor function that instantiates and returns a new ToyCon struct pointer.
func NewToyCon(s *goquery.Selection) (*ToyCon, error) {
	var (
		ok bool
	)
	ok = (s != nil)
	if !ok {
		return nil, fmt.Errorf(errorGoQuerySelectionNil)
	}
	ok = (s.Length() > 0)
	if !ok {
		return nil, fmt.Errorf(errorGoQuerySelectionEmptyHTMLNodes, s)
	}
	var (
		description string
		features    []*Feature
		name        string
		verbose     string
	)
	var (
		image, _ = NewImage(s.Find(toyConImageCSSSelector))
		icon, _  = NewImage(s.Find(toyConIconCSSSelector))
	)
	nameSelection := s.Find(toyConNameCSSSelector)
	ok = (nameSelection.Length() > 0)
	if !ok {
		return nil, fmt.Errorf(errorGoQuerySelectionEmptyHTMLNodes, nameSelection)
	}
	name = nameSelection.Text()
	ok = (len(name) > 0)
	if !ok {
		return nil, fmt.Errorf(errorToyConEmptyName, nameSelection)
	}
	name = regexpToyConMatchToyCon.ReplaceAllString(name, "")
	name = strings.TrimSpace(name)
	name = strings.ToUpper(name)
	descriptionSelection := s.Find(toyConDescriptionCSSSelector)
	ok = (descriptionSelection.Length() > 0)
	if ok {
		description = strings.TrimSpace(descriptionSelection.Text())
		description = regexpMatchLineBreaks.ReplaceAllString(description, " ")
		description = regexpMatchSequenceWhitespace.ReplaceAllString(description, "")
		description = strings.ToUpper(description)
	}

	verboseSelection := s.Find(toyConVerboseCSSSelector)
	ok = (verboseSelection.Length() > 0)
	if ok {
		verbose = strings.TrimSpace(verboseSelection.Text())
		verbose = regexpMatchLineBreaks.ReplaceAllString(verbose, " ")
		verbose = regexpMatchSequenceWhitespace.ReplaceAllString(verbose, "")
		verbose = strings.ToUpper(verbose)
	}
	s.Find(toyConFeatureRootCSSSelector).First().Each(func(i int, s *goquery.Selection) {
		s1 := s.Find(toyConFeatureS1CSSSelector)
		s2 := s.Find(toyConFeatureS2CSSSelector)
		s3 := s.Find(toyConFeatureS3CSSSelector)
		s1.Each(func(i int, _ *goquery.Selection) {
			featue, err := NewFeature(s1.Eq(i), s2.Eq(i), s3.Eq(i))
			if err != nil {
				return
			}
			features = append(features, featue)
		})
	})
	toyCon := ToyCon{
		Description: description,
		Features:    features,
		Icon:        icon,
		Image:       image,
		Name:        name,
		Verbose:     verbose}
	return &toyCon, nil
}
