package labo

import (
	"github.com/PuerkitoBio/goquery"
)

type ToyCon struct {
	Description string
	Features    []*Feature
	Image       *Image
	Name        string
}

func NewToyCon(s *goquery.Selection) {

}
