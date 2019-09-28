package labo

import (
	"fmt"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

const (
	defaultImageAttrAlt   string = "NIL"
	defaultImageAttrSizes string = "NIL"
	defaultImageAttrSrc   string = "NIL"
)

const (
	errorImageEmptyAttrAlt    string = "argument (*%p) does not contain an alt attribute"
	errorImageEmptyAttrSrc    string = "argument (*%p) does not contain a src attribute"
	errorImageEmptyAttrSrcSet string = "argument (*%p) does not contain a src-set attribute"
	errorImageEmptyFileExt    string = "argument (*%p) does not contain a file extension"
)

const (
	imageBase64Prefix string = "data:image"
)

var (
	regexpImageMatchFileExt = regexp.MustCompile(`\W`)
	regexpImageMatchFolder  = regexp.MustCompile(`\.{2}\/`)
)

// Image is a image resource that contains a related image for Nintendo Labo.
type Image struct {
	Alt      string
	Format   string
	Size     int
	Sizes    string
	Src      string
	Variants []*ImageVariant
}

// NewImage is a constructor function that instantiates and returns a new Image pointer.
func NewImage(s *goquery.Selection) (*Image, error) {
	var (
		ok bool
	)
	ok = (s != nil)
	if !ok {
		return nil, fmt.Errorf(errorGoQuerySelectionNil)
	}
	ok = (s.Length() > 0)
	if !ok {
		return nil, fmt.Errorf(errorGoQuerySelectionEmptyHTMLNodes, s)
	}
	var (
		alt      string
		format   string
		sizes    string
		src      string
		srcset   string
		variants []*ImageVariant
	)
	alt = s.AttrOr(attrAlt, defaultImageAttrAlt)
	alt = strings.ToUpper(alt)
	src, ok = s.Attr(attrSrc)
	if !ok {
		return nil, fmt.Errorf(errorImageEmptyAttrSrc, s)
	}
	ok = strings.HasPrefix(src, imageBase64Prefix)
	if _, exists := s.Attr(attrDataSrc); ok && exists {
		src, _ = s.Attr(attrDataSrc)
	}
	ok = strings.HasPrefix(src, imageBase64Prefix)
	if ok {
		return nil, fmt.Errorf(errorImageEmptyAttrSrcSet, s)
	}
	format = filepath.Ext(src)
	format = regexpImageMatchFileExt.ReplaceAllString(format, "")
	if ok = (len(format) > 0); !ok {
		return nil, fmt.Errorf(errorImageEmptyFileExt, s)
	}
	format = strings.ToUpper(format)
	src = regexpImageMatchFolder.ReplaceAllString(src, "")
	src = fmt.Sprintf("%s/%s", laboRootURL, src)
	_, ok = s.Attr(attrSrcSet)
	if ok {
		srcset, _ = s.Attr(attrSrc)
	}
	_, ok = s.Attr(attrDataSrcSet)
	ok = (ok && (len(srcset) == 0))
	if ok {
		srcset, _ = s.Attr(attrDataSrcSet)
	}
	for _, src := range strings.Split(srcset, ",") {
		imageVariant, err := NewImageVariant(src)
		if err != nil {
			continue
		}
		variants = append(variants, imageVariant)
	}
	sizes, ok = s.Attr(attrSizes)
	_, exists := s.Attr(attrDataSizes)
	ok = (!ok && exists)
	if ok {
		sizes = s.AttrOr(attrDataSizes, defaultImageAttrSizes)
	}
	sizes = strings.ToUpper(sizes)
	image := Image{
		Alt:      alt,
		Format:   format,
		Sizes:    sizes,
		Src:      src,
		Variants: variants}
	return &image, nil
}
