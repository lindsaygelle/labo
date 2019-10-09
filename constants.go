package labo

const (
	colonString      string = ":"
	emptyString      string = ""
	minusString      string = "-"
	plusString       string = "+"
	whitespaceString string = " "
)

const (
	// attrAlt is the HTML attribute namespace for alternates.
	attrAlt string = "alt"
	// attrClass is the HTML attribute namespace for classes.
	attrClass string = "class"
	// attrDataSizes is the HTML data-attribute namespace for data-sizes.
	attrDataSizes string = "data-sizes"
	// attrDataSrc is the HTML data-attribute namespace for data-src.
	attrDataSrc string = "data-src"
	// attrDataSrcSet is the HTML data-attribute namespace for data-srcset.
	attrDataSrcSet string = "data-srcset"
	// attrHref is the HTML attribute namespace for href's.
	attrHref string = "href"
	// attrSizes is the HTML attribute for namespace sizes.
	attrSizes string = "sizes"
	// attrSrc is the HTML attribute for namespace src's.
	attrSrc string = "src"
	// attrSrcSet is the HTML attribute for namespace srcset's.
	attrSrcSet string = "srcset"
	// attrTarget is the HTML attribute for namespace target's.
	attrTarget string = "target"
)

const (
	// htmlAnchor is the HTML element namespace for anchor tags.
	htmlAnchor string = "a"
	// htmlBody is the HTML element namespace for the body tag.
	htmlBody string = "body"
	// htmlImage is the HTML element namespace for image tags.
	htmlImage string = "img"
)

const (
	// laboDNS is the domain name reference for the Nintendo Labo website.
	laboDNS string = ("labo" + "." + nintendoDNS)
)

const (
	// laboURI is the URI directive to perform a Nintendo Labo site search for Nintendo Labo kits.
	laboURI string = (laboURL + "/" + "kits")
)

const (
	// laboURL is the RFC2616 compliant address for the Nintendo Labo website.
	laboURL string = ("https://" + laboDNS)
)

const (
	// nintendoDNS is the domain name reference for the Nintendo official website.
	nintendoDNS string = "nintendo.com"
)

const (
	// nintendoURL is the RFC2616 compliant address for the Nintendo official website.
	nintendoURL string = ("https://" + nintendoDNS)
)

const (
	// storeDNS is the domain name reference for the Nintendo store website.
	storeDNS string = ("store" + "." + nintendoDNS)
)

const (
	// storeURI is the URI directive to perform a Nintendo store search for Nintendo Labo kits.
	storeURI string = (storeURL + "/" + "ng3/us/po/browse/subcategory.jsp?viewAll=true&categoryId=")
)

const (
	storeProductURI string = (storeURL + "/" + "ng3/us/po/browse/productDetailColorSizePicker.jsp?productId=")
)

const (
	// storeURIKits is the URI directive to request all Nintendo Labo full kits from the Nintendo store.
	storeURIKits string = (storeURI + "cat" + kitsID)
)

const (
	// storeURILabo is the URI directive to request all Nintendo Labo full kits and parts from the Nintendo store.
	storeURILabo string = (storeURI + "cat" + laboID)
)

const (
	// storeURIParts is the URI directive to request all Nintendo Labo part kits from the Nintendo store.
	storeURIParts string = (storeURI + "cat" + partsID)
)

const (
	// storeURL is the RFC2616 compliant address for the Nintendo store website.
	storeURL string = ("https://" + storeDNS)
)

const (
	// kitsID is the product ID for all Nintendo Labo full kits.
	kitsID string = "970105"
	// laboID is the product ID for both Nintendo Labo full kits and parts.
	laboID string = "960195"
	// partsID is the product ID for all Nintendo Labo parts kits.
	partsID string = "970106"
)

const (
	patternIgnorecase string = "(?i)(%s)"
)

const (
	partAmountOne      string = "one"
	partAmountTwo      string = "two"
	partAmountThree    string = "three"
	partAmountFour     string = "four"
	partAmountFive     string = "five"
	partAmountSix      string = "six"
	partAmountSeven    string = "seven"
	partAmountEight    string = "eight"
	partAmountNine     string = "nine"
	partAmountTen      string = "ten"
	partAmountEleven   string = "eleven"
	partAmountTwelve   string = "twelve"
	partAmountThirteen string = "thirteen"
)

const (
	// partColorBlue is the namespace for blue Nintendo Labo parts.
	partColorBlue string = "blue"
	// partColorGray is the namespace for gray Nintendo Labo parts.
	partColorGray string = "gray"
	// partColorOrange is the namespace for orange Nintendo Labo parts.
	partColorOrange string = "orange"
	// partColorRed is the namespace for red Nintendo Labo parts.
	partColorRed string = "red"
	// partColorYellow is the namespace for yellow Nintendo Labo parts.
	partColorYellow string = "yellow"
)

const (
	// partGenderFemale is the namespace for Nintendo Labo parts that are of the female configuration.
	partGenderFemale string = "female"
	// partGenderMail is the alias namespace for a known typo for male parts.
	partGenderMail string = "mail"
	// partsGenderMale is the namespace for Nintendo Labo parts that of the male configuration.
	partGenderMale string = "male"
)

const (
	// defaultAttrAlt is the default namespace for HTML alt attributes.
	defaultAttrAlt string = "NIL"
	// defaultAttrTarget is the default namespace for HTML target attributes.
	defaultAttrTarget string = "NIL"
)

const (
	defaultLaboName string = "NIL"
	defaultLaboRef  string = "NIL"
)

const (
	// defaultPartColor is the default color namespace for Nintendo Labo parts.
	defaultPartColor string = "NIL"
	// defaultPartGender is the default gender namespace for Nintendo Labo parts.
	defaultPartGender string = "NEUTRAL"
	// defaultPartShape is the default shape namespace for Nintendo Labo parts.
	defaultPartShape string = "NIL"
	// defaultPartSize is the default size namespace for Nintendo Labo parts.
	defaultPartSize string = "REGULAR"
)

const (
	// partShapeOctagonal is the namespace for Nintendo Labo parts that are of octagonal shape.
	partShapeOctagonal string = "octagonal"
	// partShapeSquare is the namespace for Nintendo Labo parts that are a square shape.
	partShapeSquare string = "square"
)

const (
	// partSizeLarge is the namespace for Nintendo Labo parts that are of large size.
	partSizeLarge string = "large"
	// partSizeMedium is the namespace for Nintendo Labo parts that are of medium size.
	partSizeMedium string = "medium"
	// partSizeSmall is the namespace for Nintendo Labo parts that are of smaller size.
	partSizeSmall string = "small"
)

const (
	uriQueryParamCategoryID string = "categoryId"
	uriQueryParamProductID  string = "productId"
)
