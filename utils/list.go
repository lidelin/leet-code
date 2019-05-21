package utils

import (
	"fmt"
	"strconv"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func MakeList(values []int) *ListNode {
	head := &ListNode{Next: nil}
	prev := head

	for _, value := range values {
		prev.Next = &ListNode{Val: value, Next: nil}
		prev = prev.Next
	}

	return head
}

func PrintList(head *ListNode) {
	if head == nil {
		return
	}

	var str []string
	current := head.Next

	for current != nil {
		str = append(str, strconv.Itoa(current.Val))
		current = current.Next
	}

	fmt.Println(str)
}
