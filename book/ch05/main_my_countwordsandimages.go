package main

import (
	"fmt"
	"github.com/kbl/gopher_exercises/book/ch05/my_count"
	"golang.org/x/net/html"
	"log"
	"os"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(my_count.CountWordsAndImages(doc))
}
