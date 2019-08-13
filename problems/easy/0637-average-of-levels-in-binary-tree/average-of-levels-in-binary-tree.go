package _637_average_of_levels_in_binary_tree

import (
	"github.com/lidelin/leet-code/utils"
)

type TreeNode = utils.TreeNode

func AverageOfLevels(root *TreeNode) []float64 {
	if root == nil {
		return []float64{}
	}

	var averages []float64

	nodes := make([]*TreeNode, 0)
	nodes = append(nodes, root)

	for len(nodes) != 0 {
		count := len(nodes)
		sum := 0
		for i := 0; i < count; i++ {
			sum += nodes[0].Val
			if nodes[0].Left != nil {
				nodes = append(nodes, nodes[0].Left)
			}
			if nodes[0].Right != nil {
				nodes = append(nodes, nodes[0].Right)
			}

			nodes = nodes[1:]
		}

		averages = append(averages, float64(sum)/float64(count))
	}

	return averages
}
