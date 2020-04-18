package main

import (
	"fmt"
	"log"

	og "github.com/ammit/go-metaparser/parser"
)

func main() {
	og := og.New()

	b, err := og.FetchHTML("http://ogp.me")

	if err != nil {
		log.Fatalf("%-v", err)
	}

	err = og.ParseHTML(b)
	if err != nil {
		log.Fatalf("%-v", err)
	}

	fmt.Printf("%+v\n", og)
}
