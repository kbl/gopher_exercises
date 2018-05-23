package main

import (
	"book/ch04/github"
	"book/ch04/vim"
	"fmt"
	"log"
	"os"
)

func main() {
	if len(os.Args) != 5 {
		log.Fatalf("Usage: %s <api token> <username> <repository> <action>", os.Args[0])
	}
	apiToken := os.Args[1]
	userName := os.Args[2]
	repoName := os.Args[3]
	actionName := os.Args[4]

	client := github.NewClient(
		&github.Repository{
			Name:  repoName,
			Owner: userName,
		},
		apiToken,
	)

	action, err := github.ToAction(actionName)

	if err != nil {
		log.Fatal(err)
	}

	if action == github.CREATE {
		// create.Create(vim.Prompt("<title>"), vim.Prompt("<body>"))
		// client.Create("<title>", "<body>")
		fmt.Println(client)
	}

	ni := github.NewIssue{Body: "bodyy", Title: "titlee"}
}
