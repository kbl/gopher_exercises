package textnodes

import (
	"golang.org/x/net/html"
	"strings"
)

func TextNodes(n *html.Node, nodes []string) []string {
	if n == nil {
		return nodes
	}

	if n.Type == html.TextNode {
		if data := trim(n.Data); data != "" {
			nodes = append(nodes, data)
		}
	}

	if n.Type == html.ElementNode {
		if n.Data != "style" && n.Data != "script" {
			nodes = TextNodes(n.FirstChild, nodes)
		}
	} else {
		nodes = TextNodes(n.FirstChild, nodes)
	}

	return TextNodes(n.NextSibling, nodes)
}

func trim(s string) string {
	current := s
	for previous := ""; current != previous; {
		current, previous = strings.TrimRight(current, "\n\t "), current
		current, previous = strings.TrimLeft(current, "\n\t "), current
	}
	return current
}
