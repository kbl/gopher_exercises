package github

import (
	"errors"
	"fmt"
	"io"
	"log"
	// "net/http"
	"book/ch04/vim"
)

const graphQLEndpoint = "https://api.github.com/graphql"

type Action int

const (
	CREATE Action = iota
	READ
	UPDATE
	DELETE
)

func ToAction(name string) (Action, error) {
	if name == "update" {
		return UPDATE, nil
	} else if name == "delete" {
		return DELETE, nil
	} else if name == "create" {
		return CREATE, nil
	} else if name == "read" {
		return READ, nil
	}
	return -1, errors.New(fmt.Sprintf("Unknown action %s.", name))
}

type GraphQLClient struct {
	token string
}

func (c *GraphQLClient) Create() {
}

// type Issue struct {
// 	Title     string
// 	Body      string
// 	Assignees []string
// 	Milestone int
// 	Labels    []string
// }

func main() {
	owner := "kbl"
	repo := "gopher_exercises"
	url := fmt.Sprintf(issuesAPIURL, owner, repo)
	var requestBody io.Reader

	body := prompt("<type issue body here>")
	title := prompt("<type issue title here>")

	fmt.Println(url)
	fmt.Println(requestBody, body, title)
	/* response, err := http.Post(url, "application/json", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(response)
	*/
}

func prompt(message string) string {
	content, err := vim.Edit(message)
	if err != nil {
		log.Fatal(err)
	}
	return content
}

/*
POST /repos/:owner/:repo/issues
{
  "title": "Found a bug",
  "body": "I'm having a problem with this.",
  "assignees": [
    "octocat"
  ],
  "milestone": 1,
  "labels": [
    "bug"
  ]
}
*/
