package main

import (
	"flag"
	"fmt"
	"github.com/kbl/gopher_exercises/book/ch07/ex706"
)

var temp = ex706.CelsiusFlag("temp", 20.0, "the temperature")

func main() {
	flag.Parse()
	fmt.Println(*temp)
}
