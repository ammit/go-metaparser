package parser

import (
	"strconv"
)

type actor struct {
	URL  string `json:"url"`
	Role string `json:"role"`
}

// Video type in Open Graph
type Video struct {
	URL         string   `json:"url"`
	SecureURL   string   `json:"secure_url"`
	Type        string   `json:"type"`
	Width       int64    `json:"width"`
	Height      int64    `json:"height"`
	Actors      []*actor `json:"actors"`
	Director    string   `json:"director"`
	Writer      string   `json:"writer"`
	Duration    int64    `json:"duration"`
	ReleaseDate string   `json:"release_date"`
	Series      string   `json:"series"`
	Tags        []string `json:"tags"`
}

func (p *Parser) ensureVideo() {
	if len(p.Videos) > 0 {
		return
	}
	p.Videos = append(p.Videos, &Video{})
}

func (p *Parser) ensureVideoActor() {
	if len(p.Videos[len(p.Videos)-1].Actors) > 0 {
		return
	}

	p.Videos[len(p.Videos)-1].Actors = append(p.Videos[len(p.Videos)-1].Actors, &actor{})
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
	case "video:actor":
		p.ensureVideo()
		p.ensureVideoActor()
		p.Videos[len(p.Videos)-1].Actors[len(p.Videos[len(p.Videos)-1].Actors)-1].URL = attrs["content"]
	case "video:actor:role":
		p.ensureVideo()
		p.ensureVideoActor()
		p.Videos[len(p.Videos)-1].Actors[len(p.Videos[len(p.Videos)-1].Actors)-1].Role = attrs["content"]
	case "video:director":
		p.ensureVideo()
		p.Videos[len(p.Videos)-1].Director = attrs["content"]
	case "video:writer":
		p.ensureVideo()
		p.Videos[len(p.Videos)-1].Writer = attrs["content"]
	case "video:duration":
		h, err := strconv.ParseInt(attrs["content"], 10, 64)
		if err == nil {
			p.ensureVideo()
			p.Videos[len(p.Videos)-1].Duration = h
		}
	case "video:release_date":
		p.ensureVideo()
		p.Videos[len(p.Videos)-1].ReleaseDate = attrs["content"]
	case "video:tag":
		p.ensureVideo()
		p.Videos[len(p.Videos)-1].Tags = append(p.Videos[len(p.Videos)-1].Tags, attrs["content"])
	case "video:series":
		p.ensureVideo()
		p.Videos[len(p.Videos)-1].Series = attrs["content"]
	}
}
