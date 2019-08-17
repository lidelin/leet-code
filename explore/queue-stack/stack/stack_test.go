package stack

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStack(t *testing.T) {
	ast := assert.New(t)

	stack := NewStack()
	ast.NotNil(stack)
	ast.Nil(stack.Top())

	stack.Push(1)
	ast.Equal(1, stack.Top())

	ast.Equal(1, stack.Pop())
	ast.Nil(stack.Top())

	limit := 1023
	for i := 0; i < limit; i++ {
		stack.Push(i)
		ast.Equal(i, stack.Top())
	}

	stack.Push(1023)
	stack.Push(1024)

	for i := 1024; i >= 0; i-- {
		ast.Equal(i, stack.Top())
		ast.Equal(i, stack.Pop())
	}

	ast.Nil(stack.Top())
}
