[![Build Status](https://travis-ci.com/ammit/go-metaparser.svg?branch=master)](https://travis-ci.com/ammit/go-metaparser)

# go-metaparser

> Lightweight metadata parser written in Go

![](https://repository-images.githubusercontent.com/256327777/b9713100-8055-11ea-8d30-ca9ecc6e881b)

With go-metaparser, you can easily extract structured meta-data from HTML. The purpose of this library is to be able to obtain all types of metadata from the web page.

Currently, it supports Open Graph, Twitter Card Metadata and some general metadata that doesn't belong to a particular type, for example - title, description etc.

## Installation

Install the package with:

```go
go get github.com/ammit/go-metaparser
```

Import it with:

```go
import "github.com/ammit/go-metaparser"
```

and use parser as the package name inside the code.

## Usage example

Please check the example folder for details.

```go
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

```

## Performance

You can run the benchmarks yourself, but here's the output on my machine:

```text
    BenchmarkParser-12         19024             62147 ns/op            9858 B/op        261 allocs/op
```

## Contributing

1. Fork it (<https://github.com/ammit/go-metaparser/fork>)
2. Create your feature branch (`git checkout -b feature/fooBar`)
3. Commit your changes (`git commit -am 'Add some fooBar'`)
4. Push to the branch (`git push origin feature/fooBar`)
5. Create a new Pull Request :)
