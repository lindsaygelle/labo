package main

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type LaboVRKit struct {
	Available       bool            `json:"available"`
	BuyAt           []*LaboKitBuyAt `json:"buy_at"`
	BoxImageHref    string          `json:"box_image_href"`
	BoxImageURL     *url.URL        `json:"box_image_URL"`
	BoxImageSrcSet  []interface{}   `json:"box_image_srcset"`
	Compatibility   []interface{}   `json:"compatibility"`
	HeroImageHref   string          `json:"hero_image_href"`
	HeroImageURL    *url.URL        `json:"hero_image_URL"`
	HeroImageSrcSet []interface{}   `json:"hero_image_srcset"`
	KitContents     interface{}     `json:"kit_contents"`
	KitFeatures     []interface{}   `json:"kit_features"`
	Projects        []interface{}   `json:"projects"`
	Overview        string          `json:"overview"`
	Quotes          []interface{}   `json:"quotes"`
	VRPlaza         interface{}     `json:"VR_plaza"`
}

func NewLaboVRKit(d *goquery.Document) *LaboVRKit {

	return &LaboVRKit{
		Available:    scrapeLaboVRKitAvailable(d),
		BuyAt:        scrapeLaboVRKitBuyAt(d),
		BoxImageHref: scrapeLaboVRKitBoxImageHref(d)}
}

func scrapeLaboVRKitAvailable(d *goquery.Document) bool {
	ok := false
	CSS := "div.kit-overview h2.tenor"
	s := (d.Find(CSS).First())
	if s.Length() == 0 {
		return ok
	}
	substring := strings.TrimSpace(s.Text())
	substring = strings.ReplaceAll(substring, " ", "-")
	ok = (strings.ToUpper(substring) == "AVAILABLE-NOW")
	return ok
}

func scrapeLaboVRKitBuyAt(d *goquery.Document) []*LaboKitBuyAt {
	laboKitRetailers := []*LaboKitBuyAt{}
	CSS := "li.retailer-modal__retailers-item a[href]"
	d.Find(CSS).Each(func(i int, s *goquery.Selection) {
		laboKitBuyAt := NewLaboKitBuyAt(s)
		laboKitRetailers = append(laboKitRetailers, laboKitBuyAt)
	})
	return laboKitRetailers
}

func scrapeLaboVRKitBoxImageHref(d *goquery.Document) string {
	CSS := "div.kit-description__packshot img[src]"
	attribute := d.Find(CSS).AttrOr("src", "NIL")
	if ok := attribute != "NIL"; ok != true {
		return attribute
	}
	substring := strings.TrimSpace(attribute)
	substring = strings.ReplaceAll(substring, "../", "")
	return fmt.Sprintf("https://labo.nintendo.com/%s", substring)
}

func scrapeLaboVRKitBoxImageURL(d *goquery.Document) {}

func scrapeLaboVRKitBoxImageSrcSet(d *goquery.Document) {}

func scrapeLaboVRKitCompatibility(d *goquery.Document) {}

func scrapeLaboVRKitHeroImageURL(d *goquery.Document) {}

func scrapeLaboVRKitHeroImageSrcSet(d *goquery.Document) {}

func scrapeLaboVRKitKitContents(d *goquery.Document) {}

func scrapeLaboVRKitKitFeatures(d *goquery.Document) {}

func scrapeLaboVRKitProjects(d *goquery.Document) {}

func scrapeLaboVRKitQuotes(d *goquery.Document) {}

func scrapeLaboVRKitVRPlaza(d *goquery.Document) {}
