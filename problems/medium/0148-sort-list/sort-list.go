package _148_sort_list

import "github.com/lidelin/leet-code/utils"

type ListNode = utils.ListNode

func SortListMerge(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	prev, middle, fast := head, head, head

	for fast != nil && fast.Next != nil {
		prev = middle
		middle = middle.Next
		fast = fast.Next.Next
	}
	prev.Next = nil

	left := SortListMerge(head)
	right := SortListMerge(middle)

	newHead := &ListNode{Val: 0, Next: nil}
	current := newHead

	for left != nil && right != nil {
		if left.Val <= right.Val {
			current.Next = left
			left = left.Next
		} else {
			current.Next = right
			right = right.Next
		}

		current = current.Next
	}

	if left != nil {
		current.Next = left
	} else {
		current.Next = right
	}

	return newHead.Next
}

func SortListQuick(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	var left, leftCurrent, right, rightCurrent *ListNode = nil, nil, nil, nil
	pivot := head
	head = head.Next

	for head != nil {
		if head.Val <= pivot.Val {
			if left == nil {
				left = head
				leftCurrent = head
			} else {
				leftCurrent.Next = head
				leftCurrent = leftCurrent.Next
			}
		} else {
			if right == nil {
				right = head
				rightCurrent = head
			} else {
				rightCurrent.Next = head
				rightCurrent = rightCurrent.Next
			}
		}

		head = head.Next
	}

	if leftCurrent != nil {
		leftCurrent.Next = nil
	}
	if rightCurrent != nil {
		rightCurrent.Next = nil
	}

	left = SortListQuick(left)
	right = SortListQuick(right)

	if left != nil {
		leftCurrent = left
		for leftCurrent.Next != nil {
			leftCurrent = leftCurrent.Next
		}

		leftCurrent.Next = pivot
		pivot.Next = right

		return left
	} else {
		pivot.Next = right
		return pivot
	}
}
