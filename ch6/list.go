package main

import (
	"fmt"
	"go-learning/ch6/list"
)

func main() {
	var linkedList = list.MakeList(1, 4, 2, 3)
	linkedList.
		PushBack(5).
		PushBack(3).
		PushBack(1).
		PushBack(8).
		Print().
		PopBack().
		PushFront(0).
		Print().
		PopFront().
		Print()

	fmt.Printf("Size of list is %d\n", linkedList.Size())

	findAndPrintNodeByValue(linkedList, 5)
	findAndPrintNodeByValue(linkedList, 100)
}

func findAndPrintNodeByValue(list *list.List, value int) {
	var i, node = list.FindByValue(value)
	if node == nil {
		fmt.Printf("Node with value %d doesn't exist\n", value)
		return
	}
	fmt.Printf("index of node with value %d is %d\n", value, i)
}
