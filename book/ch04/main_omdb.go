package main

import (
	"bytes"
	"fmt"
	"github.com/kbl/gopher_exercises/book/ch04/my_omdb"
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

	poster := my_omdb.DownloadPoster(apiToken, movieTitle)
	buffer := bytes.NewBuffer(poster)

	if poster != nil {
		io.Copy(os.Stdout, buffer)
	}
}
