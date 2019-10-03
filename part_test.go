package labo

import (
	"fmt"
	"testing"

	"github.com/PuerkitoBio/goquery"
)

func TestNewPart(t *testing.T) {
	for _, URL := range []string{
		"prod940728",
		"prod940731"} {

		URL := fmt.Sprintf("https://store.nintendo.com/ng3/us/po/browse/productDetailColorSizePicker.jsp?productId=%s", URL)

		doc, _ := goquery.NewDocument(URL)

		body := doc.Find("body")

		prodDescBtm := body.Find("#prodDescBtm ul li")

		prodDescBtm.Each(func(i int, s *goquery.Selection) {
			p := newPart(s)
			fmt.Println(p)
			if p.Href != nil {
				fmt.Println(p.Href.URL)
			}
		})
	}
}

func TestPartAmount(t *testing.T) {

	s := []string{
		"One Short Straps",
		"Two Short Strap",
		"Three Short Strap",
		"Four Short Strings",
		"Five Short Strings",
		"Six Short Sheet",
		"Seven Short Sheet",
		"Eight",
		"Nine",
		"Ten",
		"Eleven",
		"Twelve Grommets",
		"Thirteen",
		"Carboard sheet x 32"}

	for _, s := range s {
		fmt.Println(getPartAmount(s))
	}
}

func TestPartColor(t *testing.T) {

	s := []string{
		"Blue",
		"(Blue)",
		"Gray",
		"One Orange",
		"Red",
		"Yellow"}

	for _, s := range s {
		fmt.Println(getPartColor(s))
	}
}

func TestPartGender(t *testing.T) {

	s := []string{
		"Short gray grommets (male) x 6",
		"Long yellow grommets (mail) x 11",
		"Long yellow grommets (female) x 11"}

	for _, s := range s {
		fmt.Println(getPartGender(s))
	}
}

func TestPartName(t *testing.T) {
	s := []string{"Long yellow grommets (female) x 11"}
	for _, s := range s {
		fmt.Println(getPartName(s))
	}
}

func TestPartSpares(t *testing.T) {

	s := []string{
		"Rubber bands (large) x 3 + spares (latex)"}

	for _, s := range s {
		fmt.Println(getPartSpares(s))
	}
}
