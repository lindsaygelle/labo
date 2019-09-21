package labo

import (
	"regexp"

	"golang.org/x/text/currency"
)

const (
	errorGoQuerySlectionEmptyHTMLNodes string = "argument (*%p) does not contain a collection of HTML elements"
	errorGoQuerySelectionNil           string = "argument *goquery.Selection cannot be nil"
)

const (
	laboRootURL string = "https://labo.nintendo.com"
)

var (
	regexpReplaceNonAlpha           = regexp.MustCompile(`\W`)
	regexpReplaceSequenceWhitespace = regexp.MustCompile(`\s{2,}`)
)

type Kit struct {
	Href      string
	ID        string
	Materials Materials
	Name      string
	Overview  Overview
	Projects  []Project
	Price     Price
	Retailers []Retailer
	Software  Software
	ToyCons   []ToyCon
}

type KitCustomization struct {
}

type KitVR struct {
	Plaza interface{}
}

type Feature struct {
	Description string
	Icon        Image
	Image       Image
	Title       string
	Video       Video
}

type Overview struct {
	Description string
	Video       Video
}

type Price struct {
	Amount   float64
	Currency currency.Unit
}

type Project struct {
	Icon        Image
	Image       Image
	Name        string
	Screenshots []Image
}

type Retailer struct {
	Href string
	Logo Image
	Name string
}

type ToyCon struct {
	Description string
	Features    []Feature
	Image       Image
	Name        string
}
