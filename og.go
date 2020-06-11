package parser

type OG struct {
	Title            string   `json:"title"`
	Type             string   `json:"type"`
	Description      string   `json:"description"`
	Determiner       string   `json:"determiner"`
	URL              string   `json:"url"`
	Locale           string   `json:"locale"`
	LocalesAlternate []string `json:"locales_alternate"`
	SiteName         string   `json:"site_name"`
}

func (p *Parser) parseBasicOGMeta(attrs map[string]string) {
	switch attrs["property"] {
	case "og:title":
		p.OpenGraph.Title = attrs["content"]
	case "og:type":
		p.OpenGraph.Type = attrs["content"]
	case "og:url":
		p.OpenGraph.URL = attrs["content"]
	case "og:description":
		p.OpenGraph.Description = attrs["content"]
	case "og:determiner":
		p.OpenGraph.Determiner = attrs["content"]
	case "og:locale":
		p.OpenGraph.Locale = attrs["content"]
	case "og:locale:alternate":
		p.OpenGraph.LocalesAlternate = append(p.OpenGraph.LocalesAlternate, attrs["content"])
	case "og:site_name":
		p.OpenGraph.SiteName = attrs["content"]

	}
}
