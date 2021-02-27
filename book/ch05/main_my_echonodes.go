package main

import (
	"fmt"
	"github.com/kbl/gopher_exercises/book/ch05/my_textnodes"
	"golang.org/x/net/html"
	"log"
	"os"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
	for _, node := range my_textnodes.TextNodes(doc, nil) {
		fmt.Println(node)
	}
}
