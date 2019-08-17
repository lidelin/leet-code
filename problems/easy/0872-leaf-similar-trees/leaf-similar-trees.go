package _872_leaf_similar_trees

import (
	"github.com/lidelin/leet-code/utils"
)

type TreeNode = utils.TreeNode

func LeafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	if root1 == nil || root2 == nil {
		return false
	}

	leaves1 := getLeaves(root1)
	leaves2 := getLeaves(root2)

	length := len(leaves1)
	if length != len(leaves2) {
		return false
	}

	for i := 0; i < length; i++ {
		if leaves1[i] != leaves2[i] {
			return false
		}
	}

	return true
}

func getLeaves(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	var leaves []int

	nodes := make([]*TreeNode, 0)
	nodes = append(nodes, root)

	length := len(nodes)

	for length != 0 {
		topIndex := length - 1
		topNode := nodes[topIndex]
		nodes = nodes[:topIndex]

		if topNode.Left != nil {
			nodes = append(nodes, topNode.Left)
		}
		if topNode.Right != nil {
			nodes = append(nodes, topNode.Right)
		}
		if topNode.Left == nil && topNode.Right == nil {
			leaves = append(leaves, topNode.Val)
		}

		length = len(nodes)
	}

	return leaves
}
