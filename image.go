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
	errorImageGoQueryNil      string = "argument goquery.Selection pointer cannot be nil"
	errorImageEmptyAttrAlt    string = "argument (*%p) does not contain an alt attribute"
	errorImageEmptyAttrSrc    string = "argument (*%p) does not contain a src attribute"
	errorImageEmptyAttrSrcSet string = "argument (*%p) does not contain a src-set attribute"
	errorImageEmptyFileExt    string = "argument (*%p) does not contain a file extension"
	errorImageEmptyHTMLNodes  string = "argument (*%p) does not contain a collection of HTML elements"
)

const (
	imageBase64Prefix string = "data:image"
)

const (
	imageAttrAlt        string = "alt"
	imageAttrDataSizes  string = "data-sizes"
	imageAttrDataSrc    string = "data-src"
	imageAttrDataSrcSet string = "data-srcset"
	imageAttrSizes      string = "sizes"
	imageAttrSrc        string = "src"
	imageAttrSrcSet     string = "srcset"
)

var (
	regexpImageReplaceFileExt     = regexp.MustCompile(`\W`)
	regexpImageReplaceFolderAlias = regexp.MustCompile(`\.{2}\/`)
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

// NewImage is a constructor function that instantiates a new Image pointer.
func NewImage(s *goquery.Selection) (*Image, error) {
	var (
		alt      string
		format   string
		sizes    string
		src      string
		srcset   string
		variants []*ImageVariant

		ok bool
	)
	if ok = (s != nil); !ok {
		return nil, fmt.Errorf(errorImageGoQueryNil)
	}
	if ok = (s.Length() > 0); !ok {
		return nil, fmt.Errorf(errorImageEmptyHTMLNodes, s)
	}
	alt = s.AttrOr(imageAttrAlt, defaultImageAttrAlt)
	alt = strings.ToUpper(alt)
	src, ok = s.Attr(imageAttrSrc)
	if !ok {
		return nil, fmt.Errorf(errorImageEmptyAttrSrc, s)
	}
	ok = strings.HasPrefix(src, imageBase64Prefix)
	if _, exists := s.Attr(imageAttrDataSrc); ok && exists {
		src, _ = s.Attr(imageAttrDataSrc)
	}
	ok = strings.HasPrefix(src, imageBase64Prefix)
	if ok {
		return nil, fmt.Errorf(errorImageEmptyAttrSrcSet, s)
	}
	format = filepath.Ext(src)
	format = regexpImageReplaceFileExt.ReplaceAllString(format, "")
	if ok = (len(format) > 0); !ok {
		return nil, fmt.Errorf(errorImageEmptyFileExt, s)
	}
	format = strings.ToUpper(format)
	src = regexpImageReplaceFolderAlias.ReplaceAllString(src, "")
	src = fmt.Sprintf("%s/%s", laboRootURL, src)
	if _, ok = s.Attr(imageAttrSrcSet); ok {
		srcset, _ = s.Attr(imageAttrSrc)
	}
	if _, ok = s.Attr(imageAttrDataSrcSet); ok && (len(srcset) == 0) {
		srcset, _ = s.Attr(imageAttrDataSrcSet)
	}
	if ok = (len(srcset) > 0); !ok {
		return nil, fmt.Errorf(errorImageEmptyAttrSrcSet, s)
	}
	for _, src := range strings.Split(srcset, ",") {
		imageVariant, err := NewImageVariant(src)
		if err != nil {
			panic(err)
		}
		variants = append(variants, imageVariant)
	}
	sizes, ok = s.Attr(imageAttrSizes)
	if _, exists := s.Attr(imageAttrDataSizes); !ok && exists {
		sizes = s.AttrOr(imageAttrDataSizes, defaultImageAttrSizes)
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
