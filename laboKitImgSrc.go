package main

import (
	"fmt"
	"net/url"
	"strconv"
	"strings"
)

type LaboKitImgSrc struct {
	At   int      `json:"at"`
	Href string   `json:"href"`
	URL  *url.URL `json:"URL"`
}

func NewLaboKitImgSrc(s string) *LaboKitImgSrc {
	substrings := strings.Split(s, " ")
	if ok := (len(substrings) == 2); ok != true {
		return nil
	}
	s1 := strings.TrimSpace(substrings[0])
	href := strings.ReplaceAll(s1, "../", "")
	href = fmt.Sprintf("%s/%s", nintendoLaboURL, href)
	URL, err := url.Parse(href)
	if err != nil {
		return nil
	}
	s2 := strings.TrimSpace(substrings[1])
	at, err := strconv.Atoi(strings.TrimSuffix(s2, "w"))
	if err != nil {
		return nil
	}
	return &LaboKitImgSrc{
		At:   at,
		Href: href,
		URL:  URL}
}
