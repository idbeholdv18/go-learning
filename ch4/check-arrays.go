package main

import "fmt"

func main() {
	array := [...]int{1, 2, 3, 4, 5}
	fmt.Println(array)

	incremented := increment(array)
	fmt.Println(*incremented)

	incrementThis(&array)
	fmt.Println(array)

}

func increment(array [5]int) *[5]int {
	for i := range array {
		array[i] += 1
	}
	return &array
}

func incrementThis(array *[5]int) {
	for i := range array {
		array[i] += 1
	}
}
