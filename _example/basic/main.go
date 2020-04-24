package main

import (
	"fmt"

	parser "github.com/ammit/go-metaparser"
)

const (
	url = "http://ogp.me"
)

func main() {
	p := parser.New()
	b, err := p.FetchHTML(url)
	if err != nil {
		fmt.Printf("Could not fetch html from given url: %v \n", url)
	}
	defer b.Close()

	err = p.ParseHTML(b)

	fmt.Printf("The parsed title is: %v \n", p.Title)
}
