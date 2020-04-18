// TODO - The album should be array.

package parser

import (
	"strconv"
)

type song struct {
	URL   string `json:"url"`
	Track int64  `json:"track"`
	Disc  int64  `json:"disc"`
}

type album struct {
	URL   string `json:"url"`
	Track int64  `json:"track"`
}

// Music type in Open Graph
type Music struct {
	Musicians   []string `json:"musicians"`
	Album       album    `json:"album"`
	Duration    int64    `json:"duration"`
	ReleaseDate string   `json:"release_date"`
	Creator     string   `json:"creator"`
	Songs       []*song  `json:"songs"`
}

func (p *Parser) ensureSongs() {
	if len(p.Music.Songs) > 0 {
		return
	}
	p.Music.Songs = append(p.Music.Songs, &song{})
}

func (p *Parser) parseMusicMeta(attrs map[string]string) {
	switch attrs["property"] {
	case "music:musician":
		p.Music.Musicians = append(p.Music.Musicians, attrs["content"])
	case "music:album":
		p.Music.Album.URL = attrs["content"]
	case "music:duration":
		t, err := strconv.ParseInt(attrs["content"], 10, 64)
		if err == nil {
			p.Music.Duration = t
		}
	case "music:release_date":
		p.Music.ReleaseDate = attrs["content"]
	case "music:creator":
		p.Music.Creator = attrs["content"]
	case "music:album:track":
		t, err := strconv.ParseInt(attrs["content"], 10, 64)
		if err == nil {
			p.Music.Album.Track = t
		}
	case "music:song":
		if len(p.Music.Songs) > 0 && len(p.Music.Songs[len(p.Music.Songs)-1].URL) == 0 {
			p.Music.Songs[len(p.Music.Songs)-1].URL = attrs["content"]
		} else {
			p.Music.Songs = append(p.Music.Songs, &song{URL: attrs["content"]})
		}
	case "music:song:disc":
		p.ensureSongs()
		t, err := strconv.ParseInt(attrs["content"], 10, 64)
		if err == nil {
			p.Music.Songs[len(p.Music.Songs)-1].Disc = t
		}
	case "music:song:track":
		p.ensureAudio()
		t, err := strconv.ParseInt(attrs["content"], 10, 64)
		if err == nil {
			p.Music.Songs[len(p.Music.Songs)-1].Track = t
		}
	}
}
