package main

import (
	"book/ch05/ex510"
	"fmt"
)

func main() {
	for _, e := range ex510.TopoSort(ex510.Prereqs) {
		fmt.Println(e)
	}
	fmt.Println()
	for _, e := range ex510.TopoSortNonDeterministic(ex510.Prereqs) {
		fmt.Println(e)
	}
}
