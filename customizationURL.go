package labo

import "net/url"

// CustomizationURL is a string that points to a Nintendo Labo Customization Kit URL.
type CustomizationURL string

func (CustomizationURL CustomizationURL) String() string {
	return string(CustomizationURL)
}

// URL is a function that parses the Nintendo Labo Customization Kit URL into a *url.URL.
func (CustomizationURL CustomizationURL) URL() (*url.URL, error) {
	return url.Parse(CustomizationURL.String())
}
