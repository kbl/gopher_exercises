package main

import "book/ch04/vim"
import "fmt"
import "log"

func main() {
	f, err := vim.Edit("Body: ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(f)
}
