package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	counts := make(map[string]int)
	names := make(map[string]map[string]int)

	files := os.Args[1:]
	for _, fileName := range files {
		data, err := ioutil.ReadFile(fileName)
		if err != nil {
			fmt.Fprintf(os.Stderr, "bum: %v", err)
			continue
		}
		for _, line := range strings.Split(string(data), "\n") {
			if len(line) > 0 {
				counts[line]++
				if names[line] == nil {
					names[line] = make(map[string]int)
				}
				names[line][fileName]++
			}
		}
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\t%v\n", n, line, getNames(names[line]))
		}
	}
}

func getNames(names map[string]int) []string {
	var fileNames []string
	for k, _ := range names {
		fileNames = append(fileNames, k)
	}
	return fileNames
}
