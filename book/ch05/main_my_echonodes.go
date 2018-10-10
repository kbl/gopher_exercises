package main

import (
	"book/ch05/textnodes"
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
	for _, node := range textnodes.TextNodes(doc, nil) {
		fmt.Println(node)
	}
}
