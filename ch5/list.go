package main

import (
	"fmt"
	"go-learning/ch5/list"
)

func main() {
	linkedList := list.List{}
	list.Add(&linkedList, 4)
	list.Add(&linkedList, 3)
	list.Add(&linkedList, 2)

	list.MapList(&linkedList, func(node *list.ListNode) {
		fmt.Printf("%d -> ", node.Value)
		if node.Next == nil {
			fmt.Println("nil")
		}
	})
}
