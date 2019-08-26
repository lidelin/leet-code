package _538_convert_bst_to_greater_tree

import "github.com/lidelin/leet-code/utils"

type TreeNode = utils.TreeNode

func convertBST(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}

	dfs(root, 0)

	return root
}

func dfs(node *TreeNode, sum int) int {
	if node == nil {
		return sum
	}

	sum = dfs(node.Right, sum)
	value := node.Val
	node.Val += sum
	sum += value

	return dfs(node.Left, sum)
}
