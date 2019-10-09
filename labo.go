package labo

import (
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

var (
	laboFn = [](func(s *goquery.Selection, l *Labo)){
		getLaboStorePageDescription,
		getLaboStorePageImages,
		getLaboStorePageName,
		getLaboStorePageParts,
		getLaboStorePagePrice,
		getLaboStorePageRef,
		getLaboStorePageTitle}
)

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
		return nil
	}
	res, err = client.Do(req)
	ok = (err == nil)
	if !ok {
		return nil
	}
	l.Status = res.Status
	l.StatusCode = res.StatusCode
	l.URL = req.URL
	ok = (res.StatusCode == http.StatusOK)
	if !ok {
		return nil
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

func GetAll(ID string) []*Labo {
	const (
		CSS string = ".product-listing .product-container > p a"
	)
	var (
		doc *goquery.Document
		err error
		ok  bool
		req *http.Request
		res *http.Response
		s   *goquery.Selection

		l = []*Labo{}
	)
	req, err = http.NewRequest(http.MethodGet, fmt.Sprintf("%scat%s", storeURI, ID), nil)
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
	return l
}

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

func newLabo(s *goquery.Selection, l *Labo) *Labo {
	for _, fn := range laboFn {
		fn(s, l)
	}
	return l
}

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
	substring = regexpMatchNonNumeric.ReplaceAllString(substring, emptyString)
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

func getLaboStorePageRef(s *goquery.Selection, l *Labo) {
	const (
		CSS string = "#main-content .results-header"
	)
	var (
		substring  string
		substrings []string
		ok         bool
	)
	ok = strings.Contains(substring, colonString)
	if !ok {
		return
	}
	substrings = strings.Split(substring, colonString)
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
	ok = (strings.Contains(substring, plusString))
	if ok {
		substrings = strings.Split(substring, plusString)
		substring = strings.ReplaceAll(substring, substrings[1], emptyString)
	}
	substring = regexpMatchNonAlphaNumericNoSpace.ReplaceAllString(substring, minusString)
	substring = strings.TrimSpace(substring)
	ok = (len(substring) > 0)
	if !ok {
		return
	}
	substring = regexpMatchMultipleSpaces.ReplaceAllString(substring, minusString)
	ok = (len(substring) > 0)
	if !ok {
		return
	}
	substring = strings.TrimSuffix(substring, minusString)
	l.Ref = substring
}

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
