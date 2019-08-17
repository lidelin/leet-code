package stack

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStack(t *testing.T) {
	ast := assert.New(t)

	stack := NewStack()
	ast.NotNil(stack)
	ast.True(stack.isEmpty())
	ast.False(stack.isFull())
	ast.Nil(stack.Peek())

	stack.Push(1)
	ast.False(stack.isEmpty())
	ast.False(stack.isFull())
	ast.Equal(1, stack.Peek())

	ast.Equal(1, stack.Pop())
	ast.True(stack.isEmpty())
	ast.False(stack.isFull())

	limit := 1023
	for i := 0; i < limit; i++ {
		stack.Push(i)
		ast.False(stack.isEmpty())
		ast.False(stack.isFull())
		ast.Equal(i, stack.Peek())
	}

	stack.Push(1023)
	ast.False(stack.isEmpty())
	ast.True(stack.isFull())

	stack.Push(1024)
	ast.False(stack.isFull())

	for i := 1024; i >= 0; i-- {
		ast.Equal(i, stack.Peek())
		ast.Equal(i, stack.Pop())
	}

	ast.True(stack.isEmpty())
	ast.False(stack.isFull())
	ast.Nil(stack.Peek())
}
