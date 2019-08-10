package _108_convert_sorted_array_to_binary_search_tree

import (
	"github.com/lidelin/leet-code/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSortedArrayToBST(t *testing.T) {
	nums := []int{-10, -3, 0, 5, 9}
	expected := []int{0, -3, -10, 9, 5}
	actual := utils.PreOrderTraversal(SortedArrayToBST(nums))
	assert.Equal(t, expected, actual)
}
