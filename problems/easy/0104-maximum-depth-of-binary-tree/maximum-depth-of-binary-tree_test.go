package _0104_maximum_depth_of_binary_tree

import (
	"github.com/lidelin/leet-code/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMaxDepth(t *testing.T) {
	root := utils.MakeTree([]int{1, 2, 3, 4, 5, 6, 7, 8})
	actual := MaxDepth(root)
	expected := 8
	assert.Equal(t, expected, actual)

	root = utils.MakeTree([]int{1})
	actual = MaxDepth(root)
	expected = 1
	assert.Equal(t, expected, actual)

	root = utils.MakeTree([]int{})
	actual = MaxDepth(root)
	expected = 0
	assert.Equal(t, expected, actual)
}
