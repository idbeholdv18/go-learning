package main

import (
	"fmt"
	"go-learning/github/issues"
	"log"
	"os"
)

func main() {
	issuesArr, err := issues.GetIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}

	dayAgoIssues := issues.FilterIssuesByAge(issuesArr.Items, 1)
	monthAgoIssues := issues.FilterIssuesByAge(issuesArr.Items, 30)
	yearAgoIssues := issues.FilterIssuesByAge(issuesArr.Items, 365)

	for n, issue := range dayAgoIssues {
		fmt.Printf("#%d\t%s\n", n, issue.Title)
	}
	fmt.Println("-----")

	for n, issue := range monthAgoIssues {
		fmt.Printf("#%d\t%s\n", n, issue.Title)
	}
	fmt.Println("-----")

	for n, issue := range yearAgoIssues {
		fmt.Printf("#%d\t%s\n", n, issue.Title)
	}
}
