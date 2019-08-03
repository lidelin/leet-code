package _235_lowest_common_ancestor

import "github.com/lidelin/leet-code/utils"

type TreeNode = utils.TreeNode

func LowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return root
	}

	if p.Val < root.Val && q.Val < root.Val && root.Left != nil {
		return LowestCommonAncestor(root.Left, p, q)
	}

	if p.Val > root.Val && q.Val > root.Val && root.Right != nil {
		return LowestCommonAncestor(root.Right, p, q)
	}

	if p.Val == root.Val {
		return p
	}

	if q.Val == root.Val {
		return q
	}

	return root
}
