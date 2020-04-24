package parser

import (
	"strconv"
)

// Image type in Open Graph
type Image struct {
	URL       string `json:"url"`
	SecureURL string `json:"secure_url"`
	Type      string `json:"type"`
	Width     int64  `json:"width"`
	Height    int64  `json:"height"`
	Alt       string `json:"alt"`
}

func (p *Parser) ensureImage() {
	if len(p.Images) > 0 {
		return
	}
	p.Images = append(p.Images, &Image{})
}

func (p *Parser) parseImageMeta(attrs map[string]string) {
	switch attrs["property"] {
	case "og:image":
		if len(p.Images) > 0 && len(p.Images[len(p.Images)-1].URL) == 0 {
			p.Images[len(p.Images)-1].URL = attrs["content"]
		} else {
			p.Images = append(p.Images, &Image{URL: attrs["content"]})
		}
	case "og:image:url":
		p.ensureImage()
		p.Images[len(p.Images)-1].URL = attrs["content"]
	case "og:image:secure_url":
		p.ensureImage()
		p.Images[len(p.Images)-1].SecureURL = attrs["content"]
	case "og:image:type":
		p.ensureImage()
		p.Images[len(p.Images)-1].Type = attrs["content"]
	case "og:image:width":
		w, err := strconv.ParseInt(attrs["content"], 10, 64)
		if err == nil {
			p.ensureImage()
			p.Images[len(p.Images)-1].Width = w
		}
	case "og:image:height":
		h, err := strconv.ParseInt(attrs["content"], 10, 64)
		if err == nil {
			p.ensureImage()
			p.Images[len(p.Images)-1].Height = h
		}
	case "og:image:alt":
		p.ensureImage()
		p.Images[len(p.Images)-1].Alt = attrs["content"]
	}
}
