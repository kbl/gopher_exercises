package github

import (
	// "encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
)

const endpointTemplate = "https://api.github.com/repos/%s/%s/issues"

const acceptHeader = "application/vnd.github.v3+json"

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

type Repository struct {
	Name, Owner string
}

type NewIssue struct {
	Body, Title string
}

type Client struct {
	endpoint string
	token    string
	client   *http.Client
}

func NewClient(repository *Repository, token string) *Client {
	return &Client{
		endpoint: fmt.Sprintf(endpointTemplate, repository.Owner, repository.Name),
		token:    token,
		client:   &http.Client{},
	}
}

func (c *Client) Create(title, content string) {
	var body io.Reader
	response := c.post(body)
	fmt.Println(response)
}

func (c *Client) post(body io.Reader) *http.Response {
	fmt.Println(c.endpoint)
	request, err := http.NewRequest("POST", c.endpoint, body)
	if err != nil {
		log.Fatal(err)
	}
	request.Header.Add("Accept", acceptHeader)
	request.Header.Add("Authorization", c.token)

	response, err := c.client.Do(request)
	if err != nil {
		log.Fatal(err)
	}
	return response
}

// type Issue struct {
// 	Title     string
// 	Body      string
// 	Assignees []string
// 	Milestone int
// 	Labels    []string
// }

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
