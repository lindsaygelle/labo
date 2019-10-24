package labo

import (
	"fmt"
	"net/http"
	"os"
	"regexp"
	"strings"
	"text/tabwriter"
	"time"
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
	categoryMap = map[string]string{
		kitsID:  categoryKit,
		laboID:  categoryLabo,
		partsID: categoryParts}
)

var (
	categoryIDMap = map[string]string{
		categoryKit:   kitsID,
		categoryLabo:  laboID,
		categoryParts: partsID}
)

var (
	// partAmountMap is the lookup used to return the correct numeric value for a string number expression.
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
	// partColorMap is the lookup used to return the correct Nintendo Labo part color value for a string.
	partColorMap = map[string]string{
		partColorBlue:   (partColorBlue),
		partColorGray:   (partColorGray),
		partColorOrange: (partColorOrange),
		partColorRed:    (partColorRed),
		partColorYellow: (partColorYellow)}
)

var (
	// partGenderMap is the lookup used to return the correct Nintendo Labo part gender value for a string.
	partGenderMap = map[string]string{
		partGenderFemale: (partGenderFemale),
		partGenderMail:   (partGenderMale),
		partGenderMale:   (partGenderMale)}
)

var (
	// partShapeMap is the lookup used to return the correct Nintendo Labo part shape value for a string.
	partShapeMap = map[string]string{
		partShapeOctagonal: (partShapeOctagonal),
		partShapeSquare:    (partShapeSquare)}
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
	regexpMatchCurrency               = regexp.MustCompile(`([+-]?[0-9]{1,3}(?:,?[0-9]{3})*(?:\.[0-9]{2})?){1}`)
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

var (
	partRegexps = []*regexp.Regexp{
		regexpMatchAmount,
		regexpMatchColor,
		regexpMatchGender,
		regexpMatchNonAlphaNumeric,
		regexpMatchNumbers,
		regexpMatchShape,
		regexpMatchSize,
		regexpMatchSpares}
)

var (
	w = new(tabwriter.Writer).Init(os.Stdout, 4, 4, 0, '\t', 0)
)
