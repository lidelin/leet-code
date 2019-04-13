package main

import (
	"leet-code/utils"
)

func removeElements(head *utils.ListNode, val int) *utils.ListNode {
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

func main() {
	list := utils.MakeList([]int{1, 2, 6, 3, 4, 5, 6})
	utils.PrintList(list)

	list = removeElements(list, 6)
	utils.PrintList(list)
}
