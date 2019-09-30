package labo

import (
	"fmt"
	"net/http"
	"net/url"
	"regexp"
	"time"

	"github.com/PuerkitoBio/goquery"
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
	errorGoQueryDocumentNil             string = "argument *goquery.Document cannot be nil"
	errorGoQueryDocumentEmptyHTMLNodes  string = "document (*%p) does not contain a collection of HTML elements"
	errorGoQuerySelectionEmptyHTMLNodes string = "argument (*%p) does not contain a collection of HTML elements"
	errorGoQuerySelectionNil            string = "argument *goquery.Selection cannot be nil"
)
const (
	errorURLHost  string = "argument (%s) host is not %s"
	errorURLMatch string = "argument (%s) is not a supported URL"
)

const (
	errorEmptyAttrClass string = "argument (*%p) does not contain a class attribute"
	errorEmptyHrefAlt   string = "argument (*%p) does not contain a href attribute"
)

const (
	kitRobot     string = "robot-kit"
	kitVariety   string = "variety-kit"
	kitVehicle   string = "vehicle-kit"
	kitVR        string = "vr-kit"
	kitVRStarter string = "vr-starter-kit"
)

const (
	kitURL string = (laboURL + "/" + "kits")
)

const (
	kitRobotURL     string = (kitURL + "/" + kitRobot)
	kitVarietyURL   string = (kitURL + "/" + kitVariety)
	kitVehicleURL   string = (kitURL + "/" + kitVehicle)
	kitVRURL        string = (kitURL + "/" + kitVR)
	kitVRStarterURL string = (kitURL + "/" + kitVRStarter)
)

const (
	laboHost string = ("labo" + "." + nintendoHost)
	laboURL  string = ("https://" + laboHost)
)

const (
	nintendoHost string = "nintendo.com"
	nintendoURL  string = ("https://" + nintendoHost)
)

const (
	clientTimeout time.Duration = (time.Second * 10)
)

var (
	client = (&http.Client{
		Timeout: clientTimeout})
)

var (
	regexpMatchNonAlpha           = regexp.MustCompile(`\W`)
	regexpMatchLineBreaks         = regexp.MustCompile(`\n`)
	regexpMatchNumeric            = regexp.MustCompile(`[0-9]+`)
	regexpMatchParenthesis        = regexp.MustCompile(`\(.+\)`)
	regexpMatchSequenceWhitespace = regexp.MustCompile(`\s{2,}`)
)

var (
	// RobotKitURL is the Nintendo Labo URL for the Nintendo Labo Robot Kit.
	RobotKitURL = URL(kitRobotURL)
)
var (
	// VarietyKitURL is the Nintendo Labo URL for the Nintendo Variety Kit.
	VarietyKitURL = URL(kitVarietyURL)
)
var (
	// VehicleKitURL is the Nintendo Labo URL for the Nintendo Labo Vehicle Kit.
	VehicleKitURL = URL(kitVehicleURL)
)

var (
	// VRKitURL is the Nintendo Labo URL for the Nintendo Labo VR Kit.
	VRKitURL = VRURL(kitVRURL)
)
var (
	// VRStarterKitURL is the Nintendo Labo URL for the Nintendo Labo VR Starter Kit.
	VRStarterKitURL = VRURL(kitVRStarterURL)
)

func net(URL *url.URL) (*goquery.Document, error) {
	ok := (URL.Host == laboHost)
	if !ok {
		return nil, fmt.Errorf(errorURLHost, URL.Host, laboHost)
	}
	req, err := http.NewRequest(http.MethodGet, URL.String(), nil)
	if err != nil {
		return nil, err
	}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	return goquery.NewDocumentFromResponse(res)
}

// GetCustomizationKIt gets a Nintendo Labo Customization Kit from the Nintendo Labo website. Assumes that the provided
// labo.URL argument contains a defined URL to a valid Nintendo Labo Customization Kit.
func GetCustomizationKIt(URL CustomizationURL) (*KitCustomization, error) {
	u, err := URL.URL()
	if err != nil {
		return nil, err
	}
	doc, err := net(u)
	return NewKitCustomization(doc)
}

// GetKit gets a Nintendo Labo Kit from the Nintendo Labo website. Assumes that the provided
// labo.URL argument contains a defined URL to a valid Nintendo Labo Kit.
func GetKit(URL URL) (*Kit, error) {
	u, err := URL.URL()
	if err != nil {
		return nil, err
	}
	doc, err := net(u)
	return NewKit(doc)
}

// GetVRKit gets a Nintendo Labo VR Kit from the Nintendo Labo website. Assumes that the provided
// labo.VRURL argument contains a defined URL to a valid Nintendo Labo VR Kit.
func GetVRKit(URL VRURL) (*KitVR, error) {
	u, err := URL.URL()
	if err != nil {
		return nil, err
	}
	doc, err := net(u)
	return NewKitVR(doc)
}
