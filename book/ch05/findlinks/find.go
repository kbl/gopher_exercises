package findlinks

import (
	"fmt"
	"golang.org/x/net/html"
	"os"
)

func Find() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}

	for _, link := range visit(nil, doc) {
		fmt.Println(link)
	}
}

func visit(links []string, n *html.Node) []string {
	if n == nil {
		return links
	}

	if n.Type == html.ElementNode {
		if n.Data == "a" || n.Data == "link" || n.Data == "script" || n.Data == "img" {
			for _, a := range n.Attr {
				if a.Key == "href" || a.Key == "src" {
					links = append(links, a.Val)
				}
			}
		}
	}
	links = visit(links, n.NextSibling)
	return visit(links, n.FirstChild)
}
