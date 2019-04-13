package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var count = 0
var nodeCount = 0
var result int
var stop = false

func kthSmallest(root *TreeNode, k int) int {
	if stop {
		return result
	}

	if root.Left != nil {
		kthSmallest(root.Left, k)
	}

	count++
	if count == k {
		result = root.Val
		stop = true
		return result
	}

	if root.Right != nil {
		kthSmallest(root.Right, k)
	}

	return result
}

func createTree() *TreeNode {
	return &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val:   1,
					Left:  nil,
					Right: nil,
				},
				Right: nil,
			},
			Right: &TreeNode{
				Val:   4,
				Left:  nil,
				Right: nil,
			},
		},
		Right: &TreeNode{
			Val:   6,
			Left:  nil,
			Right: nil,
		},
	}
}

func main() {
	tree := createTree()

	fmt.Println(kthSmallest(tree, 5))
}
