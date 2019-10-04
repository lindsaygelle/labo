package labo

import (
	"fmt"
	"testing"

	"github.com/PuerkitoBio/goquery"
)

func TestNewStore(t *testing.T) {

	URL := "https://store.nintendo.com/ng3/us/po/browse/productDetailColorSizePicker.jsp?productId=prod940728"

	doc, _ := goquery.NewDocument(URL)

	s := newStore(doc)

	fmt.Println(s)
}
