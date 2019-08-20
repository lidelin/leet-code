package _257_binary_tree_paths

import (
	"github.com/lidelin/leet-code/utils"
	"strconv"
)

type TreeNode = utils.TreeNode

func binaryTreePaths(root *TreeNode) []string {
	if root == nil {
		return []string{}
	}

	if root.Left == nil && root.Right == nil {
		return []string{strconv.FormatInt(int64(root.Val), 10)}
	}

	var paths []string

	if root.Left != nil {
		leftSubPaths := binaryTreePaths(root.Left)
		for _, path := range leftSubPaths {
			paths = append(paths, strconv.FormatInt(int64(root.Val), 10)+"->"+path)
		}
	}
	if root.Right != nil {
		rightSubPaths := binaryTreePaths(root.Right)
		for _, path := range rightSubPaths {
			paths = append(paths, strconv.FormatInt(int64(root.Val), 10)+"->"+path)
		}
	}

	return paths
}

func binaryTreePaths2(root *TreeNode) []string {
	if root == nil {
		return []string{}
	}

	var paths []string
	nodeStack := []*TreeNode{root}
	pathStack := []string{strconv.FormatInt(int64(root.Val), 10)}
	length := len(nodeStack)

	for length != 0 {
		node := nodeStack[length-1]
		path := pathStack[length-1]

		nodeStack = nodeStack[:length-1]
		pathStack = pathStack[:length-1]

		if node.Left == nil && node.Right == nil {
			paths = append(paths, path)
		}
		if node.Left != nil {
			nodeStack = append(nodeStack, node.Left)
			pathStack = append(pathStack, path+"->"+strconv.FormatInt(int64(node.Left.Val), 10))
		}
		if node.Right != nil {
			nodeStack = append(nodeStack, node.Right)
			pathStack = append(pathStack, path+"->"+strconv.FormatInt(int64(node.Right.Val), 10))
		}

		length = len(nodeStack)
	}

	return paths
}
