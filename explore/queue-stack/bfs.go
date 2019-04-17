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

func Bfs2(root *Node, target int) int {
	used := make(map[*Node]bool)
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
				if _, ok := used[current.Left]; !ok {
					q.EnQueue(current.Left)
					used[current.Left] = true
				}
			}
			if current.Right != nil {
				if _, ok := used[current.Right]; !ok {
					q.EnQueue(current.Right)
					used[current.Right] = true
				}
			}
		}

		q.DeQueue()
	}

	return -1
}
