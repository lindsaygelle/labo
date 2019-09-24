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

	fmt.Println(kit.BoxImage)

	fmt.Println(kit.Overview)

	fmt.Println(kit.Software)

	fmt.Println(kit.Retailers)

	fmt.Println(kit.ToyCons)

	for _, toycon := range kit.ToyCons {
		//fmt.Println(toycon)
		for _, f := range toycon.Features {
			fmt.Println(f)
		}
	}
}
