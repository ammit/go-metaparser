package main

import (
	"fmt"

	"github.com/ammit/go-metaparser/parser"
)

func main() {
	p := parser.New("http://css-tricks.com")

	data, err := p.Fetch()

	if err != nil {
		fmt.Printf("Error fetching url %-v", err)
	}

	fmt.Printf("%+v\n", data)
}
