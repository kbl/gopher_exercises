package main

import (
	"book/ch07/ex706"
	"flag"
	"fmt"
)

var temp = ex706.CelsiusFlag("temp", 20.0, "the temperature")

func main() {
	flag.Parse()
	fmt.Println(*temp)
}
