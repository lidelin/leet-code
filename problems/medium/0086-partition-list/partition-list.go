package _086_partition_list

import "github.com/lidelin/leet-code/utils"

type ListNode = utils.ListNode

func Partition(head *ListNode, x int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	var less, lessHead, greater, greaterHead *ListNode = nil, nil, nil, nil
	current := head

	for current != nil {
		if current.Val < x {
			if less == nil {
				less = current
				lessHead = current
			} else {
				less.Next = current
				less = less.Next
			}
		} else {
			if greater == nil {
				greater = current
				greaterHead = current
			} else {
				greater.Next = current
				greater = greater.Next
			}
		}

		current = current.Next
	}

	if less != nil {
		less.Next = greaterHead
	}

	if greater != nil {
		greater.Next = nil
	}

	if lessHead != nil {
		return lessHead
	} else {
		return greaterHead
	}
}
