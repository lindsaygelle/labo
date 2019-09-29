package labo

import (
	"fmt"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func GetVarietyKit() (*Kit, error) {
	req, err := http.NewRequest(http.MethodGet, URLVarietyKit, nil)
	if err != nil {
		return nil, err
	}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	ok := res.StatusCode == http.StatusOK
	if !ok {
		return nil, fmt.Errorf(res.Status)
	}
	doc, err := goquery.NewDocumentFromResponse(res)
	if err != nil {
		return nil, err
	}
	kit, err := NewKit(doc.Find("body"))
	return kit, err
}
