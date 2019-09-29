package labo

import (
	"fmt"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

// GetRobotKit performs a HTTP GET request to the Nintendo Labo Robot Kit URL.
// Function attempts to build a Kit struct from the
// found HTML content using the defined selectors from the Kit struct to
// populate and build the struct content.
func GetRobotKit() (*Kit, error) {
	req, err := http.NewRequest(http.MethodGet, URLRobotKit, nil)
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
	body := doc.Find("body")
	ok = (body.Length() > 0)
	if !ok {
		return nil, fmt.Errorf(errorGoQueryDocumentEmptyHTMLNodes, doc)
	}
	return NewKit(body)
}
