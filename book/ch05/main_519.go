package main

import "fmt"

func magic() {
	panic("magic")
}

func main() {
	var returnValue interface{}

	getValue := func() {
		returnValue = recover()
	}
	printValue := func() {
		fmt.Println(returnValue)
	}

	defer printValue()
	defer getValue()

	magic()
}
