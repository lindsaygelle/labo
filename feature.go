package labo

import (
	"fmt"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

const (
	featureDescriptionCSSSelector string = ".copy p"
	featureIconCSSSelector        string = "button picture img"
	featureTitleCSSSelector       string = ".header span"
	featureVideoCSSSelector       string = "video"
)

// Feature is a struct that details a unique feature that the Nintendo Labo ToyCon supports.
type Feature struct {
	Description string
	Icon        *Image
	Title       string
	Video       *Video
}

// NewFeature is a constructor function that instantiates and returns a new Feature struct pointer.
func NewFeature(a, b, c *goquery.Selection) (*Feature, error) {
	var (
		ok bool
	)
	for _, s := range []*goquery.Selection{a, b, c} {
		ok = (s != nil)
		if !ok {
			return nil, fmt.Errorf(errorGoQuerySelectionNil)
		}
		ok = (s.Length() > 0)
		if !ok {
			return nil, fmt.Errorf(errorGoQuerySelectionEmptyHTMLNodes, s)
		}
	}
	var (
		description string
		icon        *Image
		title       string
		video       *Video
	)
	iconSelection := a.Find(featureIconCSSSelector)
	ok = (iconSelection.Length() > 0)
	if !ok {
		return nil, fmt.Errorf(errorGoQuerySelectionEmptyHTMLNodes, iconSelection)
	}
	icon, err := NewImage(iconSelection)
	if err != nil {
		return nil, err
	}
	descriptionSelection := c.Find(featureDescriptionCSSSelector)
	ok = (descriptionSelection.Length() > 0)
	if !ok {
		return nil, fmt.Errorf(errorGoQuerySelectionEmptyHTMLNodes, descriptionSelection)
	}
	description = strings.TrimSpace(descriptionSelection.Text())
	titleSelection := c.Find(featureTitleCSSSelector)
	ok = (titleSelection.Length() > 0)
	if !ok {
		return nil, fmt.Errorf(errorGoQuerySelectionEmptyHTMLNodes, titleSelection)
	}
	title = strings.ToUpper(titleSelection.Text())
	videoSelection := b.Find(featureVideoCSSSelector)
	video, _ = NewVideo(videoSelection)
	feature := Feature{
		Description: description,
		Icon:        icon,
		Title:       title,
		Video:       video}
	return &feature, nil
}