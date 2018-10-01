package outline

import (
	"fmt"
	"golang.org/x/net/html"
	"os"
	"strings"
)

func Outline() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "outline: %v\n", err)
		os.Exit(1)
	}
	forEachNode(doc, startElement, endElement)
}

func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}
	if post != nil {
		post(n)
	}
}

var depth = 0

func startElement(n *html.Node) {
	switch n.Type {
	case html.ElementNode:
		attributes := printAttributes(n.Attr)
		closing := ">"
		if n.FirstChild == nil {
			closing = "/>"
		}
		fmt.Printf("%*s<%s%s%s\n", depth*2, "", n.Data, attributes, closing)
		depth++
	case html.TextNode:
		data := strings.TrimSpace(n.Data)
		if data != "" {
			fmt.Printf("%*s%s\n", depth*2, "", data)
		}
	}
}

func endElement(n *html.Node) {
	switch n.Type {
	case html.ElementNode:
		depth--
		if n.FirstChild != nil {
			fmt.Printf("%*s</%s>\n", depth*2, "", n.Data)
		}
	}
}

func printAttributes(attrs []html.Attribute) string {
	attributes := make([]string, 0)
	for _, a := range attrs {
		attributes = append(attributes, fmt.Sprintf("%s=\"%s\"", a.Key, a.Val))
	}

	strAttributes := strings.Join(attributes, " ")
	if len(strAttributes) > 0 {
		strAttributes = fmt.Sprintf(" %s", strAttributes)
	}

	return strAttributes
}
