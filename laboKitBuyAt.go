package main

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type LaboKitBuyAt struct {
	Company         string
	CompanyImageURL string
	Href            string
	Target          string
	URL             *url.URL
}

func NewLaboKitBuyAt(s *goquery.Selection) *LaboKitBuyAt {
	return &LaboKitBuyAt{
		Company:         parseLaboKitBuyAtCompany(s),
		CompanyImageURL: parseLaboKitBuyAtCompayImageURL(s),
		Href:            parseLaboKitBuyAtHref(s),
		Target:          parseLaboKitBuyAtTarget(s),
		URL:             parseLaboKitBuyAtURL(s)}
}

func parseLaboKitBuyAtCompany(s *goquery.Selection) string {
	attribute := s.Find("img").AttrOr("alt", "NIL")
	substring := strings.TrimSpace(attribute)
	substring = strings.ReplaceAll(substring, " ", "-")
	return strings.ToUpper(substring)
}

func parseLaboKitBuyAtCompayImageURL(s *goquery.Selection) string {
	attribute := s.Find("img").AttrOr("src", "NIL")
	if ok := attribute != "NIL"; ok != true {
		return attribute
	}
	substring := strings.TrimSpace(attribute)
	substring = strings.ReplaceAll(substring, "../", "")
	return fmt.Sprintf("https://labo.nintendo.com/%s", substring)
}

func parseLaboKitBuyAtHref(s *goquery.Selection) string {
	attribute := s.AttrOr("href", "NIL")
	substring := strings.TrimSpace(attribute)
	return strings.ToUpper(substring)
}

func parseLaboKitBuyAtTarget(s *goquery.Selection) string {
	attribute := s.AttrOr("target", "NIL")
	substring := strings.TrimSpace(attribute)
	return strings.ToUpper(substring)
}

func parseLaboKitBuyAtURL(s *goquery.Selection) *url.URL {
	attribute := s.AttrOr("href", "NIL")
	URL, err := url.Parse(attribute)
	if err != nil {
		return new(url.URL)
	}
	return URL
}
