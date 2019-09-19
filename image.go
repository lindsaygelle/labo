package labo

import (
	"fmt"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type Image struct {
	Alt    string
	Format string
	Size   int
	Src    string
}

func NewImage(s *goquery.Selection) (*Image, error) {
	const (
		defaultAlt string = "NIL"
		defaultSrc string = defaultAlt
	)
	var (
		alt string
		src string
	)
	if ok := (s.Length() > 0); !ok {
		return nil, fmt.Errorf("goquery.Selection is empty")
	}
	if _, ok := s.Attr("src"); !ok {
		s.SetAttr("src", s.AttrOr("data-src", defaultSrc))
	}
	src, ok := s.Attr("src")
	if ok != true {
		return nil, fmt.Errorf("goquery.Selection does not contain attribute")
	}
	src = strings.ReplaceAll(src, "../", "")
	src = fmt.Sprintf("https://labo.nintendo.com/%s", src)
	alt = s.AttrOr("alt", defaultAlt)
	alt = strings.TrimSpace(alt)
	alt = strings.ToUpper(alt)
	image := Image{
		Alt: alt,
		Src: src}
	return &image, nil
}
