package main

import (
	"book/ch05/findlinks3"
	"book/ch05/toposort"
	"fmt"
)

func main() {
	var visited []string

	traverse := func(item string) []string {
		for _, i := range visited {
			if i == item {
				return nil
			}
		}
		visited = append(visited, item)
		return toposort.Prereqs[item]
	}

	findlinks3.BreadthFirst(traverse, keys(toposort.Prereqs))

	for _, i := range visited {
		fmt.Println(i)
	}
}

func keys(mapping map[string][]string) []string {
	keys := make([]string, len(mapping))
	index := 0
	for k := range mapping {
		keys[index] = k
		index++
	}
	return keys
}
