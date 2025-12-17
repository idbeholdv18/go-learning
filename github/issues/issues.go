package issues

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const IssueURL = "https://api.github.com/search/issues"

type Issue struct {
	Title     string
	CreatedAt time.Time `json:"created_at"`
}

type IssuesResponse struct {
	Items []Issue
	Count int
}

func GetIssues(terms []string) (resp *IssuesResponse, err error) {
	query := url.QueryEscape(strings.Join(terms, " "))

	response, err := http.Get(IssueURL + "?q=" + query)

	fmt.Println(IssueURL + "?q:" + query)

	if err != nil {
		response.Body.Close()
		return nil, err
	}

	if response.StatusCode != http.StatusOK {
		response.Body.Close()
		return nil, fmt.Errorf("getting failed: %s", response.Status)
	}

	var data IssuesResponse

	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		response.Body.Close()
		return nil, err
	}

	response.Body.Close()

	return &data, nil
}

func FilterIssuesByAge(issues []Issue, age int) []Issue {
	var result []Issue
	thirtyDaysAgo := time.Now().Add(-time.Duration(age) * 24 * time.Hour)
	for _, issue := range issues {
		if issue.CreatedAt.After(thirtyDaysAgo) {
			result = append(result, issue)
		}
	}
	return result
}
