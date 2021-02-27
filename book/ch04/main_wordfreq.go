package main

import (
	"fmt"
	"github.com/kbl/gopher_exercises/book/ch04/my_count"
	"strings"
)

func main() {
	stats := my_count.Words(strings.NewReader("zażółć gęślą jaźń"))
	fmt.Printf("word\t\tcount\n")
	for w, c := range stats.Words {
		fmt.Printf("%q\t\t%d\n", w, c)
	}
	fmt.Printf("\nlength\tcount\n")
	for l, c := range stats.Lenghts {
		fmt.Printf("%d\t%d\n", l, c)
	}
}
