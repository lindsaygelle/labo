package labo

import (
	"fmt"
	"net/http"
	"regexp"
	"strings"
	"time"
)

const (
	// attrHref is the HTML href attribute.
	attrHref string = "href"
	// attrDataSrc is the HTML data-attribute src.
	attrDataSrc string = "data-src"
	// attrSrc is the HTML src attribute.
	attrSrc string = "src"
	// attrTarget is the HTML target attribute.
	attrTarget string = "target"
)

const (
	// laboDNS is the domain name reference for the Nintendo Labo website.
	laboDNS string = ("labo" + "." + nintendoDNS)
)
const (
	// laboURI is the URI directive to perform a Nintendo Labo site search for Nintendo Labo kits.
	laboURI string = (laboURL + "/" + "kits")
)
const (
	// laboURL is the RFC2616 compliant address for the Nintendo Labo website.
	laboURL string = ("https://" + laboDNS)
)
const (
	// nintendoDNS is the domain name reference for the Nintendo official website.
	nintendoDNS string = "nintendo.com"
)
const (
	// nintendoURL is the RFC2616 compliant address for the Nintendo official website.
	nintendoURL string = ("https://" + nintendoDNS)
)
const (
	// storeDNS is the domain name reference for the Nintendo store website.
	storeDNS string = ("store" + "." + nintendoDNS)
)
const (
	// storeURI is the URI directive to perform a Nintendo store search for Nintendo Labo kits.
	storeURI string = (storeURL + "/" + "ng3/us/po/browse/subcategory.jsp?viewAll=true&categoryId=")
)
const (
	// storeURIKits is the URI directive to request all Nintendo Labo full kits from the Nintendo store.
	storeURIKits string = (storeURI + "cat" + kitsID)
)
const (
	// storeURILabo is the URI directive to request all Nintendo Labo full kits and parts from the Nintendo store.
	storeURILabo string = (storeURI + "cat" + laboID)
)
const (
	// storeURIParts is the URI directive to request all Nintendo Labo part kits from the Nintendo store.
	storeURIParts string = (storeURI + "cat" + partsID)
)
const (
	// storeURL is the RFC2616 compliant address for the Nintendo store website.
	storeURL string = ("https://" + storeDNS)
)
const (
	// kitsID is the product ID for all Nintendo Labo full kits.
	kitsID string = "970105"
	// laboID is the product ID for both Nintendo Labo full kits and parts.
	laboID string = "960195"
	// partsID is the product ID for all Nintendo Labo parts kits.
	partsID string = "970106"
)

const (
	patternIgnorecase string = "(?i)(%s)"
)

const (
	partAmountOne      string = "one"
	partAmountTwo      string = "two"
	partAmountThree    string = "three"
	partAmountFour     string = "four"
	partAmountFive     string = "five"
	partAmountSix      string = "six"
	partAmountSeven    string = "seven"
	partAmountEight    string = "eight"
	partAmountNine     string = "nine"
	partAmountTen      string = "ten"
	partAmountEleven   string = "eleven"
	partAmountTwelve   string = "twelve"
	partAmountThirteen string = "thirteen"
)

const (
	partColorBlue   string = "blue"
	partColorGray   string = "gray"
	partColorOrange string = "orange"
	partColorRed    string = "red"
	partColorYellow string = "yellow"
)

const (
	// partGenderFemale is the namespace for parts that are of the female configuration.
	partGenderFemale string = "female"
	// partGenderMail is the alias namespace for a known typo for male parts.
	partGenderMail string = "mail"
	// partsGenderMale is the namespace for parts that of the male configuration.
	partGenderMale string = "male"
)

const (
	// defaultAttrAlt is the default namespace for HTML alt attributes.
	defaultAttrAlt string = "NIL"
	// defaultAttrTarget is the default namespace for HTML target attributes.
	defaultAttrTarget string = "NIL"
)

const (
	// defaultPartColor is the default color namespace for parts.
	defaultPartColor string = "NIL"
	// defaultPartGender is the default gender namespace for parts.
	defaultPartGender string = "NEUTRAL"
	// defaultPartShape is the default shape namespace for parts.
	defaultPartShape string = "NIL"
	// defaultPartSize is the default size namespace for parts.
	defaultPartSize string = "REGULAR"
)

const (
	// partShapeOctagonal is the namespace for parts that are of octagonal shape.
	partShapeOctagonal string = "octagonal"
	// partShapeSquare is the namespace for parts that are a square shape.
	partShapeSquare string = "square"
)

const (
	// partSizeLarge is the namespace for parts that are of large size.
	partSizeLarge string = "large"
	// partSizeMedium is the namespace for parts that are of medium size.
	partSizeMedium string = "medium"
	// partSizeSmall is the namespace for parts that are of smaller size.
	partSizeSmall string = "small"
)

