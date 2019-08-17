package queue_stack

import (
	s "github.com/lidelin/leet-code/explore/queue-stack/stack"
	"github.com/lidelin/leet-code/utils"
)

type TreeNode = utils.TreeNode

func Dfs(root *TreeNode, target int) bool {
	visited := make(map[*TreeNode]bool)
	stack := s.NewStack()
	stack.Push(root)
	visited[root] = true

	for stack.Top() != nil {
		current := stack.Pop().(*TreeNode)
		if current.Val == target {
			return true
		} else {
			if current.Left != nil {
				if _, ok := visited[current.Left]; !ok {
					stack.Push(current.Left)
					visited[current.Left] = true
				}
			}
			if current.Right != nil {
				if _, ok := visited[current.Right]; !ok {
					stack.Push(current.Right)
					visited[current.Right] = true
				}
			}
		}
	}

	return false
}

func Dfs2(root *TreeNode, target int) bool {
	if root == nil {
		return false
	}

	visited := make(map[*TreeNode]bool)

	return dfs(root, target, visited)
}

func dfs(root *TreeNode, target int, visited map[*TreeNode]bool) bool {
	if root == nil {
		return false
	}

	if root.Val == target {
		return true
	}

	if root.Left != nil {
		if _, ok := visited[root.Left]; !ok {
			visited[root.Left] = true
			if true == dfs(root.Left, target, visited) {
				return true
			}
		}
	}

	if root.Right != nil {
		if _, ok := visited[root.Right]; !ok {
			visited[root.Right] = true
			if true == dfs(root.Right, target, visited) {
				return true
			}
		}
	}

	return false
}
