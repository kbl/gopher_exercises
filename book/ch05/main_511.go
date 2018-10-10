package main

import (
	"book/ch05/ex511"
	"fmt"
	"log"
)

func main() {
	order, err := ex511.TopoSortWithCycles(ex511.Prereqs)
	if err != nil {
		log.Fatal(err)
	}
	for _, e := range order {
		fmt.Println(e)
	}
}
