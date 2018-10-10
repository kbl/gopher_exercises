package main

import (
	"book/ch05/outline"
	"fmt"
	"os"
)

func main() {
	fmt.Println(outline.Outline(os.Stdin))
	// fmt.Println(outline.ElementById(os.Stdin, "main"))
}
