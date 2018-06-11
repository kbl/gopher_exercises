package main

import (
	"book/ch05/count"
	"fmt"
	"golang.org/x/net/html"
	"log"
	"os"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(count.CountWordsAndImages(doc))
}
