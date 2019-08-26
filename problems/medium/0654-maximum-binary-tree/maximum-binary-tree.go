package _654_maximum_binary_tree

import "github.com/lidelin/leet-code/utils"

type TreeNode = utils.TreeNode

func constructMaximumBinaryTree(nums []int) *TreeNode {
	length := len(nums)
	if length == 0 {
		return nil
	}

	indexOfMaximum := 0

	for i := 0; i < length; i++ {
		if nums[i] > nums[indexOfMaximum] {
			indexOfMaximum = i
		}
	}

	root := &TreeNode{Val: nums[indexOfMaximum], Left: nil, Right: nil}
	if indexOfMaximum > 0 {
		root.Left = constructMaximumBinaryTree(nums[:indexOfMaximum])
	}
	if indexOfMaximum < length-1 {
		root.Right = constructMaximumBinaryTree(nums[indexOfMaximum+1:])
	}

	return root
}
