package main

import "fmt"

func magic() (returnValue string) {
	d := func() {
		recover()
		returnValue = "magic"
	}
	defer d()
	panic("panic")
}

func main() {
	fmt.Println(magic())
}
