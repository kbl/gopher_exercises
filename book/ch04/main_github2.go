package main

import (
	"book/ch04/github"
	"fmt"
	"log"
	"os"
)

func main() {
	if len(os.Args) != 5 {
		log.Fatalf("Usage: %s <api token> <username> <repository> <action>", os.Args[0])
	}
	apiToken := os.Args[1]
	username := os.Args[2]
	repository := os.Args[3]
	action := os.Args[4]

	x, err := github.ToAction(action)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(apiToken, username, repository, action, x)
}
