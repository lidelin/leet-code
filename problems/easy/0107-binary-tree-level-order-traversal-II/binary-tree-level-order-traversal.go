package _107_binary_tree_level_order_traversal_II

import (
	"github.com/lidelin/leet-code/utils"
)

type TreeNode = utils.TreeNode

func LevelOrderBottom(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	var results [][]int

	nodes := make([]*TreeNode, 0)
	nodes = append(nodes, root)

	for len(nodes) != 0 {
		count := len(nodes)
		result := make([]int, 0)
		for i := 0; i < count; i++ {
			result = append(result, nodes[0].Val)

			if nodes[0].Left != nil {
				nodes = append(nodes, nodes[0].Left)
			}
			if nodes[0].Right != nil {
				nodes = append(nodes, nodes[0].Right)
			}

			nodes = nodes[1:]
		}

		results = append(results, result)
	}

	head := 0
	tail := len(results) - 1
	for head < tail {
		temp := results[head]
		results[head] = results[tail]
		results[tail] = temp

		head++
		tail--
	}

	return results
}
