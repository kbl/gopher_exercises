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
	"time"
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

type GithubMilestone struct {
	URL          string     `json:"url"`
	HTMLURL      string     `json:"html_url"`
	LabelsURL    string     `json:"labels_url"`
	Id           int        `json:"id"`
	NodeId       string     `json:"node_id"`
	Number       int        `json:"number"`
	Title        string     `json:"title"`
	Description  string     `json:"description"`
	Creator      GithubUser `json:"creator"`
	OpenIssues   int        `json:"open_issues"`
	ClosedIssues int        `json:"closed_issues"`
	State        string     `json:"state"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at"`
	DueOn        time.Time  `json:"due_on"`
	ClosedAt     time.Time  `json:"closed_at"`
}

type GithubUser struct {
	Login             string `json:"login"`
	Id                int    `json:"id"`
	NodeId            string `json:"node_id"`
	AvatarURL         string `json:"avatar_url"`
	GravatarId        string `json:"gravatar_id"`
	URL               string `json:"url"`
	HTMLURL           string `json:"html_url"`
	FollowersURL      string `json:"followers_url"`
	FollowingURL      string `json:"following_url"`
	GistsURL          string `json:"gists_url"`
	StarredURL        string `json:"starred_url"`
	SubscriptionsURL  string `json:"subscriptions_url"`
	OrganizationsURL  string `json:"organizations_url"`
	ReposURL          string `json:"repos_url"`
	EventsURL         string `json:"events_url"`
	ReceivedEventsURL string `json:"received_events_url"`
	Type              string `json:"type"`
	SiteAdmin         bool   `json:"site_admin"`
}

type GithubLabel struct {
	Id      int    `json:"id"`
	NodeId  string `json:"node_id"`
	URL     string `json:"url"`
	Name    string `json:"name"`
	Color   string `json:"color"`
	Default bool   `json:"default"`
}

type GithubIssue struct {
	Title             string          `json:"title"`
	Body              string          `json:"body"`
	State             string          `json:"state"`
	Url               string          `json:"url"`
	RepositoryUrl     string          `json:"repository_url"`
	LabelsURL         string          `json:"labels_url"`
	CommentsURL       string          `json:"comments_url"`
	EventsURL         string          `json:"events_url"`
	HTMLURL           string          `json:"html_url"`
	Id                int             `json:"id"`
	NodeId            string          `string:"node_id"`
	Number            int             `json:"number"`
	User              GithubUser      `json:"user"`
	Labels            []GithubLabel   `json:"labels"`
	Locked            bool            `json:"locked"`
	Asignee           GithubUser      `json:"assignee"`
	Assignees         []GithubUser    `json:"assignees"`
	Milestone         GithubMilestone `json:"milestone"`
	Comments          int             `json:"comments"`
	CreatedAt         time.Time       `json:"created_at"`
	UpdatedAt         time.Time       `json:"updated_at"`
	ClosedAt          time.Time       `json:"closed_at"`
	AuthorAssociation string          `json:"author_association"`
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
	defer response.Body.Close()
	result := new(GithubIssue)
	json.NewDecoder(response.Body).Decode(result)
	return result
}

func (c *Client) Create(title, description string) int {
	var body bytes.Buffer
	json.NewEncoder(&body).Encode(GithubIssue{Body: description, Title: title})
	response := c.request("POST", c.endpoint, &body)
	defer response.Body.Close()
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
	defer response.Body.Close()

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
	defer response.Body.Close()

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
