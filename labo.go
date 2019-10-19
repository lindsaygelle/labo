package labo

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"

	"golang.org/x/text/currency"
	"golang.org/x/text/language"
)

// Labo is a Nintendo Labo product. Labo structs may either be a full Nintendo Labo kit or collections of parts.
//
// Nintendo Labo products are provided by the Nintendo store website and contain only surface
// level data about the Nintendo Labo product.
type Labo struct {
	Category         string        `json:"category"`
	CategoryID       string        `json:"category_ID"`
	Currency         currency.Unit `json:"currency"`
	ID               int           `json:"ID"`
	Language         language.Tag  `json:"language"`
	Name             string        `json:"name"`
	Parts            []*Part       `json:"parts"`
	Price            float32       `json:"price"`
	ProductID        string        `json:"product_ID"`
	Ref              string        `json:"ref"`
	Status           string        `json:"http_status"`
	StatusCode       int           `json:"http_status_code"`
	StoreDescription string        `json:"store_description"`
	StoreImages      []*Image      `json:"store_images"`
	StoreTitle       string        `json:"store_title"`
	Time             time.Time     `json:"time"`
	URL              *url.URL      `json:"URL"`
}

var (
	// laboFn are the required functions used to collect all required data points for a Nintendo Labo product
	// from the Nintendo store.
	laboFn = [](func(s *goquery.Selection, l *Labo)){
		getLaboStorePageDescription,
		getLaboStorePageImages,
		getLaboStorePageName,
		getLaboStorePageParts,
		getLaboStorePagePrice,
		getLaboStorePageRef,
		getLaboStorePageTitle}
)

// Get gets a specific Nintendo Labo product by the Nintendo stores product ID.
func Get(ID string) *Labo {
	var (
		doc *goquery.Document
		err error
		ok  bool
		q   url.Values
		req *http.Request
		res *http.Response
		s   *goquery.Selection

		l = &Labo{
			Currency:   currency.USD,
			Language:   language.AmericanEnglish,
			Name:       defaultLaboName,
			Ref:        defaultLaboRef,
			Status:     http.StatusText(http.StatusBadRequest),
			StatusCode: http.StatusBadRequest,
			Time:       time.Now()}
	)
	req, err = http.NewRequest(http.MethodGet, fmt.Sprintf("%s%s", storeProductURI, ID), nil)
	ok = (err == nil)
	if !ok {
		return l
	}
	res, err = client.Do(req)
	ok = (err == nil)
	if !ok {
		return l
	}
	l.Status = res.Status
	l.StatusCode = res.StatusCode
	l.URL = req.URL
	ok = (res.StatusCode == http.StatusOK)
	if !ok {
		return l
	}
	q = req.URL.Query()
	l.CategoryID = q.Get(uriQueryParamCategoryID)
	ok = (len(l.CategoryID) > 0)
	if !ok {

	}
	l.ProductID = q.Get(uriQueryParamProductID)
	ok = (len(l.ProductID) > 0)
	if !ok {

	}
	doc, err = goquery.NewDocumentFromResponse(res)
	ok = (err == nil)
	if !ok {
		return l
	}
	s = doc.Find(htmlBody)
	ok = (s.Length() > 0)
	if !ok {
		return l
	}
	return newLabo(s, l)
}

