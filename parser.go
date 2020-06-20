package parser

import (
	"errors"
	"io"
	"net/http"
	"strings"
	"time"

	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

const (
	httpClientTimeoutSeconds = 30
)

// Parser ...
type Parser struct {
	Result
}

// New ...
func New() *Parser {
	return &Parser{}
}

// FetchHTML returns buffer
func (p *Parser) FetchHTML(target string) (io.ReadCloser, error) {
	target = strings.TrimSpace(target)

	return fetch(target)
}

func fetch(target string) (io.ReadCloser, error) {
	var netClient = &http.Client{
		Timeout: time.Second * httpClientTimeoutSeconds,
	}
	resp, err := netClient.Get(target)
	if err != nil {
		return nil, err
	}

	if !(resp.StatusCode >= 200 && resp.StatusCode < 300) {
		return nil, errors.New("page not found")
	}

	return resp.Body, nil
}

// ParseHTMLWithResult parses given html and returns a Result
func (p *Parser) ParseHTMLWithResult(buffer io.ReadCloser) (*Result, error) {
	err := p.ParseHTML(buffer)
	if err == nil {
		result := &Result{
			Title:       p.Title,
			Description: p.Description,
			OpenGraph:   p.OpenGraph,

			Images: p.Images,
			Videos: p.Videos,
			Audios: p.Audios,

			Music:   p.Music,
			Article: p.Article,
			Book:    p.Book,
			Profile: p.Profile,

			Favicons: p.Favicons,

			Twitter: p.Twitter,
		}
		return result, nil
	} else {
		return nil, err
	}
}

// ParseHTML parses given html
func (p *Parser) ParseHTML(buffer io.ReadCloser) error {
	defer buffer.Close()

	z := html.NewTokenizer(buffer)
	extractTitle := false
	for {
		token := z.Next()
		switch token {
		case html.ErrorToken:
			if z.Err() == io.EOF {
				return nil
			}
			return z.Err()
		case html.TextToken:
			// Text cannot be extract from the tag so it must extracted here
			// extractTitle acts as a flag which is set true on opening title tag and false on closing title tag
			if extractTitle {
				p.Title = string(z.Text())
			}
		case html.StartTagToken, html.SelfClosingTagToken, html.EndTagToken:
			name, hasAttr := z.TagName()
			if atom.Lookup(name) == atom.Body {
				return nil
			}
			if atom.Lookup(name) == atom.Title {
				// Toggle title parsing
				extractTitle = !extractTitle
			} else if hasAttr {
				attrs := getAttributes(z)
				if atom.Lookup(name) == atom.Meta {
					// Parse HTML meta tag
					if _, ok := attrs["property"]; ok {
						// tag with <meta property="..." content="..." ...>
						p.ParseMetaProperty(attrs)
					} else if name, ok := attrs["name"]; ok {
						// Description meta tag
						if name == "description" {
							p.Description = attrs["content"]
						}
					}
				} else if atom.Lookup(name) == atom.Link {
					p.ParseLink(attrs)
				} else {
					continue
				}
			} else {
				continue
			}
		}
	}
}

func getAttributes(z *html.Tokenizer) map[string]string {
	m := make(map[string]string)
	var key, val []byte
	// Must be true because of the previous if
	hasAttr := true
	for hasAttr {
		key, val, hasAttr = z.TagAttr()
		m[atom.String(key)] = string(val)
	}
	return m
}

// ParseMetaProperty processes meta attributes
func (p *Parser) ParseMetaProperty(attrs map[string]string) {
	switch attrs["property"] {
	// opengraph:basic
	case "og:title", "og:type", "og:url", "og:description", "og:determiner", "og:locale", "og:locale:alternate", "og:site_name":
		p.parseBasicOGMeta(attrs)
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
	// book
	case "book:author", "book:isbn", "book:release_date", "book:tag":
		p.parseBookMeta(attrs)
	// profile
	case "profile:first_name", "profile:last_name", "profile:username", "profile:gender":
		p.parseProfileMeta(attrs)
	// twitter
	case "twitter:card", "twitter:site", "twitter:site:id", "twitter:creator", "twitter:creator:id",
		"twitter:description", "twitter:title", "twitter:image", "twitter:image:alt", "twitter:player",
		"twitter:player:height", "twitter:player:width", "twitter:player:stream", "twitter:app:name:iphone",
		"twitter:app:url:iphone", "twitter:app:id:iphone", "twitter:app:name:ipad", "twitter:app:url:ipad",
		"twitter:app:id:ipad", "twitter:app:name:googleplay", "twitter:app:url:googleplay", "twitter:app:id:googleplay":
		p.parseTwitterMeta(attrs)
	}
}

func (p *Parser) ParseLink(attrs map[string]string) {
	// Not using a switch/case because there is too much variation in naming the favicon
	// but it often includes 'icon' in the rel attribute
	if strings.Contains(attrs["rel"], "icon") {
		p.parseFaviconLink(attrs)
	}
}
