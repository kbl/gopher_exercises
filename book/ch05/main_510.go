package main

import (
	"fmt"
	"github.com/kbl/gopher_exercises/book/ch05/ex510"
	"github.com/kbl/gopher_exercises/book/ch05/toposort"
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
