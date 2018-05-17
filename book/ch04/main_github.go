package main

import (
	"book/ch04/github"
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}

	now := time.Now()
	issues := make(map[string][]*github.Issue)

	lessThanMonth, err := time.ParseDuration(fmt.Sprintf("%dh", 30*24))
	if err != nil {
		log.Fatal(err)
	}
	lessThanYear, err := time.ParseDuration(fmt.Sprintf("%dh", 365*24))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%d issues:\n", result.TotalCount)

	for _, item := range result.Items {
		age := now.Sub(item.CreatedAt)
		category := ">= 1y"
		if age < lessThanMonth {
			category = "< 30d"
		} else if age < lessThanYear {
			category = "< 1y"
		}

		if _, exists := issues[category]; !exists {
			issues[category] = make([]*github.Issue, 0)
		}

		issues[category] = append(issues[category], item)
	}

	for _, category := range [...]string{"< 30d", "< 1y", ">= 1y"} {
		fmt.Printf("%d issues %s:\n", len(issues[category]), category)
		for _, item := range issues[category] {
			fmt.Printf("#%-5d %9.9s %.55s\n", item.Number, item.User.Login, item.Title)
		}
	}
}
