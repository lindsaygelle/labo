package labo_test

import (
	"fmt"
	"testing"

	"github.com/PuerkitoBio/goquery"
	"github.com/gellel/labo"
)

func TestLabo(t *testing.T) {

	doc, err := goquery.NewDocument("https://labo.nintendo.com/kits/variety-kit/")
	if err != nil {
		panic(err)
	}

	kit, err := labo.NewKit(doc.Find("body"))

	if err != nil {
		panic(err)
	}

	fmt.Println(kit)

	doc, err = goquery.NewDocument("https://labo.nintendo.com/kits/customization-set/")
	if err != nil {
		panic(err)
	}

	customization, err := labo.NewKitCustomization(doc.Find("body"))
	if err != nil {
		panic(err)
	}
	fmt.Println(customization)

}
