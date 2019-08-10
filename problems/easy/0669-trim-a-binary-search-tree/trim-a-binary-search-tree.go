package _669_trim_a_binary_search_tree

import "github.com/lidelin/leet-code/utils"

type TreeNode = utils.TreeNode

func TrimBST(root *TreeNode, L int, R int) *TreeNode {
	if root == nil {
		return nil
	}

	if root.Val > R {
		return TrimBST(root.Left, L, R)
	}

	if root.Val < L {
		return TrimBST(root.Right, L, R)
	}

	if root.Left != nil {
		node := TrimBST(root.Left, L, R)
		if root.Left.Val < L || root.Left.Val > R {
			root.Left = node
		}
	}

	if root.Right != nil {
		node := TrimBST(root.Right, L, R)
		if root.Right.Val < L || root.Right.Val > R {
			root.Right = node
		}
	}

	return root
}
