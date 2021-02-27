package main

import "github.com/kbl/gopher_exercises/book/ch04/my_vim"
import "fmt"
import "log"

func main() {
	f, err := my_vim.Edit("Body: ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(f)
}
