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

// Software is
type Software struct {
	Image *Image
	Video *Video
}

// NewSoftware is a constructor function the instantiates a new Software pointer.
func NewSoftware(s *goquery.Selection) (*Software, error) {
	var (
		image *Image
		video *Video
	)
	if ok := (s != nil); !ok {
		return nil, fmt.Errorf(errorGoQuerySelectionNil)
	}
	if ok := (s.Length() > 0); !ok {
		return nil, fmt.Errorf(errorGoQuerySlectionEmptyHTMLNodes, s)
	}
	imageSelection := s.Find(softwareImageCSSSelector)
	fmt.Println(imageSelection.Attr("data-src"))
	if ok := (s.Length() > 0); !ok {
		return nil, fmt.Errorf(errorGoQuerySlectionEmptyHTMLNodes, imageSelection)
	}
	image, err := NewImage(s)
	fmt.Println(err)
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