// GetAll gets all available Nintendo Labo from the Nintendo store by the argument labo ID.
//
// Requires the argument ID to be one of the three exported labo.(*ID)'s.
// Each of the exported labo ID's queries and returns a specific set of Nintendo Labo products.
// Using the labo.LaboID will collect all Nintendo Labo products.
// The labo.KitsID will only collect and return Nintendo Labo products that are kits and not parts.
// The labo.PartsID will only collect and return Nintendo Labo products that are parts or accessories.
func GetAll(ID string) []*Labo {
	const (
		CSS string = ".product-listing .product-container > p a"
	)
	var (
		doc *goquery.Document
		err error
		l   []*Labo
		ok  bool
		req *http.Request
		res *http.Response
		s   *goquery.Selection
	)
	req, err = http.NewRequest(http.MethodGet, fmt.Sprintf("%s%s", storeCategoryURI, ID), nil)
	res, err = client.Do(req)
	ok = (err == nil)
	if !ok {
		return nil
	}
	ok = (res.StatusCode == http.StatusOK)
	if !ok {
		return nil
	}
	doc, err = goquery.NewDocumentFromResponse(res)
	ok = (err == nil)
	if !ok {
		return nil
	}
	s = doc.Find(CSS)
	ok = (s.Length() > 0)
	if !ok {
		return nil
	}
	s.Each(func(i int, s *goquery.Selection) {
		href, exists := s.Attr(attrHref)
		if !exists {
			return
		}
		URL, err := url.Parse(href)
		if err != nil {
			return
		}
		ID := URL.Query().Get(uriQueryParamProductID)
		labo := Get(ID)
		if labo == nil {
			return
		}
		l = append(l, labo)
	})
	return l
}

// Marshal marshals a Labo struct.
func Marshal(l *Labo) (b []byte) {
	b, _ = json.Marshal(l)
	return b
}

// newLabo is a constructor function that take an argument goquery.Selection pointer
// and runs a collection of helper functions to extract and assign all the
// required properties for the pending labo.Labo pointer. Should any of the
// helper functions fail to find the required property, the default property is not
// overriden.
func newLabo(s *goquery.Selection, l *Labo) *Labo {
	for _, fn := range laboFn {
		fn(s, l)
	}
	return l
}

// getLaboStorePageDescription searches the argument goquery.Selection pointer
// for the Nintendo Labo product description and assigns it to
// the argument labo.Labo pointer if found..
func getLaboStorePageDescription(s *goquery.Selection, l *Labo) {
	const (
		CSS string = "#main-content #prodDescBtm p:nth-child(2)"
	)
	var (
		ok        bool
		substring string
	)
	s = s.Find(CSS)
	ok = (s.Length() > 0)
	if !ok {
		return
	}
	substring = strings.TrimSpace(s.Text())
	ok = (len(substring) > 0)
	if !ok {
		return
	}
	l.StoreDescription = substring
}

// getLaboStorePageID searches the argument goquery.Selection pointer
// for the Nintendo Labo product numerical ID and assigns it to
// the argument labo.Labo pointer if found.
func getLaboStorePageID(s *goquery.Selection, l *Labo) {
	const (
		CSS string = "#main-content .results-header"
	)
	var (
		err       error
		ID        int
		substring string
		ok        bool
	)
	s = s.Find(CSS)
	ok = (s.Length() > 0)
	if !ok {
		return
	}
	substring = strings.TrimSpace(s.Text())
	ok = (len(substring) > 0)
	if !ok {
		return
	}
	substring = regexpMatchNonNumeric.ReplaceAllString(substring, stringEmpty)
	ok = (len(substring) > 0)
	if !ok {
		return
	}
	ID, err = strconv.Atoi(substring)
	ok = (err == nil)
	if !ok {
		return
	}
	l.ID = ID
}

// getLaboStorePageImages searches the argument goquery.Selection pointer
// for the Nintendo Labo product images and assigns them to
// the argument labo.Labo pointer if found.
func getLaboStorePageImages(s *goquery.Selection, l *Labo) {
	const (
		CSS string = "#main-content #product-thumbs img"
	)
	var (
		ok bool
	)
	s = s.Find(CSS)
	ok = (s.Length() > 0)
	if !ok {
		return
	}
	l.StoreImages = newImages(s)
}

// getLaboStorePageName searches the argument goquery.Selection pointer
// for the Nintendo Labo product name and assigns the unformatted string to
// the argument labo.Labo pointer if found.
func getLaboStorePageName(s *goquery.Selection, l *Labo) {
	const (
		CSS string = "#main-content .results-header"
	)
	var (
		substring string
		ok        bool
	)
	s = s.Find(CSS)
	ok = (s.Length() > 0)
	if !ok {
		return
	}
	substring = strings.TrimSpace(s.Text())
	ok = (len(substring) > 0)
	if !ok {
		return
	}
	l.Name = substring
}

