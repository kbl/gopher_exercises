package count

import (
	"golang.org/x/net/html"
)

func Count(node *html.Node) map[string]int {
	m := make(map[string]int)
	return count(m, node)
}

func count(m map[string]int, node *html.Node) map[string]int {
	if node.Type == html.ElementNode {
		m[node.Data]++
	}
	for c := node.FirstChild; c != nil; c = c.NextSibling {
		m = count(m, c)
	}
	return m
}
