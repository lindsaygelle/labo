package labo

import (
	"regexp"
)

const (
	attrAlt        string = "alt"
	attrClass      string = "class"
	attrDataSizes  string = "data-sizes"
	attrDataSrc    string = "data-src"
	attrDataSrcSet string = "data-srcset"
	attrHref       string = "href"
	attrSizes      string = "sizes"
	attrSrc        string = "src"
	attrSrcSet     string = "srcset"
	attrTarget     string = "target"
)

const (
	errorGoQuerySelectionEmptyHTMLNodes string = "argument (*%p) does not contain a collection of HTML elements"
	errorGoQuerySelectionNil            string = "argument *goquery.Selection cannot be nil"
)

const (
	errorEmptyAttrClass string = "argument (*%p) does not contain an class attribute"
	errorEmptyHrefAlt   string = "argument (*%p) does not contain an href attribute"
)

const (
	laboRootURL string = "https://labo.nintendo.com"
)

var (
	regexpMatchNonAlpha           = regexp.MustCompile(`\W`)
	regexpMatchLineBreaks         = regexp.MustCompile(`\n`)
	regexpMatchNumeric            = regexp.MustCompile(`[0-9]+`)
	regexpMatchParenthesis        = regexp.MustCompile(`\(.+\)`)
	regexpMatchSequenceWhitespace = regexp.MustCompile(`\s{2,}`)
)
