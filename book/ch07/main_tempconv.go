package main

import (
	"flag"
	"fmt"
	"github.com/kbl/gopher_exercises/book/ch07/tempconv"
)

var temp = tempconv.CelsiusFlag("temp", 20.0, "the temperature")

func main() {
	flag.Parse()
	fmt.Println(*temp)
}
