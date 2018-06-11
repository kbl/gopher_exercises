package count

import (
	"book/ch05/textnodes"
	"bufio"
	"bytes"
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

func CountWordsAndImages(node *html.Node) (words, images int) {
	for _, text := range textnodes.TextNodes(node, nil) {
		in := bufio.NewScanner(bytes.NewBufferString(text))
		in.Split(bufio.ScanWords)

		for in.Scan() {
			in.Text()
			words++
		}
	}
	images = Count(node)["img"]
	return
}
