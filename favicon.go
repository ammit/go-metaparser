package parser

// Favicon found in the head
type Favicon struct {
	Name  string `json:"name"`
	URL   string `json:"url"`
	Type  string `json:"type"`
	Sizes string `json:"sizes"`
}

func (p *Parser) parseFaviconLink(attrs map[string]string) {
	favicon := &Favicon{
		Name: attrs["rel"],
	}

	if val, ok := attrs["href"]; ok {
		favicon.URL = val
	}

	if val, ok := attrs["type"]; ok {
		favicon.Type = val
	}

	if val, ok := attrs["sizes"]; ok {
		favicon.Sizes = val
	}

	p.Favicons = append(p.Favicons, favicon)
}
