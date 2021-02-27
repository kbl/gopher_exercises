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
	for k, v := range my_count.Count(doc) {
		fmt.Println(k, v)
	}
}
