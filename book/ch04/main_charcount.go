package main

import (
	"fmt"
	"github.com/kbl/gopher_exercises/book/ch04/my_count"
	"strings"
)

func main() {
	stats := my_count.Runes(strings.NewReader("zażółć gęślą jaźń"))
	fmt.Printf("rune\tcount\n")
	for r, c := range stats.Runes {
		fmt.Printf("%q\t%d\n", r, c)
	}
	fmt.Printf("\nletter\tcount\n")
	for r, c := range stats.Letters {
		fmt.Printf("%q\t%d\n", r, c)
	}
	fmt.Printf("\nlength\tcount\n")
	for l, c := range stats.Lenghts {
		fmt.Printf("%d\t%d\n", l, c)
	}
	if stats.Invalid > 0 {
		fmt.Printf("\n%d invalid UTF-8 characters\n", stats.Invalid)
	}
}
