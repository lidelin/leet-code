package stack

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMinStack(t *testing.T) {
	ast := assert.New(t)

	stack := Constructor()
	ast.NotNil(stack)

	stack.Push(-2)
	ast.Equal(-2, stack.Top())
	ast.Equal(-2, stack.GetMin())

	stack.Push(0)
	ast.Equal(0, stack.Top())
	ast.Equal(-2, stack.GetMin())

	stack.Push(-3)
	ast.Equal(-3, stack.Top())
	ast.Equal(-3, stack.GetMin())

	stack.Pop()
	ast.Equal(0, stack.Top())
	ast.Equal(-2, stack.GetMin())
}
