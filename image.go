package labo

import (
	"fmt"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

const ()

type Image struct {
	Alt      string
	Format   string
	Size     int
	Src      string
	Variants []*ImageVariant
}

func NewImage(s *goquery.Selection) (*Image, error) {
	const (
		defaultAlt string = "NIL"
		defaultSrc string = defaultAlt
	)
	var (
		alt      string
		format   string
		src      string
		variants []*ImageVariant
	)
	if ok := (s.Length() > 0); !ok {
		return nil, fmt.Errorf("goquery.Selection is empty")
	}
	src, ok := s.Attr("data-src")
	if ok != true {
		return nil, fmt.Errorf("goquery.Selection does not contain data-src attribute")
	}
	src = strings.ReplaceAll(src, "../", "")
	src = fmt.Sprintf("https://labo.nintendo.com/%s", src)
	alt = s.AttrOr("alt", defaultAlt)
	alt = strings.TrimSpace(alt)
	alt = strings.ToUpper(alt)
	format = filepath.Ext(src)
	format = regexp.MustCompile(`\W`).ReplaceAllString(format, "")
	format = strings.ToUpper(format)
	for _, srcset := range strings.Split(s.AttrOr("data-srcset", ""), ",") {
		if imageVariant, err := NewImageVariant(srcset); err == nil {
			variants = append(variants, imageVariant)
		}
	}
	image := Image{
		Alt:      alt,
		Format:   format,
		Src:      src,
		Variants: variants}
	fmt.Println(image)
	return &image, nil
}
