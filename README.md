# go-ogp

[![Go Reference](https://pkg.go.dev/badge/github.com/shota3506/go-ogp.svg)](https://pkg.go.dev/github.com/shota3506/go-ogp)

The [Open Graph protocol](https://ogp.me/) parser for Go.

## Example

### Parse HTML

```go
package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/shota3506/go-ogp"
)

func main() {
	res, err := http.Get("https://ogp.me/")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	obj, err := ogp.Parse(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(obj.Title) // Open Graph protocol
}

```

### Render HTML meta tags

```go
package main

import (
	"os"

	"github.com/shota3506/go-ogp"
	"golang.org/x/net/html"
)

func main() {
	obj := &ogp.Object{
		Title: "Example",
		Type:  "website",
		Images: []*ogp.Image{
			{URL: "https://example.com/image.jpg"},
		},
		URL: "https://example.com",
	}

	head := &html.Node{Type: html.ElementNode, Data: "head"}
	for _, meta := range obj.HTML() {
		head.AppendChild(meta)
	}
	html.Render(os.Stdout, head)
}
```
