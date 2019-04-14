package queue_stack

import (
	"leet-code/explore/queue-stack/queue"
	"leet-code/utils"
)

type Node = utils.TreeNode

func Bfs(root *Node, target int) int {
	q := queue.NewQueue()
	step := 0

	q.EnQueue(root)

	for !q.IsEmpty() {
		step++

		size := q.Size()
		for i := 0; i < size; i++ {
			current := q.Front().(*Node)
			if current.Val == target {
				return step
			}

			if current.Left != nil {
				q.EnQueue(current.Left)
			}
			if current.Right != nil {
				q.EnQueue(current.Right)
			}
		}

		q.DeQueue()
	}

	return -1
}
