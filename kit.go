package labo

import (
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/PuerkitoBio/goquery"
	"golang.org/x/text/currency"
	"golang.org/x/text/language"
)

type Kit struct {
	Category    string        `json:"category"`
	CategoryID  string        `json:"category_ID"`
	Currency    currency.Unit `json:"currency"`
	ID          int           `json:"ID"`
	Index       int           `json:"index"`
	Language    language.Tag  `json:"language"`
	Name        string        `json:"name"`
	Parts       []*Part       `json:"parts"`
	Price       float32       `json:"price"`
	ProductID   string        `json:"product_ID"`
	Ref         string        `json:"ref"`
	Status      string        `json:"http_status"`
	StatusCode  int           `json:"http_status_code"`
	StoreImages []*Image      `json:"store_images"`
	Time        time.Time     `json:"time"`
	URL         *url.URL      `json:"URL"`
}

func getKits() []*Kit {
	const (
		CSS string = ".product-listing .product-container > p a"
	)
	var (
		d   *goquery.Document
		err error
		ok  bool
		s   *goquery.Selection
	)
	res, err := client.Do(laboRequest)
	ok = (err == nil)
	if !ok {
		return nil
	}
	ok = (res.StatusCode == http.StatusOK)
	if !ok {
		return nil
	}
	d, err = goquery.NewDocumentFromResponse(res)
	ok = (err == nil)
	if !ok {
		return nil
	}
	s = d.Find(CSS)
	ok = (s.Length() > 0)
	if !ok {
		return nil
	}
	return newKits(s)
}

func getKitStorePage(k *Kit) {
	const (
		CSS string = "body"
	)
	var (
		ok bool
	)
	req, err := http.NewRequest(http.MethodGet, k.URL.String(), nil)
	ok = (err == nil)
	if !ok {
		return
	}
	res, err := client.Do(req)
	ok = (err == nil)
	if !ok {
		return
	}
	k.Status = res.Status
	k.StatusCode = res.StatusCode
	ok = (res.StatusCode == http.StatusOK)
	if !ok {
		return
	}
	doc, err := goquery.NewDocumentFromResponse(res)
	ok = (err == nil)
	if !ok {
		return
	}
	s := doc.Find(CSS)
	ok = (s.Length() > 0)
	if !ok {
		return
	}
	getKitStorePageImages(s, k)
	getKitStorePageName(s, k)
	getKitStorePageParts(s, k)
	getKitStorePagePrice(s, k)
}

func getKitStorePageImages(s *goquery.Selection, k *Kit) {
	const (
		CSS string = "#product-thumbs img"
	)
	var (
		ok bool
	)
	s = s.Find(CSS)
	ok = (s.Length() > 0)
	if !ok {
		return
	}
	k.StoreImages = newImages(s)
}

func getKitStorePageName(s *goquery.Selection, k *Kit) {
	const (
		CSS string = "#main-content .results-header"
	)
	var (
		ref        string
		substring  string
		substrings []string
		ok         bool
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
	k.Name = substring
	ok = strings.Contains(substring, ":")
	if !ok {
		return
	}
	substrings = strings.Split(substring, ":")
	ok = (len(substrings) >= 2)
	if !ok {
		return
	}
	ref = strings.TrimSpace(substrings[1])
	ok = (len(ref) > 0)
	if !ok {
		return
	}
	ref = strings.ToLower(ref)
	ok = (strings.Contains(ref, "+"))
	if ok {
		substrings = strings.Split(ref, "+")
		ref = strings.ReplaceAll(ref, substrings[1], "")
	}
	ref = regexpMatchNonAlphaNumericNoSpace.ReplaceAllString(ref, "-")
	ref = strings.TrimSpace(ref)
	ref = regexpMatchMultipleSpaces.ReplaceAllString(ref, "-")
	k.Ref = ref
}

func getKitStorePageParts(s *goquery.Selection, k *Kit) {
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
	k.Parts = newParts(s)
}

func getKitStorePagePrice(s *goquery.Selection, k *Kit) {
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
	k.Price = float32(price)
}

func newKit(i int, s *goquery.Selection) *Kit {
	const (
		HTML            string = "a"
		queryCategoryID string = "categoryId"
		queryProductID  string = "productId"
	)
	var (
		categoryID string
		href       string
		productID  string
		ok         bool
		time       = time.Now()
		kit        = (&Kit{})
	)
	ok = (s.Length() > 0)
	if !ok {
		return kit
	}
	ok = (strings.ToLower(s.Nodes[0].Data) == HTML)
	if !ok {
		return kit
	}
	href, ok = s.Attr(attrHref)
	if !ok {
		return kit
	}
	href = (storeURL + href)
	URL, err := url.Parse(href)
	ok = (err == nil)
	if !ok {
		return kit
	}
	queryMap := URL.Query()
	ok = (len(queryMap) > 0)
	if !ok {
		return kit
	}
	categoryID = queryMap.Get(queryCategoryID)
	categoryID = regexpMatchNonNumeric.ReplaceAllString(categoryID, "")
	ok = (len(categoryID) > 0)
	if !ok {
		return kit
	}
	productID = queryMap.Get(queryProductID)
	productID = regexpMatchNonNumeric.ReplaceAllString(productID, "")
	ok = (len(productID) > 0)
	if !ok {
		return kit
	}
	kit.CategoryID = categoryID
	kit.Currency = currency.USD
	kit.Language = language.AmericanEnglish
	kit.ProductID = productID
	kit.Index = i
	kit.Time = time
	kit.URL = URL

	getKitStorePage(kit)

	return kit
}

func newKits(s *goquery.Selection) []*Kit {
	var (
		kit  *Kit
		kits []*Kit
		ok   bool
		wg   sync.WaitGroup
	)
	s.Each(func(i int, s *goquery.Selection) {
		//go func(i int, s *goquery.Selection) {
		wg.Add(1)
		kit = newKit(i, s)
		wg.Done()
		ok = (kit != nil)
		if !ok {
			return
		}
		kits = append(kits, kit)
		//}(i, s)
	})
	wg.Wait()
	return kits
}
