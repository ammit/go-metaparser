package parser

// Article type in Open Graph
type Article struct {
	PublishedTime  string   `json:"published_time"`
	ModifiedTime   string   `json:"modified_time"`
	ExpirationTime string   `json:"expiration_time"`
	Section        string   `json:"section"`
	Tags           []string `json:"tags"`
	Authors        []string `json:"authors"`
}

func (p *Parser) parseArticleMeta(attrs map[string]string) {
	switch attrs["property"] {
	case "article:published_time":
		p.Article.PublishedTime = attrs["content"]
	case "article:modified_time":
		p.Article.ModifiedTime = attrs["content"]
	case "article:expiration_time":
		p.Article.ExpirationTime = attrs["content"]
	case "article:section":
		p.Article.Section = attrs["content"]
	case "article:author":
		p.Article.Authors = append(p.Article.Authors, attrs["content"])
	case "article:tag":
		p.Article.Tags = append(p.Article.Tags, attrs["content"])
	}
}
