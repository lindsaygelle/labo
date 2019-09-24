package labo

import (
	"fmt"

	"github.com/PuerkitoBio/goquery"
)

type Feature struct {
	Description string
	Icon        *Image
	Title       string
	Video       Video
}

func NewFeature(a, b, c *goquery.Selection) (*Feature, error) {
	var (
		description string
		icon        *Image
		title       string
	)
	for _, s := range []*goquery.Selection{a, b, c} {
		if ok := (s != nil); !ok {
			return nil, fmt.Errorf(errorGoQuerySelectionNil)
		}
		if ok := (s.Length() > 0); !ok {
			return nil, fmt.Errorf(errorGoQuerySelectionEmptyHTMLNodes, s)
		}
	}
	iconSelection := a.Find("button picture img")
	if ok := (iconSelection.Length() > 0); !ok {
		return nil, fmt.Errorf(errorGoQuerySelectionEmptyHTMLNodes, iconSelection)
	}
	icon, err := NewImage(iconSelection)
	if err != nil {
		return nil, err
	}
	descriptionSelection := c.Find(".copy p")
	if ok := (descriptionSelection.Length() > 0); !ok {
		return nil, fmt.Errorf(errorGoQuerySelectionEmptyHTMLNodes, descriptionSelection)
	}
	description = descriptionSelection.Text()
	titleSelection := c.Find(".header span")
	if ok := (titleSelection.Length() > 0); !ok {
		return nil, fmt.Errorf(errorGoQuerySelectionEmptyHTMLNodes, titleSelection)
	}
	title = titleSelection.Text()
	feature := Feature{
		Description: description,
		Icon:        icon,
		Title:       title}
	return &feature, nil
}
