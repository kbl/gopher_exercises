package main

import (
	"book/ch07/tempconv"
	"flag"
	"fmt"
)

var temp = tempconv.CelsiusFlag("temp", 20.0, "the temperature")

func main() {
	flag.Parse()
	fmt.Println(*temp)
}
