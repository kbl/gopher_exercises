package echonodes

import (
	"fmt"
	"golang.org/x/net/html"
	"strings"
)

func EchoNodes(n *html.Node, debug string) {
	if n == nil {
		return
	}

	if n.Type == html.TextNode {
		if data := trim(n.Data); data != "" {
			fmt.Println(data)
		}
	}

	if n.Type == html.ElementNode {
		if n.Data != "style" && n.Data != "script" {
			EchoNodes(n.FirstChild, fmt.Sprintf("first x %v %v", n.Data, n.Type))
		}
	} else {
		EchoNodes(n.FirstChild, fmt.Sprintf("first y %v %v", n.Data, n.Type))
	}

	EchoNodes(n.NextSibling, "next")
}

func trim(s string) string {
	current := s
	for previous := ""; current != previous; {
		current, previous = strings.TrimRight(current, "\n\t "), current
		current, previous = strings.TrimLeft(current, "\n\t "), current
	}
	return current
}
