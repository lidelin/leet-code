package _0938_range_sum_of_bst

import (
	"leet-code/utils"
)

type TreeNode = utils.TreeNode

func RangeSumBST(root *TreeNode, L int, R int) int {
	sum := 0

	rangeSumBST(root, L, R, &sum)

	return sum
}

func rangeSumBST(root *TreeNode, L int, R int, sum *int) {
	if root == nil {
		return
	}

	if root.Val >= L && root.Val <= R {
		*sum += root.Val
	}

	if root.Left != nil && root.Val >= L {
		rangeSumBST(root.Left, L, R, sum)
	}

	if root.Right != nil && root.Val <= R {
		rangeSumBST(root.Right, L, R, sum)
	}
}
