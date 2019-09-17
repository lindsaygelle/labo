package main

import (
	"fmt"
	"net/url"
	"regexp"
	"strconv"
	"strings"
)

type LaboKitImgSrc struct {
	At          int      `json:"at"`
	Href        string   `json:"href"`
	Measurement string   `json:"measurement"`
	URL         *url.URL `json:"URL"`
}

func NewLaboKitImgSrc(s string) *LaboKitImgSrc {
	substrings := strings.Split(s, " ")
	if ok := (len(substrings) == 2); ok != true {
		return nil
	}
	s1 := strings.TrimSpace(substrings[0])
	href := strings.ReplaceAll(s1, "../", "")
	href = fmt.Sprintf("%s/%s", nintendoLaboURL, href)

	fmt.Println(href)
	URL, err := url.Parse(href)
	if err != nil {
		return nil
	}
	s2 := strings.TrimSpace(substrings[1])
	at, err := strconv.Atoi(string(s2[:1]))
	if err != nil {
		return nil
	}
	measurement := regexp.MustCompile(`\d+`).ReplaceAllString(s2, "")
	return &LaboKitImgSrc{
		At:          at,
		Href:        href,
		Measurement: measurement,
		URL:         URL}
}
