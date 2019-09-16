package main

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type LaboKitBuyAt struct {
	Company          string   `json:"company"`
	CompanyImageHref string   `json:"company_image_href"`
	CompanyImageURL  *url.URL `json:"company_image_URL"`
	Href             string   `json:"href"`
	Target           string   `json:"target"`
	URL              *url.URL `json:"URL"`
}

func NewLaboKitBuyAt(s *goquery.Selection) *LaboKitBuyAt {
	return &LaboKitBuyAt{
		Company:          parseLaboKitBuyAtCompany(s),
		CompanyImageHref: parseLaboKitBuyAtCompayImageHref(s),
		CompanyImageURL:  parseLaboKitBuyAtCompanyImageURL(s),
		Href:             parseLaboKitBuyAtHref(s),
		Target:           parseLaboKitBuyAtTarget(s),
		URL:              parseLaboKitBuyAtURL(s)}
}

func parseLaboKitBuyAtCompany(s *goquery.Selection) string {
	attribute := s.Find("img").AttrOr("alt", "NIL")
	substring := strings.TrimSpace(attribute)
	substring = strings.ReplaceAll(substring, " ", "-")
	return strings.ToUpper(substring)
}

func parseLaboKitBuyAtCompayImageHref(s *goquery.Selection) string {
	attribute := s.Find("img").AttrOr("src", "NIL")
	if ok := attribute != "NIL"; ok != true {
		return attribute
	}
	substring := strings.TrimSpace(attribute)
	substring = strings.ReplaceAll(substring, "../", "")
	return fmt.Sprintf("https://labo.nintendo.com/%s", substring)
}

func parseLaboKitBuyAtCompanyImageURL(s *goquery.Selection) *url.URL {
	attribute := s.Find("img").AttrOr("src", "NIL")
	URL, err := url.Parse(attribute)
	if err != nil {
		return new(url.URL)
	}
	return URL
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
