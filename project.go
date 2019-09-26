package labo

import (
	"fmt"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

const (
	projectIconCSSSelector        string = ".toycon-icon img"
	projectImageCSSSelector       string = ".toycon-image picture img"
	projectNameCSSSelector        string = ".toycon-icon p"
	projectScreenshotCSSSelector  string = "picture img"
	projectScreenshotsCSSSelector string = ".screenshot"
)

// Project is a unique build project that is bundled with a Nintendo Labo Kit.
type Project struct {
	Icon        *Image
	Image       *Image
	Name        string
	Screenshots []*Image
}

// NewProject is a constructor function that instantiates and returns a new Project struct pointer.
func NewProject(s *goquery.Selection) (*Project, error) {
	if ok := (s != nil); !ok {
		return nil, fmt.Errorf(errorGoQuerySelectionNil)
	}
	if ok := (s.Length() > 0); !ok {
		return nil, fmt.Errorf(errorGoQuerySelectionEmptyHTMLNodes, s)
	}
	var (
		name        string
		screenshots []*Image
	)
	var (
		icon, _  = NewImage(s.Find(projectIconCSSSelector))
		image, _ = NewImage(s.Find(projectImageCSSSelector))
	)
	nameSelection := s.Find(projectNameCSSSelector)
	name = strings.TrimSpace(nameSelection.Text())
	name = strings.ToUpper(name)
	screenshotsSelection := s.Find(projectScreenshotsCSSSelector)
	screenshotsSelection.Each(func(i int, s *goquery.Selection) {
		screenshot, err := NewImage(s.Find(projectScreenshotCSSSelector))
		if err != nil {
			return
		}
		screenshots = append(screenshots, screenshot)
	})
	project := Project{
		Icon:        icon,
		Image:       image,
		Name:        name,
		Screenshots: screenshots}
	return &project, nil
}
