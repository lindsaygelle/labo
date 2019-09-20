package labo

import (
	"fmt"
	"path/filepath"
	"regexp"
	"strings"
)

const (
	errorImageVariantEmptyAttr           string = "argument (%s) attributes cannot be empty"
	errorImageVariantEmptyExt            string = "argument (%s) file extension cannot be empty"
	errorImageVariantEmptySrc            string = "argument (%s) src substring cannot be empty"
	errorImageVariantEmptyString         string = "argument string cannot be empty"
	errorImageVariantIncorrectSubstrings string = "argument (%s) cannot be split into required substring"
	errorImageVariantNoAttributes        string = "argument (%s) does not contain required substring pattern"
	errorImageVariantNoDelimiter         string = "argument (%s) does not contain a whitespace delimiter"
	errorImageVariantNoFileExtension     string = "argument (%s) does not contain a file extension"
)

type ImageVariant struct {
	At     string
	Format string
	Src    string
	Units  string
	Value  string
}

func NewImageVariant(s string) (*ImageVariant, error) {
	var (
		at     string
		format string
		src    string
		units  string
	)
	if ok := (len(s) > 0); !ok {
		return nil, fmt.Errorf(errorImageVariantEmptyString)
	}
	s = strings.TrimSpace(s)
	if ok := (strings.Contains(s, " ")); !ok {
		return nil, fmt.Errorf(errorImageVariantNoDelimiter, s)
	}
	substrings := strings.Split(s, " ")
	if ok := (len(substrings) == 2); !ok {
		return nil, fmt.Errorf(errorImageVariantIncorrectSubstrings, s)
	}
	s1 := strings.TrimSpace(substrings[0])
	if ok := (len(s1) > 0); !ok {
		return nil, fmt.Errorf(errorImageVariantEmptySrc, s1)
	}
	s2 := strings.TrimSpace(substrings[1])
	if ok := (len(s2) > 0); !ok {
		return nil, fmt.Errorf(errorImageVariantEmptyAttr, s2)
	}
	at = regexp.MustCompile(`\D`).ReplaceAllString(s2, "")
	units = strings.ToUpper(strings.Replace(s2, at, "", 1))
	format = filepath.Ext(s1)
	format = regexp.MustCompile(`\W`).ReplaceAllString(format, "")
	if ok := (len(format) > 0); !ok {
		return nil, fmt.Errorf(errorImageVariantEmptyExt, format)
	}
	format = strings.ToUpper(format)
	imageVariant := ImageVariant{
		At:     at,
		Format: format,
		Src:    src,
		Units:  units}
	return &imageVariant, nil
}
