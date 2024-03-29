package labo

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
	"strings"
	"sync"
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
	Category    string        `json:"category"`
	CategoryID  string        `json:"category_ID"`
	Currency    currency.Unit `json:"currency"`
	Description string        `json:"description"`
	ID          int           `json:"ID"`
	Images      []*Image      `json:"images"`
	Language    language.Tag  `json:"language"`
	Name        string        `json:"name"`
	Parts       []*Part       `json:"parts"`
	Price       float32       `json:"price"`
	ProductID   string        `json:"product_ID"`
	Ref         string        `json:"ref"`
	Status      string        `json:"http_status"`
	StatusCode  int           `json:"http_status_code"`
	Time        time.Time     `json:"time"`
	Title       string        `json:"title"`
	URL         *URL          `json:"URL"`
}

var (
	// laboFn are the required functions used to collect all required data points for a Nintendo Labo product
	// from the Nintendo store.
	laboFn = [](func(s *goquery.Selection, l *Labo)){
		getLaboCategory,
		getLaboDescription,
		getlaboID,
		getLaboImages,
		getLaboName,
		getLaboParts,
		getLaboPrice,
		getLaboRef,
		getLaboTitle}
)

// Get gets a specific Nintendo Labo product from the Nintendo Labo store and returns
// the scraped information as a labo.Labo.
//
// To get a Labo, the Nintendo Labo ID must be provided to the function
// as a normalized ID. The ID should resemble a sequence of numbers and
// not contain the product prefix.
//
// Get will always return a labo.Labo, even if the product ID is invalid or not found.
// To identify whether the corresponding Labo was found, the HTTP status code
// or HTTP status can be checked. Successfully scraped Labo's will
// always contain a http.StatusOK.
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
	l.URL = newURL(req.URL)
	ok = (res.StatusCode == http.StatusOK)
	if !ok {
		return l
	}
	q = req.URL.Query()
	l.CategoryID = q.Get(uriQueryParamCategoryID)
	ok = (len(l.CategoryID) > 0)
	if !ok {
		l.CategoryID = stringNil
	}
	l.ProductID = q.Get(uriQueryParamProductID)
	ok = (len(l.ProductID) > 0)
	if ok {
		l.ProductID = strings.TrimPrefix(l.ProductID, "prod")
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

// GetAllLabo gets all Nintendo Labo products from the Nintendo store.
func GetAllLabo() []*Labo {
	return getAll(laboID)
}

// GetAllKits gets all Nintendo Labo products that are full kits from the Nintendo store.
func GetAllKits() []*Labo {
	return getAll(kitsID)
}

// GetAllParts get all Nintendo Labo products that are parts and accessories from the Nintendo store.
func GetAllParts() []*Labo {
	return getAll(partsID)
}

// Marshal marshals a Labo struct into an ordered byte sequence. On error returns an empty byte slice.
func Marshal(l *Labo) (b []byte) {
	b, _ = json.Marshal(l)
	return b
}

// Unmarshal unmarshals a ordered byte sequence.
func Unmarshal(b []byte) *Labo {
	var l *Labo
	json.Unmarshal(b, l)
	return l
}

// getAll gets all Nintendo Labo from the Nintendo store based on the argument ID.
func getAll(ID string) []*Labo {
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
		wg  sync.WaitGroup
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
		wg.Add(1)
		go func(ID string) {
			defer wg.Done()
			labo := Get(ID)
			if labo == nil {
				return
			}
			l = append(l, labo)
		}(ID)
	})
	wg.Wait()
	return l
}

func getLaboCategory(s *goquery.Selection, l *Labo) {
	const (
		CSS string = "#main-content .results-header"
	)
	var (
		category   = categoryLabo
		categoryID = categoryIDMap[category]
		ok         bool
		substring  string
	)
	l.Category = category
	l.CategoryID = categoryID
	s = s.Find(CSS)
	ok = (s.Length() > 0)
	if !ok {
		return
	}
	substring = strings.TrimSpace(s.Text())
	ok = regexp.MustCompile(`(?i)pack$`).MatchString(substring)
	if ok {
		category = categoryParts
	} else {
		category = categoryKit
	}
	l.CategoryID = categoryIDMap[category]
}

// getLaboDescription searches the argument goquery.Selection pointer
// for the Nintendo Labo product description and assigns it to
// the argument labo.Labo pointer if found..
func getLaboDescription(s *goquery.Selection, l *Labo) {
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
	l.Description = substring
}

// getLaboPageID searches the argument goquery.Selection pointer
// for the Nintendo Labo product numerical ID and assigns it to
// the argument labo.Labo pointer if found.
func getlaboID(s *goquery.Selection, l *Labo) {
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

// getLaboImages searches the argument goquery.Selection pointer
// for the Nintendo Labo product images and assigns them to
// the argument labo.Labo pointer if found.
func getLaboImages(s *goquery.Selection, l *Labo) {
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
	l.Images = newImages(s)
}

// getLaboName searches the argument goquery.Selection pointer
// for the Nintendo Labo product name and assigns the unformatted string to
// the argument labo.Labo pointer if found.
func getLaboName(s *goquery.Selection, l *Labo) {
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

// getLaboParts searches the argument goquery.Selection pointer
// for the Nintendo Labo product parts and components used to build
// the Nintendo Labo product and assigns them to
// the argument labo.Labo pointer if they are found.
func getLaboParts(s *goquery.Selection, l *Labo) {
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

// getLaboPrice searches the argument goquery.Selection pointer
// for the Nintendo Labo product price and assigns the floating point value to
// the argument labo.Labo pointer if found.
func getLaboPrice(s *goquery.Selection, l *Labo) {
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

// getLaboRef searches the argument goquery.Selection pointer
// for the Nintendo Labo product name alias (used to search the offical Labo site)
// and assigns the formatted string to the argument labo.Labo pointer if found.
func getLaboRef(s *goquery.Selection, l *Labo) {
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

// getLaboTitle searches the argument goquery.Selection pointer
// for the Nintendo Labo product headline and assigns the unformatted string to
// the argument labo.Labo pointer if found.
func getLaboTitle(s *goquery.Selection, l *Labo) {
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
	l.Title = substring
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
