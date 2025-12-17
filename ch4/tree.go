package main

import "fmt"

type bst_node struct {
	value int
	left  *bst_node
	right *bst_node
}

func main() {
	bst := add(nil, 8)
	bst = add(bst, 3)
	bst = add(bst, 4)
	bst = add(bst, 1)
	bst = add(bst, 12)
	bst = add(bst, 2)
	bst = add(bst, 6)

	min := getMin(bst)

	fmt.Println(min.value)
}

func add(root *bst_node, value int) *bst_node {
	if root == nil {
		root = new(bst_node)
		root.value = value
		return root
	}

	if value < root.value {
		root.left = add(root.left, value)
	} else {
		root.right = add(root.right, value)
	}
	return root
}

func getMin(root *bst_node) *bst_node {
	curr := root
	for curr.left != nil {

		curr = curr.left
	}
	return curr
}
