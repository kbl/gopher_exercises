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
	for k, v := range count.Count(doc) {
		fmt.Println(k, v)
	}
}
