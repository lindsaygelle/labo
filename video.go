package labo

import (
	"fmt"

	"github.com/PuerkitoBio/goquery"
)

// Video is a HTML5 video resource that contains content for a Nintendo Labo Kit.
type Video struct{}

// NewVideo is a constructor function that instantiates and returns a new Video struct pointer.
func NewVideo(s *goquery.Selection) (*Video, error) {
	if ok := (s != nil); !ok {
		return nil, fmt.Errorf(errorGoQuerySelectionNil)
	}
	if ok := (s.Length() > 0); !ok {
		return nil, fmt.Errorf(errorGoQuerySelectionEmptyHTMLNodes, s)
	}
	video := Video{}
	return &video, nil
}
