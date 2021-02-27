package main

import (
	"fmt"
	"github.com/kbl/gopher_exercises/book/ch04/my_github"
	"github.com/kbl/gopher_exercises/book/ch04/my_vim"
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

	client := my_github.NewClient(
		&my_github.Repository{
			Name:  repoName,
			Owner: userName,
		},
		apiToken,
	)

	action, err := my_github.ToAction(actionName)

	if err != nil {
		log.Fatal(err)
	}

	if action == my_github.CREATE {
		id := client.Create(my_vim.Prompt("<title>"), my_vim.Prompt("<body>"))
		fmt.Printf("New issue with id %d was created.\n", id)
	} else if action == my_github.READ {
		issueId := promptInt("Podaj id: ")
		fmt.Println(*client.Read(issueId))
	} else if action == my_github.EDIT {
		issueId := promptInt("Podaj id: ")
		client.Edit(issueId)
	} else if action == my_github.CLOSE {
		issueId := promptInt("Podaj id: ")
		client.Close(issueId)
	}
}

func promptInt(prompt string) int {
	fmt.Print("Podaj id: ")
	var number int
	_, err := fmt.Scanf("%d", &number)
	if err != nil {
		log.Fatal(err)
	}
	return number
}
