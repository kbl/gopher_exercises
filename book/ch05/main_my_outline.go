package main

import (
	"fmt"
	"github.com/kbl/gopher_exercises/book/ch05/my_outline"
	"os"
)

func main() {
	fmt.Println(my_outline.Outline(os.Stdin))
	// fmt.Println(my_outline.ElementById(os.Stdin, "main"))
}
