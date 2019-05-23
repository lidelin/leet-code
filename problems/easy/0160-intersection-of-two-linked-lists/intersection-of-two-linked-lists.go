package _160_intersection_of_two_linked_lists

import "leet-code/utils"

type ListNode = utils.ListNode

func GetIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}

	start := 0
	sizeA := size(headA)
	sizeB := size(headB)

	if sizeA > sizeB {
		start = sizeA - sizeB
		for i := 0; i < start; i++ {
			headA = headA.Next
		}
	} else {
		start = sizeB - sizeA
		for i := 0; i < start; i++ {
			headB = headB.Next
		}
	}

	var intersection *ListNode = nil

	for headA != nil {
		if headA == headB {
			intersection = headA
			break
		} else {
			headA = headA.Next
			headB = headB.Next
		}
	}

	return intersection
}

func size(head *ListNode) int {
	if head == nil {
		return 0
	}

	size := 1

	for head != nil {
		size++
		head = head.Next
	}

	return size
}
