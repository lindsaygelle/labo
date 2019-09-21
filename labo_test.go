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

	materials, err := labo.NewMaterials(doc.Find(".contents .right-column"))

	fmt.Println(materials, err)

	fmt.Println(materials.Image)

	for _, v := range materials.Image.Variants {
		fmt.Println(v)
	}
}
