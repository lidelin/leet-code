package _234_palindrome_linked_list

import "github.com/lidelin/leet-code/utils"

type ListNode = utils.ListNode

func IsPalindrome(head *ListNode) bool {
	middle := findMiddle(head)

	utils.PrintList(middle)

	reversed := reverse(middle)

	for head != middle {
		if head.Val != reversed.Val {
			return false
		} else {
			head = head.Next
			reversed = reversed.Next
		}
	}

	return true
}

func findMiddle(head *ListNode) *ListNode {
	slow, fast := head, head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	return slow
}

func reverse(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	prev := head
	current := head.Next
	next := current.Next

	prev.Next = nil

	for current.Next != nil {
		current.Next = prev

		prev = current
		current = next
		next = current.Next
	}

	current.Next = prev

	return current
}
