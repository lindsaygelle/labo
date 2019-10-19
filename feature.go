package labo

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type Feature struct {
	Description string `json:"description"`
	Icon        *Image `json:"icon"`
	Name        string `json:"name"`
}

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

func newFeature(a, b, c *goquery.Selection) *Feature {
	var (
		feature = &Feature{}
	)
	getFeatureDescription(c, feature)
	getFeatureIcon(a, feature)
	getFeatureName(c, feature)

	return feature
}

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
