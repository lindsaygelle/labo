package labo

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

const (
	defaultPartAmount int    = 1
	defaultPartColor  string = "NIL"
	defaultPartSize   string = "STANDARD"
)

const (
	errorPartEmptyText     string = "argument (*%p) does not contain text nodes"
	errorPartEmptyPartName string = "argument (*%p) does not contain a part name"
)

var (
	regexpPartFindColor    = regexp.MustCompile(`(\(blue|gray|orange|red\))`)
	regexpPartFindPartName = regexp.MustCompile(`^[^0-9]+`)
	regexpPartFindSize     = regexp.MustCompile(`(\(small|medium|large\))`)
	regexpPartFindSpares   = regexp.MustCompile(`\+\s{1}spares`)
)

// Part is a struct that represents a unique building component used to complete a Nintendo Labo ToyCon kit.
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
	if ok := (s != nil); !ok {
		return nil, fmt.Errorf(errorGoQuerySelectionNil)
	}
	if ok := (s.Length() > 0); !ok {
		return nil, fmt.Errorf(errorGoQuerySelectionEmptyHTMLNodes, s)
	}
	var (
		name string
		ok   bool
	)
	var (
		amount = defaultPartAmount
		color  = defaultPartColor
		size   = defaultPartSize
		spares = false
	)
	contents := strings.TrimSpace(s.Text())
	ok = (len(contents) > 0)
	if !ok {
		return nil, fmt.Errorf(errorPartEmptyText, s)
	}
	substring := regexpPartFindPartName.FindString(contents)
	ok = (len(substring) > 0)
	if !ok {
		return nil, fmt.Errorf(errorPartEmptyPartName, s)
	}
	substring = strings.TrimSpace(substring)
	substring = strings.TrimSuffix(substring, "x")
	substring = regexpMatchParenthesis.ReplaceAllString(substring, "")
	name = regexpMatchSequenceWhitespace.ReplaceAllString(substring, "")
	name = strings.TrimSpace(name)
	name = strings.ToUpper(name)
	substring = regexpMatchNumeric.FindString(contents)
	ok = (len(substring) > 0)
	if ok {
		amount, _ = strconv.Atoi(substring)
	}
	substring = regexpPartFindSize.FindString(contents)
	ok = (len(substring) > 0)
	if ok {
		substring = regexpMatchNonAlpha.ReplaceAllString(substring, "")
		size = strings.ToUpper(substring)
	}
	substring = regexpPartFindColor.FindString(contents)
	ok = (len(substring) > 0)
	if ok {
		substring = regexpMatchNonAlpha.ReplaceAllString(substring, "")
		color = strings.ToUpper(substring)
	}
	substring = regexpPartFindSpares.FindString(contents)
	ok = (len(substring) > 0)
	if ok {
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
