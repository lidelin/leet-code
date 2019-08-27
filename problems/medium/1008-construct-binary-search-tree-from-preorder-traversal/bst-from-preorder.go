package _008_construct_binary_search_tree_from_preorder_traversal

import (
	"github.com/lidelin/leet-code/utils"
)

type TreeNode = utils.TreeNode

func bstFromPreorder(preorder []int) *TreeNode {
	length := len(preorder)
	if length == 0 {
		return nil
	}

	root := &TreeNode{Val: preorder[0], Left: nil, Right: nil}

	if length > 1 {
		for _, value := range preorder[1:] {
			insert(root, value)
		}
	}

	return root
}

func insert(root *TreeNode, value int) {
	if root == nil {
		return
	}

	if value < root.Val {
		if root.Left == nil {
			root.Left = &TreeNode{Val: value, Left: nil, Right: nil}
		} else {
			insert(root.Left, value)
		}
	} else if value > root.Val {
		if root.Right == nil {
			root.Right = &TreeNode{Val: value, Left: nil, Right: nil}
		} else {
			insert(root.Right, value)
		}
	}
}
