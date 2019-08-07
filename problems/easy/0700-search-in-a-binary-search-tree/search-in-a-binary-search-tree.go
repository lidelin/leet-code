package _700_search_in_a_binary_search_tree

import (
	q "github.com/lidelin/leet-code/explore/queue-stack/queue"
	"github.com/lidelin/leet-code/utils"
)

type TreeNode = utils.TreeNode

func SearchBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return root
	}

	if root.Val == val {
		return root
	}

	if root.Val > val && root.Left != nil {
		return SearchBST(root.Left, val)
	}

	if root.Val < val && root.Right != nil {
		return SearchBST(root.Right, val)
	}

	return nil
}

func SearchBST2(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}

	queue := q.NewQueue()
	queue.EnQueue(root)

	for !queue.IsEmpty() {
		node := queue.Front().(*TreeNode)
		if node != nil {
			if node.Val == val {
				return node
			} else {
				queue.EnQueue(node.Left)
				queue.EnQueue(node.Right)
			}
		}

		queue.DeQueue()
	}

	return nil
}
