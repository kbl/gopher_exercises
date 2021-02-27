package main

import (
	"github.com/kbl/gopher_exercises/book/ch04/my_github"
	"log"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		log.Fatalf("Usage: %s <username> <repository>\n", os.Args[0])
	}
	userName := os.Args[1]
	repoName := os.Args[2]

	my_github.Browse(userName, repoName)
}
