package outline

import (
	"fmt"
	"golang.org/x/net/html"
	"io"
	"os"
	"strings"
)

var depth = 0

func Outline(reader io.Reader) string {
	doc, err := html.Parse(reader)
	if err != nil {
		fmt.Fprintf(os.Stderr, "outline: %v\n", err)
		os.Exit(1)
	}
	return forEachNode(doc, startElement, endElement)
}

func ElementById(reader io.Reader, id string) *html.Node {
	doc, err := html.Parse(reader)
	if err != nil {
		fmt.Fprintf(os.Stderr, "outline: %v\n", err)
		os.Exit(1)
	}
	return elementById(doc, id)
}

func elementById(n *html.Node, id string) *html.Node {
	fmt.Println(id)
	_, stopIteration := startElement(n, id)
	if stopIteration {
		return n
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		output := elementById(c, id)
		if output != nil {
			return output
		}
	}
	return nil
}

func forEachNode(n *html.Node, pre func(n *html.Node, id string) (string, bool), post func(n *html.Node) string) string {
	output := ""
	if pre != nil {
		part, _ := pre(n, "")
		output += part
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		output += forEachNode(c, pre, post)
	}
	if post != nil {
		output += post(n)
	}
	return output
}

func startElement(n *html.Node, id string) (output string, stopTraversal bool) {
	switch n.Type {
	case html.DoctypeNode:
		output = fmt.Sprintf("%*s<!doctype %s>\n", depth*2, "", n.Data)
	case html.ElementNode:
		if id != "" {
			for _, a := range n.Attr {
				if a.Key == "id" && a.Val == id {
					stopTraversal = true
					return
				}
			}
		}
		attributes := printAttributes(n.Attr)
		closing := ">"
		if n.FirstChild == nil {
			closing = "/>"
		}
		output = fmt.Sprintf("%*s<%s%s%s\n", depth*2, "", n.Data, attributes, closing)
		depth++
	case html.TextNode:
		data := strings.TrimSpace(n.Data)
		if data != "" {
			output = fmt.Sprintf("%*s%s\n", depth*2, "", data)
		}
	case html.CommentNode:
		data := strings.TrimSpace(n.Data)
		if data != "" {
			output = fmt.Sprintf("%*s<!-- %s -->\n", depth*2, "", data)
		}
	}
	return
}

func endElement(n *html.Node) string {
	output := ""
	switch n.Type {
	case html.ElementNode:
		depth--
		if n.FirstChild != nil {
			output = fmt.Sprintf("%*s</%s>\n", depth*2, "", n.Data)
		}
	}
	return output
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
