package labo

import (
	"github.com/PuerkitoBio/goquery"
)

type Feature struct {
	Description string
	Icon        *Image
	Name        string
}

func getFeatureSelection(CSS string, s *goquery.Selection) *goquery.Selection {
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

func getFeatureSelectionA(s *goquery.Selection) *goquery.Selection {
	const (
		CSS string = ".slider-pagination:nth-child(1) > li"
	)
	return getFeatureSelection(CSS, s)
}

func getFeatureSelectionB(s *goquery.Selection) *goquery.Selection {
	const (
		CSS string = ".slider-content:nth-child(1) > div"
	)
	return getFeatureSelection(CSS, s)
}

func getFeatureSelectionC(s *goquery.Selection) *goquery.Selection {
	const (
		CSS string = ".caption-content:nth-child(1) > div"
	)
	return getFeatureSelection(CSS, s)
}

func newFeature(a, b, c *goquery.Selection) {}

func newFeatures(s *goquery.Selection) []*Feature {
	var (
		a  = getFeatureSelectionA(s)
		b  = getFeatureSelectionB(s)
		c  = getFeatureSelectionC(s)
		ok bool

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

	})
	return features
}
