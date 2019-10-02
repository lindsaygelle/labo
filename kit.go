package labo

import (
	"net/url"
	"time"

	"golang.org/x/text/currency"
	"golang.org/x/text/language"
)

type Kit struct {
	*Store

	Currency    currency.Unit `json:"currency"`
	Description string        `json:"description"`
	Href        *Href         `json:"href"`
	Index       int           `json:"index"`
	Language    language.Tag  `json:"language"`
	Price       float32       `json:"price"`
	Timestamp   time.Time     `json:"timestamp"`
	Title       string        `json:"title"`
	URL         *url.URL      `json:"URL"`
}

func newKit(i int, URL *url.URL) *Kit {

	var (
		currency    = currency.USD
		description string
		href        *Href
		language    = language.AmericanEnglish
		price       float32
		store       *Store
		timestamp   = time.Now().UTC()
		title       string
	)

	return &Kit{
		Store:       store,
		Currency:    currency,
		Description: description,
		Href:        href,
		Index:       i,
		Language:    language,
		Price:       price,
		Timestamp:   timestamp,
		Title:       title,
		URL:         URL}
}
