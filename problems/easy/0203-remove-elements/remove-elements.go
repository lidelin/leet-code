package _0203_remove_elements

import "github.com/lidelin/leet-code/utils"

type ListNode = utils.ListNode

func RemoveElements(head *ListNode, val int) *ListNode {
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
