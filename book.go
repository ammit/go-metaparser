package parser

// Book type in Open Graph
type Book struct {
	ReleaseDate string   `json:"release_date"`
	Isbn        string   `json:"isbin"`
	Tags        []string `json:"tags"`
	Authors     []string `json:"authors"`
}

func (p *Parser) parseBookMeta(attrs map[string]string) {
	switch attrs["property"] {
	case "book:author":
		p.Book.Authors = append(p.Book.Authors, attrs["content"])
	case "book:isbn":
		p.Book.Isbn = attrs["content"]
	case "book:release_date":
		p.Book.ReleaseDate = attrs["content"]
	case "book:tag":
		p.Book.Tags = append(p.Book.Tags, attrs["content"])
	}
}
