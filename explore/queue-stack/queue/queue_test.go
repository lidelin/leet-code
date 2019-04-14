package queue

import (
	"github.com/stretchr/testify/assert"
	"leet-code/utils"
	"testing"
)

type Node = utils.TreeNode

func TestQueue(t *testing.T) {
	ast := assert.New(t)

	queue := NewQueue()

	ast.NotNil(queue)
	ast.IsType(&Queue{}, queue)

	ast.True(queue.EnQueue(5))

	ast.False(queue.IsEmpty())
	ast.Equal(5, queue.Front())

	ast.True(queue.EnQueue(3))
	ast.True(queue.DeQueue())
	ast.False(queue.IsEmpty())
	ast.Equal(3, queue.Front())

	ast.True(queue.DeQueue())
	ast.True(queue.IsEmpty())

	ast.False(queue.DeQueue())

	node := &Node{Val: 1, Left: nil, Right: nil}
	ast.True(queue.EnQueue(node))

	ast.IsType(node, queue.Front())
}
