package labo

import (
	"net/http"
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
	errorGoQueryDocumentEmptyHTMLNodes  string = "document (*%p) does not contain a collection of HTML elements"
	errorGoQuerySelectionEmptyHTMLNodes string = "argument (*%p) does not contain a collection of HTML elements"
	errorGoQuerySelectionNil            string = "argument *goquery.Selection cannot be nil"
)

const (
	errorEmptyAttrClass string = "argument (*%p) does not contain an class attribute"
	errorEmptyHrefAlt   string = "argument (*%p) does not contain an href attribute"
)

const (
	kitRobot   string = "robot-kit"
	kitVariety string = "variety-kit"
	kitVehicle string = "vehicle-kit"
)

const (
	// URL is homepage for the Nintendo Labo product.
	URL string = "https://labo.nintendo.com"
)

const (
	// URLKits is the the path for the current list of Nintendo Labo Kits.
	URLKits string = URL + "/kits"
)

const (
	// URLRobotKit is the direct URL to the Nintendo Labo Robot Kit.
	URLRobotKit string = (URLKits + "/" + kitRobot)
)
const (
	// URLVarietyKit is the direct URL to the Nintendo Labo Variety Kit.
	URLVarietyKit string = (URLKits + "/" + kitVariety)
)

const (
	// URLVehicleKit is the direct URL to the Nintendo Labo Vehicle Kit.
	URLVehicleKit string = (URLKits + "/" + kitVehicle)
)

var (
	client = (&http.Client{})
)

var (
	regexpMatchNonAlpha           = regexp.MustCompile(`\W`)
	regexpMatchLineBreaks         = regexp.MustCompile(`\n`)
	regexpMatchNumeric            = regexp.MustCompile(`[0-9]+`)
	regexpMatchParenthesis        = regexp.MustCompile(`\(.+\)`)
	regexpMatchSequenceWhitespace = regexp.MustCompile(`\s{2,}`)
)

var (
	mapURL = map[string]string{
		kitRobot:   URLRobotKit,
		kitVariety: URLVarietyKit,
		kitVehicle: URLVehicleKit}
)

func GetKit(URL string) (*Kit, error) {
	return nil, nil
}
