package labo

import (
	"github.com/PuerkitoBio/goquery"
)

// Video is a HTML5 video resource that contains content for a Nintendo Labo Kit.
type Video struct{}

// NewVideo is a constructor function that instantiates and returns a new Video struct pointer.
func NewVideo(s *goquery.Selection) (*Video, error) {

	video := Video{}
	return &video, nil
}
