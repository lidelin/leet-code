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

func createList(values []int) *ListNode {
	head := &ListNode{Next: nil}
	current := head
	length := len(values)

	for i := 0; i < length; i++ {
		current.Val = values[i]
		if i < length-1 {
			current.Next = &ListNode{Next: nil}
			current = current.Next
		} else {
			current.Next = nil
		}
	}

	return head
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
	list := createList([]int{1, 2, 6, 3, 4, 5, 6})
	printList(list)

	list = removeElements(list, 6)
	printList(list)
}
