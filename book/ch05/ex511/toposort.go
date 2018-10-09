package ex511

import (
	"fmt"
	"sort"
)

// my own

var Prereqs = map[string][]string{
	"algorithms":            {"data structures"},
	"calculus":              {"linear algebra"},
	"compilers":             {"data structures", "formal languages", "computer organization"},
	"data structures":       {"discrete math"},
	"databases":             {"data structures"},
	"discrete math":         {"intro to programming"},
	"formal languages":      {"discrete math"},
	"networks":              {"operating systems"},
	"operating systems":     {"data structures", "computer organization"},
	"programming languages": {"data structures", "computer organization"},
	//"linear algebra":        {"calculus"},
}

// TODO
func TopoSortWithCycles(m map[string][]string) []string {
	var order []string
	seen := make(map[string]bool)
	var visitAll func(items, dependencies []string)
	visitAll = func(items, dependencies []string) {
		for _, item := range items {
			for _, d := range dependencies {
				if d == item {
					fmt.Printf("cycle found! %s %s\n", dependencies, item)
					return
				}
			}
			dependencies = append(dependencies, item)
			visitAll(m[item], dependencies)
			if !seen[item] {
				seen[item] = true
				order = append(order, item)
			}
		}
	}

	var keys []string
	for key := range m {
		keys = append(keys, key)
	}

	sort.Strings(keys)
	visitAll(keys, nil)
	return order
}
