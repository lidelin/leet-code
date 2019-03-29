package main

import (
	"fmt"
	"strconv"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeElements(head *ListNode, val int) *ListNode {
	if head == nil {
		return nil
	}

	prev := head
	current := head.Next

	for {
		if current == nil {
			break
		}

		if current.Val == val {
			prev.Next = current.Next
		} else {
			prev = current
		}

		current = current.Next
	}

	return head
}

func createList() *ListNode {
	return &ListNode{Val: 1,
		Next: &ListNode{Val: 2,
			Next: &ListNode{Val: 6,
				Next: &ListNode{Val: 3,
					Next: &ListNode{Val: 4,
						Next: &ListNode{Val: 5,
							Next: &ListNode{Val: 6,
								Next: nil}}}}}}}
}

func printList(head *ListNode) {
	if head == nil {
		return
	}

	var str []string
	current := head

	for {
		if current == nil {
			break
		}

		str = append(str, strconv.Itoa(current.Val))
		current = current.Next
	}

	fmt.Println(str)
}

func main() {
	list := createList()
	printList(list)

	list = removeElements(list, 6)
	printList(list)
}
