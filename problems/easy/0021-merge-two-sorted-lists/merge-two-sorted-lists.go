package _021_merge_two_sorted_lists

import "leet-code/utils"

type ListNode = utils.ListNode

func MergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	head := &ListNode{Val: 0, Next: nil}
	current := head

	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			current.Next = l1
			l1 = l1.Next
		} else {
			current.Next = l2
			l2 = l2.Next
		}

		current = current.Next
	}

	if l1 != nil {
		current.Next = l1
	} else {
		current.Next = l2
	}

	return head.Next
}

func MergeTwoListsRecursively(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	if l1.Val <= l2.Val {
		l1.Next = MergeTwoListsRecursively(l1.Next, l2)
		return l1
	} else {
		l2.Next = MergeTwoListsRecursively(l1, l2.Next)
		return l2
	}
}
