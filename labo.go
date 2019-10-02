package labo

import (
	"fmt"
	"regexp"
	"strings"
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
		partGenderFemale: partGenderFemale,
		partGenderMail:   partGenderMale,
		partGenderMale:   partGenderMale}
)

var (
	partAmountExpression = fmt.Sprintf("(?i)(%s)", strings.Join(partAmounts, "|"))
)

var (
	partColorsExpression = fmt.Sprintf("(?i)(%s)", strings.Join(partColors, "|"))
)

var (
	partSizeExpression = fmt.Sprintf("(?i)(%s)", strings.Join(partSizes, "|"))
)

var (
	regexpMatchAmount = regexp.MustCompile(partAmountExpression)
)

var (
	regexpMatchColor = regexp.MustCompile(partColorsExpression)
)

var (
	regexpMatchNumbers = regexp.MustCompile(`(\d+)`)
)

var (
	regexpMatchSize = regexp.MustCompile(partSizeExpression)
)

var (
	regexpMatchSpares = regexp.MustCompile(`(?i)spares`)
)
