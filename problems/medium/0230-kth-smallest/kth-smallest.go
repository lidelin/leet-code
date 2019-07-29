package _0230_kth_smallest

import "github.com/lidelin/leet-code/utils"

type TreeNode = utils.TreeNode

var count = 0
var result int
var stop = false

func KthSmallest(root *TreeNode, k int) int {
	if stop {
		return result
	}

	if root.Left != nil {
		KthSmallest(root.Left, k)
	}

	count++
	if count == k {
		result = root.Val
		stop = true
		return result
	}

	if root.Right != nil {
		KthSmallest(root.Right, k)
	}

	return result
}
