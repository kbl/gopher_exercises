package main

import (
	"fmt"
	"github.com/kbl/gopher_exercises/book/ch04/my_xkcd"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Please provide search term!")
		os.Exit(1)
	}
	searchTerm := os.Args[1]
	archivePath := "./archive"
	my_xkcd.ArchiveTo(archivePath)
	for _, comic := range my_xkcd.Search(archivePath, searchTerm) {
		fmt.Println(comic.URL())
		fmt.Println(comic.Title)
		fmt.Println(comic.Transcript)
		fmt.Println()
	}
}
