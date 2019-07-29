package _206_reverse_linked_list

import (
	"github.com/lidelin/leet-code/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReverseList(t *testing.T) {
	head := utils.MakeList([]int{1, 2, 3, 4, 5, 6, 6})
	reversed := utils.MakeList([]int{6, 6, 5, 4, 3, 2, 1})
	head = ReverseList(head)
	assert.Equal(t, reversed, head)

	head = utils.MakeList([]int{1, 2})
	reversed = utils.MakeList([]int{2, 1})
	head = ReverseList(head)
	assert.Equal(t, reversed, head)

	head = utils.MakeList([]int{1})
	reversed = utils.MakeList([]int{1})
	head = ReverseList(head)
	assert.Equal(t, reversed, head)
}

func TestReverseListRecursively(t *testing.T) {
	head := utils.MakeList([]int{1, 2, 3, 4, 5, 6, 6})
	reversed := utils.MakeList([]int{6, 6, 5, 4, 3, 2, 1})
	head = ReverseListRecursively(head)
	utils.PrintList(head)
	assert.Equal(t, reversed, head)

	head = utils.MakeList([]int{1, 2})
	reversed = utils.MakeList([]int{2, 1})
	head = ReverseListRecursively(head)
	assert.Equal(t, reversed, head)

	head = utils.MakeList([]int{1})
	reversed = utils.MakeList([]int{1})
	head = ReverseListRecursively(head)
	assert.Equal(t, reversed, head)
}
