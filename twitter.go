package parser

import "strconv"

type player struct {
	URL    string `json:"url"`
	Width  int64  `json:"width"`
	Height int64  `json:"height"`
	Stream string `json:"stream"`
}

type app struct {
	Name string `json:"name"`
	ID   int64  `json:"id"`
	URL  string `json:"url"`
	Type string `json:"type"`
}

// Twitter type in Open Graph
type Twitter struct {
	Card        string `json:"card"`
	Site        string `json:"site"`
	SiteID      string `json:"site_id"`
	Creator     string `json:"creator"`
	CreatorID   string `json:"creator_id"`
	Description string `json:"description"`
	Title       string `json:"title"`
	Image       string `json:"image"`
	ImageAlt    string `json:"image_alt"`
	Player      player `json:"player"`
	Apps        []app  `json:"apps"`
}

func (p *Parser) ensureApp(appType string) {
	exist := false
	for i := 0; i < len(p.Twitter.Apps); i++ {
		attr := &p.Twitter.Apps[i]
		if attr.Type == appType {
			exist = true
			break
		}
	}

	if !exist {
		p.Twitter.Apps = append(p.Twitter.Apps, app{
			Type: appType,
		})
	}
}

func (p *Parser) parseTwitterMeta(attrs map[string]string) {
	switch attrs["property"] {
	case "twitter:card":
		p.Twitter.Card = attrs["content"]
	case "twitter:site":
		p.Twitter.Site = attrs["content"]
	case "twitter:site:id":
		p.Twitter.SiteID = attrs["content"]
	case "twitter:creator":
		p.Twitter.Creator = attrs["content"]
	case "twitter:creator:id":
		p.Twitter.CreatorID = attrs["content"]
	case "twitter:description":
		p.Twitter.Description = attrs["content"]
	case "twitter:title":
		p.Twitter.Title = attrs["content"]
	case "twitter:image":
		p.Twitter.Image = attrs["content"]
	case "twitter:image:alt":
		p.Twitter.ImageAlt = attrs["content"]
	case "twitter:player":
		p.Twitter.Player.URL = attrs["content"]
	case "twitter:player:height":
		w, err := strconv.ParseInt(attrs["content"], 10, 64)
		if err == nil {
			p.Twitter.Player.Height = w
		}
	case "twitter:player:width":
		w, err := strconv.ParseInt(attrs["content"], 10, 64)
		if err == nil {
			p.Twitter.Player.Width = w
		}
	case "twitter:player:stream":
		p.Twitter.Player.Stream = attrs["content"]
	case "twitter:app:name:iphone":
		p.ensureApp("iphone")

		for i := range p.Twitter.Apps {
			attr := &p.Twitter.Apps[i]
			if attr.Type == "iphone" {
				attr.Name = attrs["content"]
				break
			}
		}
	case "twitter:app:url:iphone":
		p.ensureApp("iphone")

		for i := range p.Twitter.Apps {
			attr := &p.Twitter.Apps[i]
			if attr.Type == "iphone" {
				attr.URL = attrs["content"]
				break
			}
		}
	case "twitter:app:id:iphone":
		p.ensureApp("iphone")
		w, err := strconv.ParseInt(attrs["content"], 10, 64)
		if err == nil {
			for i := range p.Twitter.Apps {
				attr := &p.Twitter.Apps[i]
				if attr.Type == "iphone" {
					attr.ID = w
					break
				}
			}
		}
	case "twitter:app:name:ipad":
		p.ensureApp("ipad")

		for i := range p.Twitter.Apps {
			attr := &p.Twitter.Apps[i]
			if attr.Type == "ipad" {
				attr.Name = attrs["content"]
				break
			}
		}
	case "twitter:app:url:ipad":
		p.ensureApp("ipad")

		for i := range p.Twitter.Apps {
			attr := &p.Twitter.Apps[i]
			if attr.Type == "ipad" {
				attr.URL = attrs["content"]
				break
			}
		}
	case "twitter:app:id:ipad":
		w, err := strconv.ParseInt(attrs["content"], 10, 64)
		if err == nil {
			p.ensureApp("ipad")

			for i := range p.Twitter.Apps {
				attr := &p.Twitter.Apps[i]
				if attr.Type == "ipad" {
					attr.ID = w
					break
				}
			}
		}
	case "twitter:app:name:googleplay":
		p.ensureApp("googleplay")

		for i := range p.Twitter.Apps {
			attr := &p.Twitter.Apps[i]
			if attr.Type == "googleplay" {
				attr.Name = attrs["content"]
				break
			}
		}
	case "twitter:app:url:googleplay":
		p.ensureApp("googleplay")

		for i := range p.Twitter.Apps {
			attr := &p.Twitter.Apps[i]
			if attr.Type == "googleplay" {
				attr.URL = attrs["content"]
				break
			}
		}
	case "twitter:app:id:googleplay":
		w, err := strconv.ParseInt(attrs["content"], 10, 64)
		if err == nil {
			p.ensureApp("googleplay")

			for i := range p.Twitter.Apps {
				attr := &p.Twitter.Apps[i]
				if attr.Type == "googleplay" {
					attr.ID = w
					break
				}
			}
		}
	}
}

// TODO: mapping if not exisits
// twitter:card <= og:type
// twitter:description <= og:description
// twitter:title <= og:title
// twitter:image <= og:image
