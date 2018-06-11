package github

import (
	"encoding/json"
	//"fmt"
	"html/template"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

var issues []GithubIssue
var milestones map[string]GithubMilestone = make(map[string]GithubMilestone)
var users map[string]GithubUser = make(map[string]GithubUser)

func Browse() {
	prepareData()

	http.HandleFunc("/user/", browseUsers)
	http.HandleFunc("/milestone/", browseMilestones)
	http.HandleFunc("/", allIssues)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

func prepareData() {
	issuesFile, err := os.Open("issues.json")
	if err != nil {
		log.Fatal(err)
	}
	json.NewDecoder(issuesFile).Decode(&issues)

	for _, issue := range issues {
		m := issue.Milestone
		if m.Title != "" {
			milestones[m.Title] = m
			users[m.Creator.Login] = m.Creator
		}
		users[issue.User.Login] = issue.User
		for _, a := range issue.Assignees {
			users[a.Login] = a
		}
	}
}

func browseMilestones(w http.ResponseWriter, r *http.Request) {
	milestone, err := url.PathUnescape(strings.SplitN(r.URL.Path, "/", 3)[2])
	if err != nil {
		log.Fatal(err)
	}
	t := template.Must(template.New("milestone view").Parse(`<html>
	<body>
		<h1>Milestone {{ .Title }}</h1>
		<ul>
			<li>description: {{ .Description }}</li>
			<li>number: {{ .Number }}</li>
			<li>open issues: {{ .OpenIssues }}</li>
			<li>state: {{ .State }}</li>
			<li>creator <a href="/user/{{ .Creator.Login }}">{{ .Creator.Login }}</a></li>
			<li>created at {{ .CreatedAt }}</li>
		</ul>
		<p><a href="/">back to the issues list</a></p>
	</body>
</html>`))
	t.Execute(w, milestones[milestone])
}

func browseUsers(w http.ResponseWriter, r *http.Request) {
	username, err := url.PathUnescape(strings.SplitN(r.URL.Path, "/", 3)[2])
	if err != nil {
		log.Fatal(err)
	}
	t := template.Must(template.New("user view").Parse(`<html>
	<body>
		<h1>User {{ .Login }}</h1>
		<img src="{{ .AvatarURL }}" />
		<ul>
			<li><a href="{{ .FollowersURL }}">followers</a></li>
			<li><a href="{{ .FollowingURL }}">following</a></li>
			<li><a href="{{ .GistsURL }}">gists</a></li>
			<li><a href="{{ .ReposURL }}">repos</a></li>
		</ul>
		<p><a href="/">back to the issues list</a></p>
	</body>
</html>`))
	t.Execute(w, users[username])
}

func allIssues(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.New("main page").Parse(`<html>
	<body>
		<h1>Issues</h1>
		<table>
			<tr>
				<th>Id</th>
				<th>State</th>
				<th>Title</th>
				<th>User</th>
				<th>Assignees</th>
				<th>Milestone</th>
				<th>CreatedAt</th>
			</tr>
		{{ range . }}
			<tr>
				<td>{{ .Id }}</td>
				<td>{{ .State }}</td>
				<td>{{ .Title }}</td>
				<td><a href="/user/{{ .User.Login }}">{{ .User.Login }}</a></td>
				<td>
				{{ range .Assignees }}
				  <a href="/user/{{ .Login }}">{{ .Login }}</a>&nbsp;
			    {{ end }}
				</td>
				<td><a href="/milestone/{{ .Milestone.Title }}">{{ .Milestone.Title }}</a></td>
				<td>{{ .CreatedAt }}</td>
			</tr>
		{{ end }}
		</table>
	</body>
</html>`))
	t.Execute(w, issues)
}
