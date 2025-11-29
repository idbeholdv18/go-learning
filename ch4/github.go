package main

import (
	"fmt"
	"go-learning/ch4/github"
	"log"
	"os"
)

func main() {
	result, err := github.SearchIssues(os.Args[1:])

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%d тем: \n", result.TotalCount)

	for _, term := range result.Items {
		fmt.Printf("#%-5d %9.9s %.55s\n",
			term.Number, term.User.Login, term.Title)
	}
}
