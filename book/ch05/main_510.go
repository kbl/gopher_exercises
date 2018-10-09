package main

import (
	"book/ch05/ex510"
	"fmt"
)

func main() {
	for _, e := range toposort.TopoSort(toposort.Prereqs) {
		fmt.Println(e)
	}
}
