package main

import (
	"book/ch05/ex511"
	"fmt"
)

func main() {
	for _, e := range ex511.TopoSortWithCycles(ex511.Prereqs) {
		fmt.Println(e)
	}
}
