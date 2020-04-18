package parser

import (
	"strconv"
)

// Video type in Open Graph
type Video struct {
	URL       string `json:"url"`
	SecureURL string `json:"secure_url"`
	Type      string `json:"type"`
	Width     int64  `json:"width"`
	Height    int64  `json:"height"`
}

func (p *Parser) ensureVideo() {
	if len(p.Videos) > 0 {
		return
	}
	p.Videos = append(p.Videos, &Video{})
}

func (p *Parser) parseVideoMeta(attrs map[string]string) {
	switch attrs["property"] {
	case "og:video":
		if len(p.Videos) > 0 && len(p.Videos[len(p.Videos)-1].URL) == 0 {
			p.Videos[len(p.Videos)-1].URL = attrs["content"]
		} else {
			p.Videos = append(p.Videos, &Video{URL: attrs["content"]})
		}
	case "og:video:url":
		p.ensureVideo()
		p.Videos[len(p.Videos)-1].URL = attrs["content"]
	case "og:video:secure_url":
		p.ensureVideo()
		p.Videos[len(p.Videos)-1].SecureURL = attrs["content"]
	case "og:video:type":
		p.ensureVideo()
		p.Videos[len(p.Videos)-1].Type = attrs["content"]
	case "og:video:width":
		w, err := strconv.ParseInt(attrs["content"], 10, 64)
		if err == nil {
			p.ensureVideo()
			p.Videos[len(p.Videos)-1].Width = w
		}
	case "og:video:height":
		h, err := strconv.ParseInt(attrs["content"], 10, 64)
		if err == nil {
			p.ensureVideo()
			p.Videos[len(p.Videos)-1].Height = h
		}
	}
}
