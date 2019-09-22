package labo

import (
	"github.com/PuerkitoBio/goquery"
)

type Overview struct {
	Description string
	Video       *Video
}

func NewOverview(s *goquery.Selection) {

}
