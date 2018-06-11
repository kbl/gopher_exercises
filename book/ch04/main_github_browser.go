package main

import (
	"book/ch04/github"
	"log"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		log.Fatalf("Usage: %s <username> <repository>\n", os.Args[0])
	}
	userName := os.Args[1]
	repoName := os.Args[2]

	github.Browse(userName, repoName)
}
