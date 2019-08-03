package _235_lowest_common_ancestor

import "github.com/lidelin/leet-code/utils"

type TreeNode = utils.TreeNode

func LowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return root
	}

	if p.Val < root.Val && q.Val < root.Val && root.Left != nil {
		return LowestCommonAncestor(root.Left, p, q)
	} else if p.Val > root.Val && q.Val > root.Val && root.Right != nil {
		return LowestCommonAncestor(root.Right, p, q)
	}

	return root
}

func LowestCommonAncestorIteration(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return root
	}

	node := root

	for node != nil {
		if p.Val < node.Val && q.Val < node.Val && node.Left != nil {
			node = node.Left
		} else if p.Val > node.Val && q.Val > node.Val && node.Right != nil {
			node = node.Right
		} else {
			return node
		}
	}

	return root
}
