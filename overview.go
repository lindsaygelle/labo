package labo

import (
	"fmt"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

const (
	defaultOverviewDescription string = "NIL"
)

const (
	overviewDescriptionCSSSelector string = "p.soprano"
	overviewRootCSSSelector        string = ".hero-content .kit.column"
	overviewVideoCSSSelector       string = ".content[data-id]"
)

// Overview is a struct that gives a short textual overview (with an optional Video) about a Nintendo Labo Kit.
type Overview struct {
	Description string
	Video       *Video
}

// NewOverview is a constructor function that instantiates and returns a new Overview struct pointer.
func NewOverview(s *goquery.Selection) (*Overview, error) {
	if ok := (s != nil); !ok {
		return nil, fmt.Errorf(errorGoQuerySelectionNil)
	}
	if ok := (s.Length() > 0); !ok {
		return nil, fmt.Errorf(errorGoQuerySelectionEmptyHTMLNodes, s)
	}
	var (
		description string
		video       *Video
	)
	descriptionSelection := s.Find(overviewDescriptionCSSSelector)
	if ok := (descriptionSelection.Length() > 0); ok {
		description = strings.TrimSpace(descriptionSelection.Text())
	}
	if ok := (len(description) > 0); !ok {
		description = defaultOverviewDescription
	}
	videoSelection := s.Find(overviewVideoCSSSelector)
	if ok := (videoSelection.Length() > 0); ok {
		video, _ = NewVideo(videoSelection)
	}
	overview := Overview{
		Description: description,
		Video:       video}
	return &overview, nil
}
