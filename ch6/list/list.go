package list

import (
	"fmt"
)

type ListNode struct {
	Value int
	Next  *ListNode
	Prev  *ListNode
}

type List struct {
	Head *ListNode
	Tail *ListNode
}

func MakeList(values ...int) *List {
	var list = &List{}
	for _, n := range values {
		list.PushBack(n)
	}
	return list
}

func (list *List) PushBack(value int) *List {
	var newNode = &ListNode{Value: value, Next: nil, Prev: nil}
	if list.Head == nil {
		list.Head = newNode
		list.Tail = newNode
		return list
	}

	list.Tail.Next = newNode
	newNode.Prev = list.Tail
	list.Tail = newNode
	return list
}

func (list *List) PushFront(value int) *List {
	var node = &ListNode{Value: value}
	if list.Head == nil {
		list.Head = node
		list.Tail = node
		return list
	}

	list.Head.Prev = node
	node.Next = list.Head
	list.Head = node
	return list
}

func (list *List) PopBack() *List {
	if list.Head == nil {
		return list
	}
	if list.Head == list.Tail {
		list.Tail = nil
		list.Head = nil
		return list
	}

	list.Tail = list.Tail.Prev
	list.Tail.Next.Prev = nil
	list.Tail.Next = nil
	return list
}

func (list *List) PopFront() *List {
	if list.Head == nil {
		return list
	}
	if list.Head == list.Tail {
		list.Head = nil
		list.Tail = nil
		return list
	}
	list.Head = list.Head.Next
	list.Head.Prev.Next = nil
	list.Head.Prev = nil
	return list
}

func (list List) For(cb func(node ListNode)) *List {
	for n := list.Head; n != nil; n = n.Next {
		cb(*n)
	}
	return &list
}

func (list List) Print() *List {
	List.For(list, func(node ListNode) {
		fmt.Printf("%d -> ", node.Value)
	})
	fmt.Println("nil")
	return &list
}

func (list List) Size() int {
	var result = 0
	for node := list.Head; node != nil; node = node.Next {
		result += 1
	}
	return result
}

func (list List) FindByValue(value int) (int, *ListNode) {
	var i = 0
	for node := list.Head; node != nil; node = node.Next {
		if node.Value == value {
			return i, node
		}
		i++
	}
	return -1, nil
}
