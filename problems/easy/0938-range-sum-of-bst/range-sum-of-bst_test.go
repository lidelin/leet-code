package _0938_range_sum_of_bst

import (
	"github.com/stretchr/testify/assert"
	"leet-code/utils"
	"testing"
)

func TestRangeSumBST(t *testing.T) {
	root := utils.MakeTree([]int{10, 5, 15, 3, 7, 18})
	actual := RangeSumBST(root, 7, 15)
	expected := 32
	assert.Equal(t, expected, actual)

	root = utils.MakeTree([]int{10, 5, 15, 3, 7, 13, 18, 1, 6})
	actual = RangeSumBST(root, 6, 10)
	expected = 23
	assert.Equal(t, expected, actual)
}
