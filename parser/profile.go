package parser

// Profile type in Open Graph
type Profile struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Username  string `json:"username"`
	Gender    string `json:"gender"`
}

func (p *Parser) parseProfileMeta(attrs map[string]string) {
	switch attrs["property"] {
	case "profile:first_name":
		p.Profile.FirstName = attrs["content"]
	case "profile:last_name":
		p.Profile.LastName = attrs["content"]
	case "profile:username":
		p.Profile.Username = attrs["content"]
	case "profile:gender":
		p.Profile.Gender = attrs["content"]
	}
}
