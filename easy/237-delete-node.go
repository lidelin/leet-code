package main

import "leet-code/utils"

func deleteNode(node *utils.ListNode) {
	node.Val = node.Next.Val
	node.Next = node.Next.Next
}

func main() {
	list := utils.CreateList([]int{1, 3, 5, 7, 9})

	deleteNode(list.Next.Next.Next)

	utils.PrintList(list)
}
