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

type createGithubIssue struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}

type ReadGithubIssue struct {
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

func (c *Client) Read(issueId int) *ReadGithubIssue {
	url := fmt.Sprintf("%s/%d", c.endpoint, issueId)
	response := c.request("GET", url, nil)
	result := new(ReadGithubIssue)
	json.NewDecoder(response.Body).Decode(result)
	return result
}

func (c *Client) Create(title, content string) int {
	var body bytes.Buffer
	json.NewEncoder(&body).Encode(createGithubIssue{Body: content, Title: title})
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

func (c *Client) Update(id int) {
	issue := c.Read(id)
	title := vim.Prompt(issue.Title)
	body := vim.Prompt(issue.Body)
	fmt.Println(title, body)
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
