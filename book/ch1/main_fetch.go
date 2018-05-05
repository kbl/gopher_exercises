package main

import (
	"book/ch1/fetch"
	"fmt"
	"os"
)

func main() {
	for _, url := range os.Args[1:] {
		body, err := fetch.Fetch(url)
		if err != nil {
			fmt.Printf("%s", err)
			os.Exit(1)
		}
		fmt.Printf("%s\n", body)
	}
}
