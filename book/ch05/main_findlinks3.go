package main

import (
	"github.com/kbl/gopher_exercises/book/ch05/findlinks3"
	"os"
)

func main() {
	findlinks3.BreadthFirst(findlinks3.Crawl, os.Args[1:])
}
