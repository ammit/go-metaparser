package parser

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"golang.org/x/net/html/charset"
)

const (
	maxResponseBodySize = 10485760 // 10MB
)

var (
	cssSelectors = strings.Join([]string{
		"link[rel='icon']",
		"link[rel='shortcut icon']",
		"link[rel='apple-touch-icon']",
		"link[rel='apple-touch-icon-precomposed']",
		"link[rel='ICON']",
		"link[rel='SHORTCUT ICON']",
		"link[rel='APPLE-TOUCH-ICON']",
		"link[rel='APPLE-TOUCH-ICON-PRECOMPOSED']",
	}, ", ")
)

// Parser ...
type Parser struct {
	URL string
}

// ParsedData ...
type ParsedData struct {
	BaseURL *url.URL
	Icons   []string
}

// New ...
func New(url string) *Parser {
	return &Parser{
		URL: url,
	}
}

// Fetch ...
func (p *Parser) Fetch() (*ParsedData, error) {
	u := strings.TrimSpace(p.URL)
	if !strings.HasPrefix(u, "http:") && !strings.HasPrefix(u, "https:") {
		u = "http://" + u
	}

	response, err := http.Get(u)

	if err != nil {
		fmt.Printf("%s", err)
		return nil, err
	}

	if !(response.StatusCode >= 200 && response.StatusCode < 300) {
		return nil, errors.New("page not found")
	}

	html, siteURL, err := getHTML(response)
	if err != nil {
		fmt.Printf("%s", err)
		return nil, err
	}

	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(html))
	if err != nil {
		fmt.Printf("%s", err)
		return nil, err
	}

	baseURL := extractBaseURL(siteURL, doc)
	icons := extractIcons(baseURL, doc)

	pd := ParsedData{
		BaseURL: baseURL,
		Icons:   icons,
	}

	return &pd, nil
}

func getHTML(r *http.Response) ([]byte, *url.URL, error) {
	limitReader := io.LimitReader(r.Body, maxResponseBodySize)
	b, err := ioutil.ReadAll(limitReader)
	r.Body.Close()

	if len(b) >= maxResponseBodySize {
		return nil, nil, errors.New("body too large")
	}

	if err != nil {
		fmt.Printf("%s", err)
		return nil, nil, err
	}

	if len(b) == 0 {
		return nil, nil, errors.New("empty response")
	}

	reader := bytes.NewReader(b)
	contentType := r.Header.Get("Content-Type")
	utf8reader, err := charset.NewReader(reader, contentType)
	if err != nil {
		return nil, nil, err
	}
	utf8bytes, err := ioutil.ReadAll(utf8reader)
	if err != nil {
		return nil, nil, err
	}

	return utf8bytes, r.Request.URL, nil
}

func absURL(baseURL *url.URL, path string) (string, error) {
	u, err := url.Parse(path)
	if err != nil {
		return "", err
	}

	u.Scheme = baseURL.Scheme
	if u.Scheme == "" {
		u.Scheme = "http"
	}

	if u.Host == "" {
		u.Host = baseURL.Host
	}
	return baseURL.ResolveReference(u).String(), nil
}

func extractBaseURL(siteURL *url.URL, doc *goquery.Document) *url.URL {
	href := ""
	doc.Find("head base[href]").First().Each(func(i int, s *goquery.Selection) {
		href, _ = s.Attr("href")
	})

	if href != "" {
		baseTagURL, err := url.Parse(href)
		if err != nil {
			return siteURL
		}
		return baseTagURL
	}

	return siteURL
}

func extractIcons(baseURL *url.URL, doc *goquery.Document) []string {
	var hits []string
	doc.Find(cssSelectors).Each(func(i int, s *goquery.Selection) {
		href, ok := s.Attr("href")
		if ok && href != "" {
			var err error

			href, err = absURL(baseURL, href)

			if err != nil {
				fmt.Printf("%s", err)
			} else {
				hits = append(hits, href)
			}
		}
	})

	return hits
}
