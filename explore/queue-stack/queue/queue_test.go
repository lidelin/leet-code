package queue

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestQueue(t *testing.T) {
	ast := assert.New(t)

	queue := NewQueue()

	ast.NotNil(queue)
	ast.IsType(&Queue{}, queue)

	queue.EnQueue(5)

	ast.False(queue.IsEmpty())
	ast.Equal(5, queue.Front())

	queue.EnQueue(3)
	queue.DeQueue()
	ast.False(queue.IsEmpty())
	ast.Equal(3, queue.Front())

	queue.DeQueue()
	ast.True(queue.IsEmpty())
}
