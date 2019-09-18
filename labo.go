package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/PuerkitoBio/goquery"
	"golang.org/x/text/currency"
)

type Kit struct {
	Href      string
	ID        string
	Materials Materials
	Name      string
	Overview  Overview
	Projects  []Project
	Price     Price
	Retailers []Retailer
	Software  Software
	ToyCons   []ToyCon
}

type KitCustomization struct {
}

type KitVR struct {
}

type Feature struct {
	Description string
	Icon        Image
	Image       Image
	Title       string
	Video       Video
}

type Image struct {
	Alt    string
	Format string
	Size   int
	Src    string
}

type Materials struct {
	Image Image
	Parts []Part
}

type Overview struct {
	Description string
	Video       Video
}

type Part struct {
	Amount int
	Name   string
	Size   string
}

type Price struct {
	Amount   float64
	Currency currency.Unit
}

type Project struct {
	Icon        Image
	Image       Image
	Name        string
	Screenshots []Image
}

type Retailer struct {
	Href string
	Logo Image
	Name string
}

type Software struct {
	Image Image
	Video Video
}

type ToyCon struct {
	Description string
	Features    []Feature
	Image       Image
	Name        string
}

type Video struct{}

const nintendoLaboURL string = "https://labo.nintendo.com"

func main() {
	c := (&http.Client{Timeout: time.Second * 10})
	req, err := http.NewRequest(http.MethodGet, (nintendoLaboURL + "/kits/vr-kit/"), nil)
	if err != nil {
		panic(err)
	}
	res, err := c.Do(req)
	if err != nil {
		panic(err)
	}
	if res.StatusCode != http.StatusOK {
		panic(res.Status)
	}
	doc, err := goquery.NewDocumentFromResponse(res)
	if err != nil {
		panic(err)
	}
	laboVRKit := NewLaboVRKit(doc)

	fmt.Println(laboVRKit.BoxImageURL.Hostname())
}
