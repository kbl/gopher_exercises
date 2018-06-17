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

type Speaker struct {
	Id int
	DataURL,
	Name,
	Stage,
	ImageURL,
	Bio,
	Title string
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
		speakersMap[s.Id] = &s
	}

	go submitAllJobs(jobs)

	for w := 0; w < 30; w++ {
		go worker(w, jobs, results)
	}

	for _, _ = range speakersMap {
		_ = <-results
	}

	for _, v := range speakersMap {
		fmt.Println(v)
	}
}

func submitAllJobs(jobs chan<- int) {
	for id, _ := range speakersMap {
		jobs <- id
	}
}

func worker(workerId int, jobs <-chan int, results chan<- bool) {
	fmt.Println("Started worker ", workerId)
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

	bioNode := findBioNode(speakerId, doc)
	s := speakersMap[speakerId]
	s.Bio = fetchText(bioNode, bytes.NewBufferString("")).String()
}

func findBioNode(speakerId int, n *html.Node) *html.Node {
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
	if x := findBioNode(speakerId, n.FirstChild); x != nil {
		return x
	}
	return findBioNode(speakerId, n.NextSibling)
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

func parseSpeakers(n *html.Node) []Speaker {
	var speakers []Speaker
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

func parseSpeaker(n *html.Node) Speaker {
	s := Speaker{}
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
