package main

import (
	"book/ch04/xkcd"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Please provide search term!")
		os.Exit(1)
	}
	searchTerm := os.Args[1]
	archivePath := "./archive"
	xkcd.ArchiveTo(archivePath)
	for _, comic := range xkcd.Search(archivePath, searchTerm) {
		fmt.Println(comic.URL())
		fmt.Println(comic.Title)
		fmt.Println(comic.Transcript)
		fmt.Println()
	}
}
