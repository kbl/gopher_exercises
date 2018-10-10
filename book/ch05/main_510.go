package main

import (
	"book/ch05/ex510"
	"book/ch05/toposort"
	"fmt"
)

func main() {
	for _, e := range toposort.TopoSort(toposort.Prereqs) {
		fmt.Println(e)
	}
	fmt.Println()
	for _, e := range ex510.TopoSortNonDeterministic(toposort.Prereqs) {
		fmt.Println(e)
	}
}
