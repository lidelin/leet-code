package queue

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCircularQueue(t *testing.T) {
	ast := assert.New(t)

	queue := NewCircularQueue(3)

	ast.True(queue.IsEmpty())
	ast.False(queue.IsFull())
	ast.Equal(-1, queue.Front())
	ast.Equal(-1, queue.Rear())

	ast.True(queue.EnQueue(9))
	ast.False(queue.IsEmpty())
	ast.False(queue.IsFull())
	ast.Equal(9, queue.Front())
	ast.Equal(9, queue.Rear())

	ast.True(queue.EnQueue(100))
	ast.Equal(9, queue.Front())
	ast.Equal(100, queue.Rear())

	ast.True(queue.EnQueue(200))
	ast.Equal(9, queue.Front())
	ast.Equal(200, queue.Rear())
	ast.True(queue.IsFull())

	ast.False(queue.EnQueue(300))

	ast.True(queue.DeQueue())
	ast.Equal(100, queue.Front())
	ast.Equal(200, queue.Rear())
	ast.False(queue.IsFull())

	ast.True(queue.DeQueue())
	ast.Equal(200, queue.Front())
	ast.Equal(200, queue.Rear())
	ast.False(queue.IsFull())

	ast.True(queue.DeQueue())
	ast.Equal(-1, queue.Front())
	ast.Equal(-1, queue.Rear())
	ast.False(queue.IsFull())
	ast.True(queue.IsEmpty())

	ast.False(queue.DeQueue())
}
