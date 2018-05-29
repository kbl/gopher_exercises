package main

import (
	"book/ch04/omdb"
	"bytes"
	"fmt"
	"io"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: %s <api key> <movie title>")
		os.Exit(1)
	}

	apiToken := os.Args[1]
	movieTitle := os.Args[2]

	poster := omdb.DownloadPoster(apiToken, movieTitle)
	buffer := bytes.NewBuffer(poster)

	if poster != nil {
		io.Copy(os.Stdout, buffer)
	}
}
