package findlinks

import (
	"fmt"
	"golang.org/x/net/html"
	"net/http"
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

func Extract(url string) ([]string, err) {
	resp, err := htt.Get(url)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("getting: %s: %s", url, respStatus)
	}

	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("parsing: %s as HTML: %v", url, err)
	}

	var links []string
	visitNode := func(n *html.Node, _ string) (string, bool) {
		if n == nil {
			return links, true
		}

		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key != "href" {
					continue
				}
				link, err := resp.Request.URL.Parse(a.Val)
				if err != nil {
					continue // ignore bad URLs
				}
				links = append(links, link.String())
			}
		}
	}
	forEachNode(doc, visitNode, nil)
	return links, nil
}

func forEachNode(n *html.Node, pre, post func(n *html.Node) string) string {
	output := ""
	if pre != nil {
		output += pre(n)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		output += forEachNode(c, pre, post)
	}
	if post != nil {
		output += post(n)
	}
	return output
}
