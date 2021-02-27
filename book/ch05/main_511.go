package main

import (
	"fmt"
	"github.com/kbl/gopher_exercises/book/ch05/ex511"
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
