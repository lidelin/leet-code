package _147_insertion_sort_list

import "github.com/lidelin/leet-code/utils"

type ListNode = utils.ListNode

func InsertionSortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	newHead := &ListNode{Val: 0, Next: head}
	compared := head.Next
	nextCompared := compared.Next
	prev := head
	current := newHead.Next

	head.Next = nil

	for compared != nil {
		nextCompared = compared.Next
		current = newHead.Next
		prev = newHead

		for current != nil {
			if current.Val < compared.Val {
				prev = current
				current = current.Next
			} else {
				compared.Next = current
				break
			}
		}

		prev.Next = compared
		if current == nil {
			compared.Next = nil
		}

		compared = nextCompared
	}

	return newHead.Next
}

//func InsertionSortList(head *ListNode) *ListNode {
//	if head == nil || head.Next == nil {
//		return head
//	} else {
//		newHead := InsertionSortList(head.Next)
//
//		head.Next = nil
//
//		if head.Val <= newHead.Val {
//			head.Next = newHead
//			return head
//		} else {
//			current := newHead.Next
//			prev := newHead
//			for current != nil {
//				if head.Val > current.Val {
//					prev = current
//					current = current.Next
//				} else {
//					head.Next = current
//					break
//				}
//			}
//
//			prev.Next = head
//
//			return newHead
//		}
//	}
//}
