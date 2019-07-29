package _876_middle_of_the_linked_list

import "github.com/lidelin/leet-code/utils"

type ListNode = utils.ListNode

func MiddleNode(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	slow := head
	fast := head

	for fast != nil {
		if fast.Next != nil {
			fast = fast.Next.Next
			slow = slow.Next
		} else {
			fast = fast.Next
		}
	}

	return slow
}
