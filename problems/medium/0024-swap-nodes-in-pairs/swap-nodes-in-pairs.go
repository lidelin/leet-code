package _024_swap_nodes_in_pairs

import "github.com/lidelin/leet-code/utils"

type ListNode = utils.ListNode

func SwapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	first, second := head, head.Next
	var prev *ListNode = nil
	newHead := second

	for first != nil && second != nil {
		first.Next = second.Next
		second.Next = first
		if prev != nil {
			prev.Next = second
		}

		prev = first
		first = first.Next
		if first != nil {
			second = first.Next
		}
	}

	return newHead
}
