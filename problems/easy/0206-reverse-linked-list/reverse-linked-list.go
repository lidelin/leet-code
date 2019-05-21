package _206_reverse_linked_list

import "leet-code/utils"

type ListNode = utils.ListNode

func ReverseList(head *ListNode) *ListNode {
	current := head.Next
	next := current.Next
	nextNext := next.Next

	current.Next = nil

	for next != nil {
		next.Next = current
		current = next
		next = nextNext

		if next != nil {
			nextNext = next.Next
		}
	}

	head.Next = current

	return head
}
