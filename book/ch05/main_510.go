package main

import (
	"book/ch05/ex510"
	"fmt"
)

func main() {
	for _, e := range toposort.TopoSort(toposort.Prereqs) {
		fmt.Println(e)
	}
	fmt.Println()
	for _, e := range toposort.TopoSortNonDeterministic(toposort.Prereqs) {
		fmt.Println(e)
	}
}
