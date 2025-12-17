package list

type ListNode struct {
	Value int
	Next  *ListNode
}

type List struct {
	Head *ListNode
}

func MapList(list *List, cb func(node *ListNode)) {
	for n := list.Head; n != nil; n = n.Next {
		cb(n)
	}
}

func Add(list *List, value int) {
	newNode := &ListNode{Value: value, Next: nil}
	if list.Head == nil {
		list.Head = newNode
		return
	}

	curr := list.Head
	for curr.Next != nil {
		curr = curr.Next
	}

	curr.Next = newNode
}
