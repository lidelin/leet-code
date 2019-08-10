package _965_univalued_binary_tree

import "github.com/lidelin/leet-code/utils"

type TreeNode = utils.TreeNode

func IsUnivalTree(root *TreeNode) bool {
	if root == nil {
		return true
	}

	if (root.Left != nil && root.Val != root.Left.Val) ||
		(root.Right != nil && root.Val != root.Right.Val) {
		return false
	}

	return IsUnivalTree(root.Left) && IsUnivalTree(root.Right)
}