var (
	// client is the HTTP client used throughout the Nintendo Labo package.
	client = &http.Client{Timeout: (time.Second * 10)}
)

var (
	// kitsRequest is a HTTP GET request for all Nintendo Labo full kits.
	kitsRequest, _ = http.NewRequest(http.MethodGet, storeURIKits, nil)
	// laboRequest is a HTTP GET request for both Nintendo Labo full kits and parts.
	laboRequest, _ = http.NewRequest(http.MethodGet, storeURILabo, nil)
	// partsRequest is a HTTP GET request for all Nintendo Labo parts kits.
	partsRequest, _ = http.NewRequest(http.MethodGet, storeURIParts, nil)
)

var (
	// partsAmounts is the collection of all defined part amount namespaces.
	partAmounts = []string{
		partAmountOne,
		partAmountTwo,
		partAmountThree,
		partAmountFour,
		partAmountFive,
		partAmountSix,
		partAmountSeven,
		partAmountEight,
		partAmountNine,
		partAmountTen,
		partAmountEleven,
		partAmountTwelve,
		partAmountThirteen}
)

var (
	// partColors is the collection of all defined part color namespaces.
	partColors = []string{
		partColorBlue,
		partColorGray,
		partColorOrange,
		partColorRed,
		partColorYellow}
)

var (
	// partsGenders is the collection of all defined part gender namespaces.
	partGenders = []string{
		partGenderFemale,
		partGenderMail,
		partGenderMale}
)

var (
	// partShapes is the collection all defined part shape namespaces.
	partShapes = []string{
		partShapeOctagonal,
		partShapeSquare}
)

var (
	// partSizes is the collection of all defined part size namespaces.
	partSizes = []string{
		partSizeLarge,
		partSizeMedium,
		partSizeSmall}
)

var (
	partAmountMap = map[string]int{
		partAmountOne:      1,
		partAmountTwo:      2,
		partAmountThree:    3,
		partAmountFour:     4,
		partAmountFive:     5,
		partAmountSix:      6,
		partAmountSeven:    7,
		partAmountEight:    8,
		partAmountNine:     9,
		partAmountTen:      10,
		partAmountEleven:   11,
		partAmountTwelve:   12,
		partAmountThirteen: 13}
)

var (
	partColorMap = map[string]string{
		partColorBlue:   strings.ToUpper(partColorBlue),
		partColorGray:   strings.ToUpper(partColorGray),
		partColorOrange: strings.ToUpper(partColorOrange),
		partColorRed:    strings.ToUpper(partColorRed),
		partColorYellow: strings.ToUpper(partColorYellow)}
)

var (
	partGenderMap = map[string]string{
		partGenderFemale: strings.ToUpper(partGenderFemale),
		partGenderMail:   strings.ToUpper(partGenderMale),
		partGenderMale:   strings.ToUpper(partGenderMale)}
)

var (
	partShapeMap = map[string]string{
		partShapeOctagonal: strings.ToUpper(partShapeOctagonal),
		partShapeSquare:    strings.ToUpper(partShapeSquare)}
)

var (
	partAmountExpression = fmt.Sprintf(patternIgnorecase, strings.Join(partAmounts, "|"))
	partColorsExpression = fmt.Sprintf(patternIgnorecase, strings.Join(partColors, "|"))
	partGenderExpression = fmt.Sprintf(patternIgnorecase, strings.Join(partGenders, "|"))
	partShapeExpression  = fmt.Sprintf(patternIgnorecase, strings.Join(partShapes, "|"))
	partSizeExpression   = fmt.Sprintf(patternIgnorecase, strings.Join(partSizes, "|"))
	partSparesExpression = fmt.Sprintf(patternIgnorecase, "spares")
)
var (
	regexpMatchAmount                 = regexp.MustCompile(partAmountExpression)
	regexpMatchColor                  = regexp.MustCompile(partColorsExpression)
	regexpMatchGender                 = regexp.MustCompile(partGenderExpression)
	regexpMatchNonAlphaNumeric        = regexp.MustCompile(`[^a-zA-Z0-9\s]+`)
	regexpMatchNonAlphaNumericNoSpace = regexp.MustCompile(`[^a-zA-Z0-9]+`)
	regexpMatchNonNumeric             = regexp.MustCompile(`[^0-9]+`)
	regexpMatchNumbers                = regexp.MustCompile(`(\d+)`)
	regexpMatchShape                  = regexp.MustCompile(partShapeExpression)
	regexpMatchSize                   = regexp.MustCompile(partSizeExpression)
	regexpMatchMultipleSpaces         = regexp.MustCompile(`\s{2,}`)
	regexpMatchSpares                 = regexp.MustCompile(partSparesExpression)
)
