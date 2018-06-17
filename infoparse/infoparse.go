package main

import (
	"bytes"
	"fmt"
	"golang.org/x/net/html"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

var stageIds = map[string]int{
	"TECH":      0,
	"INSPIRE":   1,
	"MARKETING": 2,
	"STARTUP":   3,
	"LEADERS":   4,
	"WORKSHO":   5,
}

type Speaker struct {
	Id int
	DataURL,
	Name,
	Stage string
	StageId int
	ImageURL,
	Bio,
	Title,
	GithubURL,
	LinkedInURL,
	FacebookURL,
	TwitterURL string
}

var speakersMap map[int]*Speaker

func main() {
	speakersMap = make(map[int]*Speaker)

	doc, err := html.Parse(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	speakersNode := findSpeakersNode(doc)
	speakers := parseSpeakers(speakersNode)

	jobs := make(chan int)
	results := make(chan bool)

	for _, s := range speakers {
		speakersMap[s.Id] = s
	}

	go submitAllJobs(jobs)

	for w := 0; w < 30; w++ {
		go worker(w, jobs, results)
	}

	for _, _ = range speakersMap {
		_ = <-results
	}

	for _, s := range speakersMap {
		fmt.Printf(
			"INSERT INTO speaker (id, infoshareid, category, description, facebookprofile, githubprofile, linkedinprofile, twitterprofile, name) VALUES(nextval('speaker_seq'), %d, %d, %s, %s, %s, %s, %s, %s);\n",
			s.Id,
			s.StageId,
			nullValue(s.Bio),
			nullValue(s.FacebookURL),
			nullValue(s.GithubURL),
			nullValue(s.LinkedInURL),
			nullValue(s.TwitterURL),
			nullValue(s.Name),
		)
	}
}

func nullValue(v string) string {
	v = strings.Trim(v, " \n\r\t")
	if v == "" {
		return "NULL"
	}
	return quote(v)
}

func quote(v string) string {
	return fmt.Sprintf("'%s'", strings.Trim(strings.Replace(v, "'", "\\'", -1), " \n\r\t"))
}

func submitAllJobs(jobs chan<- int) {
	for id, _ := range speakersMap {
		jobs <- id
	}
}

func worker(workerId int, jobs <-chan int, results chan<- bool) {
	for speakerId := range jobs {
		readBio(speakerId)
		results <- true
	}
}

func readBio(speakerId int) {
	url := fmt.Sprintf("https://infoshare.pl/speaker2.php?cid=48&id=%d&year=2018&agenda_id=99999&fancybox=true", speakerId)
	r, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer r.Body.Close()
	doc, err := html.Parse(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	fillBio(speakersMap[speakerId], doc)
	fillSocialURLs(speakersMap[speakerId], doc)
}

func fillBio(s *Speaker, doc *html.Node) {
	bioNode := findBioNode(doc)
	s.Bio = fetchText(bioNode, bytes.NewBufferString("")).String()
}

func findBioNode(n *html.Node) *html.Node {
	if n == nil {
		return nil
	}
	if n.Type == html.ElementNode && n.Data == "div" {
		if attr, ok := getAttr(n, "class"); ok && attr == "user-profil-desc" {
			n = n.FirstChild
			foundNestedNode := false
			for {
				if foundNestedNode {
					if n == nil || n.Type == html.ElementNode {
						break
					}
				}
				if attr, ok = getAttr(n, "class"); ok && attr == "clear visible-xs" {
					foundNestedNode = true
				}
				n = n.NextSibling
			}
			return n
		}
	}
	if x := findBioNode(n.FirstChild); x != nil {
		return x
	}
	return findBioNode(n.NextSibling)
}

func fillSocialURLs(s *Speaker, doc *html.Node) {
	n := findSocialNode(doc).FirstChild
	for ; n != nil; n = n.NextSibling {
		if n.Type != html.ElementNode || n.Data != "a" {
			continue
		}
		if attr, ok := getAttr(n, "title"); ok {
			url, _ := getAttr(n, "href")
			if attr == "Facebook" {
				s.FacebookURL = url
			} else if attr == "Twitter" {
				s.TwitterURL = url
			} else if attr == "LinkedIn" {
				s.LinkedInURL = url
			} else if attr == "Github" {
				s.GithubURL = url
			} else {
				log.Fatalf("Unknown type %q", attr)
			}
		}
	}
}

func findSocialNode(n *html.Node) *html.Node {
	if n == nil {
		return nil
	}
	if attr, ok := getAttr(n, "class"); ok && attr == "no-wrap" {
		return n
	}
	if x := findSocialNode(n.FirstChild); x != nil {
		return x
	}
	return findSocialNode(n.NextSibling)
}

func fetchText(n *html.Node, buffer *bytes.Buffer) *bytes.Buffer {
	if n == nil {
		return buffer
	}
	if n.Type == html.TextNode {
		content := strings.Trim(n.Data, " \n\r\t")
		buffer.WriteString(content)
	}
	buffer = fetchText(n.FirstChild, buffer)
	return fetchText(n.NextSibling, buffer)
}

func findSpeakersNode(n *html.Node) *html.Node {
	if n == nil {
		return nil
	}
	if n.Type == html.ElementNode && n.Data == "section" {
		containerNode := n.FirstChild
		for {
			if containerNode.Type == html.ElementNode && containerNode.Data == "div" {
				break
			}
			containerNode = containerNode.NextSibling
		}
		flexContainerNode := containerNode.FirstChild
		for {
			if flexContainerNode.Type == html.ElementNode && flexContainerNode.Data == "div" {
				break
			}
			flexContainerNode = flexContainerNode.NextSibling
		}
		return flexContainerNode
	}
	if x := findSpeakersNode(n.FirstChild); x != nil {
		return x
	}
	return findSpeakersNode(n.NextSibling)
}

func parseSpeakers(n *html.Node) []*Speaker {
	var speakers []*Speaker
	for nestedNode := n.FirstChild; nestedNode != nil; nestedNode = nestedNode.NextSibling {
		if _, hasAttr := getAttr(nestedNode, "data-fancybox"); nestedNode.Type == html.ElementNode && hasAttr {
			speakers = append(speakers, parseSpeaker(nestedNode))
		}
	}

	return speakers
}

func getAttr(n *html.Node, attrName string) (string, bool) {
	for _, a := range n.Attr {
		if a.Key == attrName {
			return a.Val, true
		}
	}
	return "", false
}

func parseSpeaker(n *html.Node) *Speaker {
	s := new(Speaker)
	v, _ := getAttr(n, "data-fancybox")
	id, err := strconv.Atoi(strings.Split(v, "speaker")[1])
	if err != nil {
		log.Fatal(err)
	}
	s.Id = id
	v, _ = getAttr(n, "data-src")
	s.DataURL = v

	for childNode := n.FirstChild; childNode != nil; childNode = childNode.NextSibling {
		if childNode.Type != html.ElementNode {
			continue
		}
		if childNode.Data == "span" {
			s.Stage = childNode.FirstChild.Data
			s.StageId = stageIds[s.Stage]
		} else if childNode.Data == "img" {
			url, ok := getAttr(childNode, "src")
			if !ok {
				log.Fatalf("Can't find src attribute in img node")
			}
			s.ImageURL = url
		} else if childNode.Data == "div" {
			for titleNode := childNode.FirstChild; titleNode != nil; titleNode = titleNode.NextSibling {
				if titleNode.Type != html.ElementNode {
					continue
				}
				if titleNode.Data == "h4" {
					s.Name = titleNode.FirstChild.Data
				} else if titleNode.Data == "span" {
					s.Title = titleNode.FirstChild.Data
				}
			}
		}
	}
	return s
}
