package _817_linked_list_components

import "github.com/lidelin/leet-code/utils"

type ListNode = utils.ListNode

func NumComponents(head *ListNode, G []int) int {
	if head == nil {
		return 0
	}

	gMap := make(map[int]bool)
	for _, value := range G {
		gMap[value] = true
	}

	count := 0
	shouldCount := false

	for head != nil {
		if _, exists := gMap[head.Val]; !exists {
			if shouldCount == true {
				count++
				shouldCount = false
			}
		} else {
			shouldCount = true
		}

		head = head.Next
	}

	if shouldCount == true {
		count++
	}

	return count
}
