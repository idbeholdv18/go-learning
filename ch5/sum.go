package main

import (
	"fmt"
)

func min(values ...int) (error, int) {
	if len(values) == 0 {
		return fmt.Errorf("min should contain at least 1 value"), 0
	}
	result := values[0]

	for _, value := range values[1:] {
		if value < result {
			result = value
		}
	}
	return nil, result
}

func main() {
	err, value := min(1, 5, -1, 6)

	if err != nil {
		fmt.Printf("error: %v", err)
	} else {
		fmt.Printf("min value is: %d\n", value)
	}
}
