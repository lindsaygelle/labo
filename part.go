package labo

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

const (
	defaultPartColor  string = "NIL"
	defaultPartGender string = "NEUTRAL"
	defaultPartShape  string = "NIL"
	defaultPartSize   string = "REGULAR"
)

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

func getPartAmount(s string) int {
	var (
		amount    int
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

func getPartName(s string) string {

	for _, r := range []*regexp.Regexp{
		regexpMatchAmount,
		regexpMatchColor,
		regexpMatchGender,
		regexpMatchNonAlphaNumeric,
		regexpMatchNumbers,
		regexpMatchShape,
		regexpMatchSize,
		regexpMatchSpares} {

		s = r.ReplaceAllString(s, "")
	}
	s = regexpMatchMultipleSpaces.ReplaceAllString(s, " ")
	s = regexp.MustCompile(`(?i)(\sx\s$)`).ReplaceAllString(s, "")
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

func newPart(s *goquery.Selection) Part {
	var (
		substring = strings.TrimSpace(s.Text())
	)
	return Part{
		Amount: getPartAmount(substring),
		Color:  getPartColor(substring),
		Gender: getPartGender(substring),
		Href:   newHref(s),
		Name:   getPartName(substring),
		Shape:  getPartShape(substring),
		Size:   getPartSize(substring),
		Spares: getPartSpares(substring)}
}
