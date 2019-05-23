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

func PrintList(head *ListNode) {
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

func CatList(l1, l2 *ListNode) {
	for l1.Next != nil {
		l1 = l1.Next
	}

	l1.Next = l2
}