// getLaboStorePageParts searches the argument goquery.Selection pointer
// for the Nintendo Labo product parts and components used to build
// the Nintendo Labo product and assigns them to
// the argument labo.Labo pointer if they are found.
func getLaboStorePageParts(s *goquery.Selection, l *Labo) {
	const (
		CSS string = "#main-content #prodDescBtm ul li"
	)
	var (
		ok bool
	)
	s = s.Find(CSS)
	ok = (s.Length() > 0)
	if !ok {
		return
	}
	l.Parts = newParts(s)
}

// getLaboStorePagePrice searches the argument goquery.Selection pointer
// for the Nintendo Labo product price and assigns the floating point value to
// the argument labo.Labo pointer if found.
func getLaboStorePagePrice(s *goquery.Selection, l *Labo) {
	const (
		CSS string = "#main-content #addToCart > p .txt-bold"
	)
	var (
		err       error
		price     float64
		substring string
		ok        bool
	)
	s = s.Find(CSS)
	ok = (s.Length() > 0)
	if !ok {
		return
	}
	substring = strings.TrimSpace(s.Text())
	ok = (len(substring) > 0)
	if !ok {
		return
	}
	substring = regexpMatchNumbers.FindString(substring)
	ok = (len(substring) > 0)
	if !ok {
		return
	}
	price, err = strconv.ParseFloat(substring, 32)
	ok = (err == nil)
	if !ok {
		return
	}
	l.Price = float32(price)
}

// getLaboStorePageRef searches the argument goquery.Selection pointer
// for the Nintendo Labo product name alias (used to search the offical Labo site)
// and assigns the formatted string to the argument labo.Labo pointer if found.
func getLaboStorePageRef(s *goquery.Selection, l *Labo) {
	const (
		CSS string = "#main-content .results-header"
	)
	var (
		substring  string
		substrings []string
		ok         bool
	)
	s = s.Find(CSS)
	ok = (s.Length() > 0)
	if !ok {
		return
	}
	substring = s.Text()
	ok = (len(substring) > 0)
	if !ok {
		return
	}
	ok = strings.Contains(substring, stringColon)
	if !ok {
		return
	}
	substrings = strings.Split(substring, stringColon)
	ok = (len(substrings) >= 2)
	if !ok {
		return
	}
	substring = strings.TrimSpace(substrings[1])
	ok = (len(substring) > 0)
	if !ok {
		return
	}
	substring = strings.ToLower(substring)
	ok = (strings.Contains(substring, stringPlus))
	if ok {
		substrings = strings.Split(substring, stringPlus)
		substring = strings.ReplaceAll(substring, substrings[1], stringEmpty)
	}
	substring = regexpMatchNonAlphaNumericNoSpace.ReplaceAllString(substring, stringMinus)
	substring = strings.TrimSpace(substring)
	ok = (len(substring) > 0)
	if !ok {
		return
	}
	substring = regexpMatchMultipleSpaces.ReplaceAllString(substring, stringMinus)
	ok = (len(substring) > 0)
	if !ok {
		return
	}
	substring = strings.TrimSuffix(substring, stringMinus)
	l.Ref = substring
}

// getLaboStorePageTitle searches the argument goquery.Selection pointer
// for the Nintendo Labo product headline and assigns the unformatted string to
// the argument labo.Labo pointer if found.
func getLaboStorePageTitle(s *goquery.Selection, l *Labo) {
	const (
		CSS string = "#main-content #prodDescBtm p:nth-child(1)"
	)
	var (
		ok        bool
		substring string
	)
	s = s.Find(CSS)
	ok = (s.Length() > 0)
	if !ok {
		return
	}
	substring = strings.TrimSpace(s.Text())
	ok = (len(substring) > 0)
	if !ok {
		return
	}
	l.StoreTitle = substring
}
