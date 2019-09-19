package labo

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// Part is a struct that represents the building component uses to complete
// a Nintendo Labo ToyCon kit.
type Part struct {
	Amount int    `json:"amount"`
	Color  string `json:"color"`
	Index  int    `json:"index"`
	Name   string `json:"name"`
	Size   string `json:"size"`
	Spares bool   `json:"spares"`
}

// NewPart is a Part constructor that takes a single goquery.Selection pointer. It scrubs
// the contents of the HTML element contents and attempts to parse the required
// and optional fields that describes a Nintendo Labo kit part.
func NewPart(s *goquery.Selection) (*Part, error) {
	const (
		defaultColor string = "NIL"
		defaultSize  string = "STANDARD"
	)
	var (
		amount = 1
		color  = defaultColor
		size   = defaultSize
		spares = false
	)
	if ok := (s.Length() > 0); !ok {
		return nil, fmt.Errorf("goquery.Selection is empty")
	}
	contents := strings.TrimSpace(s.Text())
	if ok := (len(contents) > 0); !ok {
		return nil, fmt.Errorf("goquery.Selection does not contain text")
	}
	substring := regexp.MustCompile(`^[^0-9]+`).FindString(contents)
	if ok := (len(substring) > 0); !ok {
		return nil, fmt.Errorf("goquery.Selection does not contain a valid part name")
	}
	substring = strings.TrimSpace(substring)
	substring = strings.TrimSuffix(substring, "x")
	substring = regexp.MustCompile(`\(.+\)`).ReplaceAllString(substring, "")
	name := regexp.MustCompile(`\s{2,}`).ReplaceAllString(substring, "")
	name = strings.TrimSpace(name)
	name = strings.ToUpper(name)
	substring = regexp.MustCompile("[0-9]+").FindString(contents)
	if ok := (len(substring) > 0); ok {
		amount, _ = strconv.Atoi(substring)
	}
	substring = regexp.MustCompile(`(\(small|medium|large\))`).FindString(contents)
	if ok := (len(substring) > 0); ok {
		substring = regexp.MustCompile(`\W`).ReplaceAllString(substring, "")
		size = strings.ToUpper(substring)
	}
	substring = regexp.MustCompile(`(\(blue|gray|orange|red\))`).FindString(contents)
	if ok := (len(substring) > 0); ok {
		substring = regexp.MustCompile(`\W`).ReplaceAllString(substring, "")
		color = strings.ToUpper(substring)
	}
	substring = regexp.MustCompile(`\+\s{1}spares`).FindString(contents)
	if ok := (len(substring) > 0); ok {
		spares = true
	}
	part := Part{
		Amount: amount,
		Color:  color,
		Name:   name,
		Size:   size,
		Spares: spares}
	return &part, nil
}
