package _0938_range_sum_of_bst

import (
	"leet-code/utils"
)

type TreeNode = utils.TreeNode

func RangeSumBST(root *TreeNode, L int, R int) int {
	if root == nil {
		return 0
	}

	sum := 0

	if root.Val >= L && root.Val <= R {
		sum += root.Val
	}

	if root.Val > L {
		sum += RangeSumBST(root.Left, L, R)
	}

	if root.Val < R {
		sum += RangeSumBST(root.Right, L, R)
	}

	return sum
}
