package _024_swap_nodes_in_pairs

import (
	"github.com/stretchr/testify/assert"
	"leet-code/utils"
	"testing"
)

func TestSwapPairs(t *testing.T) {
	head := utils.MakeList([]int{2, 1, 4, 3})
	expected := utils.MakeList([]int{1, 2, 3, 4})
	actual := SwapPairs(head)
	assert.Equal(t, expected, actual)

	head = utils.MakeList([]int{1, 2, 3, 4})
	expected = utils.MakeList([]int{2, 1, 4, 3})
	actual = SwapPairs(head)
	assert.Equal(t, expected, actual)

	head = utils.MakeList([]int{1, 2, 3})
	expected = utils.MakeList([]int{2, 1, 3})
	actual = SwapPairs(head)
	assert.Equal(t, expected, actual)

	head = utils.MakeList([]int{1, 2})
	expected = utils.MakeList([]int{2, 1})
	actual = SwapPairs(head)
	assert.Equal(t, expected, actual)

	head = utils.MakeList([]int{1})
	expected = utils.MakeList([]int{1})
	actual = SwapPairs(head)
	assert.Equal(t, expected, actual)
}
