package _897_increasing_order_search_tree

import "github.com/lidelin/leet-code/utils"

type TreeNode = utils.TreeNode

func IncreasingBST(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}

	values := inOrderTraversal(root)
	length := len(values)
	newRoot := &TreeNode{Val: values[0]}
	current := newRoot

	for i := 1; i < length; i++ {
		current.Right = &TreeNode{Val: values[i]}
		current = current.Right
	}

	return newRoot
}

func inOrderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	var result []int

	left := inOrderTraversal(root.Left)
	result = append(result, left...)
	result = append(result, root.Val)
	right := inOrderTraversal(root.Right)
	result = append(result, right...)

	return result
}
