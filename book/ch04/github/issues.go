package github

import (
	"book/ch04/vim"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
)

const (
	baseURLTemplate = "https://api.github.com/repos/%s/%s/issues"
	acceptHeader    = "application/vnd.github.v3+json"
)

type Action int

const (
	CREATE Action = iota
	READ
	EDIT
	CLOSE
)

func ToAction(name string) (Action, error) {
	if name == "edit" {
		return EDIT, nil
	} else if name == "close" {
		return CLOSE, nil
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

type GithubIssue struct {
	Title string `json:"title"`
	Body  string `json:"body"`
	State string `json:"state"`
}

type Client struct {
	endpoint string
	token    string
	client   *http.Client
}

func NewClient(repository *Repository, token string) *Client {
	return &Client{
		endpoint: fmt.Sprintf(baseURLTemplate, repository.Owner, repository.Name),
		token:    token,
		client:   &http.Client{},
	}
}

func (c *Client) Read(issueId int) *GithubIssue {
	url := fmt.Sprintf("%s/%d", c.endpoint, issueId)
	response := c.request("GET", url, nil)
	result := new(GithubIssue)
	json.NewDecoder(response.Body).Decode(result)
	return result
}

func (c *Client) Create(title, description string) int {
	var body bytes.Buffer
	json.NewEncoder(&body).Encode(GithubIssue{Body: description, Title: title})
	response := c.request("POST", c.endpoint, &body)
	locationURL, err := response.Location()
	if err != nil {
		log.Fatal(err)
	}
	location := strings.Split(locationURL.String(), "/")
	id, err := strconv.Atoi(location[len(location)-1])
	if err != nil {
		log.Fatal(err)
	}
	return id
}

func (c *Client) Edit(issueId int) {
	issue := c.Read(issueId)
	issue.Title = vim.Prompt(issue.Title)
	issue.Body = vim.Prompt(issue.Body)

	var body bytes.Buffer
	url := fmt.Sprintf("%s/%d", c.endpoint, issueId)
	json.NewEncoder(&body).Encode(issue)

	response := c.request("PATCH", url, &body)

	if response.StatusCode != http.StatusOK {
		log.Fatalf("something went wrong during editing an issue %v", response)
	}
}

func (c *Client) Close(issueId int) {
	issue := c.Read(issueId)
	issue.State = "closed"

	var body bytes.Buffer
	url := fmt.Sprintf("%s/%d", c.endpoint, issueId)
	json.NewEncoder(&body).Encode(issue)

	response := c.request("PATCH", url, &body)

	if response.StatusCode != http.StatusOK {
		log.Fatalf("something went wrong during closing an issue %v", response)
	}
}

func (c *Client) request(requestType, url string, requestBody io.Reader) *http.Response {
	request, err := http.NewRequest(requestType, url, requestBody)
	if err != nil {
		log.Fatal(err)
	}
	request.Header.Add("Accept", acceptHeader)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Authorization", fmt.Sprintf("token %s", c.token))

	response, err := c.client.Do(request)
	if err != nil {
		log.Fatal(err)
	}
	return response
}
