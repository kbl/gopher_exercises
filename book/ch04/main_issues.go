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
		id := client.Create(vim.Prompt("<title>"), vim.Prompt("<body>"))
		fmt.Printf("New issue with id %d was created.\n", id)
	} else if action == github.READ {
		fmt.Print("Podaj id: ")
		var issueId int
		_, err := fmt.Scanf("%d", &issueId)
		if err != nil {
			log.Fatal(err)
		}
		client.Read(issueId)
	}
}
