package labo

import (
	"fmt"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type LaboBoxImage struct {
	Alt       string   `json:"alt"`
	Source    string   `json:"src"`
	SourceSet []string `json:"source_set"`
}

func NewLaboBoxImage(s *goquery.Selection) LaboBoxImage {
	return LaboBoxImage{
		Alt:       parseLaboBoxImageAlt(s),
		Source:    parseLaboBoxImageSource(s),
		SourceSet: parseLaboBoxImageSourceSet(s)}
}

func parseLaboBoxImageAlt(s *goquery.Selection) string {
	alt := s.AttrOr("alt", "NIL")
	if ok := alt != "NIL"; ok != true {
		return alt
	}
	return strings.TrimSpace(alt)
}

func parseLaboBoxImageSource(s *goquery.Selection) string {
	source := s.AttrOr("src", "NIL")
	if ok := source != "NIL"; ok != true {
		return source
	}
	substring := strings.ReplaceAll(source, "../", "")
	return fmt.Sprintf("https://labo.nintendo.com/%s", substring)
}

func parseLaboBoxImageSourceSet(s *goquery.Selection) []string {
	var sourceset []string
	srcset := strings.TrimSpace(s.AttrOr("srcset", "NIL"))
	if ok := srcset != "NIL"; ok != true {
		return sourceset
	}
	for _, src := range strings.Split(srcset, ",") {
		i := strings.Index(src, " ")
		if ok := (i != -1); ok != true {
			continue
		}
		source := src[:i]
		substring := strings.ReplaceAll(source, "../", "")
		sourceset = append(sourceset, fmt.Sprintf("https://labo.nintendo.com/%s", substring))
	}
	return sourceset
}
