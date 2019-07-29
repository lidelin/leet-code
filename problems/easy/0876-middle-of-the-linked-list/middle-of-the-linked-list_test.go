package _876_middle_of_the_linked_list

import (
	"github.com/lidelin/leet-code/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMiddleNode(t *testing.T) {
	head := utils.MakeList([]int{1, 2, 3, 4, 5})
	middle := utils.MakeList([]int{3, 4, 5})
	result := MiddleNode(head)
	assert.Equal(t, result, middle)

	head = utils.MakeList([]int{1, 2, 3, 4, 5, 6})
	middle = utils.MakeList([]int{4, 5, 6})
	result = MiddleNode(head)
	assert.Equal(t, result, middle)

	head = utils.MakeList([]int{1, 2})
	middle = utils.MakeList([]int{2})
	result = MiddleNode(head)
	assert.Equal(t, result, middle)

	head = utils.MakeList([]int{1})
	middle = utils.MakeList([]int{1})
	result = MiddleNode(head)
	assert.Equal(t, result, middle)
}
