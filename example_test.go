package ogp_test

import (
	"fmt"
	"log"
	"net/http"

	"github.com/shota3506/go-ogp"
)

func ExampleParse() {
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
