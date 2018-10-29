package main

import (
	"book/ch05/ex513"
	"log"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatalf("usage: %s <url to download>", os.Args[0])
	}
	ex513.Crawl(os.Args[1])
}
