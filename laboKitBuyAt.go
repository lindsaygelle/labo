package main

import (
	"errors"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type LaboKitBuyAt struct {
	Name   string
	Target string
}

func NewLaboKitBuyAt(s *goquery.Selection) (*LaboKitBuyAt, error) {
	if ok := len(s.Nodes) != 0; ok != true {
		return nil, errors.New("HTML childNodes cannot be empty")
	}
	if ok := strings.ToUpper(s.Nodes[0].Data) == "A"; ok != true {
		return nil, errors.New("HTML must be an anchor tag")
	}
	return nil, nil
}
