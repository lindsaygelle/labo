package labo

import (
	"fmt"

	"github.com/PuerkitoBio/goquery"
)

const (
	softwareImageCSSSelector string = ".packshot picture img"
	softwareRootCSSSelector  string = ".contents .left-column"
	softwareVideoCSSSelector string = ".software-video"
)

// Software is a struct that contains a brief overview about the software provided with a Nintendo Labo Kit.
type Software struct {
	Image *Image
	Video *Video
}

// NewSoftware is a constructor function that instantiates and returns a new Software pointer.
func NewSoftware(s *goquery.Selection) (*Software, error) {
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
		image *Image
		video *Video
	)
	image, err := NewImage(s.Find(softwareImageCSSSelector))
	if err != nil {
		return nil, err
	}
	videoSelection := s.Find(softwareVideoCSSSelector)
	if ok := (videoSelection.Length() > 0); ok {
		video, _ = NewVideo(videoSelection)
	}
	software := Software{
		Image: image,
		Video: video}
	return &software, nil
}
