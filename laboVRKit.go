package labo

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type LaboVRKit struct {
	BoxImage      LaboBoxImage  `json:"box_image"`
	BoxImageURL   string        `json:"box_image_URL"`
	BuyAt         []interface{} `json:"buy_at"`
	Overview      string        `json:"overview"`
	Price         float64       `json:"price"`
	PriceCurrency string        `json:"price_currency"`
}

func NewLaboVRKit(s *goquery.Selection) LaboVRKit {

	return LaboVRKit{
		Overview:      parseLaboVRKitOverview(s),
		Price:         parseLaboVRKitPrice(s),
		PriceCurrency: "USD"}
}

func parseLaboVRKitBoxImageURL(s *goquery.Selection) string {
	s.Find("div.kit-description picture img")
	return ""
}

func parseLaboVRKitOverview(s *goquery.Selection) string {
	var paragraphs []string
	s.Find("div.kit-overview > p").Each(func(i int, s *goquery.Selection) {
		paragraphs = append(paragraphs, strings.TrimSpace(s.Text()))
	})
	overview := strings.Join(paragraphs, " ")
	return overview
}

func parseLaboVRKitPrice(s *goquery.Selection) float64 {
	substring := strings.TrimSpace(s.Find("div.kit-pricing p.price").First().Text())
	substring = regexp.MustCompile(`([0-9]+\.[0-9]+)`).FindString(substring)
	price, err := strconv.ParseFloat(substring, 64)
	if err != nil {
		return -1
	}
	return price
}
