package main

import (
	"fmt"
	"github.com/kbl/gopher_exercises/book/ch04/my_github"
	"log"
	"os"
	"time"
)

func main() {
	result, err := my_github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}

	now := time.Now()
	issues := make(map[string][]*my_github.Issue)

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
			issues[category] = make([]*my_github.Issue, 0)
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
