package ex511

import (
	"errors"
	"fmt"
	"sort"
)

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
	"intro to programming":  {"networks"}, // cycle
}

func TopoSortWithCycles(m map[string][]string) ([]string, error) {
	var order []string
	seen := make(map[string]bool)
	var visitAll func(items, dependencies []string) error
	visitAll = func(items, dependencies []string) error {
		for _, item := range items {
			for _, d := range dependencies {
				if d == item {
					return errors.New(fmt.Sprintf("cycle found! %s %s", dependencies, item))
				}
			}
			item_dependencies := append(dependencies, item)
			err := visitAll(m[item], item_dependencies)
			if err != nil {
				return err
			}
			if !seen[item] {
				seen[item] = true
				order = append(order, item)
			}
		}
		return nil
	}

	var keys []string
	for key := range m {
		keys = append(keys, key)
	}

	sort.Strings(keys)
	err := visitAll(keys, nil)
	if err == nil {
		return order, nil
	}
	return nil, err
}
