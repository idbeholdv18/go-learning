package main

import "fmt"

func main() {
	var slice []int
	printSlice(slice)
	slice = append(slice, 1)
	printSlice(slice)
	slice = append(slice, 2)
	printSlice(slice)
	slice = append(slice, 3)
	printSlice(slice)
	slice = append(slice, 4)
	printSlice(slice)
	slice = append(slice, 5)
	printSlice(slice)
}

func printSlice(slice []int) {
	fmt.Printf("%v\t\t%d\n", slice, cap(slice))
}
