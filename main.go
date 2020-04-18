package main

import (
	"fmt"
	"log"

	og "github.com/ammit/go-metaparser/parser"
)

func main() {
	og := og.New("ogp.me")

	b, err := og.FetchHTML()
	if err != nil {
		log.Fatalf("%-v", err)
	}

	err = og.ParseHTML(b)
	if err != nil {
		log.Fatalf("%-v", err)
	}

	fmt.Printf("%+v\n", og)
}
