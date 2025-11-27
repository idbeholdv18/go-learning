package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5}

	reverse(arr)

	fmt.Println(arr)
}

func reverse(array []int) {
	for i, j := 0, len(array)-1; i < j; i, j = i+1, j-1 {
		array[i], array[j] = array[j], array[i]
	}
}
