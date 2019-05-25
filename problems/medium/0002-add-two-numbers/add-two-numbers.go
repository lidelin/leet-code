package _002_add_two_numbers

import "leet-code/utils"

type ListNode = utils.ListNode

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil && l2 == nil {
		return nil
	}

	result := &ListNode{Val: 0, Next: nil}
	current := result
	carry := 0
	sum := 0

	for {
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		sum += carry
		value := sum % 10
		carry = sum / 10
		sum = 0

		current.Next = &ListNode{Val: value, Next: nil}
		current = current.Next

		if l1 == nil && l2 == nil && carry == 0 {
			break
		}
	}

	return result.Next
}
