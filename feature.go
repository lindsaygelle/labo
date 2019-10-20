package labo

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// Feature is a snapshot of a unique Nintendo Labo kit mechanic provided by a Nintendo Labo Toycon.
//
// Features are provided from the official Nintendo Labo website.
type Feature struct {
	Description string `json:"description"`
	Icon        *Image `json:"icon"`
	Name        string `json:"name"`
}

// getFeatureDescription searches the *goquery.Selection for the description required for a labo.Feature.
//
// Description is a short overview about the Nintendo Labo Toycon integration and the
// gameplay mechanics it offers. The description is provided from the official Nintendo Labo website.
//
// getFeatureDescription assigns a default description placeholder string if a description substring
// cannot be found.
func getFeatureDescription(s *goquery.Selection, f *Feature) {
	const (
		CSS string = ".copy > p"
	)
	var (
		description = defaultFeatureDescription
		ok          bool
		substring   string
	)
	s = s.Find(CSS)
	substring = strings.TrimSpace(s.Text())
	ok = (len(substring) > 0)
	if ok {
		description = substring
	}
	f.Description = description
}

// getFeatureIcon searches the *goquery.Selection for the *labo.Image required for a labo.Feature.
//
// Image is a HTML img reference that is provided from the official Nintendo Labo website. Icons
// are generally SVG image files.
//
// getFeatureIcon does not assign an empty *labo.Image should no HTML img tag can be found.
func getFeatureIcon(s *goquery.Selection, f *Feature) {
	const (
		CSS string = "picture:nth-child(1) > img:nth-child(1)"
	)
	var (
		ok bool
	)
	s = s.Find(CSS)
	ok = (s.Length() > 0)
	if !ok {
		return
	}
	f.Icon = newImage(s)
}

// getFeatureName searches the *goquery.Selection for the name required for a labo.Feature.
//
// Name is the namespace of the Nintendo Labo Toycon feature. Name may be a shorthand caption
// that gives a broad summary of the feature but not as verbose as the description.
//
// getFeatureName assigns a default name placeholder string if a name substring
// cannot be found.
func getFeatureName(s *goquery.Selection, f *Feature) {
	const (
		CSS string = ".header:nth-child(1) > span:nth-child(1)"
	)
	var (
		name      = defaultFeatureName
		ok        bool
		substring string
	)
	s = s.Find(CSS)
	substring = strings.TrimSpace(s.Text())
	ok = (len(substring) > 0)
	if ok {
		name = substring
	}
	f.Name = name
}

// getFeatureSelectionA searches the *goquery.Selection for the HTML content
// needed to get the Nintendo Labo feature icons.
func getFeatureSelectionA(s *goquery.Selection) *goquery.Selection {
	const (
		CSS string = ".slider-pagination:nth-child(1) > li"
	)
	var (
		ok bool
	)
	s = s.Find(CSS)
	ok = (s.Length() > 0)
	if !ok {
		return nil
	}
	return s
}

// getFeatureSelectionB searches the *goquery.Selection for the HTML content
// needed to get the Nintendo Labo feature main body.
func getFeatureSelectionB(s *goquery.Selection) *goquery.Selection {
	const (
		CSS string = ".slider-content > div"
	)
	var (
		ok bool
	)
	s = s.Find(CSS)
	ok = (s.Length() > 0)
	if !ok {
		return nil
	}
	return s
}

// getFeatureSelectionC searches the *goquery.Selection for the HTML content
// needed to get the Nintendo Labo feature name and description.
func getFeatureSelectionC(s *goquery.Selection) *goquery.Selection {
	const (
		CSS string = ".caption-content > div"
	)
	var (
		ok bool
	)
	s = s.Find(CSS)
	ok = (s.Length() > 0)
	if !ok {
		return nil
	}
	return s
}

// newFeature is a constructor function that instantiates and returns a new Feature struct pointer from
// the Nintendo Labo official website.
//
// newFeature requires a collection of specific HTML fragments that make up the entire contents
// of a Nintendo Labo Toycon feature. These are built from the getFeatureSelection lookups.
func newFeature(a, b, c *goquery.Selection) *Feature {
	var (
		feature = &Feature{}
	)
	getFeatureDescription(c, feature)
	getFeatureIcon(a, feature)
	getFeatureName(c, feature)

	return feature
}

// newFeatures is a constructor function that instantiates and returns a slice of Feature struct pointers.
func newFeatures(s *goquery.Selection) []*Feature {
	var (
		a  = getFeatureSelectionA(s)
		b  = getFeatureSelectionB(s)
		c  = getFeatureSelectionC(s)
		ok bool

		feature  *Feature
		features []*Feature
	)
	ok = (a != nil && b != nil && c != nil)
	if !ok {
		return features
	}
	ok = (a.Length() == b.Length())
	if !ok {
		return features
	}
	ok = (b.Length() == c.Length())
	if !ok {
		return features
	}
	a.Each(func(i int, _ *goquery.Selection) {
		feature = newFeature(a.Eq(i), b.Eq(i), c.Eq(i))
		ok = (feature != nil)
		if !ok {
			return
		}
		features = append(features, feature)
	})
	return features
}
