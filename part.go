package labo

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// Part is a snapshot of a Nintendo Labo part that is used in the construction of a Nintendo Labo kit.
//
// Parts are built from reading Nintendo Labo product descriptions and contain varying levels
// of detail and verbosity. A part, depending on the content read, may contain mostly
// default part amounts, colors, genders, shapes and sizes.
type Part struct {
	Amount int    `json:"amount"`
	Color  string `json:"color"`
	Gender string `json:"gender"`
	Href   *Href  `json:"href"`
	Name   string `json:"name"`
	Shape  string `json:"shape"`
	Size   string `json:"size"`
	Spares bool   `json:"spares"`
}

// getPartAmount searches the argument string for substrings that describe the quantity
// of a Nintendo Labo kit part.
//
// getPartAmount relies on the argument string containing some
// form of numeric pattern or numeric namespace that can be used
// to determine the value of the parts provided. When the argument string contains
// no numeric pattern, the function performs a lookup against the current known
// max number of Nintendo Labo parts per kit. Should a value exceed the known range,
// the default value of one is assigned.
func getPartAmount(s string) int {
	var (
		amount    = defaultPartAmount
		ok        bool
		substring string
	)
	substring = regexpMatchNumbers.FindString(s)
	ok = (len(substring) > 0)
	if ok {
		amount, _ = strconv.Atoi(substring)
		return amount
	}
	substring = regexpMatchAmount.FindString(s)
	substring = strings.ToLower(substring)
	ok = (len(substring) > 0)
	if ok {
		amount = partAmountMap[substring]
	}
	return amount
}

// getPartColor searches the argument string for substrings that describe the color
// of a Nintendo Labo kit part.
//
// getPartColor relies on the argument string containing some
// form of color namespace that can be used to determine the color of the part.
// When the argument string does not contain a known Nintendo Labo part color,
// the default part color is assigned.
func getPartColor(s string) string {
	var (
		color     = defaultPartColor
		ok        bool
		substring string
	)
	substring = regexpMatchColor.FindString(s)
	substring = strings.ToLower(substring)
	ok = (len(substring) > 0)
	if ok {
		color = partColorMap[substring]
	}
	return color
}

// getPartGender searches the argument string for substrings that describe the gender
// of a Nintendo Labo kit part.
//
// getPartGender relies on the argument string containing some
// form of gender namespace that can be used to determine the gender of the part.
// When the argument string does not contain a known Nintendo Labo part gender,
// the default part gender is assigned.
func getPartGender(s string) string {
	var (
		gender    = defaultPartGender
		ok        bool
		substring string
	)
	substring = regexpMatchGender.FindString(s)
	substring = strings.ToLower(substring)
	ok = (len(substring) > 0)
	if ok {
		gender = partGenderMap[substring]
	}
	return gender
}

// getPartName returns the Nintendo Labo part name.
//
// getPartName works by substituting all potential part properties from
// within the Nintendo Labo kit name.
func getPartName(s string) string {
	for _, r := range partRegexps {
		s = r.ReplaceAllString(s, stringEmpty)
	}
	s = regexpMatchMultipleSpaces.ReplaceAllString(s, stringWhitespace)
	s = regexp.MustCompile(`(?i)(\sx\s$)`).ReplaceAllString(s, stringEmpty)
	s = strings.ToUpper(s)
	s = strings.TrimSpace(s)
	return s
}

func getPartShape(s string) string {
	var (
		ok        bool
		shape     = defaultPartShape
		substring string
	)
	substring = regexpMatchShape.FindString(s)
	substring = strings.ToLower(substring)
	ok = (len(substring) > 0)
	if ok {
		shape = partShapeMap[substring]
	}
	return shape
}

func getPartSize(s string) string {
	var (
		ok        bool
		size      = defaultPartSize
		substring string
	)
	substring = regexpMatchSize.FindString(s)
	substring = strings.ToLower(substring)
	ok = (len(substring) > 0)
	if ok {
		size = partShapeMap[substring]
	}
	return size
}

func getPartSpares(s string) bool {
	var (
		ok        bool
		substring string
	)
	substring = regexpMatchSpares.FindString(s)
	ok = (len(substring) > 0)
	return ok
}

func newPart(s *goquery.Selection) *Part {
	var (
		substring = strings.TrimSpace(s.Text())
	)
	return &Part{
		Amount: getPartAmount(substring),
		Color:  getPartColor(substring),
		Gender: getPartGender(substring),
		Href:   newHref(s),
		Name:   getPartName(substring),
		Shape:  getPartShape(substring),
		Size:   getPartSize(substring),
		Spares: getPartSpares(substring)}
}

func newParts(s *goquery.Selection) []*Part {
	var (
		part  *Part
		parts []*Part
		ok    bool
	)
	s.Each(func(i int, s *goquery.Selection) {
		part = newPart(s)
		ok = (part != nil)
		if !ok {
			return
		}
		parts = append(parts, part)
	})
	return parts
}
