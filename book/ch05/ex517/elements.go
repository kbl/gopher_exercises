package ex517

import (
	"github.com/kbl/gopher_exercises/book/ch05/outline2"
	"golang.org/x/net/html"
)

func ElementsByTagName(doc *html.Node, names ...string) []*html.Node {
	var nodes []*html.Node

	findNodes := func(n *html.Node) {
		if n.Type == html.ElementNode && contains(n.Data, names) {
			nodes = append(nodes, n)
		}
	}

	outline2.ForEachNode(doc, findNodes, nil)

	return nodes
}

func contains(what string, where []string) bool {
	for _, e := range where {
		if e == what {
			return true
		}
	}
	return false
}
