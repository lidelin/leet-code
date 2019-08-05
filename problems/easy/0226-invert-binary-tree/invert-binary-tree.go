package _226_invert_binary_tree

import "github.com/lidelin/leet-code/utils"

type TreeNode = utils.TreeNode

func InvertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}

	if root.Left != nil {
		root.Left = InvertTree(root.Left)
	}

	if root.Right != nil {
		root.Right = InvertTree(root.Right)
	}

	temp := root.Right
	root.Right = root.Left
	root.Left = temp

	return root
}
