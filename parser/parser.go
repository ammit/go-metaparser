package parser

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

// Parser ...
type Parser struct {
	Title            string   `json:"title"`
	Type             string   `json:"type"`
	Description      string   `json:"description"`
	Determiner       string   `json:"determiner"`
	URL              string   `json:"url"`
	Locale           string   `json:"locale"`
	LocalesAlternate []string `json:"locales_alternate"`
	SiteName         string   `json:"site_name"`

	Images []*Image `json:"images"`
	Videos []*Video `json:"videos"`
	Audios []*Audio `json:"audios"`

	// New
	Music   Music   `json:"music"`
	Article Article `json:"article"`
}

// New ...
func New(url string) *Parser {
	return &Parser{
		URL: url,
	}
}

// FetchHTML returns buffer
func (p *Parser) FetchHTML() (buffer io.Reader, err error) {
	u := strings.TrimSpace(p.URL)
	if !strings.HasPrefix(u, "http:") && !strings.HasPrefix(u, "https:") {
		u = "http://" + u
	}

	client := &http.Client{Timeout: time.Second * 10}

	req, err := http.NewRequest("GET", u, nil)
	req.Header.Add("Accept", `text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8`)
	req.Header.Add("User-Agent", `Mozilla/5.0 (Macintosh; Intel Mac OS X 10_7_5) AppleWebKit/537.11 (KHTML, like Gecko) Chrome/23.0.1271.64 Safari/537.11`)
	r, err := client.Do(req)

	if err != nil {
		fmt.Printf("%s", err)
		return nil, err
	}

	if !(r.StatusCode >= 200 && r.StatusCode < 300) {
		return nil, errors.New("page not found")
	}

	rd := io.Reader(r.Body)

	return rd, err
}

// ParseHTML parses given html
func (p *Parser) ParseHTML(buffer io.Reader) error {
	z := html.NewTokenizer(buffer)
	for {
		token := z.Next()
		switch token {
		case html.ErrorToken:
			if z.Err() == io.EOF {
				return nil
			}
			return z.Err()
		case html.StartTagToken, html.SelfClosingTagToken, html.EndTagToken:
			name, hasAttr := z.TagName()
			if atom.Lookup(name) == atom.Body {
				return nil
			}
			if atom.Lookup(name) != atom.Meta || !hasAttr {
				continue
			}
			m := make(map[string]string)
			var key, val []byte
			for hasAttr {
				key, val, hasAttr = z.TagAttr()
				m[atom.String(key)] = string(val)
			}
			p.ParseMeta(m)
		}
	}
}

// ParseMeta processes meta attributes
func (p *Parser) ParseMeta(attrs map[string]string) {
	switch attrs["property"] {
	// opengraph:basic
	case "og:title", "og:type", "og:url", "og:description", "og:determiner", "og:locale", "og:locale:alternate", "og:site_name":
		p.parseBasicMeta(attrs)
	// opengraph:image
	case "og:image", "og:image:url", "og:image:secure_url", "og:image:type", "og:image:width", "og:image:height", "og:image:alt":
		p.parseImageMeta(attrs)
	// opengraph:video
	case "og:video", "og:video:url", "og:video:secure_url", "og:video:type", "og:video:width", "og:video:height",
		"video:actor", "video:actor:role", "video:director", "video:writer", "video:duration", "video:release_date", "video:tag", "video:series":
		p.parseVideoMeta(attrs)
	// opengraph:audio
	case "og:audio", "og:audio:secure_url", "og:audio:type":
		p.parseAudioMeta(attrs)
	// music
	case "music:musician", "music:album", "music:album:disc", "music:album:track", "music:song",
		"music:song:disc", "music:song:track", "music:release_date", "music:creator", "music:duration":
		p.parseMusicMeta(attrs)
	// article
	case "article:published_time", "article:modified_time", "article:expiration_time", "article:author",
		"article:section", "article:tag":
		p.parseArticleMeta(attrs)
	}
}
