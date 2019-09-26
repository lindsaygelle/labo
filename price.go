package labo

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"golang.org/x/text/currency"
)

const (
	errorPriceEmptyPrice string = "argument (*%p) does not contain a currency substring"
)

var (
	regexpPriceMatchCurrency = regexp.MustCompile(`([+-]?[0-9]{1,3}(?:,?[0-9]{3})*(?:\.[0-9]{2})?){1}`)
)

// Price is a struct that expresses the monetary cost for a Nintendo Labo Kit.
type Price struct {
	Amount   float64
	Currency currency.Unit
}

// NewPrice is a constructor function that instantiates and returns a new Price struct pointer.
func NewPrice(s *goquery.Selection) (*Price, error) {
	if ok := (s != nil); !ok {
		return nil, fmt.Errorf(errorGoQuerySelectionNil)
	}
	if ok := (s.Length() > 0); !ok {
		return nil, fmt.Errorf(errorGoQuerySelectionEmptyHTMLNodes, s)
	}
	substring := s.Text()
	if ok := (len(substring) > 0); !ok {
		return nil, fmt.Errorf(errorPriceEmptyPrice, s)
	}
	substring = strings.TrimSpace(substring)
	amount, err := strconv.ParseFloat(regexpPriceMatchCurrency.FindString(substring), 64)
	if err != nil {
		return nil, err
	}
	price := Price{
		Amount:   amount,
		Currency: currency.USD}
	return &price, nil
}
