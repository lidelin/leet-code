package _206_reverse_linked_list

import "github.com/lidelin/leet-code/utils"

type ListNode = utils.ListNode

func ReverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	current := head
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

	return current
}

func ReverseListRecursively(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	} else {
		newHead := ReverseListRecursively(head.Next)

		head.Next.Next = head
		head.Next = nil

		return newHead
	}
}
