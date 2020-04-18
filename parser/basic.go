package parser

func (p *Parser) parseBasicMeta(attrs map[string]string) {
	switch attrs["property"] {
	case "og:title":
		p.Title = attrs["content"]
	case "og:type":
		p.Type = attrs["content"]
	case "og:url":
		p.URL = attrs["content"]
	case "og:description":
		p.Description = attrs["content"]
	case "og:determiner":
		p.Determiner = attrs["content"]
	case "og:locale":
		p.Locale = attrs["content"]
	case "og:locale:alternate":
		p.LocalesAlternate = append(p.LocalesAlternate, attrs["content"])
	case "og:site_name":
		p.SiteName = attrs["content"]

	}
}
