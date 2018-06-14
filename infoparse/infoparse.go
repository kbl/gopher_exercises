package main

import (
	"fmt"
	"golang.org/x/net/html"
	"log"
	"os"
	"strconv"
	"strings"
)

type Speaker struct {
	Id int
	Name,
	Stage,
	ImageURL,
	Title string
}

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	speakersNode := findSpeakersNode(doc)
	speakers := parseSpeakers(speakersNode)

	for _, s := range speakers {
		fmt.Println(s)
	}
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
