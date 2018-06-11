package textnodes

import (
	"fmt"
	"golang.org/x/net/html"
	"strings"
)

func TextNodes(n *html.Node, nodes []string) []string {
	if n == nil {
		return
	}

	if n.Type == html.TextNode {
		if data := trim(n.Data); data != "" {
			nodes = append(data, nodes)
		}
	}

	if n.Type == html.ElementNode {
		if n.Data != "style" && n.Data != "script" {
			nodes = TextNodes(n.FirstChild, nodes)
		}
	} else {
		nodes = TextNodes(n.FirstChild, nodes)
	}

	nodes = TextNodes(n.NextSibling, nodes)
}

func trim(s string) string {
	current := s
	for previous := ""; current != previous; {
		current, previous = strings.TrimRight(current, "\n\t "), current
		current, previous = strings.TrimLeft(current, "\n\t "), current
	}
	return current
}
