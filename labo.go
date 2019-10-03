package labo

import (
	"fmt"
	"regexp"
	"strings"
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
	partGenderFemale string = "female"
	partGenderMail   string = "mail"
	partGenderMale   string = "male"
)

const (
	partShapeOctagonal string = "octagonal"
	partShapeSquare    string = "square"
)

const (
	partSizeLarge  string = "large"
	partSizeMedium string = "medium"
	partSizeSmall  string = "small"
)

var (
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
	partColors = []string{
		partColorBlue,
		partColorGray,
		partColorOrange,
		partColorRed,
		partColorYellow}
)

var (
	partGenders = []string{
		partGenderFemale,
		partGenderMail,
		partGenderMale}
)

var (
	partShapes = []string{
		partShapeOctagonal,
		partShapeSquare}
)

var (
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
)

var (
	partColorsExpression = fmt.Sprintf(patternIgnorecase, strings.Join(partColors, "|"))
)

var (
	partGenderExpression = fmt.Sprintf(patternIgnorecase, strings.Join(partGenders, "|"))
)

var (
	partShapeExpression = fmt.Sprintf(patternIgnorecase, strings.Join(partShapes, "|"))
)

var (
	partSizeExpression = fmt.Sprintf(patternIgnorecase, strings.Join(partSizes, "|"))
)

var (
	partSparesExpression = fmt.Sprintf(patternIgnorecase, "spares")
)

var (
	regexpMatchAmount = regexp.MustCompile(partAmountExpression)
)

var (
	regexpMatchColor = regexp.MustCompile(partColorsExpression)
)

var (
	regexpMatchGender = regexp.MustCompile(partGenderExpression)
)

var (
	regexpMatchNumbers = regexp.MustCompile(`(\d+)`)
)

var (
	regexpMatchNonAlphaNumeric = regexp.MustCompile(`[^a-zA-Z0-9\s]+`)
)

var (
	regexpMatchShape = regexp.MustCompile(partShapeExpression)
)

var (
	regexpMatchSize = regexp.MustCompile(partSizeExpression)
)

var (
	regexpMatchMultipleSpaces = regexp.MustCompile(`\s{2,}`)
)

var (
	regexpMatchSpares = regexp.MustCompile(partSparesExpression)
)
