package parser

type Result struct {
	Title       string `json:"title"`
	Description string `json:"description"`

	OpenGraph OG `json:"open_graph"`

	Images []*Image `json:"images"`
	Videos []*Video `json:"videos"`
	Audios []*Audio `json:"audios"`

	// New
	Music   Music   `json:"music"`
	Article Article `json:"article"`
	Book    Book    `json:"book"`
	Profile Profile `json:"profile"`

	Favicons []*Favicon `json:"favicons"`

	// Twitter
	Twitter Twitter `json:"twitter"`
}

// GetTitle returns either Open Graph title or standard title as fallback
func (result *Result) GetTitle() string {
	if len(result.OpenGraph.Title) > 0 {
		return result.OpenGraph.Title
	} else {
		return result.Title
	}
}

// GetDescription returns either Open Graph description or standard description as fallback
func (result *Result) GetDescription() string {
	if len(result.OpenGraph.Description) > 0 {
		return result.OpenGraph.Description
	} else {
		return result.Description
	}
}
