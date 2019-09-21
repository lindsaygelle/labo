package labo

import (
	"github.com/PuerkitoBio/goquery"
)

type Video struct{}

func NewVideo(s *goquery.Selection) (*Video, error) {

	video := Video{}
	return &video, nil
}
