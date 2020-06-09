package parser

type Result struct {
	Title            string   `json:"title"`
	Type             string   `json:"type"`
	Description      string   `json:"description"`
	Determiner       string   `json:"determiner"`
	URL              string   `json:"url"`
	Locale           string   `json:"locale"`
	LocalesAlternate []string `json:"locales_alternate"`
	SiteName         string   `json:"site_name"`

	Images []*Image `json:"images"`
	Videos []*Video `json:"videos"`
	Audios []*Audio `json:"audios"`

	// New
	Music   Music   `json:"music"`
	Article Article `json:"article"`
	Book    Book    `json:"book"`
	Profile Profile `json:"profile"`

	// Twitter
	Twitter Twitter `json:"twitter"`
}
