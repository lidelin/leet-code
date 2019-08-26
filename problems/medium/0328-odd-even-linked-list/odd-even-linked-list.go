package _328_odd_even_linked_list

import "github.com/lidelin/leet-code/utils"

type ListNode = utils.ListNode

func OddEvenList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	odd := head
	even, evenHead := head.Next, head.Next
	current := head.Next.Next
	count := 1

	for current != nil {
		if count%2 == 0 {
			even.Next = current
			even = even.Next
		} else {
			odd.Next = current
			odd = odd.Next
		}

		count++
		current = current.Next
	}

	even.Next = nil
	odd.Next = evenHead

	return head
}
