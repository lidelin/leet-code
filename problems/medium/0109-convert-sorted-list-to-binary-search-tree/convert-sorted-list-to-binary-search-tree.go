package _109_convert_sorted_list_to_binary_search_tree

import (
	"leet-code/utils"
)

type ListNode = utils.ListNode
type TreeNode = utils.TreeNode

func SortedListToBST(head *ListNode) *TreeNode {
	if head == nil {
		return nil
	}

	var prev *ListNode = nil
	middle, fast := head, head

	for fast != nil && fast.Next != nil {
		prev = middle
		middle = middle.Next
		fast = fast.Next.Next
	}

	if prev != nil {
		prev.Next = nil
	}

	root := &TreeNode{Val: middle.Val, Left: nil, Right: nil}

	if head != middle {
		root.Left = SortedListToBST(head)
		root.Right = SortedListToBST(middle.Next)
	}

	return root
}
